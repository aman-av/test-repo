package main

import (
    "fmt"
)

func add(a, b int) int {
    var c int
    var d int
    c=a
    d=b
    e:=c+d
    return e
}

func main() {
    fmt.Println("Hello, World!")
    sum := add(3, 5)
    fmt.Printf("Sum: %d\n", sum)
}
