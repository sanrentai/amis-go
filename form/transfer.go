package form

// 穿梭器
func Transfer(opts ...opt) map[string]interface{} {
	return newForm("transfer", opts...)
}



// [选项组](./options#%E9%9D%99%E6%80%81%E9%80%89%E9%A1%B9%E7%BB%84-options)
func Transfer_options(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["options"] = p
	}
} 

// [动态选项组](./options#%E5%8A%A8%E6%80%81%E9%80%89%E9%A1%B9%E7%BB%84-source)
func Transfer_source(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["source"] = p
	}
} 

// [拼接符](./options#%E6%8B%BC%E6%8E%A5%E7%AC%A6-delimiter)
func Transfer_delimeter(p string) opt {
	return func(o map[string]interface{}) {
		o["delimeter"] = p
	}
} 

// [拼接值](./options#%E6%8B%BC%E6%8E%A5%E5%80%BC-joinvalues)
func Transfer_joinValues(p bool) opt {
	return func(o map[string]interface{}) {
		o["joinValues"] = p
	}
} 

// [提取值](./options#%E6%8F%90%E5%8F%96%E5%A4%9A%E9%80%89%E5%80%BC-extractvalue)
func Transfer_extractValue(p bool) opt {
	return func(o map[string]interface{}) {
		o["extractValue"] = p
	}
} 

// 如果想通过接口检索，可以设置这个 api。
func Transfer_searchApi(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["searchApi"] = p
	}
} 

// 结果面板跟随模式，目前只支持`list`、`table`、`tree`（tree 目前只支持非延时加载的`tree`）
func Transfer_resultListModeFollowSelect(p bool) opt {
	return func(o map[string]interface{}) {
		o["resultListModeFollowSelect"] = p
	}
} 

// 是否显示统计数据
func Transfer_statistics(p bool) opt {
	return func(o map[string]interface{}) {
		o["statistics"] = p
	}
} 

// 左侧的标题文字
func Transfer_selectTitle(p string) opt {
	return func(o map[string]interface{}) {
		o["selectTitle"] = p
	}
} 

// 右侧结果的标题文字
func Transfer_resultTitle(p string) opt {
	return func(o map[string]interface{}) {
		o["resultTitle"] = p
	}
} 

// 结果可以进行拖拽排序（结果列表为树时，不支持排序）
func Transfer_sortable(p bool) opt {
	return func(o map[string]interface{}) {
		o["sortable"] = p
	}
} 

// 可选：`list`、`table`、`tree`、`chained`、`associated`。分别为：列表形式、表格形式、树形选择形式、级联选择形式，关联选择形式（与级联选择的区别在于，级联是无限极，而关联只有一级，关联左边可以是个 tree）。
func Transfer_selectMode(p string) opt {
	return func(o map[string]interface{}) {
		o["selectMode"] = p
	}
} 

// 如果不设置将采用 `selectMode` 的值，可以单独配置，参考 `selectMode`，决定搜索结果的展示形式。
func Transfer_searchResultMode(p string) opt {
	return func(o map[string]interface{}) {
		o["searchResultMode"] = p
	}
} 

// 左侧列表搜索功能，当设置为  true  时表示可以通过输入部分内容检索出选项项。
func Transfer_searchable(p bool) opt {
	return func(o map[string]interface{}) {
		o["searchable"] = p
	}
} 

// 左侧列表搜索框提示
func Transfer_searchPlaceholder(p string) opt {
	return func(o map[string]interface{}) {
		o["searchPlaceholder"] = p
	}
} 

// 当展示形式为 `table` 可以用来配置展示哪些列，跟 table 中的 columns 配置相似，只是只有展示功能。
func Transfer_columns(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["columns"] = p
	}
} 

// 当展示形式为 `associated` 时用来配置左边的选项集。
func Transfer_leftOptions(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["leftOptions"] = p
	}
} 

// 当展示形式为 `associated` 时用来配置左边的选择形式，支持 `list` 或者 `tree`。默认为 `list`。
func Transfer_leftMode(p string) opt {
	return func(o map[string]interface{}) {
		o["leftMode"] = p
	}
} 

// 当展示形式为 `associated` 时用来配置右边的选择形式，可选：`list`、`table`、`tree`、`chained`。
func Transfer_rightMode(p string) opt {
	return func(o map[string]interface{}) {
		o["rightMode"] = p
	}
} 

// 结果（右则）列表的检索功能，当设置为 true 时，可以通过输入检索模糊匹配检索内容（目前树的延时加载不支持结果搜索功能）
func Transfer_resultSearchable(p bool) opt {
	return func(o map[string]interface{}) {
		o["resultSearchable"] = p
	}
} 

// 右侧列表搜索框提示
func Transfer_resultSearchPlaceholder(p string) opt {
	return func(o map[string]interface{}) {
		o["resultSearchPlaceholder"] = p
	}
} 

// 每个选项的高度，用于虚拟渲染
func Transfer_itemHeight(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["itemHeight"] = p
	}
} 

// 在选项数量超过多少时开启虚拟渲染
func Transfer_virtualThreshold(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["virtualThreshold"] = p
	}
} 

// 更新数据，开启`multiple`支持设置多项，开启`joinValues`时，多值用`,`分隔，否则多值用数组
func Transfer_setValue(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["setValue"] = p
	}
} 