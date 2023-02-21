package form

// 树形选择框
func InputTree(opts ...opt) map[string]interface{} {
	return newForm("input-tree", opts...)
}



// [选项组](./options#%E9%9D%99%E6%80%81%E9%80%89%E9%A1%B9%E7%BB%84-options)
func InputTree_options(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["options"] = p
	}
} 

// [动态选项组](./options#%E5%8A%A8%E6%80%81%E9%80%89%E9%A1%B9%E7%BB%84-source)
func InputTree_source(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["source"] = p
	}
} 

// [自动提示补全](./options#%E8%87%AA%E5%8A%A8%E8%A1%A5%E5%85%A8-autocomplete)
func InputTree_autoComplete(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["autoComplete"] = p
	}
} 

// 是否多选
func InputTree_multiple(p bool) opt {
	return func(o map[string]interface{}) {
		o["multiple"] = p
	}
} 

// [拼接符](./options#%E6%8B%BC%E6%8E%A5%E7%AC%A6-delimiter)
func InputTree_delimiter(p string) opt {
	return func(o map[string]interface{}) {
		o["delimiter"] = p
	}
} 

// [选项标签字段](./options#%E9%80%89%E9%A1%B9%E6%A0%87%E7%AD%BE%E5%AD%97%E6%AE%B5-labelfield)
func InputTree_labelField(p string) opt {
	return func(o map[string]interface{}) {
		o["labelField"] = p
	}
} 

// [选项值字段](./options#%E9%80%89%E9%A1%B9%E5%80%BC%E5%AD%97%E6%AE%B5-valuefield)
func InputTree_valueField(p string) opt {
	return func(o map[string]interface{}) {
		o["valueField"] = p
	}
} 

// 图标值字段
func InputTree_iconField(p string) opt {
	return func(o map[string]interface{}) {
		o["iconField"] = p
	}
} 

// [拼接值](./options#%E6%8B%BC%E6%8E%A5%E5%80%BC-joinvalues)
func InputTree_joinValues(p bool) opt {
	return func(o map[string]interface{}) {
		o["joinValues"] = p
	}
} 

// [提取值](./options#%E6%8F%90%E5%8F%96%E5%A4%9A%E9%80%89%E5%80%BC-extractvalue)
func InputTree_extractValue(p bool) opt {
	return func(o map[string]interface{}) {
		o["extractValue"] = p
	}
} 

// [新增选项](./options#%E5%89%8D%E7%AB%AF%E6%96%B0%E5%A2%9E-creatable)
func InputTree_creatable(p bool) opt {
	return func(o map[string]interface{}) {
		o["creatable"] = p
	}
} 

// [自定义新增表单项](./options#%E8%87%AA%E5%AE%9A%E4%B9%89%E6%96%B0%E5%A2%9E%E8%A1%A8%E5%8D%95%E9%A1%B9-addcontrols)
func InputTree_addControls(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["addControls"] = p
	}
} 

// [配置新增选项接口](./options#%E9%85%8D%E7%BD%AE%E6%96%B0%E5%A2%9E%E6%8E%A5%E5%8F%A3-addapi)
func InputTree_addApi(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["addApi"] = p
	}
} 

// [编辑选项](./options#%E5%89%8D%E7%AB%AF%E7%BC%96%E8%BE%91-editable)
func InputTree_editable(p bool) opt {
	return func(o map[string]interface{}) {
		o["editable"] = p
	}
} 

// [自定义编辑表单项](./options#%E8%87%AA%E5%AE%9A%E4%B9%89%E7%BC%96%E8%BE%91%E8%A1%A8%E5%8D%95%E9%A1%B9-editcontrols)
func InputTree_editControls(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["editControls"] = p
	}
} 

// [配置编辑选项接口](./options#%E9%85%8D%E7%BD%AE%E7%BC%96%E8%BE%91%E6%8E%A5%E5%8F%A3-editapi)
func InputTree_editApi(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["editApi"] = p
	}
} 

// [删除选项](./options#%E5%88%A0%E9%99%A4%E9%80%89%E9%A1%B9)
func InputTree_removable(p bool) opt {
	return func(o map[string]interface{}) {
		o["removable"] = p
	}
} 

// [配置删除选项接口](./options#%E9%85%8D%E7%BD%AE%E5%88%A0%E9%99%A4%E6%8E%A5%E5%8F%A3-deleteapi)
func InputTree_deleteApi(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["deleteApi"] = p
	}
} 

// 是否可检索，仅在 type 为 `tree-select` 的时候生效
func InputTree_searchable(p bool) opt {
	return func(o map[string]interface{}) {
		o["searchable"] = p
	}
} 

// 如果想要显示个顶级节点，请设置为 `false`
func InputTree_hideRoot(p bool) opt {
	return func(o map[string]interface{}) {
		o["hideRoot"] = p
	}
} 

// 当 `hideRoot` 不为 `false` 时有用，用来设置顶级节点的文字。
func InputTree_rootLabel(p bool) opt {
	return func(o map[string]interface{}) {
		o["rootLabel"] = p
	}
} 

// 是否显示图标
func InputTree_showIcon(p bool) opt {
	return func(o map[string]interface{}) {
		o["showIcon"] = p
	}
} 

// 是否显示单选按钮，`multiple` 为 `false` 是有效。
func InputTree_showRadio(p bool) opt {
	return func(o map[string]interface{}) {
		o["showRadio"] = p
	}
} 

// 是否显示树层级展开线
func InputTree_showOutline(p bool) opt {
	return func(o map[string]interface{}) {
		o["showOutline"] = p
	}
} 

// 设置是否默认展开所有层级。
func InputTree_initiallyOpen(p bool) opt {
	return func(o map[string]interface{}) {
		o["initiallyOpen"] = p
	}
} 

// 设置默认展开的级数，只有`initiallyOpen`不是`true`时生效。
func InputTree_unfoldedLevel(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["unfoldedLevel"] = p
	}
} 

// 当选中父节点时级联选择子节点。
func InputTree_autoCheckChildren(p bool) opt {
	return func(o map[string]interface{}) {
		o["autoCheckChildren"] = p
	}
} 

// autoCheckChildren 为 true 时生效；默认行为：子节点禁用，值只包含父节点值；设置为 true 时，子节点可反选，值包含父子节点值。
func InputTree_cascade(p bool) opt {
	return func(o map[string]interface{}) {
		o["cascade"] = p
	}
} 

// cascade 为 false 时生效，选中父节点时，值里面将包含父子节点的值，否则只会保留父节点的值。
func InputTree_withChildren(p bool) opt {
	return func(o map[string]interface{}) {
		o["withChildren"] = p
	}
} 

// autoCheckChildren 为 true 时生效，不受 cascade 影响；onlyChildren 为 true，ui 行为级联选中子节点，子节点可反选，值只包含子节点的值。
func InputTree_onlyChildren(p bool) opt {
	return func(o map[string]interface{}) {
		o["onlyChildren"] = p
	}
} 

// 只允许选择叶子节点
func InputTree_onlyLeaf(p bool) opt {
	return func(o map[string]interface{}) {
		o["onlyLeaf"] = p
	}
} 

// 是否可以创建顶级节点
func InputTree_rootCreatable(p bool) opt {
	return func(o map[string]interface{}) {
		o["rootCreatable"] = p
	}
} 

// 创建顶级节点的悬浮提示
func InputTree_rootCreateTip(p string) opt {
	return func(o map[string]interface{}) {
		o["rootCreateTip"] = p
	}
} 

// 最少选中的节点数
func InputTree_minLength(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["minLength"] = p
	}
} 

// 最多选中的节点数
func InputTree_maxLength(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["maxLength"] = p
	}
} 

// tree 最外层容器类名
func InputTree_treeContainerClassName(p string) opt {
	return func(o map[string]interface{}) {
		o["treeContainerClassName"] = p
	}
} 

// 是否开启节点路径模式
func InputTree_enableNodePath(p bool) opt {
	return func(o map[string]interface{}) {
		o["enableNodePath"] = p
	}
} 

// 节点路径的分隔符，`enableNodePath`为`true`时生效
func InputTree_pathSeparator(p string) opt {
	return func(o map[string]interface{}) {
		o["pathSeparator"] = p
	}
} 

// 标签中需要高亮的字符，支持变量
func InputTree_highlightTxt(p string) opt {
	return func(o map[string]interface{}) {
		o["highlightTxt"] = p
	}
} 

// 每个选项的高度，用于虚拟渲染
func InputTree_itemHeight(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["itemHeight"] = p
	}
} 

// 在选项数量超过多少时开启虚拟渲染
func InputTree_virtualThreshold(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["virtualThreshold"] = p
	}
} 

// 更新数据，开启`multiple`支持设置多项，开启`joinValues`时，多值用`,`分隔，否则多值用数组
func InputTree_setValue(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["setValue"] = p
	}
} 