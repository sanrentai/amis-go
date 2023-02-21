package form

// 子表单
func InputSubForm(opts ...opt) map[string]interface{} {
	return newForm("input-sub-form", opts...)
}



// 是否为多选模式
func InputSubForm_multiple(p bool) opt {
	return func(o map[string]interface{}) {
		o["multiple"] = p
	}
} 

// 当值中存在这个字段，则按钮名称将使用此字段的值来展示。
func InputSubForm_labelField(p string) opt {
	return func(o map[string]interface{}) {
		o["labelField"] = p
	}
} 

// 按钮默认名称
func InputSubForm_btnLabel(p string) opt {
	return func(o map[string]interface{}) {
		o["btnLabel"] = p
	}
} 

// 限制最小个数。
func InputSubForm_minLength(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["minLength"] = p
	}
} 

// 限制最大个数。
func InputSubForm_maxLength(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["maxLength"] = p
	}
} 

// 是否可拖拽排序
func InputSubForm_draggable(p bool) opt {
	return func(o map[string]interface{}) {
		o["draggable"] = p
	}
} 

// 是否可新增
func InputSubForm_addable(p bool) opt {
	return func(o map[string]interface{}) {
		o["addable"] = p
	}
} 

// 是否可删除
func InputSubForm_removable(p bool) opt {
	return func(o map[string]interface{}) {
		o["removable"] = p
	}
} 

// 新增按钮 CSS 类名
func InputSubForm_addButtonClassName(p string) opt {
	return func(o map[string]interface{}) {
		o["addButtonClassName"] = p
	}
} 

// 值元素 CSS 类名
func InputSubForm_itemClassName(p string) opt {
	return func(o map[string]interface{}) {
		o["itemClassName"] = p
	}
} 

// 值包裹元素 CSS 类名
func InputSubForm_itemsClassName(p string) opt {
	return func(o map[string]interface{}) {
		o["itemsClassName"] = p
	}
} 

// 子表单配置，同 [Form](./index)
func InputSubForm_form(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["form"] = p
	}
} 

// 自定义新增一项的文本
func InputSubForm_addButtonText(p string) opt {
	return func(o map[string]interface{}) {
		o["addButtonText"] = p
	}
} 

// 是否在左下角显示报错信息
func InputSubForm_showErrorMsg(p bool) opt {
	return func(o map[string]interface{}) {
		o["showErrorMsg"] = p
	}
} 