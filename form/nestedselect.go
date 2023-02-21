package form

// 级联选择器
func NestedSelect(opts ...opt) map[string]interface{} {
	return newForm("nestedselect", opts...)
}



// [选项组](./options#%E9%9D%99%E6%80%81%E9%80%89%E9%A1%B9%E7%BB%84-options)
func NestedSelect_options(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["options"] = p
	}
} 

// [动态选项组](./options#%E5%8A%A8%E6%80%81%E9%80%89%E9%A1%B9%E7%BB%84-source)
func NestedSelect_source(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["source"] = p
	}
} 

// [拼接符](./options#%E6%8B%BC%E6%8E%A5%E7%AC%A6-delimiter)
func NestedSelect_delimiter(p bool) opt {
	return func(o map[string]interface{}) {
		o["delimiter"] = p
	}
} 

// [选项标签字段](./options#%E9%80%89%E9%A1%B9%E6%A0%87%E7%AD%BE%E5%AD%97%E6%AE%B5-labelfield)
func NestedSelect_labelField(p bool) opt {
	return func(o map[string]interface{}) {
		o["labelField"] = p
	}
} 

// [选项值字段](./options#%E9%80%89%E9%A1%B9%E5%80%BC%E5%AD%97%E6%AE%B5-valuefield)
func NestedSelect_valueField(p bool) opt {
	return func(o map[string]interface{}) {
		o["valueField"] = p
	}
} 

// [拼接值](./options#%E6%8B%BC%E6%8E%A5%E5%80%BC-joinvalues)
func NestedSelect_joinValues(p bool) opt {
	return func(o map[string]interface{}) {
		o["joinValues"] = p
	}
} 

// [提取值](./options#%E6%8F%90%E5%8F%96%E5%A4%9A%E9%80%89%E5%80%BC-extractvalue)
func NestedSelect_extractValue(p bool) opt {
	return func(o map[string]interface{}) {
		o["extractValue"] = p
	}
} 

// [自动填充](./options#%E8%87%AA%E5%8A%A8%E5%A1%AB%E5%85%85-autofill)
func NestedSelect_autoFill(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["autoFill"] = p
	}
} 

// 设置 `true`时，当选中父节点时不自动选择子节点。
func NestedSelect_cascade(p bool) opt {
	return func(o map[string]interface{}) {
		o["cascade"] = p
	}
} 

// 设置 `true`时，选中父节点时，值里面将包含子节点的值，否则只会保留父节点的值。
func NestedSelect_withChildren(p bool) opt {
	return func(o map[string]interface{}) {
		o["withChildren"] = p
	}
} 

// 多选时，选中父节点时，是否只将其子节点加入到值中。
func NestedSelect_onlyChildren(p bool) opt {
	return func(o map[string]interface{}) {
		o["onlyChildren"] = p
	}
} 

// 可否搜索
func NestedSelect_searchable(p bool) opt {
	return func(o map[string]interface{}) {
		o["searchable"] = p
	}
} 

// 搜索框占位文本
func NestedSelect_searchPromptText(p string) opt {
	return func(o map[string]interface{}) {
		o["searchPromptText"] = p
	}
} 

// 无结果时的文本
func NestedSelect_noResultsText(p string) opt {
	return func(o map[string]interface{}) {
		o["noResultsText"] = p
	}
} 

// 可否多选
func NestedSelect_multiple(p bool) opt {
	return func(o map[string]interface{}) {
		o["multiple"] = p
	}
} 

// 是否隐藏选择框中已选择节点的路径 label 信息
func NestedSelect_hideNodePathLabel(p bool) opt {
	return func(o map[string]interface{}) {
		o["hideNodePathLabel"] = p
	}
} 

// 只允许选择叶子节点
func NestedSelect_onlyLeaf(p bool) opt {
	return func(o map[string]interface{}) {
		o["onlyLeaf"] = p
	}
} 