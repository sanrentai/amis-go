package form

// 勾选框
func Checkbox(opts ...opt) map[string]interface{} {
	return newForm("checkbox", opts...)
}



// 选项说明
func Checkbox_option(p string) opt {
	return func(o map[string]interface{}) {
		o["option"] = p
	}
} 

// 标识真值
func Checkbox_trueValue(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["trueValue"] = p
	}
} 

// 标识假值
func Checkbox_falseValue(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["falseValue"] = p
	}
} 

// 设置 option 类型
func Checkbox_optionType(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["optionType"] = p
	}
} 