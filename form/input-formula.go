package form

// 公式编辑器
func InputFormula(opts ...opt) map[string]interface{} {
	return newForm("input-formula", opts...)
}



// 弹框标题
func InputFormula_title(p string) opt {
	return func(o map[string]interface{}) {
		o["title"] = p
	}
} 

// 编辑器 header 标题，如果不设置，默认使用表单项`label`字段
func InputFormula_header(p string) opt {
	return func(o map[string]interface{}) {
		o["header"] = p
	}
} 

// 表达式模式 或者 模板模式，模板模式则需要将表达式写在 `${` 和 `}` 中间。
func InputFormula_evalMode(p bool) opt {
	return func(o map[string]interface{}) {
		o["evalMode"] = p
	}
} 

// 可用变量
func InputFormula_variables(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["variables"] = p
	}
} 

// 可配置成 `tabs` 或者 `tree` 默认为列表，支持分组。
func InputFormula_variableMode(p string) opt {
	return func(o map[string]interface{}) {
		o["variableMode"] = p
	}
} 

// 可以不设置，默认就是 amis-formula 里面定义的函数，如果扩充了新的函数则需要指定
func InputFormula_functions(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["functions"] = p
	}
} 

// 按钮图标，例如`fa fa-list`
func InputFormula_icon(p string) opt {
	return func(o map[string]interface{}) {
		o["icon"] = p
	}
} 

// 按钮文本，`inputMode`为`button`时生效
func InputFormula_btnLabel(p string) opt {
	return func(o map[string]interface{}) {
		o["btnLabel"] = p
	}
} 

// 输入框是否可输入
func InputFormula_allowInput(p bool) opt {
	return func(o map[string]interface{}) {
		o["allowInput"] = p
	}
} 

// 输入框占位符
func InputFormula_placeholder(p string) opt {
	return func(o map[string]interface{}) {
		o["placeholder"] = p
	}
} 

// 控件外层 CSS 样式类名
func InputFormula_className(p string) opt {
	return func(o map[string]interface{}) {
		o["className"] = p
	}
} 

// 变量面板 CSS 样式类名
func InputFormula_variableClassName(p string) opt {
	return func(o map[string]interface{}) {
		o["variableClassName"] = p
	}
} 

// 函数面板 CSS 样式类名
func InputFormula_functionClassName(p string) opt {
	return func(o map[string]interface{}) {
		o["functionClassName"] = p
	}
} 