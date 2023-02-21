package amis

// 任务操作集合
func Tasks(opts ...opt) map[string]interface{} {
	return newCompent("tasks", opts...)
}

// 外层 Dom 的类名
func Tasks_className(p string) opt {
	return func(o map[string]interface{}) {
		o["className"] = p
	}
}

// table Dom 的类名
func Tasks_tableClassName(p string) opt {
	return func(o map[string]interface{}) {
		o["tableClassName"] = p
	}
}

// 任务列表
func Tasks_items(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["items"] = p
	}
}

// // 任务名称
// func Tasks_items[x].label(p string) opt {
// 	return func(o map[string]interface{}) {
// 		o["items[x].label"] = p
// 	}
// }

// // 任务键值，请唯一区分
// func Tasks_items[x].key(p string) opt {
// 	return func(o map[string]interface{}) {
// 		o["items[x].key"] = p
// 	}
// }

// // 当前任务状态，支持 html
// func Tasks_items[x].remark(p string) opt {
// 	return func(o map[string]interface{}) {
// 		o["items[x].remark"] = p
// 	}
// }

// // 任务状态： 0: 初始状态，不可操作。1: 就绪，可操作状态。2: 进行中，还没有结束。3：有错误，不可重试。4: 已正常结束。5：有错误，且可以重试。
// func Tasks_items[x].status(p string) opt {
// 	return func(o map[string]interface{}) {
// 		o["items[x].status"] = p
// 	}
// }

// 返回任务列表，返回的数据请参考 items。
func Tasks_checkApi(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["checkApi"] = p
	}
}

// 提交任务使用的 API
func Tasks_submitApi(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["submitApi"] = p
	}
}

// 如果任务失败，且可以重试，提交的时候会使用此 API
func Tasks_reSubmitApi(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["reSubmitApi"] = p
	}
}

// 当有任务进行中，会每隔一段时间再次检测，而时间间隔就是通过此项配置，默认 3s。
func Tasks_interval(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["interval"] = p
	}
}

// 任务名称列说明
func Tasks_taskNameLabel(p string) opt {
	return func(o map[string]interface{}) {
		o["taskNameLabel"] = p
	}
}

// 操作列说明
func Tasks_operationLabel(p string) opt {
	return func(o map[string]interface{}) {
		o["operationLabel"] = p
	}
}

// 状态列说明
func Tasks_statusLabel(p string) opt {
	return func(o map[string]interface{}) {
		o["statusLabel"] = p
	}
}

// 备注列说明
func Tasks_remarkLabel(p string) opt {
	return func(o map[string]interface{}) {
		o["remarkLabel"] = p
	}
}

// 操作按钮文字
func Tasks_btnText(p string) opt {
	return func(o map[string]interface{}) {
		o["btnText"] = p
	}
}

// 重试操作按钮文字
func Tasks_retryBtnText(p string) opt {
	return func(o map[string]interface{}) {
		o["retryBtnText"] = p
	}
}

// 配置容器按钮 className
func Tasks_btnClassName(p string) opt {
	return func(o map[string]interface{}) {
		o["btnClassName"] = p
	}
}

// 配置容器重试按钮 className
func Tasks_retryBtnClassName(p string) opt {
	return func(o map[string]interface{}) {
		o["retryBtnClassName"] = p
	}
}

// 状态显示对应的类名配置
func Tasks_statusLabelMap(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["statusLabelMap"] = p
	}
}

// 状态显示对应的文字显示配置
func Tasks_statusTextMap(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["statusTextMap"] = p
	}
}
