package form

// 公式
func Formula(opts ...opt) map[string]interface{} {
	return newForm("formula", opts...)
}



// 需要应用的表单项`name`值，公式结果将作用到此处指定的变量中去。
func Formula_name(p string) opt {
	return func(o map[string]interface{}) {
		o["name"] = p
	}
} 

// 应用的公式
func Formula_formula(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["formula"] = p
	}
} 

// 公式作用条件
func Formula_condition(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["condition"] = p
	}
} 

// 初始化时是否设置
func Formula_initSet(p bool) opt {
	return func(o map[string]interface{}) {
		o["initSet"] = p
	}
} 

// 观察公式结果，如果计算结果有变化，则自动应用到变量上
func Formula_autoSet(p bool) opt {
	return func(o map[string]interface{}) {
		o["autoSet"] = p
	}
} 

// 定义个名字，当某个按钮的目标指定为此值后，会触发一次公式应用。这个机制可以在 `autoSet` 为 false 时用来手动触发
func Formula_id(p bool) opt {
	return func(o map[string]interface{}) {
		o["id"] = p
	}
} 