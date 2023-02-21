package form

// 开关
func Switch(opts ...opt) map[string]interface{} {
	return newForm("switch", opts...)
}



// 选项说明
func Switch_option(p string) opt {
	return func(o map[string]interface{}) {
		o["option"] = p
	}
} 

// 开启时开关显示的内容
func Switch_onText(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["onText"] = p
	}
} 

// 关闭时开关显示的内容
func Switch_offText(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["offText"] = p
	}
} 

// 标识真值
func Switch_trueValue(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["trueValue"] = p
	}
} 

// 标识假值
func Switch_falseValue(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["falseValue"] = p
	}
} 

// 说明
func Switch_属性名(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["属性名"] = p
	}
} 

// icon 的类型
func Switch_icon(p string) opt {
	return func(o map[string]interface{}) {
		o["icon"] = p
	}
} 

// 开关值变化时触发
func Switch_change(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["change"] = p
	}
} 

// 更新数据
func Switch_setValue(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["setValue"] = p
	}
} 