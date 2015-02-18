package parserreflection

import (
    "errors"
    "strings"
    "log"
    "strconv"
    "reflect"
)

func (model *Model) Imprint(inputModel interface{}) error {
    base := model.Map()
    return imprint(base, inputModel)
}

func imprint(model *map[string]interface{}, inputModel interface{}) error {

    ref := reflect.ValueOf(inputModel)

    if ref.Kind() == reflect.Ptr {
        ref = ref.Elem()
    }

    refType := ref.Type()

    for i := 0; i < ref.NumField(); i++ {

        fieldName := extractFieldName(refType.Field(i).Tag)

        currentValue, ok := (*model)[fieldName]
        if !ok {
            continue // Nothing to do here, we don't care about this value
        }

        if currentValue == nil {
            continue // This value wasn't set in our model, nothing to imprint
        }

        field := ref.Field(i)
        if field.Kind() == reflect.Ptr {
            field = field.Elem()
        }

        switch field.Kind() {

            case reflect.Slice:
                innerModel, ok := (*model)[fieldName].([]interface{})
                if !ok {
                    return errors.New("Input model does not match the parsed document structure (document field '" + fieldName + "' as " + field.Type().String() + ")")
                }
                err := imprintSlice(&innerModel, field.Interface())
                log.Println(innerModel, fieldName)
                (*model)[fieldName] = innerModel
                if err != nil {
                    return err
                }

            case reflect.Struct:
                innerModel, ok := (*model)[fieldName].(map[string]interface{})
                if !ok {
                    return errors.New("Input model does not match the parsed document structure (document field '" + fieldName + "' as " + field.Type().String() + ")")
                }
                err := imprint(&innerModel, field.Interface())
                (*model)[fieldName] = innerModel
                if err != nil {
                    return err
                }

            default:
                (*model)[fieldName] = field.Interface()

        }

    }

    return nil

}

func imprintSlice(slice interface{}, inputModel interface{}) error {

    refSlice := reflect.ValueOf(slice)
    if refSlice.Kind() == reflect.Ptr {
        refSlice = refSlice.Elem()
    }
    refInputModel := reflect.ValueOf(inputModel)
    if refInputModel.Kind() == reflect.Ptr {
        refInputModel = refInputModel.Elem()
    }
    if refSlice.Kind() != reflect.Slice {
        return errors.New("ImprintSlice can only deal with slices")
    }
    if refInputModel.Kind() != reflect.Slice {
        return errors.New("ImprintSlice can only deal with slices")
    }

    for refInputModel.Len() > refSlice.Len() && refSlice.Len() > 0 {
        newItem := reflect.New(refSlice.Index(0).Elem().Type())
        refSlice.Set(reflect.Append(refSlice, newItem))
    }

    for idx := 0; idx < refInputModel.Len(); idx++ {

        model := refSlice.Index(idx)
        innerInputModel := refInputModel.Index(idx)

        //refType := innerInputModel.Type()
        switch innerInputModel.Kind() {

            case reflect.Slice:
                value, ok := model.Interface().([]interface{})
                if !ok {
                    return errors.New("Input model does not match the parsed document structure (document field index '" + strconv.Itoa(idx) + "' as " + model.Kind().String() + ")")
                }
                return imprintSlice(&value, innerInputModel.Interface())

            case reflect.Struct:
                value, ok := model.Interface().(map[string]interface{})
                if !ok {
                    return errors.New("Input model does not match the parsed document structure (document field index '" + strconv.Itoa(idx) + "' as " + model.Kind().String() + ")")
                }
                return imprint(&value, innerInputModel.Interface())

            default:
                model.Set(innerInputModel)
                return nil

        }

    }

    return errors.New("The slice seems to have been empty")

}

func extractFieldName(tag reflect.StructTag) string {
    json := tag.Get("json")
    parts := strings.Split(json, ",")
    if len(parts) > 0 {
        return parts[0]
    }
    if json != "" {
        return json
    }
    yaml := tag.Get("yaml")
    parts = strings.Split(yaml, ",")
    if len(parts) > 0 {
        return parts[0]
    }
    return yaml
}
