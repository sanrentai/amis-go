package form

// 选择器
func Select(opts ...opt) map[string]interface{} {
	return newForm("select", opts...)
}



// [选项组](./options#%E9%9D%99%E6%80%81%E9%80%89%E9%A1%B9%E7%BB%84-options)
func Select_options(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["options"] = p
	}
} 

// [动态选项组](./options#%E5%8A%A8%E6%80%81%E9%80%89%E9%A1%B9%E7%BB%84-source)
func Select_source(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["source"] = p
	}
} 

// [自动提示补全](./options#%E8%87%AA%E5%8A%A8%E8%A1%A5%E5%85%A8-autocomplete)
func Select_autoComplete(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["autoComplete"] = p
	}
} 

// [拼接符](./options#%E6%8B%BC%E6%8E%A5%E7%AC%A6-delimiter)
func Select_delimiter(p string) opt {
	return func(o map[string]interface{}) {
		o["delimiter"] = p
	}
} 

// [选项标签字段](./options#%E9%80%89%E9%A1%B9%E6%A0%87%E7%AD%BE%E5%AD%97%E6%AE%B5-labelfield)
func Select_labelField(p string) opt {
	return func(o map[string]interface{}) {
		o["labelField"] = p
	}
} 

// [选项值字段](./options#%E9%80%89%E9%A1%B9%E5%80%BC%E5%AD%97%E6%AE%B5-valuefield)
func Select_valueField(p string) opt {
	return func(o map[string]interface{}) {
		o["valueField"] = p
	}
} 

// [拼接值](./options#%E6%8B%BC%E6%8E%A5%E5%80%BC-joinvalues)
func Select_joinValues(p bool) opt {
	return func(o map[string]interface{}) {
		o["joinValues"] = p
	}
} 

// [提取值](./options#%E6%8F%90%E5%8F%96%E5%A4%9A%E9%80%89%E5%80%BC-extractvalue)
func Select_extractValue(p bool) opt {
	return func(o map[string]interface{}) {
		o["extractValue"] = p
	}
} 

// 是否支持全选
func Select_checkAll(p bool) opt {
	return func(o map[string]interface{}) {
		o["checkAll"] = p
	}
} 

// 全选的文字
func Select_checkAllLabel(p string) opt {
	return func(o map[string]interface{}) {
		o["checkAllLabel"] = p
	}
} 

// 有检索时只全选检索命中的项
func Select_checkAllBySearch(p bool) opt {
	return func(o map[string]interface{}) {
		o["checkAllBySearch"] = p
	}
} 

// 默认是否全选
func Select_defaultCheckAll(p bool) opt {
	return func(o map[string]interface{}) {
		o["defaultCheckAll"] = p
	}
} 

// [新增选项](./options#%E5%89%8D%E7%AB%AF%E6%96%B0%E5%A2%9E-creatable)
func Select_creatable(p bool) opt {
	return func(o map[string]interface{}) {
		o["creatable"] = p
	}
} 

// [多选](./options#多选-multiple)
func Select_multiple(p bool) opt {
	return func(o map[string]interface{}) {
		o["multiple"] = p
	}
} 

// [检索](./options#检索-searchable)
func Select_searchable(p bool) opt {
	return func(o map[string]interface{}) {
		o["searchable"] = p
	}
} 

// [新增选项](./options#%E6%96%B0%E5%A2%9E%E9%80%89%E9%A1%B9)
func Select_createBtnLabel(p string) opt {
	return func(o map[string]interface{}) {
		o["createBtnLabel"] = p
	}
} 

// [自定义新增表单项](./options#%E8%87%AA%E5%AE%9A%E4%B9%89%E6%96%B0%E5%A2%9E%E8%A1%A8%E5%8D%95%E9%A1%B9-addcontrols)
func Select_addControls(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["addControls"] = p
	}
} 

// [配置新增选项接口](./options#%E9%85%8D%E7%BD%AE%E6%96%B0%E5%A2%9E%E6%8E%A5%E5%8F%A3-addapi)
func Select_addApi(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["addApi"] = p
	}
} 

// [编辑选项](./options#%E5%89%8D%E7%AB%AF%E7%BC%96%E8%BE%91-editable)
func Select_editable(p bool) opt {
	return func(o map[string]interface{}) {
		o["editable"] = p
	}
} 

// [自定义编辑表单项](./options#%E8%87%AA%E5%AE%9A%E4%B9%89%E7%BC%96%E8%BE%91%E8%A1%A8%E5%8D%95%E9%A1%B9-editcontrols)
func Select_editControls(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["editControls"] = p
	}
} 

// [配置编辑选项接口](./options#%E9%85%8D%E7%BD%AE%E7%BC%96%E8%BE%91%E6%8E%A5%E5%8F%A3-editapi)
func Select_editApi(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["editApi"] = p
	}
} 

// [删除选项](./options#%E5%88%A0%E9%99%A4%E9%80%89%E9%A1%B9)
func Select_removable(p bool) opt {
	return func(o map[string]interface{}) {
		o["removable"] = p
	}
} 

// [配置删除选项接口](./options#%E9%85%8D%E7%BD%AE%E5%88%A0%E9%99%A4%E6%8E%A5%E5%8F%A3-deleteapi)
func Select_deleteApi(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["deleteApi"] = p
	}
} 

// [自动填充](./options#%E8%87%AA%E5%8A%A8%E5%A1%AB%E5%85%85-autofill)
func Select_autoFill(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["autoFill"] = p
	}
} 

// 支持配置自定义菜单
func Select_menuTpl(p string) opt {
	return func(o map[string]interface{}) {
		o["menuTpl"] = p
	}
} 

// 单选模式下是否支持清空
func Select_clearable(p bool) opt {
	return func(o map[string]interface{}) {
		o["clearable"] = p
	}
} 

// 隐藏已选选项
func Select_hideSelected(p bool) opt {
	return func(o map[string]interface{}) {
		o["hideSelected"] = p
	}
} 

// 移动端浮层类名
func Select_mobileClassName(p string) opt {
	return func(o map[string]interface{}) {
		o["mobileClassName"] = p
	}
} 

// 可选：`group`、`table`、`tree`、`chained`、`associated`。分别为：列表形式、表格形式、树形选择形式、级联选择形式，关联选择形式（与级联选择的区别在于，级联是无限极，而关联只有一级，关联左边可以是个 tree）。
func Select_selectMode(p string) opt {
	return func(o map[string]interface{}) {
		o["selectMode"] = p
	}
} 

// 如果不设置将采用 `selectMode` 的值，可以单独配置，参考 `selectMode`，决定搜索结果的展示形式。
func Select_searchResultMode(p string) opt {
	return func(o map[string]interface{}) {
		o["searchResultMode"] = p
	}
} 

// 当展示形式为 `table` 可以用来配置展示哪些列，跟 table 中的 columns 配置相似，只是只有展示功能。
func Select_columns(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["columns"] = p
	}
} 

// 当展示形式为 `associated` 时用来配置左边的选项集。
func Select_leftOptions(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["leftOptions"] = p
	}
} 

// 当展示形式为 `associated` 时用来配置左边的选择形式，支持 `list` 或者 `tree`。默认为 `list`。
func Select_leftMode(p string) opt {
	return func(o map[string]interface{}) {
		o["leftMode"] = p
	}
} 

// 当展示形式为 `associated` 时用来配置右边的选择形式，可选：`list`、`table`、`tree`、`chained`。
func Select_rightMode(p string) opt {
	return func(o map[string]interface{}) {
		o["rightMode"] = p
	}
} 

// 标签的最大展示数量，超出数量后以收纳浮层的方式展示，仅在多选模式开启后生效
func Select_maxTagCount(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["maxTagCount"] = p
	}
} 

// 收纳浮层的配置属性，详细配置参考[Tooltip](../tooltip#属性表)
func Select_overflowTagPopover(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["overflowTagPopover"] = p
	}
} 

// 选项 CSS 类名
func Select_optionClassName(p string) opt {
	return func(o map[string]interface{}) {
		o["optionClassName"] = p
	}
} 

// 弹层挂载位置选择器，会通过`querySelector`获取
func Select_popOverContainerSelector(p string) opt {
	return func(o map[string]interface{}) {
		o["popOverContainerSelector"] = p
	}
} 

// 选中值变化时触发
func Select_change(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["change"] = p
	}
} 

// 更新数据，开启`multiple`支持设置多项，开启`joinValues`时，多值用`,`分隔，否则多值用数组
func Select_setValue(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["setValue"] = p
	}
} 