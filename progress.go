package amis

// 进度条
func Progress(opts ...opt) map[string]interface{} {
	return newCompent("progress", opts...)
}

// 进度「条」的类型，可选`line circle dashboard`
func Progress_mode(p string) opt {
	return func(o map[string]interface{}) {
		o["mode"] = p
	}
}

// 外层 CSS 类名
func Progress_className(p string) opt {
	return func(o map[string]interface{}) {
		o["className"] = p
	}
}

// 进度值
func Progress_value(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["value"] = p
	}
}

// 占位文本
func Progress_placeholder(p string) opt {
	return func(o map[string]interface{}) {
		o["placeholder"] = p
	}
}

// 是否展示进度文本
func Progress_showLabel(p bool) opt {
	return func(o map[string]interface{}) {
		o["showLabel"] = p
	}
}

// 背景是否显示条纹
func Progress_stripe(p bool) opt {
	return func(o map[string]interface{}) {
		o["stripe"] = p
	}
}

// type 为 line，可支持动画
func Progress_animate(p bool) opt {
	return func(o map[string]interface{}) {
		o["animate"] = p
	}
}

// 是否显示阈值（刻度）数值
func Progress_showThresholdText(p bool) opt {
	return func(o map[string]interface{}) {
		o["showThresholdText"] = p
	}
}

// 自定义格式化内容
func Progress_valueTpl(p string) opt {
	return func(o map[string]interface{}) {
		o["valueTpl"] = p
	}
}

// 进度条线宽度
func Progress_strokeWidth(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["strokeWidth"] = p
	}
}

// 仪表盘缺角角度，可取值 0 ~ 295
func Progress_gapDegree(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["gapDegree"] = p
	}
}

// 仪表盘进度条缺口位置，可选`top bottom left right`
func Progress_gapPosition(p string) opt {
	return func(o map[string]interface{}) {
		o["gapPosition"] = p
	}
}
