package form

// 选择器表单项
func Options(opts ...opt) map[string]interface{} {
	return newForm("options", opts...)
}



// 选项组，供用户选择
func Options_options(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["options"] = p
	}
} 

// 选项组源，可通过数据映射获取当前数据域变量、或者配置 API 对象
func Options_source(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["source"] = p
	}
} 

// 是否支持多选
func Options_multiple(p bool) opt {
	return func(o map[string]interface{}) {
		o["multiple"] = p
	}
} 

// 标识选项中哪个字段是`label`值
func Options_labelField(p bool) opt {
	return func(o map[string]interface{}) {
		o["labelField"] = p
	}
} 

// 标识选项中哪个字段是`value`值
func Options_valueField(p bool) opt {
	return func(o map[string]interface{}) {
		o["valueField"] = p
	}
} 

// 是否拼接`value`值
func Options_joinValues(p bool) opt {
	return func(o map[string]interface{}) {
		o["joinValues"] = p
	}
} 

// 是否将`value`值抽取出来组成新的数组，只有在`joinValues`是`false`是生效
func Options_extractValue(p bool) opt {
	return func(o map[string]interface{}) {
		o["extractValue"] = p
	}
} 

// 每个选项的高度，用于虚拟渲染
func Options_itemHeight(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["itemHeight"] = p
	}
} 

// 在选项数量超过多少时开启虚拟渲染
func Options_virtualThreshold(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["virtualThreshold"] = p
	}
} 

// 默认情况下多选所有选项都会显示，通过这个可以最多显示一行，超出的部分变成 ...
func Options_valuesNoWrap(p bool) opt {
	return func(o map[string]interface{}) {
		o["valuesNoWrap"] = p
	}
} 