package amis

// 布局
func HBox(opts ...opt) map[string]interface{} {
	return newCompent("hbox", opts...)
}

// 外层 Dom 的类名
func HBox_className(p string) opt {
	return func(o map[string]interface{}) {
		o["className"] = p
	}
}

// 列集合
func HBox_columns(p ...interface{}) opt {
	return func(o map[string]interface{}) {
		o["columns"] = p
	}
}

// // 成员可以是其他渲染器
// func HBox_columns[x](p interface{}) opt {
// 	return func(o map[string]interface{}) {
// 		o["columns[x]"] = p
// 	}
// }

// // 列上类名
// func HBox_columns[x].columnClassName(p string) opt {
// 	return func(o map[string]interface{}) {
// 		o["columns[x].columnClassName"] = p
// 	}
// }
