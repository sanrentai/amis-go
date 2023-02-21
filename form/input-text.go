package form

// 输入框
func InputText(opts ...opt) map[string]interface{} {
	return newForm("input-text", opts...)
}

// [选项组](./options#%E9%9D%99%E6%80%81%E9%80%89%E9%A1%B9%E7%BB%84-options)
func InputText_options(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["options"] = p
	}
}

// [动态选项组](./options#%E5%8A%A8%E6%80%81%E9%80%89%E9%A1%B9%E7%BB%84-source)
func InputText_source(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["source"] = p
	}
}

// [自动补全](./options#%E8%87%AA%E5%8A%A8%E8%A1%A5%E5%85%A8-autocomplete)
func InputText_autoComplete(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["autoComplete"] = p
	}
}

// [是否多选](./options#%E5%A4%9A%E9%80%89-multiple)
func InputText_multiple(p bool) opt {
	return func(o map[string]interface{}) {
		o["multiple"] = p
	}
}

// [拼接符](./options#%E6%8B%BC%E6%8E%A5%E7%AC%A6-delimiter)
func InputText_delimiter(p string) opt {
	return func(o map[string]interface{}) {
		o["delimiter"] = p
	}
}

// [选项标签字段](./options#%E9%80%89%E9%A1%B9%E6%A0%87%E7%AD%BE%E5%AD%97%E6%AE%B5-labelfield)
func InputText_labelField(p string) opt {
	return func(o map[string]interface{}) {
		o["labelField"] = p
	}
}

// [选项值字段](./options#%E9%80%89%E9%A1%B9%E5%80%BC%E5%AD%97%E6%AE%B5-valuefield)
func InputText_valueField(p string) opt {
	return func(o map[string]interface{}) {
		o["valueField"] = p
	}
}

// [拼接值](./options#%E6%8B%BC%E6%8E%A5%E5%80%BC-joinvalues)
func InputText_joinValues(p bool) opt {
	return func(o map[string]interface{}) {
		o["joinValues"] = p
	}
}

// [提取值](./options#%E6%8F%90%E5%8F%96%E5%A4%9A%E9%80%89%E5%80%BC-extractvalue)
func InputText_extractValue(p bool) opt {
	return func(o map[string]interface{}) {
		o["extractValue"] = p
	}
}

// 输入框附加组件，比如附带一个提示文字，或者附带一个提交按钮。
func InputText_addOn(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["addOn"] = p
	}
}

// 请选择 `text` 、`button` 或者 `submit`。
// func InputText_addOn.type(p string) opt {
// 	return func(o map[string]interface{}) {
// 		o["addOn.type"] = p
// 	}
// }

// 文字说明
// func InputText_addOn.label(p string) opt {
// 	return func(o map[string]interface{}) {
// 		o["addOn.label"] = p
// 	}
// }

// 其他参数请参考按钮文档
// func InputText_addOn.xxx(p string) opt {
// 	return func(o map[string]interface{}) {
// 		o["addOn.xxx"] = p
// 	}
// }

// 是否去除首尾空白文本。
func InputText_trimContents(p bool) opt {
	return func(o map[string]interface{}) {
		o["trimContents"] = p
	}
}

// 文本内容为空时去掉这个值
func InputText_clearValueOnEmpty(p bool) opt {
	return func(o map[string]interface{}) {
		o["clearValueOnEmpty"] = p
	}
}

// 是否可以创建，默认为可以，除非设置为 false 即只能选择选项中的值
func InputText_creatable(p bool) opt {
	return func(o map[string]interface{}) {
		o["creatable"] = p
	}
}

// 是否可清除
func InputText_clearable(p bool) opt {
	return func(o map[string]interface{}) {
		o["clearable"] = p
	}
}

// 清除后设置此配置项给定的值。
func InputText_resetValue(p string) opt {
	return func(o map[string]interface{}) {
		o["resetValue"] = p
	}
}

// 前缀
func InputText_prefix(p string) opt {
	return func(o map[string]interface{}) {
		o["prefix"] = p
	}
}

// 后缀
func InputText_suffix(p string) opt {
	return func(o map[string]interface{}) {
		o["suffix"] = p
	}
}

// 是否显示计数器
func InputText_showCounter(p bool) opt {
	return func(o map[string]interface{}) {
		o["showCounter"] = p
	}
}

// 限制最小字数
func InputText_minLength(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["minLength"] = p
	}
}

// 限制最大字数
func InputText_maxLength(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["maxLength"] = p
	}
}

// 自动转换值，可选 `transform: { lowerCase: true, upperCase: true }`
func InputText_transform(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["transform"] = p
	}
}

// control 节点的 CSS 类名
func InputText_inputControlClassName(p string) opt {
	return func(o map[string]interface{}) {
		o["inputControlClassName"] = p
	}
}

// 原生 input 标签的 CSS 类名
func InputText_nativeInputClassName(p string) opt {
	return func(o map[string]interface{}) {
		o["nativeInputClassName"] = p
	}
}
