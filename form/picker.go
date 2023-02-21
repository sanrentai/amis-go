package form

// 列表选择器
func Picker(opts ...opt) map[string]interface{} {
	return newForm("picker", opts...)
}



// [选项组](./options#%E9%9D%99%E6%80%81%E9%80%89%E9%A1%B9%E7%BB%84-options)
func Picker_options(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["options"] = p
	}
} 

// [动态选项组](./options#%E5%8A%A8%E6%80%81%E9%80%89%E9%A1%B9%E7%BB%84-source)
func Picker_source(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["source"] = p
	}
} 

// 是否为多选。
func Picker_multiple(p bool) opt {
	return func(o map[string]interface{}) {
		o["multiple"] = p
	}
} 

// [拼接符](./options#%E6%8B%BC%E6%8E%A5%E7%AC%A6-delimiter)
func Picker_delimiter(p bool) opt {
	return func(o map[string]interface{}) {
		o["delimiter"] = p
	}
} 

// [选项标签字段](./options#%E9%80%89%E9%A1%B9%E6%A0%87%E7%AD%BE%E5%AD%97%E6%AE%B5-labelfield)
func Picker_labelField(p bool) opt {
	return func(o map[string]interface{}) {
		o["labelField"] = p
	}
} 

// [选项值字段](./options#%E9%80%89%E9%A1%B9%E5%80%BC%E5%AD%97%E6%AE%B5-valuefield)
func Picker_valueField(p bool) opt {
	return func(o map[string]interface{}) {
		o["valueField"] = p
	}
} 

// [拼接值](./options#%E6%8B%BC%E6%8E%A5%E5%80%BC-joinvalues)
func Picker_joinValues(p bool) opt {
	return func(o map[string]interface{}) {
		o["joinValues"] = p
	}
} 

// [提取值](./options#%E6%8F%90%E5%8F%96%E5%A4%9A%E9%80%89%E5%80%BC-extractvalue)
func Picker_extractValue(p bool) opt {
	return func(o map[string]interface{}) {
		o["extractValue"] = p
	}
} 

// [自动填充](./options#%E8%87%AA%E5%8A%A8%E5%A1%AB%E5%85%85-autofill)
func Picker_autoFill(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["autoFill"] = p
	}
} 

// 设置 `dialog` 或者 `drawer`，用来配置弹出方式。
func Picker_modalMode(p string) opt {
	return func(o map[string]interface{}) {
		o["modalMode"] = p
	}
} 

// 即用 List 类型的渲染，来展示列表信息。更多配置参考 [CRUD](../crud)
func Picker_pickerSchema(p string) opt {
	return func(o map[string]interface{}) {
		o["pickerSchema"] = p
	}
} 

// 是否使用内嵌模式
func Picker_embed(p bool) opt {
	return func(o map[string]interface{}) {
		o["embed"] = p
	}
} 

// 选中值变化时触发
func Picker_change(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["change"] = p
	}
} 