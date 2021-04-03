package main

import (
        "fmt"
        "github.com/kounkou/crotal"
)

func main() {
    account := "019876543210"
    r, err := crotal.Crotal(account)

    if err != nil {
        fmt.Println(err)
    }

    fmt.Println(r)
}
