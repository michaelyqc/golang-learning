/*

brief : use select statement to wait channel readable
author : michaelyqc
date : 2017/08/02

 */

package main

import (
    "fmt"
    "time"
)

func main() {
    c1 := make(chan string)
    c2 := make(chan string)

    go func () {
        time.Sleep(time.Second)
        c1<- "func1"
    }()

    go func () {
        time.Sleep(time.Second)
        c2<- "func2"
    }()

    for i := 0; i < 2; i++ {
        select {
            case msg1 := <-c1:
                fmt.Println("received : ", msg1)
            case msg2 := <-c2:
                fmt.Println("received : ", msg2)
        }
    }
}
