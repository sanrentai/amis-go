package amis

// 日历日程
func Calendar(opts ...opt) map[string]interface{} {
	return newCompent("calendar", opts...)
}
