package form

// 标签选择器
func InputTag(opts ...opt) map[string]interface{} {
	return newForm("input-tag", opts...)
}



// [选项组](./options#%E9%9D%99%E6%80%81%E9%80%89%E9%A1%B9%E7%BB%84-options)
func InputTag_options(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["options"] = p
	}
} 

// 选项提示
func InputTag_optionsTip(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["optionsTip"] = p
	}
} 

// [动态选项组](./options#%E5%8A%A8%E6%80%81%E9%80%89%E9%A1%B9%E7%BB%84-source)
func InputTag_source(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["source"] = p
	}
} 

// [拼接符](./options#%E6%8B%BC%E6%8E%A5%E7%AC%A6-delimiter)
func InputTag_delimiter(p string) opt {
	return func(o map[string]interface{}) {
		o["delimiter"] = p
	}
} 

// [选项标签字段](./options#%E9%80%89%E9%A1%B9%E6%A0%87%E7%AD%BE%E5%AD%97%E6%AE%B5-labelfield)
func InputTag_labelField(p string) opt {
	return func(o map[string]interface{}) {
		o["labelField"] = p
	}
} 

// [选项值字段](./options#%E9%80%89%E9%A1%B9%E5%80%BC%E5%AD%97%E6%AE%B5-valuefield)
func InputTag_valueField(p string) opt {
	return func(o map[string]interface{}) {
		o["valueField"] = p
	}
} 

// [拼接值](./options#%E6%8B%BC%E6%8E%A5%E5%80%BC-joinvalues)
func InputTag_joinValues(p bool) opt {
	return func(o map[string]interface{}) {
		o["joinValues"] = p
	}
} 

// [提取值](./options#%E6%8F%90%E5%8F%96%E5%A4%9A%E9%80%89%E5%80%BC-extractvalue)
func InputTag_extractValue(p bool) opt {
	return func(o map[string]interface{}) {
		o["extractValue"] = p
	}
} 

// 在有值的时候是否显示一个删除图标在右侧。
func InputTag_clearable(p bool) opt {
	return func(o map[string]interface{}) {
		o["clearable"] = p
	}
} 

// 删除后设置此配置项给定的值。
func InputTag_resetValue(p string) opt {
	return func(o map[string]interface{}) {
		o["resetValue"] = p
	}
} 

// 允许添加的标签的最大数量
func InputTag_max(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["max"] = p
	}
} 

// 单个标签的最大文本长度
func InputTag_maxTagLength(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["maxTagLength"] = p
	}
} 

// 标签的最大展示数量，超出数量后以收纳浮层的方式展示，仅在多选模式开启后生效
func InputTag_maxTagCount(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["maxTagCount"] = p
	}
} 

// 收纳浮层的配置属性，详细配置参考[Tooltip](../tooltip#属性表)
func InputTag_overflowTagPopover(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["overflowTagPopover"] = p
	}
} 

// 是否开启批量添加模式
func InputTag_enableBatchAdd(p bool) opt {
	return func(o map[string]interface{}) {
		o["enableBatchAdd"] = p
	}
} 

// 开启批量添加后，输入多个标签的分隔符，支持传入多个符号，默认为"-"
func InputTag_separator(p string) opt {
	return func(o map[string]interface{}) {
		o["separator"] = p
	}
} 