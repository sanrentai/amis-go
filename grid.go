package amis

// 水平分栏
func Grid(opts ...opt) map[string]interface{} {
	return newCompent("grid", opts...)
}

// 外层 Dom 的类名
func Grid_className(p string) opt {
	return func(o map[string]interface{}) {
		o["className"] = p
	}
}

// 列集合
func Grid_columns(p ...interface{}) opt {
	return func(o map[string]interface{}) {
		o["columns"] = p
	}
}

// // 成员可以是其他渲染器
// func Grid_columns[x](p interface{}) opt {
// 	return func(o map[string]interface{}) {
// 		o["columns[x]"] = p
// 	}
// }

// // 宽度占比： 1 - 12
// func Grid_columns[x].xs(p interface{}) opt {
// 	return func(o map[string]interface{}) {
// 		o["columns[x].xs"] = p
// 	}
// }

// // 列类名
// func Grid_columns[x].columnClassName(p interface{}) opt {
// 	return func(o map[string]interface{}) {
// 		o["columns[x].columnClassName"] = p
// 	}
// }

// // 宽度占比： 1 - 12
// func Grid_columns[x].sm(p interface{}) opt {
// 	return func(o map[string]interface{}) {
// 		o["columns[x].sm"] = p
// 	}
// }

// // 宽度占比： 1 - 12
// func Grid_columns[x].md(p interface{}) opt {
// 	return func(o map[string]interface{}) {
// 		o["columns[x].md"] = p
// 	}
// }

// // 宽度占比： 1 - 12
// func Grid_columns[x].lg(p interface{}) opt {
// 	return func(o map[string]interface{}) {
// 		o["columns[x].lg"] = p
// 	}
// }
