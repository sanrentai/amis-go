package amis

// 图表
func Chart(opts ...opt) map[string]interface{} {
	return newCompent("chart", opts...)
}

// 外层 Dom 的类名
func Chart_className(p string) opt {
	return func(o map[string]interface{}) {
		o["className"] = p
	}
}

// 内容容器
func Chart_body(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["body"] = p
	}
}

// 配置项接口地址
func Chart_api(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["api"] = p
	}
}

// 通过数据映射获取数据链中变量值作为配置
func Chart_source(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["source"] = p
	}
}

// 组件初始化时，是否请求接口
func Chart_initFetch(p bool) opt {
	return func(o map[string]interface{}) {
		o["initFetch"] = p
	}
}

// 刷新时间(最小 1000)
func Chart_interval(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["interval"] = p
	}
}

// 设置 eschars 的配置项,当为`string`的时候可以设置 function 等配置项
func Chart_config(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["config"] = p
	}
}

// 设置根元素的 style
func Chart_style(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["style"] = p
	}
}

// 设置根元素的宽度
func Chart_width(p string) opt {
	return func(o map[string]interface{}) {
		o["width"] = p
	}
}

// 设置根元素的高度
func Chart_height(p string) opt {
	return func(o map[string]interface{}) {
		o["height"] = p
	}
}

// 每次更新是完全覆盖配置项还是追加？
func Chart_replaceChartOption(p bool) opt {
	return func(o map[string]interface{}) {
		o["replaceChartOption"] = p
	}
}

// 当这个表达式的值有变化时更新图表
func Chart_trackExpression(p string) opt {
	return func(o map[string]interface{}) {
		o["trackExpression"] = p
	}
}

// 自定义 echart config 转换，函数签名：function(config, echarts, data) {return config;} 配置时直接写函数体。其中 config 是当前 echart 配置，echarts 就是 echarts 对象，data 为上下文数据。
func Chart_dataFilter(p string) opt {
	return func(o map[string]interface{}) {
		o["dataFilter"] = p
	}
}

// 地图 geo json 地址
func Chart_mapURL(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["mapURL"] = p
	}
}

// 地图名称
func Chart_mapName(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["mapName"] = p
	}
}

// 加载百度地图
func Chart_loadBaiduMap(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["loadBaiduMap"] = p
	}
}
