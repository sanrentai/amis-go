package amis

// 功能型容器
func Service(opts ...opt) map[string]interface{} {
	return newCompent("service", opts...)
}

// 外层 Dom 的类名
func Service_className(p string) opt {
	return func(o map[string]interface{}) {
		o["className"] = p
	}
}

// 内容容器
func Service_body(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["body"] = p
	}
}

// 初始化数据域接口地址
func Service_api(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["api"] = p
	}
}

// WebScocket 地址
func Service_ws(p string) opt {
	return func(o map[string]interface{}) {
		o["ws"] = p
	}
}

// 是否默认拉取
func Service_initFetch(p bool) opt {
	return func(o map[string]interface{}) {
		o["initFetch"] = p
	}
}

// 用来获取远程 Schema 接口地址
func Service_schemaApi(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["schemaApi"] = p
	}
}

// 是否默认拉取 Schema
func Service_initFetchSchema(p bool) opt {
	return func(o map[string]interface{}) {
		o["initFetchSchema"] = p
	}
}

// 消息提示覆写，默认消息读取的是接口返回的 toast 提示文字，但是在此可以覆写它。
func Service_messages(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["messages"] = p
	}
}

// // 接口请求成功时的 toast 提示文字
// func Service_messages.fetchSuccess(p string) opt {
// 	return func(o map[string]interface{}) {
// 		o["messages.fetchSuccess"] = p
// 	}
// }

// // 接口请求失败时 toast 提示文字
// func Service_messages.fetchFailed(p string) opt {
// 	return func(o map[string]interface{}) {
// 		o["messages.fetchFailed"] = p
// 	}
// }

// 轮询时间间隔，单位 ms(最低 1000)
func Service_interval(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["interval"] = p
	}
}

// 配置轮询时是否显示加载动画
func Service_silentPolling(p bool) opt {
	return func(o map[string]interface{}) {
		o["silentPolling"] = p
	}
}

// 配置停止轮询的条件
func Service_stopAutoRefreshWhen(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["stopAutoRefreshWhen"] = p
	}
}
