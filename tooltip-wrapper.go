package amis

// 文字提示容器
func TooltipWrapper(opts ...opt) map[string]interface{} {
	return newCompent("tooltip-wrapper", opts...)
}

// 文字提示标题
func TooltipWrapper_title(p string) opt {
	return func(o map[string]interface{}) {
		o["title"] = p
	}
}

// 文字提示内容, 兼容之前的 tooltip 属性
func TooltipWrapper_content(p string) opt {
	return func(o map[string]interface{}) {
		o["content"] = p
	}
}

// 文字提示浮层位置相对偏移量，单位 px
func TooltipWrapper_offset(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["offset"] = p
	}
}

// 是否展示浮层指向箭头
func TooltipWrapper_showArrow(p bool) opt {
	return func(o map[string]interface{}) {
		o["showArrow"] = p
	}
}

// 是否鼠标可以移入到浮层中
func TooltipWrapper_enterable(p bool) opt {
	return func(o map[string]interface{}) {
		o["enterable"] = p
	}
}

// 是否禁用浮层提示
func TooltipWrapper_disabled(p bool) opt {
	return func(o map[string]interface{}) {
		o["disabled"] = p
	}
}

// 浮层延迟展示时间，单位 ms
func TooltipWrapper_mouseEnterDelay(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["mouseEnterDelay"] = p
	}
}

// 浮层延迟隐藏时间，单位 ms
func TooltipWrapper_mouseLeaveDelay(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["mouseLeaveDelay"] = p
	}
}

// 是否点击非内容区域关闭提示
func TooltipWrapper_rootClose(p bool) opt {
	return func(o map[string]interface{}) {
		o["rootClose"] = p
	}
}

// 内容区是否内联显示
func TooltipWrapper_inline(p bool) opt {
	return func(o map[string]interface{}) {
		o["inline"] = p
	}
}

// 内容容器
func TooltipWrapper_body(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["body"] = p
	}
}

// 内容区类名
func TooltipWrapper_className(p string) opt {
	return func(o map[string]interface{}) {
		o["className"] = p
	}
}

// 文字提示浮层类名
func TooltipWrapper_tooltipClassName(p string) opt {
	return func(o map[string]interface{}) {
		o["tooltipClassName"] = p
	}
}
