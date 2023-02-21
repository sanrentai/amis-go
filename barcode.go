package amis

// 条形码
func BarCode(opts ...opt) map[string]interface{} {
	return newCompent("barcode", opts...)
}

// 外层 CSS 类名
func BarCode_className(p string) opt {
	return func(o map[string]interface{}) {
		o["className"] = p
	}
}

// 显示的颜色值
func BarCode_value(p string) opt {
	return func(o map[string]interface{}) {
		o["value"] = p
	}
}

// 在其他组件中，时，用作变量映射
func BarCode_name(p string) opt {
	return func(o map[string]interface{}) {
		o["name"] = p
	}
}
