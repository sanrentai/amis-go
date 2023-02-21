package form

// 滑块
func InputRange(opts ...opt) map[string]interface{} {
	return newForm("input-range", opts...)
}



// css 类名
func InputRange_className(p string) opt {
	return func(o map[string]interface{}) {
		o["className"] = p
	}
} 

// 
func InputRange_value(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["value"] = p
	}
} 

// 最小值
func InputRange_min(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["min"] = p
	}
} 

// 最大值
func InputRange_max(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["max"] = p
	}
} 

// 是否禁用
func InputRange_disabled(p bool) opt {
	return func(o map[string]interface{}) {
		o["disabled"] = p
	}
} 

// 步长
func InputRange_step(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["step"] = p
	}
} 

// 是否显示步长
func InputRange_showSteps(p bool) opt {
	return func(o map[string]interface{}) {
		o["showSteps"] = p
	}
} 

// 分割的块数<br/>主持数组传入分块的节点
func InputRange_parts(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["parts"] = p
	}
} 

// 刻度标记<br/>- 支持自定义样式<br/>- 设置百分比
func InputRange_marks(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["marks"] = p
	}
} 

// 是否显示滑块标签
func InputRange_tooltipVisible(p bool) opt {
	return func(o map[string]interface{}) {
		o["tooltipVisible"] = p
	}
} 

// 滑块标签的位置，默认`auto`，方向自适应<br/>前置条件：tooltipVisible 不为 false 时有效
func InputRange_tooltipPlacement(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["tooltipPlacement"] = p
	}
} 

// 控制滑块标签显隐函数<br/>前置条件：tooltipVisible 不为 false 时有效
func InputRange_tipFormatter(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["tipFormatter"] = p
	}
} 

// 支持选择范围
func InputRange_multiple(p bool) opt {
	return func(o map[string]interface{}) {
		o["multiple"] = p
	}
} 

// 默认为 `true`，选择的 `value` 会通过 `delimiter` 连接起来，否则直接将以`{min: 1, max: 100}`的形式提交<br/>前置条件：开启`multiple`时有效
func InputRange_joinValues(p bool) opt {
	return func(o map[string]interface{}) {
		o["joinValues"] = p
	}
} 

// 分隔符
func InputRange_delimiter(p string) opt {
	return func(o map[string]interface{}) {
		o["delimiter"] = p
	}
} 

// 单位
func InputRange_unit(p string) opt {
	return func(o map[string]interface{}) {
		o["unit"] = p
	}
} 

// 是否可清除<br/>前置条件：开启`showInput`时有效
func InputRange_clearable(p bool) opt {
	return func(o map[string]interface{}) {
		o["clearable"] = p
	}
} 

// 是否显示输入框
func InputRange_showInput(p bool) opt {
	return func(o map[string]interface{}) {
		o["showInput"] = p
	}
} 

// 当 组件 的值发生改变时，会触发 onChange 事件，并把改变后的值作为参数传入
func InputRange_onChange(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["onChange"] = p
	}
} 

// 与 `onmouseup` 触发时机一致，把当前值作为参数传入
func InputRange_onAfterChange(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["onAfterChange"] = p
	}
} 