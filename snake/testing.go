package main

import "fmt"


func main() {
    data := []string{"Hey", "lekker", "my" , "bra"}
    
    for i := 0; i < 4; i++ {
        fmt.Println(data[i])
    }
}
