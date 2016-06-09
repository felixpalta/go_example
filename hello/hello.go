package main

import (
    "fmt"
    "github.com/lolportal/go_example/stringutil"
)

func main() {

    s := "Hello, няша!"
    fmt.Println(s)
    fmt.Println("Reversed:", stringutil.Reverse(s))
}
