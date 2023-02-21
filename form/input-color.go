package form

// 颜色选择器
func InputColor(opts ...opt) map[string]interface{} {
	return newForm("input-color", opts...)
}



// 请选择 `hex`、`hls`、`rgb`或者`rgba`。
func InputColor_format(p string) opt {
	return func(o map[string]interface{}) {
		o["format"] = p
	}
} 

// 选择器底部的默认颜色，数组内为空则不显示默认颜色
func InputColor_presetColors(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["presetColors"] = p
	}
} 

// 为`false`时只能选择颜色，使用 `presetColors` 设定颜色选择范围
func InputColor_allowCustomColor(p bool) opt {
	return func(o map[string]interface{}) {
		o["allowCustomColor"] = p
	}
} 

// 是否显示清除按钮
func InputColor_clearable(p bool) opt {
	return func(o map[string]interface{}) {
		o["clearable"] = p
	}
} 

// 清除后，表单项值调整成该值
func InputColor_resetValue(p string) opt {
	return func(o map[string]interface{}) {
		o["resetValue"] = p
	}
} 