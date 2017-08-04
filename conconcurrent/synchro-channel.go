/*

brief : use channel to simulate synchronization
author : michaelyqc
date : 2017/08/02

 */

package main

import (
    "fmt"
    "time"
    "runtime"
)

func Producer(cs chan string) {
    watchword := string("synchro")
    cs<- watchword
}

func Consumer(cs chan string) {
    watchword := <-cs
    fmt.Printf("watchword : %s\n", watchword)
}

func main() {
    cs := make(chan string, 0)

    go Consumer(cs)
    go Producer(cs)

    runtime.Gosched()
    time.Sleep(time.Second*10)
}
