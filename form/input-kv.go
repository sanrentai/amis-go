package form

// 键值对
func InputKV(opts ...opt) map[string]interface{} {
	return newForm("input-kv", opts...)
}



// 值类型
func InputKV_valueType(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["valueType"] = p
	}
} 

// key 的提示信息的
func InputKV_keyPlaceholder(p string) opt {
	return func(o map[string]interface{}) {
		o["keyPlaceholder"] = p
	}
} 

// value 的提示信息的
func InputKV_valuePlaceholder(p string) opt {
	return func(o map[string]interface{}) {
		o["valuePlaceholder"] = p
	}
} 

// 是否可拖拽排序
func InputKV_draggable(p bool) opt {
	return func(o map[string]interface{}) {
		o["draggable"] = p
	}
} 

// 默认值
func InputKV_defaultValue(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["defaultValue"] = p
	}
} 

// 添加组合项时触发
func InputKV_add(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["add"] = p
	}
} 