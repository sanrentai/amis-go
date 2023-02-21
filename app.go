package amis

// 多页应用
func App(opts ...opt) map[string]interface{} {
	return newCompent("app", opts...)
}
