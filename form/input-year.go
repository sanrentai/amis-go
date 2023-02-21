package form

// 年份选择
func Year(opts ...opt) map[string]interface{} {
	return newForm("input-year", opts...)
}

