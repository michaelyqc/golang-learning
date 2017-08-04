/* brief : use channel to simulate synchronization author : michaelyqc date : 2017/08/02 */ 
package main

import (
    "fmt"
    "runtime"
    "sync"
    "time"
)

var (
    mu  sync.Mutex  // guards balance
    balance int
)

func Deposit(amount int) {
    mu.Lock()
    defer mu.Unlock()
    balance = balance + amount
}

func Balance() int {
    mu.Lock()
    defer mu.Unlock()
    b := balance

    return b
}

func Withdraw(amount int) bool {
    mu.Lock()
    defer mu.Unlock()
    balance = balance - amount
    if balance < 0 {
        balance = balance + amount
        return false
    } else {
        return true
    }
}

/*
func Withdraw(amount int) bool {
    Deposit(-amount)
    if Balance() < 0 {
        Deposit(amount)
        return false
    } else {
        return true
    }
}
*/

func main() {
    runtime.GOMAXPROCS(3)

    go Deposit(100)
    go Withdraw(200)

    fmt.Printf("Balance : %d\n", Balance() )
    time.Sleep(time.Minute)
}







