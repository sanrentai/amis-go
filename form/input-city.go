package form

// 城市选择器
func InputCity(opts ...opt) map[string]interface{} {
	return newForm("input-city", opts...)
}



// 允许选择城市
func InputCity_allowCity(p bool) opt {
	return func(o map[string]interface{}) {
		o["allowCity"] = p
	}
} 

// 允许选择区域
func InputCity_allowDistrict(p bool) opt {
	return func(o map[string]interface{}) {
		o["allowDistrict"] = p
	}
} 

// 是否出搜索框
func InputCity_searchable(p bool) opt {
	return func(o map[string]interface{}) {
		o["searchable"] = p
	}
} 

// 默认 `true` 是否抽取值，如果设置成 `false` 值格式会变成对象，包含 `code`、`province`、`city` 和 `district` 文字信息。
func InputCity_extractValue(p bool) opt {
	return func(o map[string]interface{}) {
		o["extractValue"] = p
	}
} 

// 选中值变化时触发
func InputCity_change(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["change"] = p
	}
} 

// 更新数据
func InputCity_setValue(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["setValue"] = p
	}
} 