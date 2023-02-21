package form

// JSONSchema
func JSONSchema(opts ...opt) map[string]interface{} {
	return newForm("json-schema", opts...)
}

