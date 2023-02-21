package form

// 字段
func UUID(opts ...opt) map[string]interface{} {
	return newForm("uuid", opts...)
}

