# json-more
golang json validator is used for validating missing data in JSON. The way it works is to notify user of missing data.

# License
This library is distributed with [GNU GPLv3](https://spdx.org/licenses/GPL-3.0.html)

## Installation
```bash
go get github.com/songpollee/json-more
```

## Example
```go
import (
  "fmt"
  "github.com/songpollee/json-more"
)

type Chlid struct {
 D string `json:"d" jsonMore:"required"`
}

type Parent struct {
 A string `json:"a"`
 B int `json:"b" jsonMore:"required"`
 C Chlid `json:"c" jsonMore:"required"`
 E []int `json:"e" jsonMore:"required"`
 F []Chlid `json:"f" jsonMore:"required"`
}

func main() {
  rawJson := []byte(`{
    "a": "abc",
    "b": 0,
    "c": { "d": "xyz" },
    "e": [1, 2],
    "f": [{ "d": "c1" }, { "d": "c2" }]
  }`)
  err := jsonMore.ValidateJson(Parent{}, rawJson)
  fmt.Println(err);

  missingBJson := []byte(`{
    "a": "abc",
    "c": { "d": "xyz" },
    "e": [1, 2],
    "f": [{ "d": "c1" }, { "d": "c2" }]
  }`)
  err = jsonMore.ValidateJson(Parent{}, missingBJson)
  fmt.Println(err);
}
```
