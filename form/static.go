package form

// 静态展示
func Static(opts ...opt) map[string]interface{} {
	return newForm("static", opts...)
}

