package form

// 单选框
func Radios(opts ...opt) map[string]interface{} {
	return newForm("radios", opts...)
}



// [选项组](./options#%E9%9D%99%E6%80%81%E9%80%89%E9%A1%B9%E7%BB%84-options)
func Radios_options(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["options"] = p
	}
} 

// [动态选项组](./options#%E5%8A%A8%E6%80%81%E9%80%89%E9%A1%B9%E7%BB%84-source)
func Radios_source(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["source"] = p
	}
} 

// [选项标签字段](./options#%E9%80%89%E9%A1%B9%E6%A0%87%E7%AD%BE%E5%AD%97%E6%AE%B5-labelfield)
func Radios_labelField(p string) opt {
	return func(o map[string]interface{}) {
		o["labelField"] = p
	}
} 

// [选项值字段](./options#%E9%80%89%E9%A1%B9%E5%80%BC%E5%AD%97%E6%AE%B5-valuefield)
func Radios_valueField(p string) opt {
	return func(o map[string]interface{}) {
		o["valueField"] = p
	}
} 

// 选项按几列显示，默认为一列
func Radios_columnsCount(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["columnsCount"] = p
	}
} 

// 是否显示为一行
func Radios_inline(p bool) opt {
	return func(o map[string]interface{}) {
		o["inline"] = p
	}
} 

// 是否默认选中第一个
func Radios_selectFirst(p bool) opt {
	return func(o map[string]interface{}) {
		o["selectFirst"] = p
	}
} 

// [自动填充](./options#%E8%87%AA%E5%8A%A8%E5%A1%AB%E5%85%85-autofill)
func Radios_autoFill(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["autoFill"] = p
	}
} 