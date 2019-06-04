package main
import (
    "fmt"
    "test-go/json-more"
)

type Chlid struct {
  D string `json:"d" jsonValidate:"required"`
}

type Parent struct {
  A string `json:"a"`
  B int `json:"b" jsonValidate:"required"`
  C Chlid `json:"c" jsonValidate:"required"`
  E []int `json:"e" jsonValidate:"required"`
  F []Chlid `json:"f" jsonValidate:"required"`
}

func main() {
  rawJson := []byte(`{"a": "abc", "b": 10, "c": { "d": "xyz" }, "e": [1, 2], "f": [{ "d": "c1" }, { "d": "c2" }]}`)
  err := jsonMore.ValidateJson(Parent{}, rawJson)
  fmt.Println(err);
}
