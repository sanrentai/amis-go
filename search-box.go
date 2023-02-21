package amis

// Search
func Search(opts ...opt) map[string]interface{} {
	return newCompent("search-box", opts...)
}

// 外层 CSS 类名
func Search_className(p string) opt {
	return func(o map[string]interface{}) {
		o["className"] = p
	}
}

// 是否为 mini 模式
func Search_mini(p bool) opt {
	return func(o map[string]interface{}) {
		o["mini"] = p
	}
}

// 是否立即搜索
func Search_searchImediately(p bool) opt {
	return func(o map[string]interface{}) {
		o["searchImediately"] = p
	}
}
