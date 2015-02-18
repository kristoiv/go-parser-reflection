package parserreflection

type Model map[string]interface{}

func NewModel() *Model {
    return &Model{}
}
