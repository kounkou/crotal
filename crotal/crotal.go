package crotal

import (
        "encoding/json"
        "os"
        "io/ioutil"
)

const CONFIG_FILENAME string = "crotal.json"
const NOT_FOUND_ERROR int    = -1

func readJSONToken(filename string, keyId string) (int, error) {
    data, err := ioutil.ReadFile(filename)

    if err != nil {
        return NOT_FOUND_ERROR, err
    }

    obj := map[string]int{}
    if err := json.Unmarshal(data, &obj); err != nil {
        return NOT_FOUND_ERROR, err
    }

    if v, s := obj[keyId]; s {
        return v, err
    }

    return NOT_FOUND_ERROR, err
}

func GetLimit(keyId string) (int, error) {
    if _, err := os.Stat(CONFIG_FILENAME); os.IsNotExist(err) {
        return NOT_FOUND_ERROR, err
    }

    return readJSONToken(CONFIG_FILENAME, keyId)
}

// crotal takes your keyId and provides the limit
func Crotal(keyId string) (int, error) {
    return GetLimit(keyId)
}
