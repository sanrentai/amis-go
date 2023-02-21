package form

// 组合
func Combo(opts ...opt) map[string]interface{} {
	return newForm("combo", opts...)
}

// 单组表单项的类名
func Combo_formClassName(p string) opt {
	return func(o map[string]interface{}) {
		o["formClassName"] = p
	}
}

// 组合展示的表单项
func Combo_items(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["items"] = p
	}
}

// // 列的类名，可以用它配置列宽度。默认平均分配。
// func Combo_items[x].columnClassName(p string) opt {
// 	return func(o map[string]interface{}) {
// 		o["items[x].columnClassName"] = p
// 	}
// }

// // 设置当前列值是否唯一，即不允许重复选择。
// func Combo_items[x].unique(p bool) opt {
// 	return func(o map[string]interface{}) {
// 		o["items[x].unique"] = p
// 	}
// }

// 单组表单项是否显示边框
func Combo_noBorder(p bool) opt {
	return func(o map[string]interface{}) {
		o["noBorder"] = p
	}
}

// 单组表单项初始值
func Combo_scaffold(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["scaffold"] = p
	}
}

// 是否多选
func Combo_multiple(p bool) opt {
	return func(o map[string]interface{}) {
		o["multiple"] = p
	}
}

// 默认是横着展示一排，设置以后竖着展示
func Combo_multiLine(p bool) opt {
	return func(o map[string]interface{}) {
		o["multiLine"] = p
	}
}

// 最少添加的条数，`2.4.1` 版本后支持变量
func Combo_minLength(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["minLength"] = p
	}
}

// 最多添加的条数，`2.4.1` 版本后支持变量
func Combo_maxLength(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["maxLength"] = p
	}
}

// 是否将结果扁平化(去掉 name),只有当 items 的 length 为 1 且 multiple 为 true 的时候才有效。
func Combo_flat(p bool) opt {
	return func(o map[string]interface{}) {
		o["flat"] = p
	}
}

// 默认为 `true` 当扁平化开启的时候，是否用分隔符的形式发送给后端，否则采用 array 的方式。
func Combo_joinValues(p bool) opt {
	return func(o map[string]interface{}) {
		o["joinValues"] = p
	}
}

// 当扁平化开启并且 joinValues 为 true 时，用什么分隔符。
func Combo_delimiter(p string) opt {
	return func(o map[string]interface{}) {
		o["delimiter"] = p
	}
}

// 是否可新增
func Combo_addable(p bool) opt {
	return func(o map[string]interface{}) {
		o["addable"] = p
	}
}

// 在顶部添加
func Combo_addattop(p bool) opt {
	return func(o map[string]interface{}) {
		o["addattop"] = p
	}
}

// 是否可删除
func Combo_removable(p bool) opt {
	return func(o map[string]interface{}) {
		o["removable"] = p
	}
}

// 如果配置了，则删除前会发送一个 api，请求成功才完成删除
func Combo_deleteApi(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["deleteApi"] = p
	}
}

// 当配置 `deleteApi` 才生效！删除时用来做用户确认
func Combo_deleteConfirmText(p string) opt {
	return func(o map[string]interface{}) {
		o["deleteConfirmText"] = p
	}
}

// 是否可以拖动排序, 需要注意的是当启用拖动排序的时候，会多一个\$id 字段
func Combo_draggable(p bool) opt {
	return func(o map[string]interface{}) {
		o["draggable"] = p
	}
}

// 可拖拽的提示文字
func Combo_draggableTip(p string) opt {
	return func(o map[string]interface{}) {
		o["draggableTip"] = p
	}
}

// 可选`normal`、`horizontal`、`inline`
func Combo_subFormMode(p string) opt {
	return func(o map[string]interface{}) {
		o["subFormMode"] = p
	}
}

// 没有成员时显示。
func Combo_placeholder(p string) opt {
	return func(o map[string]interface{}) {
		o["placeholder"] = p
	}
}

// 指定是否可以自动获取上层的数据并映射到表单项上
func Combo_canAccessSuperData(p bool) opt {
	return func(o map[string]interface{}) {
		o["canAccessSuperData"] = p
	}
}

// 数组的形式包含所有条件的渲染类型，单个数组内的`test` 为判断条件，数组内的`items`为符合该条件后渲染的`schema`
func Combo_conditions(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["conditions"] = p
	}
}

// 是否可切换条件，配合`conditions`使用
func Combo_typeSwitchable(p bool) opt {
	return func(o map[string]interface{}) {
		o["typeSwitchable"] = p
	}
}

// 默认为严格模式，设置为 false 时，当其他表单项更新是，里面的表单项也可以及时获取，否则不会。
func Combo_strictMode(p bool) opt {
	return func(o map[string]interface{}) {
		o["strictMode"] = p
	}
}

// 配置同步字段。只有 `strictMode` 为 `false` 时有效。如果 Combo 层级比较深，底层的获取外层的数据可能不同步。但是给 combo 配置这个属性就能同步下来。输入格式：`["os"]`
func Combo_syncFields(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["syncFields"] = p
	}
}

// 允许为空，如果子表单项里面配置验证器，且又是单条模式。可以允许用户选择清空（不填）。
func Combo_nullable(p bool) opt {
	return func(o map[string]interface{}) {
		o["nullable"] = p
	}
}

// 单组 CSS 类
func Combo_itemClassName(p string) opt {
	return func(o map[string]interface{}) {
		o["itemClassName"] = p
	}
}

// 组合区域 CSS 类
func Combo_itemsWrapperClassName(p string) opt {
	return func(o map[string]interface{}) {
		o["itemsWrapperClassName"] = p
	}
}

// 只有当`removable`为 `true` 时有效; 如果为`string`则为按钮的文本；如果为`Button`则根据配置渲染删除按钮。
func Combo_deleteBtn(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["deleteBtn"] = p
	}
}

// 可新增自定义配置渲染新增按钮，在`tabsMode: true`下不生效。
func Combo_addBtn(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["addBtn"] = p
	}
}

// 新增按钮 CSS 类名
func Combo_addButtonClassName(p string) opt {
	return func(o map[string]interface{}) {
		o["addButtonClassName"] = p
	}
}

// 新增按钮文字
func Combo_addButtonText(p string) opt {
	return func(o map[string]interface{}) {
		o["addButtonText"] = p
	}
}

// 添加组合项时触发
func Combo_add(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["add"] = p
	}
}

// 删除组合项时触发
func Combo_delete(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["delete"] = p
	}
}

// 当设置 tabsMode 为 true 时，切换选项卡时触发
func Combo_tabsChange(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["tabsChange"] = p
	}
}

// 更新数据，对象数组针对开启`multiple`模式, `multiple`模式下可以通过指定`index`来更新指定索引的数据
func Combo_setValue(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["setValue"] = p
	}
}
