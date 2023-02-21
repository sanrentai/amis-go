package form

// 多行文本输入框
func Textarea(opts ...opt) map[string]interface{} {
	return newForm("textarea", opts...)
}



// 最小行数
func Textarea_minRows(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["minRows"] = p
	}
} 

// 最大行数
func Textarea_maxRows(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["maxRows"] = p
	}
} 

// 是否去除首尾空白文本
func Textarea_trimContents(p bool) opt {
	return func(o map[string]interface{}) {
		o["trimContents"] = p
	}
} 

// 是否只读
func Textarea_readOnly(p bool) opt {
	return func(o map[string]interface{}) {
		o["readOnly"] = p
	}
} 

// 是否显示计数器
func Textarea_showCounter(p bool) opt {
	return func(o map[string]interface{}) {
		o["showCounter"] = p
	}
} 

// 限制最大字数
func Textarea_maxLength(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["maxLength"] = p
	}
} 

// 是否可清除
func Textarea_clearable(p bool) opt {
	return func(o map[string]interface{}) {
		o["clearable"] = p
	}
} 

// 清除后设置此配置项给定的值。
func Textarea_resetValue(p string) opt {
	return func(o map[string]interface{}) {
		o["resetValue"] = p
	}
} 