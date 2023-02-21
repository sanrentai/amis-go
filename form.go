package amis

// 表单
func Form(opts ...opt) map[string]interface{} {
	return newCompent("form", opts...)
}

// 设置一个名字后，方便其他组件与其通信
func Form_name(p string) opt {
	return func(o map[string]interface{}) {
		o["name"] = p
	}
}

// 表单展示方式，可以是：`normal`、`horizontal` 或者 `inline`
func Form_mode(p string) opt {
	return func(o map[string]interface{}) {
		o["mode"] = p
	}
}

// 当 mode 为 `horizontal` 时有用，用来控制 label
func Form_horizontal(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["horizontal"] = p
	}
}

// Form 的标题
func Form_title(p string) opt {
	return func(o map[string]interface{}) {
		o["title"] = p
	}
}

// 默认的提交按钮名称，如果设置成空，则可以把默认按钮去掉。
func Form_submitText(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["submitText"] = p
	}
}

// 外层 Dom 的类名
func Form_className(p string) opt {
	return func(o map[string]interface{}) {
		o["className"] = p
	}
}

// Form 表单项集合
func Form_body(p ...interface{}) opt {
	return func(o map[string]interface{}) {
		o["body"] = p
	}
}

// Form 提交按钮，成员为 Action
func Form_actions(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["actions"] = p
	}
}

// 消息提示覆写，默认消息读取的是 API 返回的消息，但是在此可以覆写它。
func Form_messages(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["messages"] = p
	}
}

// // 获取成功时提示
// func Form_messages.fetchSuccess(p string) opt {
// 	return func(o map[string]interface{}) {
// 		o["messages.fetchSuccess"] = p
// 	}
// }

// // 获取失败时提示
// func Form_messages.fetchFailed(p string) opt {
// 	return func(o map[string]interface{}) {
// 		o["messages.fetchFailed"] = p
// 	}
// }

// // 保存成功时提示
// func Form_messages.saveSuccess(p string) opt {
// 	return func(o map[string]interface{}) {
// 		o["messages.saveSuccess"] = p
// 	}
// }

// // 保存失败时提示
// func Form_messages.saveFailed(p string) opt {
// 	return func(o map[string]interface{}) {
// 		o["messages.saveFailed"] = p
// 	}
// }

// 是否让 Form 用 panel 包起来，设置为 false 后，actions 将无效。
func Form_wrapWithPanel(p bool) opt {
	return func(o map[string]interface{}) {
		o["wrapWithPanel"] = p
	}
}

// 外层 panel 的类名
func Form_panelClassName(p string) opt {
	return func(o map[string]interface{}) {
		o["panelClassName"] = p
	}
}

// Form 用来保存数据的 api。
func Form_api(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["api"] = p
	}
}

// Form 用来获取初始数据的 api。
func Form_initApi(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["initApi"] = p
	}
}

// 表单组合校验规则
func Form_rules(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["rules"] = p
	}
}

// 刷新时间(最低 3000)
func Form_interval(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["interval"] = p
	}
}

// 配置刷新时是否显示加载动画
func Form_silentPolling(p bool) opt {
	return func(o map[string]interface{}) {
		o["silentPolling"] = p
	}
}

// 通过[表达式](./Types.md#表达式) 来配置停止刷新的条件
func Form_stopAutoRefreshWhen(p string) opt {
	return func(o map[string]interface{}) {
		o["stopAutoRefreshWhen"] = p
	}
}

// Form 用来获取初始数据的 api,与 initApi 不同的是，会一直轮询请求该接口，直到返回 finished 属性为 true 才 结束。
func Form_initAsyncApi(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["initAsyncApi"] = p
	}
}

// 设置了 initApi 或者 initAsyncApi 后，默认会开始就发请求，设置为 false 后就不会起始就请求接口
func Form_initFetch(p bool) opt {
	return func(o map[string]interface{}) {
		o["initFetch"] = p
	}
}

// 用表达式来配置
func Form_initFetchOn(p string) opt {
	return func(o map[string]interface{}) {
		o["initFetchOn"] = p
	}
}

// 设置了 initAsyncApi 后，默认会从返回数据的 data.finished 来判断是否完成，也可以设置成其他的 xxx，就会从 data.xxx 中获取
func Form_initFinishedField(p string) opt {
	return func(o map[string]interface{}) {
		o["initFinishedField"] = p
	}
}

// 设置了 initAsyncApi 以后，默认拉取的时间间隔
func Form_initCheckInterval(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["initCheckInterval"] = p
	}
}

// 设置此属性后，表单提交发送保存接口后，还会继续轮询请求该接口，直到返回 `finished` 属性为 `true` 才 结束。
func Form_asyncApi(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["asyncApi"] = p
	}
}

// 轮询请求的时间间隔，默认为 3 秒。设置 `asyncApi` 才有效
func Form_checkInterval(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["checkInterval"] = p
	}
}

// 如果决定结束的字段名不是 `finished` 请设置此属性，比如 `is_success`
func Form_finishedField(p string) opt {
	return func(o map[string]interface{}) {
		o["finishedField"] = p
	}
}

// 表单修改即提交
func Form_submitOnChange(p bool) opt {
	return func(o map[string]interface{}) {
		o["submitOnChange"] = p
	}
}

// 初始就提交一次
func Form_submitOnInit(p bool) opt {
	return func(o map[string]interface{}) {
		o["submitOnInit"] = p
	}
}

// 提交后是否重置表单
func Form_resetAfterSubmit(p bool) opt {
	return func(o map[string]interface{}) {
		o["resetAfterSubmit"] = p
	}
}

// 设置主键 id, 当设置后，检测表单是否完成时（asyncApi），只会携带此数据。
func Form_primaryField(p string) opt {
	return func(o map[string]interface{}) {
		o["primaryField"] = p
	}
}

// 默认表单提交自己会通过发送 api 保存数据，但是也可以设定另外一个 form 的 name 值，或者另外一个 `CRUD` 模型的 name 值。 如果 target 目标是一个 `Form` ，则目标 `Form` 会重新触发 `initApi`，api 可以拿到当前 form 数据。如果目标是一个 `CRUD` 模型，则目标模型会重新触发搜索，参数为当前 Form 数据。当目标是 `window` 时，会把当前表单的数据附带到页面地址上。
func Form_target(p string) opt {
	return func(o map[string]interface{}) {
		o["target"] = p
	}
}

// 设置此属性后，Form 保存成功后，自动跳转到指定页面。支持相对地址，和绝对地址（相对于组内的）。
func Form_redirect(p string) opt {
	return func(o map[string]interface{}) {
		o["redirect"] = p
	}
}

// 操作完后刷新目标对象。请填写目标组件设置的 name 值，如果填写为 `window` 则让当前页面整体刷新。
func Form_reload(p string) opt {
	return func(o map[string]interface{}) {
		o["reload"] = p
	}
}

// 是否自动聚焦。
func Form_autoFocus(p bool) opt {
	return func(o map[string]interface{}) {
		o["autoFocus"] = p
	}
}

// 指定是否可以自动获取上层的数据并映射到表单项上
func Form_canAccessSuperData(p bool) opt {
	return func(o map[string]interface{}) {
		o["canAccessSuperData"] = p
	}
}

// 指定一个唯一的 key，来配置当前表单是否开启本地缓存
func Form_persistData(p string) opt {
	return func(o map[string]interface{}) {
		o["persistData"] = p
	}
}

// 指指定只有哪些 key 缓存
func Form_persistDataKeys(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["persistDataKeys"] = p
	}
}

// 指定表单提交成功后是否清除本地缓存
func Form_clearPersistDataAfterSubmit(p bool) opt {
	return func(o map[string]interface{}) {
		o["clearPersistDataAfterSubmit"] = p
	}
}

// 禁用回车提交表单
func Form_preventEnterSubmit(p bool) opt {
	return func(o map[string]interface{}) {
		o["preventEnterSubmit"] = p
	}
}

// trim 当前表单项的每一个值
func Form_trimValues(p bool) opt {
	return func(o map[string]interface{}) {
		o["trimValues"] = p
	}
}

// form 还没保存，即将离开页面前是否弹框确认。
func Form_promptPageLeave(p bool) opt {
	return func(o map[string]interface{}) {
		o["promptPageLeave"] = p
	}
}

// 表单项显示为几列
func Form_columnCount(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["columnCount"] = p
	}
}

// 默认表单是采用数据链的形式创建个自己的数据域，表单提交的时候只会发送自己这个数据域的数据，如果希望共用上层数据域可以设置这个属性为 false，这样上层数据域的数据不需要在表单中用隐藏域或者显式映射才能发送了。
func Form_inheritData(p bool) opt {
	return func(o map[string]interface{}) {
		o["inheritData"] = p
	}
}

// `2.4.0` 整个表单静态方式展示，详情请查看[示例页](../../../examples/form/switchDisplay)
func Form_static(p bool) opt {
	return func(o map[string]interface{}) {
		o["static"] = p
	}
}

// `2.4.0` 表单静态展示时使用的类名
func Form_staticClassName(p string) opt {
	return func(o map[string]interface{}) {
		o["staticClassName"] = p
	}
}
