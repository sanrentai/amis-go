package amis

// 对话框
func Dialog(opts ...opt) map[string]interface{} {
	return newCompent("dialog", opts...)
}

// 弹出层标题
func Dialog_title(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["title"] = p
	}
}

// 往 Dialog 内容区加内容
func Dialog_body(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["body"] = p
	}
}

// 指定 dialog 大小，支持: `xs`、`sm`、`md`、`lg`、`xl`、`full`
func Dialog_size(p string) opt {
	return func(o map[string]interface{}) {
		o["size"] = p
	}
}

// Dialog body 区域的样式类名
func Dialog_bodyClassName(p string) opt {
	return func(o map[string]interface{}) {
		o["bodyClassName"] = p
	}
}

// 是否支持按 `Esc` 关闭 Dialog
func Dialog_closeOnEsc(p bool) opt {
	return func(o map[string]interface{}) {
		o["closeOnEsc"] = p
	}
}

// 是否显示右上角的关闭按钮
func Dialog_showCloseButton(p bool) opt {
	return func(o map[string]interface{}) {
		o["showCloseButton"] = p
	}
}

// 是否在弹框左下角显示报错信息
func Dialog_showErrorMsg(p bool) opt {
	return func(o map[string]interface{}) {
		o["showErrorMsg"] = p
	}
}

// 是否在弹框左下角显示 loading 动画
func Dialog_showLoading(p bool) opt {
	return func(o map[string]interface{}) {
		o["showLoading"] = p
	}
}

// 如果设置此属性，则该 Dialog 只读没有提交操作。
func Dialog_disabled(p bool) opt {
	return func(o map[string]interface{}) {
		o["disabled"] = p
	}
}

// 如果想不显示底部按钮，可以配置：`[]`
func Dialog_actions(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["actions"] = p
	}
}

// 支持[数据映射](../../docs/concepts/data-mapping)，如果不设定将默认将触发按钮的上下文中继承数据。
func Dialog_data(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["data"] = p
	}
}
