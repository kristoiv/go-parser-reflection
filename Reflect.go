package parserreflection

import (
    "reflect"
)

func (model *Model) Map() *map[string]interface{} {
    typ := reflect.TypeOf(&map[string]interface{}{})
    return reflect.ValueOf(model).Convert(typ).Interface().(*map[string]interface{})
}
