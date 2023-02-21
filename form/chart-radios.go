package form

// 图表单选框
func ChartRadios(opts ...opt) map[string]interface{} {
	return newForm("chart-radios", opts...)
}



// echart 图表配置
func ChartRadios_config(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["config"] = p
	}
} 

// 高亮的时候是否显示 tooltip
func ChartRadios_showTooltipOnHighlight(p bool) opt {
	return func(o map[string]interface{}) {
		o["showTooltipOnHighlight"] = p
	}
} 

// 图表数值字段名
func ChartRadios_chartValueField(p string) opt {
	return func(o map[string]interface{}) {
		o["chartValueField"] = p
	}
} 