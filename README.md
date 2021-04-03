[![license](https://img.shields.io/badge/License-MIT-blue.svg)](https://github.com/kounkou/crotal/blob/master/LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/kounkou/crotal)](https://goreportcard.com/report/github.com/kounkou/crotal)

# Crotal
Crotal makes it easy to retrieve the configuration value for a given key.

# Installation
```
go get github.com/kounkou/crotal
go build
```

# Usage
Add a json file called `crotal.json` with your configurations to your 
root directory.

```
{
    "212345678910":1000,
    "119876543210":3,
    "612345677910":5000
}
```

Then use one of you account numbers to query the configuration value.

```
package main

import (
        "fmt"
        "github.com/kounkou/crotal"
)

func main() {
    keyId := "119876543210"
    r, err := crotal.Crotal(keyId)

    if err != nil {
        fmt.Println(err)
    }

    // Use the configuration value in your application
    // ...
}
```
