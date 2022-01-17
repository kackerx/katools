package mq

import (
    "fmt"
    "github.com/streadway/amqp"
    "log"
)

const MQURL = "amqp://admin:123456@10.0.12.3:5672"

type RabbitMQ struct {
    conn         *amqp.Connection
    channel      *amqp.Channel
    Mqurl        string
    ExchangeName string
    QueueName    string
    Key          string
}

func (r *RabbitMQ) destroy() (err error) {
    if err = r.channel.Close(); err != nil {
        return err
    }
    if err = r.conn.Close(); err != nil {
        return err
    }
    return nil
}

func NewRabbitMQ(exchangeName string, queueName string, key string) *RabbitMQ {
    conn, err := amqp.Dial(MQURL)
    if err != nil {
        log.Fatalln(err)
    }
    
    channel, err := conn.Channel()
    if err != nil {
        log.Fatalln(err)
    }
    
    return &RabbitMQ{
        conn:         conn,
        channel:      channel,
        Mqurl:        MQURL,
        ExchangeName: exchangeName,
        QueueName:    queueName,
        Key:          key,
    }
}

func NewSimple(queueName string) *RabbitMQ {
    return NewRabbitMQ("", queueName, "")
}

func (mq *RabbitMQ) PublishSimple(msg string) {
    // 防止队列不存在
    _, err := mq.channel.QueueDeclare(mq.QueueName, false, false, false, false, nil)
    if err != nil {
        log.Fatalln(err)
    }
    
    mq.channel.Publish(
        mq.ExchangeName,
        mq.QueueName,
        false, // true: ex和key找不到, 队列, 消息返回发送者
        false, // true: 消息到队列, 队列没有消费者, 消息返回发送者
        amqp.Publishing{
        ContentType:     "application/plain",
        Body:            []byte(msg),
    },
    )
}

func (mq *RabbitMQ) ConsumeSimple() {
    // 防止队列不存在
    _, err := mq.channel.QueueDeclare(mq.QueueName, false, false, false, false, nil)
    if err != nil {
        log.Fatalln(err)
    }
    
    deliveryCh, err := mq.channel.Consume(mq.QueueName, "", true, false, false, false, nil)
    if err != nil {
        log.Fatalln(err)
    }
    
    forever := make(chan struct{})
    go func() {
        for msg := range deliveryCh {
            fmt.Println("get msg: ", string(msg.Body))
        }
    }()
    <-forever
}

