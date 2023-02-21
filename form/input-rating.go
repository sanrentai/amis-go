package form

// 评分
func InputRating(opts ...opt) map[string]interface{} {
	return newForm("input-rating", opts...)
}



// 当前值
func InputRating_value(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["value"] = p
	}
} 

// 是否使用半星选择
func InputRating_half(p bool) opt {
	return func(o map[string]interface{}) {
		o["half"] = p
	}
} 

// 总星数
func InputRating_count(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["count"] = p
	}
} 

// 只读
func InputRating_readOnly(p bool) opt {
	return func(o map[string]interface{}) {
		o["readOnly"] = p
	}
} 

// 是否允许再次点击后清除
func InputRating_allowClear(p bool) opt {
	return func(o map[string]interface{}) {
		o["allowClear"] = p
	}
} 

// 星星被选中的颜色。 若传入字符串，则只有一种颜色。若传入对象，可自定义分段，键名为分段的界限值，键值为对应的类名
func InputRating_colors(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["colors"] = p
	}
} 

// 未被选中的星星的颜色
func InputRating_inactiveColor(p string) opt {
	return func(o map[string]interface{}) {
		o["inactiveColor"] = p
	}
} 

// 星星被选中时的提示文字。可自定义分段，键名为分段的界限值，键值为对应的类名
func InputRating_texts(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["texts"] = p
	}
} 

// 文字的位置
func InputRating_textPosition(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["textPosition"] = p
	}
} 

// 自定义字符
func InputRating_char(p string) opt {
	return func(o map[string]interface{}) {
		o["char"] = p
	}
} 

// 自定义样式类名
func InputRating_className(p string) opt {
	return func(o map[string]interface{}) {
		o["className"] = p
	}
} 

// 自定义字符类名
func InputRating_charClassName(p string) opt {
	return func(o map[string]interface{}) {
		o["charClassName"] = p
	}
} 

// 自定义文字类名
func InputRating_textClassName(p string) opt {
	return func(o map[string]interface{}) {
		o["textClassName"] = p
	}
} 