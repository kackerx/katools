package main

type hashtable struct {
    array [7]empLink
}

func (t *hashtable) insert(e *Emp) {
    hashNo := e.id % 7
    
    t.array[hashNo].insert(e)
}

func (t *hashtable) showAll() {
    for i := 0; i < len(t.array); i++ {
        t.array[i].show(i)
    }
}

func main() {
    h := hashtable{}
    h.insert(&Emp{
        id:   7,
        name: "kacker",
        next: nil,
    })
    
    h.insert(&Emp{
        id:   2,
        name: "liyong",
        next: nil,
    })
    
    h.insert(&Emp{
        id:   1,
        name: "xia",
        next: nil,
    })
    
    h.insert(&Emp{
        id:   0,
        name: "xia",
        next: nil,
    })
    h.showAll()
}
