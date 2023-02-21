package amis

// 状态
func Status(opts ...opt) map[string]interface{} {
	return newCompent("status", opts...)
}
