package main

import "fmt"

var c, python, java bool
var c1, python1 int = 1, 2

func main() {
    var i int
    var i1, j1 = 1, "ok"
    m1, m2 := false, "m2"
    fmt.Println(i, c, python, java)
    fmt.Println(c1, python1, i1, j1)
    fmt.Println(m1, m2)
}
