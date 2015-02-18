# go-parser-reflection
Have you ever wanted to Unmarshal a whole JSON / Yaml document in Golang and only update parts of it using your model (struct), before Marshaling it back to disk again?

In this lib we work with map[string]interface{} and use the reflection package to "imprint" or "extract" your models (struct) onto / off of.
