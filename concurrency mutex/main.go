package main

import (
    "fmt"
    "sync"
)

var counter = 0
var mu sync.Mutex

func increment(wg *sync.WaitGroup) {
    defer wg.Done()

    mu.Lock()         //  Lock start
    counter++         // safe update
    mu.Unlock()       //  Unlock end
}

func main() {
    var wg sync.WaitGroup

    for i := 0; i < 100; i++ {
        wg.Add(1)
        go increment(&wg)
    }

    wg.Wait()
    fmt.Println("Final Counter:", counter)
}
