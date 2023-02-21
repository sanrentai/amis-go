package form

// 对比编辑器
func DiffEditor(opts ...opt) map[string]interface{} {
	return newForm("diff-editor", opts...)
}



// 编辑器高亮的语言，可选 [支持的语言](./editor#%E6%94%AF%E6%8C%81%E7%9A%84%E8%AF%AD%E8%A8%80)
func DiffEditor_language(p string) opt {
	return func(o map[string]interface{}) {
		o["language"] = p
	}
} 

// 左侧值
func DiffEditor_diffValue(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["diffValue"] = p
	}
} 