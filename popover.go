package amis

// 弹出提示
func PopOver(opts ...opt) map[string]interface{} {
	return newCompent("popover", opts...)
}
