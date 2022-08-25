package kafka

import (
    "fmt"
    "github.com/Shopify/sarama"
    "github.com/eapache/go-resiliency/breaker"
    "github.com/pkg/errors"
    "log"
    "os"
    "os/signal"
    "sync"
    "syscall"
    "time"
)

type kafkaProducer struct {
    Name       string
    Hosts      []string
    Config     *sarama.Config
    Status     string
    Breaker    *breaker.Breaker
    ReConnect  chan bool
    StatusLock sync.Mutex
}

type Per interface {
    n()
}

type KafkaMsg struct {
    Topic     string
    KeyBytes  []byte
    DataBytes []byte
}

// 同步生产者
type SyncProducer struct {
    kafkaProducer
    SyncProducer *sarama.SyncProducer
}

// 同步生产者
type AsyncProducer struct {
    kafkaProducer
    AsyncProducer *sarama.AsyncProducer
}

type stdLogger interface {
    Print(v ...any)
    Printf(format string, v ...any)
    Println(v ...any)
}

const (
    // 生产者已连接
    KafkaProducerConnected string = "connected"
    // 生产者断开
    KafkaProducerDisconnected string = "disconnected"
    // 生产者关闭
    KafkaProducerClosed string = "closed"

    DefaultKafkaSyncProducer  = "default-kafka-sync-producer"
    DefaultKafkaAsyncProducer = "default-kafka-async-producer"
)

var (
    ErrProducerTimeout  = errors.New("push msg timeout")
    kafkaSyncProducers  = make(map[string]*SyncProducer)
    kafkaAsyncProducers = make(map[string]*AsyncProducer)
    KafkaStdLogger      stdLogger
)

func init() {
    KafkaStdLogger = log.New(os.Stdout, "[kafka] ", log.LstdFlags|log.Lshortfile)
}

//kafka默认生产者配置
func getDefaultProducerConfig(clientID string) (config *sarama.Config) {
    config = sarama.NewConfig()
    config.ClientID = clientID
    config.Version = sarama.V2_0_0_0

    config.Net.DialTimeout = time.Second * 30
    config.Net.WriteTimeout = time.Second * 30
    config.Net.ReadTimeout = time.Second * 30

    config.Producer.Retry.Backoff = time.Millisecond * 500
    config.Producer.Retry.Max = 3

    config.Producer.Return.Successes = true
    config.Producer.Return.Errors = true

    //需要小于broker的 `message.max.bytes`配置，默认是1000000
    config.Producer.MaxMessageBytes = 1000000 * 2

    config.Producer.RequiredAcks = sarama.WaitForLocal
    //config.Producer.Partitioner = sarama.NewRandomPartitioner
    //config.Producer.Partitioner = sarama.NewRoundRobinPartitioner
    config.Producer.Partitioner = sarama.NewHashPartitioner
    //config.Producer.Partitioner = sarama.NewReferenceHashPartitioner

    // zstd 算法有着最高的压缩比，而在吞吐量上的表现只能说中规中矩，LZ4 > Snappy > zstd 和 GZIP
    //LZ4 具有最高的吞吐性能，压缩比zstd > LZ4 > GZIP > Snappy
    //综上LZ4性价比最高
    config.Producer.Compression = sarama.CompressionLZ4
    return
}

func InitSyncKafkaProducer(name string, hosts []string, config *sarama.Config) error {
    syncProducer := &SyncProducer{}
    syncProducer.Name = name
    syncProducer.Hosts = hosts
    syncProducer.Status = KafkaProducerDisconnected
    if config == nil {
        config = getDefaultProducerConfig(name)
    }
    syncProducer.Config = config

    if producer, err := sarama.NewSyncProducer(hosts, config); err != nil {
        return errors.Wrap(err, "NewSyncProducer error name"+name)
    } else {

        syncProducer.Breaker = breaker.New(3, 1, 2*time.Second)
        syncProducer.ReConnect = make(chan bool)
        syncProducer.SyncProducer = &producer
        syncProducer.Status = KafkaProducerConnected
        fmt.Println("SyncKakfaProducer connected name " + name)
    }
    go syncProducer.keepConnect()
    go syncProducer.check()
    kafkaSyncProducers[name] = syncProducer
    return nil
}

//初始化异步生产者
func InitAsyncKafkaProducer(name string, hosts []string, config *sarama.Config) error {
    asyncProducer := &AsyncProducer{}
    asyncProducer.Name = name
    asyncProducer.Hosts = hosts
    asyncProducer.Status = KafkaProducerDisconnected
    if config == nil {
        config = getDefaultProducerConfig(name)
    }
    asyncProducer.Config = config

    if producer, err := sarama.NewAsyncProducer(hosts, config); err != nil {
        return errors.Wrap(err, "NewAsyncProducer error name"+name)
    } else {

        asyncProducer.Breaker = breaker.New(3, 1, 5*time.Second)
        asyncProducer.ReConnect = make(chan bool)
        asyncProducer.AsyncProducer = &producer
        asyncProducer.Status = KafkaProducerConnected
        KafkaStdLogger.Println("AsyncKakfaProducer  connected name ", name)
    }

    go asyncProducer.keepConnect()
    go asyncProducer.check()
    kafkaAsyncProducers[name] = asyncProducer
    return nil
}

func GetKafkaSyncProducer(name string) *SyncProducer {
    if producer, ok := kafkaSyncProducers[name]; ok {
        return producer
    } else {
        KafkaStdLogger.Println("InitSyncKafkaProducer must be called !")
        return nil
    }
}

func GetKafkaAsyncProducer(name string) *AsyncProducer {
    if producer, ok := kafkaAsyncProducers[name]; ok {
        return producer
    } else {
        KafkaStdLogger.Println("InitAsyncKafkaProducer must be called !")
        return nil
    }
}

//检查同步生产者的连接状态,如果断开链接则尝试重连
func (syncProducer *SyncProducer) keepConnect() {
    defer func() {
        KafkaStdLogger.Println("syncProducer keepConnect exited")
    }()
    signals := make(chan os.Signal, 1)
    signal.Notify(signals, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
    for {
        if syncProducer.Status == KafkaProducerClosed {
            return
        }
        select {
        case <-signals:
            syncProducer.StatusLock.Lock()
            syncProducer.Status = KafkaProducerClosed
            syncProducer.StatusLock.Unlock()
            return
        case <-syncProducer.ReConnect:
            if syncProducer.Status != KafkaProducerDisconnected {
                break
            }

            KafkaStdLogger.Println("kafka syncProducer ReConnecting... name " + syncProducer.Name)
            var producer sarama.SyncProducer
        syncBreakLoop:
            for {
                //利用熔断器给集群以恢复时间，避免不断的发送重联
                err := syncProducer.Breaker.Run(func() (err error) {
                    producer, err = sarama.NewSyncProducer(syncProducer.Hosts, syncProducer.Config)
                    return
                })

                switch err {
                case nil:
                    syncProducer.StatusLock.Lock()
                    if syncProducer.Status == KafkaProducerDisconnected {
                        syncProducer.SyncProducer = &producer
                        syncProducer.Status = KafkaProducerConnected
                    }
                    syncProducer.StatusLock.Unlock()
                    KafkaStdLogger.Println("kafka syncProducer ReConnected, name:", syncProducer.Name)
                    break syncBreakLoop
                case breaker.ErrBreakerOpen:
                    KafkaStdLogger.Println("kafka connect fail, broker is open")
                    //2s后重连，此时breaker刚好 half close
                    if syncProducer.Status == KafkaProducerDisconnected {
                        time.AfterFunc(2*time.Second, func() {
                            KafkaStdLogger.Println("kafka begin to ReConnect ,because of  ErrBreakerOpen ")
                            syncProducer.ReConnect <- true
                        })
                    }
                    break syncBreakLoop
                default:
                    KafkaStdLogger.Println("kafka ReConnect error, name:", syncProducer.Name, err)
                }
            }
        }
    }
}
