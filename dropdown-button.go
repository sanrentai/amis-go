package amis

// 下拉菜单
func DropDownButton(opts ...opt) map[string]interface{} {
	return newCompent("dropdown-button", opts...)
}

// 按钮文本
func DropDownButton_label(p string) opt {
	return func(o map[string]interface{}) {
		o["label"] = p
	}
}

// 外层 CSS 类名
func DropDownButton_className(p string) opt {
	return func(o map[string]interface{}) {
		o["className"] = p
	}
}

// 按钮 CSS 类名
func DropDownButton_btnClassName(p string) opt {
	return func(o map[string]interface{}) {
		o["btnClassName"] = p
	}
}

// 下拉菜单 CSS 类名
func DropDownButton_menuClassName(p string) opt {
	return func(o map[string]interface{}) {
		o["menuClassName"] = p
	}
}

// 块状样式
func DropDownButton_block(p bool) opt {
	return func(o map[string]interface{}) {
		o["block"] = p
	}
}

// 尺寸，支持`'xs'`、`'sm'`、`'md'` 、`'lg'`
func DropDownButton_size(p string) opt {
	return func(o map[string]interface{}) {
		o["size"] = p
	}
}

// 位置，可选`'left'`或`'right'`
func DropDownButton_align(p string) opt {
	return func(o map[string]interface{}) {
		o["align"] = p
	}
}

// 配置下拉按钮
func DropDownButton_buttons(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["buttons"] = p
	}
}

// 只显示 icon
func DropDownButton_iconOnly(p bool) opt {
	return func(o map[string]interface{}) {
		o["iconOnly"] = p
	}
}

// 默认是否打开
func DropDownButton_defaultIsOpened(p bool) opt {
	return func(o map[string]interface{}) {
		o["defaultIsOpened"] = p
	}
}

// 点击外侧区域是否收起
func DropDownButton_closeOnOutside(p bool) opt {
	return func(o map[string]interface{}) {
		o["closeOnOutside"] = p
	}
}

// 点击按钮后自动关闭下拉菜单
func DropDownButton_closeOnClick(p bool) opt {
	return func(o map[string]interface{}) {
		o["closeOnClick"] = p
	}
}

// 触发方式
func DropDownButton_trigger(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["trigger"] = p
	}
}

// 隐藏下拉图标
func DropDownButton_hideCaret(p bool) opt {
	return func(o map[string]interface{}) {
		o["hideCaret"] = p
	}
}
