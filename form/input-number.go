package form

// 数字输入框
func InputNumber(opts ...opt) map[string]interface{} {
	return newForm("input-number", opts...)
}



// 最小值
func InputNumber_min(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["min"] = p
	}
} 

// 最大值
func InputNumber_max(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["max"] = p
	}
} 

// 步长
func InputNumber_step(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["step"] = p
	}
} 

// 精度，即小数点后几位，支持 0 和正整数
func InputNumber_precision(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["precision"] = p
	}
} 

// 是否显示上下点击按钮
func InputNumber_showSteps(p bool) opt {
	return func(o map[string]interface{}) {
		o["showSteps"] = p
	}
} 

// 前缀
func InputNumber_prefix(p string) opt {
	return func(o map[string]interface{}) {
		o["prefix"] = p
	}
} 

// 后缀
func InputNumber_suffix(p string) opt {
	return func(o map[string]interface{}) {
		o["suffix"] = p
	}
} 

// 千分分隔
func InputNumber_kilobitSeparator(p bool) opt {
	return func(o map[string]interface{}) {
		o["kilobitSeparator"] = p
	}
} 

// 键盘事件（方向上下）
func InputNumber_keyboard(p bool) opt {
	return func(o map[string]interface{}) {
		o["keyboard"] = p
	}
} 

// 是否使用大数
func InputNumber_big(p bool) opt {
	return func(o map[string]interface{}) {
		o["big"] = p
	}
} 

// 样式类型
func InputNumber_displayMode(p string) opt {
	return func(o map[string]interface{}) {
		o["displayMode"] = p
	}
} 