package form

// 重复频率选择器
func InputRepeat(opts ...opt) map[string]interface{} {
	return newForm("input-repeat", opts...)
}



// 可用配置 `secondly,minutely,hourly,daily,weekdays,weekly,monthly,yearly`
func InputRepeat_options(p string) opt {
	return func(o map[string]interface{}) {
		o["options"] = p
	}
} 

// 当不指定值时的说明。
func InputRepeat_placeholder(p string) opt {
	return func(o map[string]interface{}) {
		o["placeholder"] = p
	}
} 