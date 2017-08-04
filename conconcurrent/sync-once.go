/*
    brief : use channel to simulate semaphore
    author : michaelyqc
    date : 2017/08/04
 */

package main

import (
    "fmt"
    "sync"
)

var (
    once sync.Once
)

func initial() {
    fmt.Println("having initialized")
}

func InitializeOnce() {
    once.Do(initial)
}


func main() {
    InitializeOnce()
    InitializeOnce()
}


