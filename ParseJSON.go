package parserreflection

import (
    "io/ioutil"
    "encoding/json"
)

func ParseJSON(filename string) (*Model, error) {
    buf, err := ioutil.ReadFile(filename)
    if err != nil {
        return nil, err
    }
    model := NewModel()
    err = json.Unmarshal(buf, model)
    return model, err
}
