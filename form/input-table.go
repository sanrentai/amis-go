package form

// 表格
func InputTable(opts ...opt) map[string]interface{} {
	return newForm("input-table", opts...)
}

// 是否可增加一行
func InputTable_addable(p bool) opt {
	return func(o map[string]interface{}) {
		o["addable"] = p
	}
}

// 是否可编辑
func InputTable_editable(p bool) opt {
	return func(o map[string]interface{}) {
		o["editable"] = p
	}
}

// 是否可删除
func InputTable_removable(p bool) opt {
	return func(o map[string]interface{}) {
		o["removable"] = p
	}
}

// 是否显示添加按钮
func InputTable_showAddBtn(p bool) opt {
	return func(o map[string]interface{}) {
		o["showAddBtn"] = p
	}
}

// 新增时提交的 API
func InputTable_addApi(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["addApi"] = p
	}
}

// 修改时提交的 API
func InputTable_updateApi(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["updateApi"] = p
	}
}

// 删除时提交的 API
func InputTable_deleteApi(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["deleteApi"] = p
	}
}

// 增加按钮名称
func InputTable_addBtnLabel(p string) opt {
	return func(o map[string]interface{}) {
		o["addBtnLabel"] = p
	}
}

// 增加按钮图标
func InputTable_addBtnIcon(p string) opt {
	return func(o map[string]interface{}) {
		o["addBtnIcon"] = p
	}
}

// 复制按钮文字
func InputTable_copyBtnLabel(p string) opt {
	return func(o map[string]interface{}) {
		o["copyBtnLabel"] = p
	}
}

// 复制按钮图标
func InputTable_copyBtnIcon(p string) opt {
	return func(o map[string]interface{}) {
		o["copyBtnIcon"] = p
	}
}

// 编辑按钮名称
func InputTable_editBtnLabel(p string) opt {
	return func(o map[string]interface{}) {
		o["editBtnLabel"] = p
	}
}

// 编辑按钮图标
func InputTable_editBtnIcon(p string) opt {
	return func(o map[string]interface{}) {
		o["editBtnIcon"] = p
	}
}

// 删除按钮名称
func InputTable_deleteBtnLabel(p string) opt {
	return func(o map[string]interface{}) {
		o["deleteBtnLabel"] = p
	}
}

// 删除按钮图标
func InputTable_deleteBtnIcon(p string) opt {
	return func(o map[string]interface{}) {
		o["deleteBtnIcon"] = p
	}
}

// 确认编辑按钮名称
func InputTable_confirmBtnLabel(p string) opt {
	return func(o map[string]interface{}) {
		o["confirmBtnLabel"] = p
	}
}

// 确认编辑按钮图标
func InputTable_confirmBtnIcon(p string) opt {
	return func(o map[string]interface{}) {
		o["confirmBtnIcon"] = p
	}
}

// 取消编辑按钮名称
func InputTable_cancelBtnLabel(p string) opt {
	return func(o map[string]interface{}) {
		o["cancelBtnLabel"] = p
	}
}

// 取消编辑按钮图标
func InputTable_cancelBtnIcon(p string) opt {
	return func(o map[string]interface{}) {
		o["cancelBtnIcon"] = p
	}
}

// 是否需要确认操作，，可用来控控制表格的操作交互
func InputTable_needConfirm(p bool) opt {
	return func(o map[string]interface{}) {
		o["needConfirm"] = p
	}
}

// 是否可以访问父级数据，也就是表单中的同级数据，通常需要跟 strictMode 搭配使用
func InputTable_canAccessSuperData(p bool) opt {
	return func(o map[string]interface{}) {
		o["canAccessSuperData"] = p
	}
}

// 为了性能，默认其他表单项项值变化不会让当前表格更新，有时候为了同步获取其他表单项字段，需要开启这个。
func InputTable_strictMode(p bool) opt {
	return func(o map[string]interface{}) {
		o["strictMode"] = p
	}
}

// 最小行数, `2.4.1`版本后支持变量
func InputTable_minLength(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["minLength"] = p
	}
}

// 最大行数, `2.4.1`版本后支持变量
func InputTable_maxLength(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["maxLength"] = p
	}
}

// 列信息
func InputTable_columns(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["columns"] = p
	}
}

// // 配合 editable 为 true 一起使用
// func InputTable_columns[x].quickEdit(p interface{}) opt {
// 	return func(o map[string]interface{}) {
// 		o["columns[x].quickEdit"] = p
// 	}
// }

// // 可以用来区分新建模式和更新模式的编辑配置
// func InputTable_columns[x].quickEditOnUpdate(p interface{}) opt {
// 	return func(o map[string]interface{}) {
// 		o["columns[x].quickEditOnUpdate"] = p
// 	}
// }
