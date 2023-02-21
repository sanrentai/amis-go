package amis

// 轻提示
func Toast(opts ...opt) map[string]interface{} {
	return newCompent("toast", opts...)
}

// 指定为 toast 轻提示组件
func Toast_actionType(p string) opt {
	return func(o map[string]interface{}) {
		o["actionType"] = p
	}
}

// 轻提示内容
func Toast_items(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["items"] = p
	}
}

// 是否展示关闭按钮，移动端不展示
func Toast_closeButton(p bool) opt {
	return func(o map[string]interface{}) {
		o["closeButton"] = p
	}
}

// 持续时间
func Toast_timeout(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["timeout"] = p
	}
}

// 说明
func Toast_属性名(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["属性名"] = p
	}
}

// 展示图标，可选'info'、'success'、'error'、'warning'
func Toast_level(p string) opt {
	return func(o map[string]interface{}) {
		o["level"] = p
	}
}

// 提示显示位置，可用'top-right'、'top-center'、'top-left'、'bottom-center'、'bottom-left'、'bottom-right'、'center'
func Toast_position(p string) opt {
	return func(o map[string]interface{}) {
		o["position"] = p
	}
}

// 是否展示图标
func Toast_showIcon(p bool) opt {
	return func(o map[string]interface{}) {
		o["showIcon"] = p
	}
}
