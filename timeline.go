package amis

// 时间轴
func Timeline(opts ...opt) map[string]interface{} {
	return newCompent("timeline", opts...)
}

// 配置节点数据
func Timeline_items(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["items"] = p
	}
}

// 数据源，可通过数据映射获取当前数据域变量、或者配置 API 对象
func Timeline_source(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["source"] = p
	}
}

// 根据时间倒序显示
func Timeline_reverse(p bool) opt {
	return func(o map[string]interface{}) {
		o["reverse"] = p
	}
}

// 说明
func Timeline_属性名(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["属性名"] = p
	}
}

// 节点时间
func Timeline_time(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["time"] = p
	}
}

// 节点详细描述（折叠）
func Timeline_detail(p string) opt {
	return func(o map[string]interface{}) {
		o["detail"] = p
	}
}

// 详细内容折叠时按钮文案
func Timeline_detailCollapsedText(p string) opt {
	return func(o map[string]interface{}) {
		o["detailCollapsedText"] = p
	}
}

// 详细内容展开时按钮文案
func Timeline_detailExpandedText(p string) opt {
	return func(o map[string]interface{}) {
		o["detailExpandedText"] = p
	}
}

// icon 名，支持 fontawesome v4 或使用 url（优先级高于 color）
func Timeline_icon(p string) opt {
	return func(o map[string]interface{}) {
		o["icon"] = p
	}
}
