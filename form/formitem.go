package form

// 普通表单项
func FormItem(opts ...opt) map[string]interface{} {
	return newForm("formitem", opts...)
}

// 表单最外层类名
func FormItem_className(p string) opt {
	return func(o map[string]interface{}) {
		o["className"] = p
	}
}

// 表单控制器类名
func FormItem_inputClassName(p string) opt {
	return func(o map[string]interface{}) {
		o["inputClassName"] = p
	}
}

// label 的类名
func FormItem_labelClassName(p string) opt {
	return func(o map[string]interface{}) {
		o["labelClassName"] = p
	}
}

// 字段名，指定该表单项提交时的 key
func FormItem_name(p string) opt {
	return func(o map[string]interface{}) {
		o["name"] = p
	}
}

// 表单默认值
func FormItem_value(p string) opt {
	return func(o map[string]interface{}) {
		o["value"] = p
	}
}

// 表单项标签
func FormItem_label(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["label"] = p
	}
}

// 表单项标签描述
func FormItem_labelRemark(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["labelRemark"] = p
	}
}

// 表单项描述
func FormItem_description(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["description"] = p
	}
}

// 表单项描述
func FormItem_placeholder(p string) opt {
	return func(o map[string]interface{}) {
		o["placeholder"] = p
	}
}

// 是否为 内联 模式
func FormItem_inline(p bool) opt {
	return func(o map[string]interface{}) {
		o["inline"] = p
	}
}

// 是否该表单项值发生变化时就提交当前表单。
func FormItem_submitOnChange(p bool) opt {
	return func(o map[string]interface{}) {
		o["submitOnChange"] = p
	}
}

// 当前表单项是否是禁用状态
func FormItem_disabled(p bool) opt {
	return func(o map[string]interface{}) {
		o["disabled"] = p
	}
}

// 当前表单项是否禁用的条件
func FormItem_disabledOn(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["disabledOn"] = p
	}
}

// 当前表单项是否禁用的条件
func FormItem_visible(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["visible"] = p
	}
}

// 当前表单项是否禁用的条件
func FormItem_visibleOn(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["visibleOn"] = p
	}
}

// 是否为必填。
func FormItem_required(p bool) opt {
	return func(o map[string]interface{}) {
		o["required"] = p
	}
}

// 通过[表达式](../Types.md#表达式)来配置当前表单项是否为必填。
func FormItem_requiredOn(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["requiredOn"] = p
	}
}

// 表单项值格式验证，支持设置多个，多个规则用英文逗号隔开。
func FormItem_validations(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["validations"] = p
	}
}

// 表单校验接口
func FormItem_validateApi(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["validateApi"] = p
	}
}

// 数据录入配置，自动填充或者参照录入
func FormItem_autoFill(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["autoFill"] = p
	}
}

// // true 为参照录入，false 自动填充
// func FormItem_autoFill.showSuggestion(p bool) opt {
// 	return func(o map[string]interface{}) {
// 		o["autoFill.showSuggestion"] = p
// 	}
// }

// // 自动填充接口/参照录入筛选 CRUD 请求配置
// func FormItem_autoFill.api(p interface{}) opt {
// 	return func(o map[string]interface{}) {
// 		o["autoFill.api"] = p
// 	}
// }

// // 是否展示数据格式错误提示，默认为 true
// func FormItem_autoFill.silent(p bool) opt {
// 	return func(o map[string]interface{}) {
// 		o["autoFill.silent"] = p
// 	}
// }

// // 自动填充/参照录入数据映射配置，键值对形式，值支持变量获取及表达式
// func FormItem_autoFill.fillMappinng(p interface{}) opt {
// 	return func(o map[string]interface{}) {
// 		o["autoFill.fillMappinng"] = p
// 	}
// }

// // showSuggestion 为 true 时，参照录入支持的触发方式，目前支持 change「值变化」｜ focus 「表单项聚焦」
// func FormItem_autoFill.trigger(p string) opt {
// 	return func(o map[string]interface{}) {
// 		o["autoFill.trigger"] = p
// 	}
// }

// // showSuggestion 为 true 时，参照弹出方式 dialog, drawer, popOver
// func FormItem_autoFill.mode(p string) opt {
// 	return func(o map[string]interface{}) {
// 		o["autoFill.mode"] = p
// 	}
// }

// // showSuggestion 为 true 时，设置弹出 dialog,drawer,popOver 中 picker 的 labelField
// func FormItem_autoFill.labelField(p string) opt {
// 	return func(o map[string]interface{}) {
// 		o["autoFill.labelField"] = p
// 	}
// }

// // showSuggestion 为 true 时，参照录入 mode 为 popOver 时，可配置弹出位置
// func FormItem_autoFill.position(p string) opt {
// 	return func(o map[string]interface{}) {
// 		o["autoFill.position"] = p
// 	}
// }

// // showSuggestion 为 true 时，参照录入 mode 为 dialog 时，可设置大小
// func FormItem_autoFill.size(p string) opt {
// 	return func(o map[string]interface{}) {
// 		o["autoFill.size"] = p
// 	}
// }

// // showSuggestion 为 true 时，数据展示列配置
// func FormItem_autoFill.columns(p interface{}) opt {
// 	return func(o map[string]interface{}) {
// 		o["autoFill.columns"] = p
// 	}
// }

// // showSuggestion 为 true 时，数据查询过滤条件
// func FormItem_autoFill.filter(p interface{}) opt {
// 	return func(o map[string]interface{}) {
// 		o["autoFill.filter"] = p
// 	}
// }

// `2.4.0` 当前表单项是否是静态展示，目前支持静[支持静态展示的表单项](#支持静态展示的表单项)
func FormItem_static(p bool) opt {
	return func(o map[string]interface{}) {
		o["static"] = p
	}
}

// `2.4.0` 静态展示时的类名
func FormItem_staticClassName(p string) opt {
	return func(o map[string]interface{}) {
		o["staticClassName"] = p
	}
}

// `2.4.0` 静态展示时的 Label 的类名
func FormItem_staticLabelClassName(p string) opt {
	return func(o map[string]interface{}) {
		o["staticLabelClassName"] = p
	}
}

// `2.4.0` 静态展示时的 value 的类名
func FormItem_staticInputClassName(p string) opt {
	return func(o map[string]interface{}) {
		o["staticInputClassName"] = p
	}
}

// `2.4.0` 自定义静态展示方式
func FormItem_staticSchema(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["staticSchema"] = p
	}
}

// // `2.4.0` select、checkboxes 等选择类组件多选时展示态展示的数量
// func FormItem_staticSchema.limit(p interface{}) opt {
// 	return func(o map[string]interface{}) {
// 		o["staticSchema.limit"] = p
// 	}
// }
