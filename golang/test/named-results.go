package main

import "fmt"

func split(sum int) (x, y int) {
    x = sum * 4 / 7
    y = sum + 2
    return
}

func main() {
    fmt.Println(split(63))
    fmt.Println(split(63))
}
