package main

import (
    "log"

    "github.com/kristoiv/go-parser-reflection"
)

type DocumentModel struct {
    Test1 *string `json:"test1,omitempty"`
    Test3 *[]string `json:"test3,omitempty"`
    Test4 *DocumentModelNested `json:"test4,omitempty"`
}

type DocumentModelNested struct {
    Test41 *[]DocumentModelNested2 `json:"test42,omitempty"`
}

type DocumentModelNested2 struct {
    Test421 *string `json:"test421,omitempty"`
}

func main() {

    model, err := parserreflection.ParseJSON("example.json")
    if err != nil {
        log.Fatalln(err)
    }

    log.Println(model)

    var1 := "Hellu"
    var2 := []string{"Noe, only one!", "Test32New", "Test33"}
    var3 := "Yeap, me too!"
    documentModel := DocumentModel{
        Test1: &var1,
        Test3: &var2,
        Test4: &DocumentModelNested{
            Test41: &[]DocumentModelNested2{
                DocumentModelNested2{Test421: &var3},
            },
        },
    }

    err = model.Imprint(documentModel)
    if err != nil {
        log.Fatalln(err)
    }

    log.Println(model)

}
