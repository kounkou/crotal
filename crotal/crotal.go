package crotal

import (
        "os"
)

func GetLimit(keyId string) (int, error) {
    crt := "crt.conf"
    var err error

    if _, err := os.Stat(crt); os.IsNotExist(err) {
        return -1, err
    }
    return 0, err
}

// crotal takes your keyId and provides the limit
func Crotal(keyId string) (int, error) {
    return GetLimit(keyId)
}
