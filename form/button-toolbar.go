package form

// 按钮工具栏
func ButtonToolbar(opts ...opt) map[string]interface{} {
	return newForm("button-toolbar", opts...)
}



// 按钮组
func ButtonToolbar_buttons(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["buttons"] = p
	}
} 