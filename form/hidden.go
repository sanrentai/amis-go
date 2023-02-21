package form

// 隐藏字段
func Hidden(opts ...opt) map[string]interface{} {
	return newForm("hidden", opts...)
}

