package amis

// 折叠器
func Collapse(opts ...opt) map[string]interface{} {
	return newCompent("collapse", opts...)
}
