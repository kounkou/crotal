package main

import (
        "fmt"
        "github.com/kounkou/crotal"
)

func main() {
    keyId := "012345678910"
    r, err := crotal.Crotal(keyId)

    if err != nil {
        fmt.Println(err)
    }

    fmt.Println(r)
}
