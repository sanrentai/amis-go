package amis

// 分页组件
func Pagination(opts ...opt) map[string]interface{} {
	return newCompent("pagination", opts...)
}

// 最多显示多少个分页按钮，最小为5
func Pagination_maxButtons(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["maxButtons"] = p
	}
}

// 总页数 （设置总条数total的时候，lastPage会重新计算）
func Pagination_lastPage(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["lastPage"] = p
	}
}

// 总条数
func Pagination_total(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["total"] = p
	}
}

// 当前页数
func Pagination_activePage(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["activePage"] = p
	}
}

// 每页显示多条数据
func Pagination_perPage(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["perPage"] = p
	}
}

// 是否展示perPage切换器 layout和showPerPage都可以控制
func Pagination_showPerPage(p bool) opt {
	return func(o map[string]interface{}) {
		o["showPerPage"] = p
	}
}

// 指定每页可以显示多少条
func Pagination_perPageAvailable(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["perPageAvailable"] = p
	}
}

// 是否显示快速跳转输入框  layout和showPageInput都可以控制
func Pagination_showPageInput(p bool) opt {
	return func(o map[string]interface{}) {
		o["showPageInput"] = p
	}
}

// 是否禁用
func Pagination_disabled(p bool) opt {
	return func(o map[string]interface{}) {
		o["disabled"] = p
	}
}

// 分页改变触发
func Pagination_onPageChange(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["onPageChange"] = p
	}
}
