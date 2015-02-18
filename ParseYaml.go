package parserreflection

import (
    "io/ioutil"
    "gopkg.in/yaml.v2"
)

func ParseYaml(filename string) (*Model, error) {
    buf, err := ioutil.ReadFile(filename)
    if err != nil {
        return nil, err
    }
    model := NewModel()
    err = yaml.Unmarshal(buf, model)
    return model, err
}
