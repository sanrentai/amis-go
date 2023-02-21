package form

// 编辑器
func Editor(opts ...opt) map[string]interface{} {
	return newForm("editor", opts...)
}



// 编辑器高亮的语言，支持通过 `${xxx}` 变量获取
func Editor_language(p string) opt {
	return func(o map[string]interface{}) {
		o["language"] = p
	}
} 

// 编辑器高度，取值可以是 `md`、`lg`、`xl`、`xxl`
func Editor_size(p string) opt {
	return func(o map[string]interface{}) {
		o["size"] = p
	}
} 

// 是否显示全屏模式开关
func Editor_allowFullscreen(p bool) opt {
	return func(o map[string]interface{}) {
		o["allowFullscreen"] = p
	}
} 

// monaco 编辑器的其它配置，比如是否显示行号等，请参考[这里](https://microsoft.github.io/monaco-editor/api/enums/monaco.editor.EditorOption.html)，不过无法设置 readOnly，只读模式需要使用 `disabled: true`
func Editor_options(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["options"] = p
	}
} 

// 占位描述，没有值的时候展示
func Editor_placeholder(p string) opt {
	return func(o map[string]interface{}) {
		o["placeholder"] = p
	}
} 