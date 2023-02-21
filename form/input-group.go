package form

// 输入框组合
func InputGroup(opts ...opt) map[string]interface{} {
	return newForm("input-group", opts...)
}



// CSS 类名
func InputGroup_className(p string) opt {
	return func(o map[string]interface{}) {
		o["className"] = p
	}
} 

// 表单项集合
func InputGroup_body(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["body"] = p
	}
} 