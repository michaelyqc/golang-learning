/*

brief : use channel to simulate semaphore
author : michaelyqc
date : 2017/08/02

 */

package main

import (
    "fmt"
    "runtime"
)

var (
    sema = make(chan struct{}, 1)
    balance int
)

func Deposit(amount int) {
    sema <- struct{}{}
    balance = balance + amount
    <-sema
}

func Balance() int {
    sema <- struct{}{}
    b := balance
    <-sema

    return b
}

func main() {
    go Deposit(100)
    runtime.Gosched()
    b := Balance()

    fmt.Printf("Balance := %d\n", b)
}
