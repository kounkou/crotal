package crotal

import (
        "encoding/json"
        "os"
        "io/ioutil"
)

const CONFIG_FILENAME string = "crotal.json"

func readJSONToken(filename string, keyId string) (int, error) {
    data, err := ioutil.ReadFile(filename)

    if err != nil {
        panic(err)
    }

    obj := map[string]int{}
    if err := json.Unmarshal(data, &obj); err != nil {
        panic(err)
    }

    return obj[keyId], err
}

func GetLimit(keyId string) (int, error) {
    if _, err := os.Stat(CONFIG_FILENAME); os.IsNotExist(err) {
        return -1, err
    }

    return readJSONToken(CONFIG_FILENAME, keyId)
}

// crotal takes your keyId and provides the limit
func Crotal(keyId string) (int, error) {
    return GetLimit(keyId)
}
