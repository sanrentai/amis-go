package form

// 树形选择器
func TreeSelect(opts ...opt) map[string]interface{} {
	return newForm("treeselect", opts...)
}



// 是否隐藏选择框中已选择节点的路径 label 信息
func TreeSelect_hideNodePathLabel(p bool) opt {
	return func(o map[string]interface{}) {
		o["hideNodePathLabel"] = p
	}
} 

// 只允许选择叶子节点
func TreeSelect_onlyLeaf(p bool) opt {
	return func(o map[string]interface{}) {
		o["onlyLeaf"] = p
	}
} 

// 更新数据，开启`multiple`支持设置多项，开启`joinValues`时，多值用`,`分隔，否则多值用数组
func TreeSelect_setValue(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["setValue"] = p
	}
} 