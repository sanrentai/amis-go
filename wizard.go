package amis

// 向导
func Wizard(opts ...opt) map[string]interface{} {
	return newCompent("wizard", opts...)
}

// 最后一步保存的接口。
func Wizard_api(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["api"] = p
	}
}

// 初始是否拉取数据，通过表达式来配置
func Wizard_initFetchOn(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["initFetchOn"] = p
	}
}

// 上一步按钮文本
func Wizard_actionPrevLabel(p string) opt {
	return func(o map[string]interface{}) {
		o["actionPrevLabel"] = p
	}
}

// 下一步按钮文本
func Wizard_actionNextLabel(p string) opt {
	return func(o map[string]interface{}) {
		o["actionNextLabel"] = p
	}
}

// 保存并下一步按钮文本
func Wizard_actionNextSaveLabel(p string) opt {
	return func(o map[string]interface{}) {
		o["actionNextSaveLabel"] = p
	}
}

// 完成按钮文本
func Wizard_actionFinishLabel(p string) opt {
	return func(o map[string]interface{}) {
		o["actionFinishLabel"] = p
	}
}

// 外层 CSS 类名
func Wizard_className(p string) opt {
	return func(o map[string]interface{}) {
		o["className"] = p
	}
}

// 按钮 CSS 类名
func Wizard_actionClassName(p string) opt {
	return func(o map[string]interface{}) {
		o["actionClassName"] = p
	}
}

// 操作完后刷新目标对象。请填写目标组件设置的 name 值，如果填写为 `window` 则让当前页面整体刷新。
func Wizard_reload(p string) opt {
	return func(o map[string]interface{}) {
		o["reload"] = p
	}
}

// 操作完后跳转。
func Wizard_redirect(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["redirect"] = p
	}
}

// 可以把数据提交给别的组件而不是自己保存。请填写目标组件设置的 name 值，如果填写为 `window` 则把数据同步到地址栏上，同时依赖这些数据的组件会自动重新刷新。
func Wizard_target(p string) opt {
	return func(o map[string]interface{}) {
		o["target"] = p
	}
}

// 数组，配置步骤信息
func Wizard_steps(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["steps"] = p
	}
}

// 起始默认值，从第几步开始。可支持模版，但是只有在组件创建时渲染模版并设置当前步数，在之后组件被刷新时，当前 step 不会根据 startStep 改变
func Wizard_startStep(p string) opt {
	return func(o map[string]interface{}) {
		o["startStep"] = p
	}
}

// 说明
func Wizard_属性名(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["属性名"] = p
	}
}

// 步骤标题
func Wizard_title(p string) opt {
	return func(o map[string]interface{}) {
		o["title"] = p
	}
}

// 展示默认，跟 [Form](./Form/Form) 中的模式一样，选择： `normal`、`horizontal`或者`inline`。
func Wizard_mode(p string) opt {
	return func(o map[string]interface{}) {
		o["mode"] = p
	}
}

// 当为水平模式时，用来控制左右占比
func Wizard_horizontal(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["horizontal"] = p
	}
}

// // 左边 label 的宽度占比
// func Wizard_horizontal.label(p interface{}) opt {
// 	return func(o map[string]interface{}) {
// 		o["horizontal.label"] = p
// 	}
// }

// // 右边控制器的宽度占比。
// func Wizard_horizontal.right(p interface{}) opt {
// 	return func(o map[string]interface{}) {
// 		o["horizontal.right"] = p
// 	}
// }

// // 当没有设置 label 时，右边控制器的偏移量
// func Wizard_horizontal.offset(p interface{}) opt {
// 	return func(o map[string]interface{}) {
// 		o["horizontal.offset"] = p
// 	}
// }

// 当前步骤数据初始化接口。
func Wizard_initApi(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["initApi"] = p
	}
}

// 当前步骤数据初始化接口是否初始拉取。
func Wizard_initFetch(p bool) opt {
	return func(o map[string]interface{}) {
		o["initFetch"] = p
	}
}

// 当前步骤的表单项集合，请参考 [FormItem](./form/formItem)。
func Wizard_body(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["body"] = p
	}
}
