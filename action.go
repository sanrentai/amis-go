package amis

// 行为按钮
func Action(opts ...opt) map[string]interface{} {
	return newCompent("action", opts...)
}
