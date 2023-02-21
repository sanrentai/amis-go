package form

// 组合条件
func 组合条件(opts ...opt) map[string]interface{} {
	return newForm("condition-builder", opts...)
}



// 外层 dom 类名
func 组合条件_className(p string) opt {
	return func(o map[string]interface{}) {
		o["className"] = p
	}
} 

// 输入字段的类名
func 组合条件_fieldClassName(p string) opt {
	return func(o map[string]interface{}) {
		o["fieldClassName"] = p
	}
} 

// 通过远程拉取配置项
func 组合条件_source(p string) opt {
	return func(o map[string]interface{}) {
		o["source"] = p
	}
} 

// 内嵌展示
func 组合条件_embed(p bool) opt {
	return func(o map[string]interface{}) {
		o["embed"] = p
	}
} 

// 弹窗配置的顶部标题
func 组合条件_title(p string) opt {
	return func(o map[string]interface{}) {
		o["title"] = p
	}
} 

// 字段配置
func 组合条件_fields(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["fields"] = p
	}
} 

// 用于 simple 模式下显示切换按钮
func 组合条件_showANDOR(p bool) opt {
	return func(o map[string]interface{}) {
		o["showANDOR"] = p
	}
} 

// 是否显示「非」按钮
func 组合条件_showNot(p bool) opt {
	return func(o map[string]interface{}) {
		o["showNot"] = p
	}
} 

// 字段是否可搜索
func 组合条件_searchable(p bool) opt {
	return func(o map[string]interface{}) {
		o["searchable"] = p
	}
} 

// 组合条件左侧选项类型
func 组合条件_selectMode(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["selectMode"] = p
	}
} 