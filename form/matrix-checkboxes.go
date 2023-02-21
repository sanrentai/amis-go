package form

// 矩阵
func MatrixCheckboxes(opts ...opt) map[string]interface{} {
	return newForm("matrix-checkboxes", opts...)
}



// 列信息，数组中 `label` 字段是必须给出的
func MatrixCheckboxes_columns(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["columns"] = p
	}
} 

// 行信息， 数组中 `label` 字段是必须给出的
func MatrixCheckboxes_rows(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["rows"] = p
	}
} 

// 行标题说明
func MatrixCheckboxes_rowLabel(p string) opt {
	return func(o map[string]interface{}) {
		o["rowLabel"] = p
	}
} 

// Api 地址，如果选项组不固定，可以通过配置 `source` 动态拉取。
func MatrixCheckboxes_source(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["source"] = p
	}
} 

// 是否多选
func MatrixCheckboxes_multiple(p bool) opt {
	return func(o map[string]interface{}) {
		o["multiple"] = p
	}
} 

// 设置单选模式，`multiple`为`false`时有效，可设置为`cell`, `row`, `column` 分别为全部选项中只能单选某个单元格、每行只能单选某个单元格，每列只能单选某个单元格
func MatrixCheckboxes_singleSelectMode(p string) opt {
	return func(o map[string]interface{}) {
		o["singleSelectMode"] = p
	}
} 