package amis

func Api(opts ...opt) map[string]interface{} {
	var o = map[string]interface{}{}
	for _, op := range opts {
		op(o)
	}
	return o
}

// 请求方式     | 字符串 支持：get、post、put、delete
func Api_method(p string) opt {
	return func(o map[string]interface{}) {
		o["method"] = p
	}
}

// 请求地址     |
// [模板字符串](https://suda.bce.baidu.com/docs/template#%E6%A8%A1%E6%9D%BF%E5%AD%97%E7%AC%A6%E4%B8%B2) |
func Api_url(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["url"] = p
	}
}

//  请求数据     | 对象或字符串       |
func Api_data(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["data"] = p
	}
}

// 数据体格式   | 字符串
// 默认为 `json` 可以配置成 `form` 或者 `form-data`。当 `data` 中包含文件时，自动会采用 `form-data（multipart/form-data）` 格式。当配置为 `form` 时为 `application/x-www-form-urlencoded` 格式。 |
func Api_dataType(p string) opt {
	return func(o map[string]interface{}) {
		o["dataType"] = p
	}
}

// 对象或字符串
//  当 dataType 为 form 或者 form-data 的时候有用。具体参数请参考这里，默认设置为: `{ arrayFormat: 'indices', encodeValuesOnly: true }`                                                           |
func Api_qsOptions(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["qsOptions"] = p
	}
}

// 请求头       | 对象
func Api_headers(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["headers"] = p
	}
}

// 请求条件     | [表达式](../concepts/expression)
func Api_sendOn(p string) opt {
	return func(o map[string]interface{}) {
		o["sendOn"] = p
	}
}

// 接口缓存时间 | 整型数字
func Api_cache(p int64) opt {
	return func(o map[string]interface{}) {
		o["cache"] = p
	}
}

// 发送适配器   | 字符串
func Api_requestAdaptor(p string) opt {
	return func(o map[string]interface{}) {
		o["requestAdaptor"] = p
	}
}

// 接收适配器   | 字符串
func Api_adaptor(p string) opt {
	return func(o map[string]interface{}) {
		o["adaptor"] = p
	}
}

// 替换当前数据 | 布尔
func Api_replaceData(p bool) opt {
	return func(o map[string]interface{}) {
		o["replaceData"] = p
	}
}

// 返回类型     | 字符串
func Api_responseType(p string) opt {
	return func(o map[string]interface{}) {
		o["responseType"] = p
	}
}

// 是否自动刷新 | 布尔
func Api_autoRefresh(p bool) opt {
	return func(o map[string]interface{}) {
		o["autoRefresh"] = p
	}
}

// 配置返回数据 | 对象   对返回结果做个映射                                                                                                                                                                            |
func Api_responseData(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["responseData"] = p
	}
}

//
func Api_trackExpression(p string) opt {
	return func(o map[string]interface{}) {
		o["trackExpression"] = p
	}
}

// 提示信息
// 配置接口请求的提示信息，messages.success 表示请求成功提示信息、messages.failed 表示请求失败提示信息，2.4.1 及以上版本                                                                         |
func Api_messages(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["messages"] = p
	}
}
