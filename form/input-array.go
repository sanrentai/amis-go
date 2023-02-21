package form

// 数组输入框
func InputArray(opts ...opt) map[string]interface{} {
	return newForm("input-array", opts...)
}



// 配置单项表单类型
func InputArray_items(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["items"] = p
	}
} 

// 是否可新增。
func InputArray_addable(p bool) opt {
	return func(o map[string]interface{}) {
		o["addable"] = p
	}
} 

// 是否可删除
func InputArray_removable(p bool) opt {
	return func(o map[string]interface{}) {
		o["removable"] = p
	}
} 

// 是否可以拖动排序, 需要注意的是当启用拖动排序的时候，会多一个\$id 字段
func InputArray_draggable(p bool) opt {
	return func(o map[string]interface{}) {
		o["draggable"] = p
	}
} 

// 可拖拽的提示文字，默认为：`"可通过拖动每行中的【交换】按钮进行顺序调整"`
func InputArray_draggableTip(p string) opt {
	return func(o map[string]interface{}) {
		o["draggableTip"] = p
	}
} 

// 新增按钮文字
func InputArray_addButtonText(p string) opt {
	return func(o map[string]interface{}) {
		o["addButtonText"] = p
	}
} 

// 限制最小长度
func InputArray_minLength(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["minLength"] = p
	}
} 

// 限制最大长度
func InputArray_maxLength(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["maxLength"] = p
	}
} 

// 新增成员时的默认值，一般根据`items`的数据类型指定需要的默认值
func InputArray_scaffold(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["scaffold"] = p
	}
} 