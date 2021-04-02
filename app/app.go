package main

import (
        "fmt"
        "github.com/kounkou/crotal"
)

func main() {
    keyId := "user_1"
    r, err := crotal.Crotal(keyId)

    if err != nil {
        fmt.Println(err)
    }

    fmt.Println(r)
}
