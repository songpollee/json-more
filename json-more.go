package jsonMore
// Example
//
// type Chlid struct {
//   D string `json:"d" jsonMore:"required"`
// }
//
// type Parent struct {
//   A string `json:"a"`
//   B int `json:"b" jsonMore:"required"`
//   C Chlid `json:"c" jsonMore:"required"`
//   E []int `json:"e" jsonMore:"required"`
//   F []Chlid `json:"f" jsonMore:"required"`
// }
//
// func main() {
//   rawJson := []byte(`{"a": "abc", "b": 0, "c": { "d": "xyz" }, "e": [1, 2], "f": [{ "d": "d1" }, { "d": "d2" }]}`)
//   err := jsonMore.ValidateJson(Parent{}, rawJson)
//   fmt.Println(err);
// }


import (
  "fmt"
  "encoding/json"
  "reflect"
  "strings"
  "errors"
)

func GetJsonKey(reflectStructField reflect.StructField) string {
  nameField := reflectStructField.Name
  jsonTag := reflectStructField.Tag.Get("json")
  jsonKey := strings.Split(jsonTag, ",")[0]
  if(jsonKey != "") { return jsonKey }
  return nameField
}

func validate(reflectType reflect.Type, targetMap map[string]interface{}) error {
  for i := 0; i < reflectType.NumField(); i++ {
    currentType := reflectType.Field(i)
    kind := currentType.Type.Kind()
    jsonValidateTag := currentType.Tag.Get("jsonMore")
    jsonKey := GetJsonKey(currentType)
    currentMap, ok := targetMap[jsonKey]
    if(jsonValidateTag == "required") {
      if(!ok) { return errors.New(fmt.Sprintf("invalid missing '%s.%s'", reflectType.Name(), jsonKey)) }
    }
    if(kind.String() == "struct" && ok) {
      err := validate(currentType.Type, currentMap.(map[string]interface{}))
      if(err != nil) { return err }
    } else if(kind.String() == "slice" && ok) {
      kindElem := currentType.Type.Elem().Kind()
      if(kindElem.String() == "struct" && ok) {
        typeElem := currentType.Type.Elem()
        arrayMap := currentMap.([]interface{})
        for j := 0; j < len(arrayMap); j++ {
          err := validate(typeElem, arrayMap[j].(map[string]interface{}))
          if(err != nil) { return err }
        }
      }
    }
  }
  return nil
}

func ValidateMap(targetStruct interface{}, targetMap map[string]interface{}) error {
  targetRealStruct := targetStruct
  if(reflect.TypeOf(targetStruct).Kind() == reflect.Ptr) {
    targetRealStruct = reflect.ValueOf(targetStruct).Elem().Interface()
  }
  reflectStruct := reflect.TypeOf(targetRealStruct)
  return validate(reflectStruct, targetMap)
}

func ValidateJson(targetStruct interface{}, myJson []byte) error {
  var targetMap map[string]interface{}
  err := json.Unmarshal(myJson, &targetMap)
  if(err != nil) { return err }
  targetRealStruct := targetStruct
  if(reflect.TypeOf(targetStruct).Kind() == reflect.Ptr) {
    targetRealStruct = reflect.ValueOf(targetStruct).Elem().Interface()
  }
  return ValidateMap(targetRealStruct, targetMap)
}

func ToStruct(targetStruct interface{}, myJson []byte) error {
  if(reflect.TypeOf(targetStruct).Kind() != reflect.Ptr) { return errors.New("invalid must pass params with address struct") }
  targetRealStruct := reflect.ValueOf(targetStruct).Elem().Interface()
  notMissingField := ValidateJson(targetRealStruct, myJson)
  if(notMissingField != nil) { return notMissingField }
  err := json.Unmarshal(myJson, targetStruct)
  if(err != nil) { return err }
  return nil
}
