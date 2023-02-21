package form

// 表单项组
func Group(opts ...opt) map[string]interface{} {
	return newForm("group", opts...)
}



// CSS 类名
func Group_className(p string) opt {
	return func(o map[string]interface{}) {
		o["className"] = p
	}
} 

// group 的标签
func Group_label(p string) opt {
	return func(o map[string]interface{}) {
		o["label"] = p
	}
} 

// 表单项集合
func Group_body(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["body"] = p
	}
} 

// 展示默认，同 [Form](./index#%E8%A1%A8%E5%8D%95%E5%B1%95%E7%A4%BA) 中的模式
func Group_mode(p string) opt {
	return func(o map[string]interface{}) {
		o["mode"] = p
	}
} 

// 表单项之间的间距，可选：`xs`、`sm`、`normal`
func Group_gap(p string) opt {
	return func(o map[string]interface{}) {
		o["gap"] = p
	}
} 

// 可以配置水平展示还是垂直展示。对应的配置项分别是：`vertical`、`horizontal`
func Group_direction(p string) opt {
	return func(o map[string]interface{}) {
		o["direction"] = p
	}
} 