package form

// 表单项集合
func FieldSet(opts ...opt) map[string]interface{} {
	return newForm("fieldset", opts...)
}



// CSS 类名
func FieldSet_className(p string) opt {
	return func(o map[string]interface{}) {
		o["className"] = p
	}
} 

// 标题 CSS 类名
func FieldSet_headingClassName(p string) opt {
	return func(o map[string]interface{}) {
		o["headingClassName"] = p
	}
} 

// 内容区域 CSS 类名
func FieldSet_bodyClassName(p string) opt {
	return func(o map[string]interface{}) {
		o["bodyClassName"] = p
	}
} 

// 标题
func FieldSet_title(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["title"] = p
	}
} 

// 表单项集合
func FieldSet_body(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["body"] = p
	}
} 

// 展示默认，同 [Form](./index#%E8%A1%A8%E5%8D%95%E5%B1%95%E7%A4%BA) 中的模式
func FieldSet_mode(p string) opt {
	return func(o map[string]interface{}) {
		o["mode"] = p
	}
} 

// 是否可折叠
func FieldSet_collapsable(p bool) opt {
	return func(o map[string]interface{}) {
		o["collapsable"] = p
	}
} 

// 默认是否折叠
func FieldSet_collapsed(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["collapsed"] = p
	}
} 

// 收起的标题
func FieldSet_collapseTitle(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["collapseTitle"] = p
	}
} 

// 大小，支持 xs、sm、base、lg、xl
func FieldSet_size(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["size"] = p
	}
} 