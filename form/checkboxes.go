package form

// 复选框
func Checkboxes(opts ...opt) map[string]interface{} {
	return newForm("checkboxes", opts...)
}



// [选项组](./#%E9%9D%99%E6%80%81%E9%80%89%E9%A1%B9%E7%BB%84-options)
func Checkboxes_options(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["options"] = p
	}
} 

// [动态选项组](./options#%E5%8A%A8%E6%80%81%E9%80%89%E9%A1%B9%E7%BB%84-source)
func Checkboxes_source(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["source"] = p
	}
} 

// [拼接符](./options#%E6%8B%BC%E6%8E%A5%E7%AC%A6-delimiter)
func Checkboxes_delimiter(p string) opt {
	return func(o map[string]interface{}) {
		o["delimiter"] = p
	}
} 

// [选项标签字段](./options#%E9%80%89%E9%A1%B9%E6%A0%87%E7%AD%BE%E5%AD%97%E6%AE%B5-labelfield)
func Checkboxes_labelField(p string) opt {
	return func(o map[string]interface{}) {
		o["labelField"] = p
	}
} 

// [选项值字段](./options#%E9%80%89%E9%A1%B9%E5%80%BC%E5%AD%97%E6%AE%B5-valuefield)
func Checkboxes_valueField(p string) opt {
	return func(o map[string]interface{}) {
		o["valueField"] = p
	}
} 

// [拼接值](./options#%E6%8B%BC%E6%8E%A5%E5%80%BC-joinvalues)
func Checkboxes_joinValues(p bool) opt {
	return func(o map[string]interface{}) {
		o["joinValues"] = p
	}
} 

// [提取值](./options#%E6%8F%90%E5%8F%96%E5%A4%9A%E9%80%89%E5%80%BC-extractvalue)
func Checkboxes_extractValue(p bool) opt {
	return func(o map[string]interface{}) {
		o["extractValue"] = p
	}
} 

// 选项按几列显示，默认为一列
func Checkboxes_columnsCount(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["columnsCount"] = p
	}
} 

// 支持自定义选项渲染
func Checkboxes_menuTpl(p string) opt {
	return func(o map[string]interface{}) {
		o["menuTpl"] = p
	}
} 

// 是否支持全选
func Checkboxes_checkAll(p bool) opt {
	return func(o map[string]interface{}) {
		o["checkAll"] = p
	}
} 

// 是否显示为一行
func Checkboxes_inline(p bool) opt {
	return func(o map[string]interface{}) {
		o["inline"] = p
	}
} 

// 默认是否全选
func Checkboxes_defaultCheckAll(p bool) opt {
	return func(o map[string]interface{}) {
		o["defaultCheckAll"] = p
	}
} 

// [新增选项](./options#%E5%89%8D%E7%AB%AF%E6%96%B0%E5%A2%9E-creatable)
func Checkboxes_creatable(p bool) opt {
	return func(o map[string]interface{}) {
		o["creatable"] = p
	}
} 

// [新增选项](./options#%E6%96%B0%E5%A2%9E%E9%80%89%E9%A1%B9)
func Checkboxes_createBtnLabel(p string) opt {
	return func(o map[string]interface{}) {
		o["createBtnLabel"] = p
	}
} 

// [自定义新增表单项](./options#%E8%87%AA%E5%AE%9A%E4%B9%89%E6%96%B0%E5%A2%9E%E8%A1%A8%E5%8D%95%E9%A1%B9-addcontrols)
func Checkboxes_addControls(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["addControls"] = p
	}
} 

// [配置新增选项接口](./options#%E9%85%8D%E7%BD%AE%E6%96%B0%E5%A2%9E%E6%8E%A5%E5%8F%A3-addapi)
func Checkboxes_addApi(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["addApi"] = p
	}
} 

// [编辑选项](./options#%E5%89%8D%E7%AB%AF%E7%BC%96%E8%BE%91-editable)
func Checkboxes_editable(p bool) opt {
	return func(o map[string]interface{}) {
		o["editable"] = p
	}
} 

// [自定义编辑表单项](./options#%E8%87%AA%E5%AE%9A%E4%B9%89%E7%BC%96%E8%BE%91%E8%A1%A8%E5%8D%95%E9%A1%B9-editcontrols)
func Checkboxes_editControls(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["editControls"] = p
	}
} 

// [配置编辑选项接口](./options#%E9%85%8D%E7%BD%AE%E7%BC%96%E8%BE%91%E6%8E%A5%E5%8F%A3-editapi)
func Checkboxes_editApi(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["editApi"] = p
	}
} 

// [删除选项](./options#%E5%88%A0%E9%99%A4%E9%80%89%E9%A1%B9)
func Checkboxes_removable(p bool) opt {
	return func(o map[string]interface{}) {
		o["removable"] = p
	}
} 

// [配置删除选项接口](./options#%E9%85%8D%E7%BD%AE%E5%88%A0%E9%99%A4%E6%8E%A5%E5%8F%A3-deleteapi)
func Checkboxes_deleteApi(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["deleteApi"] = p
	}
} 

// 选项样式类名
func Checkboxes_itemClassName(p string) opt {
	return func(o map[string]interface{}) {
		o["itemClassName"] = p
	}
} 

// 选项标签样式类名
func Checkboxes_labelClassName(p string) opt {
	return func(o map[string]interface{}) {
		o["labelClassName"] = p
	}
} 