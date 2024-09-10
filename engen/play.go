package main

import (
//    "fmt"
    "time"
    "math/rand"
)

func main() {
    screenHeight := 49
    screenWidth := 99

    screen := InitScreen45100()

    charr := make(chan []byte)
    go ReadKey(charr) 

    for true {

        select {
        case key :=  <-charr:
            // logic to button reactions
        
        default: 

            time.Sleep(160*time.Millisecond)
        }
    }   
}

