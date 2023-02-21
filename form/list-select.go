package form

// 列表
func ListSelect(opts ...opt) map[string]interface{} {
	return newForm("list-select", opts...)
}



// [选项组](./options#%E9%9D%99%E6%80%81%E9%80%89%E9%A1%B9%E7%BB%84-options)
func ListSelect_options(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["options"] = p
	}
} 

// [动态选项组](./options#%E5%8A%A8%E6%80%81%E9%80%89%E9%A1%B9%E7%BB%84-source)
func ListSelect_source(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["source"] = p
	}
} 

// [多选](./options#%E5%A4%9A%E9%80%89-multiple)
func ListSelect_multiple(p bool) opt {
	return func(o map[string]interface{}) {
		o["multiple"] = p
	}
} 

// [选项标签字段](./options#%E9%80%89%E9%A1%B9%E6%A0%87%E7%AD%BE%E5%AD%97%E6%AE%B5-labelfield)
func ListSelect_labelField(p string) opt {
	return func(o map[string]interface{}) {
		o["labelField"] = p
	}
} 

// [选项值字段](./options#%E9%80%89%E9%A1%B9%E5%80%BC%E5%AD%97%E6%AE%B5-valuefield)
func ListSelect_valueField(p string) opt {
	return func(o map[string]interface{}) {
		o["valueField"] = p
	}
} 

// [拼接值](./options#%E6%8B%BC%E6%8E%A5%E5%80%BC-joinvalues)
func ListSelect_joinValues(p bool) opt {
	return func(o map[string]interface{}) {
		o["joinValues"] = p
	}
} 

// [提取值](./options#%E6%8F%90%E5%8F%96%E5%A4%9A%E9%80%89%E5%80%BC-extractvalue)
func ListSelect_extractValue(p bool) opt {
	return func(o map[string]interface{}) {
		o["extractValue"] = p
	}
} 

// [自动填充](./options#%E8%87%AA%E5%8A%A8%E5%A1%AB%E5%85%85-autofill)
func ListSelect_autoFill(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["autoFill"] = p
	}
} 

// 支持配置 list div 的 css 类名。比如: `flex justify-between`
func ListSelect_listClassName(p string) opt {
	return func(o map[string]interface{}) {
		o["listClassName"] = p
	}
} 

// 更新数据，开启`multiple`支持设置多项，开启`joinValues`时，多值用`,`分隔，否则多值用数组
func ListSelect_setValue(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["setValue"] = p
	}
} 