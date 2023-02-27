package amis

// 增删改查
func CRUD(opts ...opt) map[string]interface{} {
	return newCompent("crud", opts...)
}

// `"table" 、 "cards" 或者 "list"`
func CRUD_mode(p string) opt {
	return func(o map[string]interface{}) {
		o["mode"] = p
	}
}

// 可设置成空，当设置成空时，没有标题栏
func CRUD_title(p string) opt {
	return func(o map[string]interface{}) {
		o["title"] = p
	}
}

// 表格外层 Dom 的类名
func CRUD_className(p string) opt {
	return func(o map[string]interface{}) {
		o["className"] = p
	}
}

// CRUD 用来获取列表数据的 api。
func CRUD_api(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["api"] = p
	}
}

// 是否一次性加载所有数据（前端分页）
func CRUD_loadDataOnce(p bool) opt {
	return func(o map[string]interface{}) {
		o["loadDataOnce"] = p
	}
}

// 在开启 loadDataOnce 时，filter 时是否去重新请求 api
func CRUD_loadDataOnceFetchOnFilter(p bool) opt {
	return func(o map[string]interface{}) {
		o["loadDataOnceFetchOnFilter"] = p
	}
}

// 数据映射接口返回某字段的值，不设置会默认使用接口返回的`${items}`或者`${rows}`，也可以设置成上层数据源的内容
func CRUD_source(p string) opt {
	return func(o map[string]interface{}) {
		o["source"] = p
	}
}

// 设置过滤器，当该表单提交后，会把数据带给当前 `mode` 刷新列表。
func CRUD_filter(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["filter"] = p
	}
}

// 是否可显隐过滤器
func CRUD_filterTogglable(p bool) opt {
	return func(o map[string]interface{}) {
		o["filterTogglable"] = p
	}
}

// 设置过滤器默认是否可见。
func CRUD_filterDefaultVisible(p bool) opt {
	return func(o map[string]interface{}) {
		o["filterDefaultVisible"] = p
	}
}

// 是否初始化的时候拉取数据, 只针对有 filter 的情况, 没有 filter 初始都会拉取数据
func CRUD_initFetch(p bool) opt {
	return func(o map[string]interface{}) {
		o["initFetch"] = p
	}
}

// 刷新时间(最低 1000)
func CRUD_interval(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["interval"] = p
	}
}

// 配置刷新时是否隐藏加载动画
func CRUD_silentPolling(p bool) opt {
	return func(o map[string]interface{}) {
		o["silentPolling"] = p
	}
}

// 通过[表达式](../../docs/concepts/expression)来配置停止刷新的条件
func CRUD_stopAutoRefreshWhen(p string) opt {
	return func(o map[string]interface{}) {
		o["stopAutoRefreshWhen"] = p
	}
}

// 当有弹框时关闭自动刷新，关闭弹框又恢复
func CRUD_stopAutoRefreshWhenModalIsOpen(p bool) opt {
	return func(o map[string]interface{}) {
		o["stopAutoRefreshWhenModalIsOpen"] = p
	}
}

// 是否将过滤条件的参数同步到地址栏
func CRUD_syncLocation(p bool) opt {
	return func(o map[string]interface{}) {
		o["syncLocation"] = p
	}
}

// // 是否可通过拖拽排序
// func CRUD_draggable(p bool) opt {
// 	return func(o map[string]interface{}) {
// 		o["draggable"] = p
// 	}
// }

// 是否可以调整列宽度
func CRUD_resizable(p bool) opt {
	return func(o map[string]interface{}) {
		o["resizable"] = p
	}
}

// 用[表达式](../../docs/concepts/expression)来配置是否可拖拽排序
func CRUD_itemDraggableOn(p bool) opt {
	return func(o map[string]interface{}) {
		o["itemDraggableOn"] = p
	}
}

// // 保存排序的 api。
// func CRUD_[saveOrderApi](#saveOrderApi)(p interface{}) opt {
// 	return func(o map[string]interface{}) {
// 		o["[saveOrderApi](#saveOrderApi)"] = p
// 	}
// }

// // 快速编辑后用来批量保存的 API。
// func CRUD_[quickSaveApi](#quickSaveApi)(p interface{}) opt {
// 	return func(o map[string]interface{}) {
// 		o["[quickSaveApi](#quickSaveApi)"] = p
// 	}
// }

// // 快速编辑配置成及时保存时使用的 API。
// func CRUD_[quickSaveItemApi](#quickSaveItemApi)(p interface{}) opt {
// 	return func(o map[string]interface{}) {
// 		o["[quickSaveItemApi](#quickSaveItemApi)"] = p
// 	}
// }

// 批量操作列表，配置后，表格可进行选中操作。
func CRUD_bulkActions(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["bulkActions"] = p
	}
}

// 覆盖消息提示，如果不指定，将采用 api 返回的 message
func CRUD_messages(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["messages"] = p
	}
}

// // 获取失败时提示
// func CRUD_messages.fetchFailed(p string) opt {
// 	return func(o map[string]interface{}) {
// 		o["messages.fetchFailed"] = p
// 	}
// }

// // 保存顺序失败提示
// func CRUD_messages.saveOrderFailed(p string) opt {
// 	return func(o map[string]interface{}) {
// 		o["messages.saveOrderFailed"] = p
// 	}
// }

// // 保存顺序成功提示
// func CRUD_messages.saveOrderSuccess(p string) opt {
// 	return func(o map[string]interface{}) {
// 		o["messages.saveOrderSuccess"] = p
// 	}
// }

// // 快速保存失败提示
// func CRUD_messages.quickSaveFailed(p string) opt {
// 	return func(o map[string]interface{}) {
// 		o["messages.quickSaveFailed"] = p
// 	}
// }

// // 快速保存成功提示
// func CRUD_messages.quickSaveSuccess(p string) opt {
// 	return func(o map[string]interface{}) {
// 		o["messages.quickSaveSuccess"] = p
// 	}
// }

// 设置 ID 字段名。
func CRUD_primaryField(p string) opt {
	return func(o map[string]interface{}) {
		o["primaryField"] = p
	}
}

// 设置一页显示多少条数据。
func CRUD_perPage(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["perPage"] = p
	}
}

// 默认排序字段，这个是传给后端，需要后端接口实现
func CRUD_orderBy(p string) opt {
	return func(o map[string]interface{}) {
		o["orderBy"] = p
	}
}

// 设置默认 filter 默认参数，会在查询的时候一起发给后端
func CRUD_defaultParams(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["defaultParams"] = p
	}
}

// 设置分页页码字段名。
func CRUD_pageField(p string) opt {
	return func(o map[string]interface{}) {
		o["pageField"] = p
	}
}

// 设置分页一页显示的多少条数据的字段名。注意：最好与 defaultParams 一起使用，请看下面例子。
func CRUD_perPageField(p string) opt {
	return func(o map[string]interface{}) {
		o["perPageField"] = p
	}
}

// 设置一页显示多少条数据下拉框可选条数。
func CRUD_perPageAvailable(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["perPageAvailable"] = p
	}
}

// 设置用来确定位置的字段名，设置后新的顺序将被赋值到该字段中。
func CRUD_orderField(p string) opt {
	return func(o map[string]interface{}) {
		o["orderField"] = p
	}
}

// 隐藏顶部快速保存提示
func CRUD_hideQuickSaveBtn(p bool) opt {
	return func(o map[string]interface{}) {
		o["hideQuickSaveBtn"] = p
	}
}

// 当切分页的时候，是否自动跳顶部。
func CRUD_autoJumpToTopOnPagerChange(p bool) opt {
	return func(o map[string]interface{}) {
		o["autoJumpToTopOnPagerChange"] = p
	}
}

// 将返回数据同步到过滤器上。
func CRUD_syncResponse2Query(p bool) opt {
	return func(o map[string]interface{}) {
		o["syncResponse2Query"] = p
	}
}

// 保留条目选择，默认分页、搜素后，用户选择条目会被清空，开启此选项后会保留用户选择，可以实现跨页面批量操作。
func CRUD_keepItemSelectionOnPageChange(p bool) opt {
	return func(o map[string]interface{}) {
		o["keepItemSelectionOnPageChange"] = p
	}
}

// 单条描述模板，`keepItemSelectionOnPageChange`设置为`true`后会把所有已选择条目列出来，此选项可以用来定制条目展示文案。
func CRUD_labelTpl(p string) opt {
	return func(o map[string]interface{}) {
		o["labelTpl"] = p
	}
}

// 顶部工具栏配置
func CRUD_headerToolbar(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["headerToolbar"] = p
	}
}

// 底部工具栏配置
func CRUD_footerToolbar(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["footerToolbar"] = p
	}
}

// 是否总是显示分页
func CRUD_alwaysShowPagination(p bool) opt {
	return func(o map[string]interface{}) {
		o["alwaysShowPagination"] = p
	}
}

// 是否固定表头(table 下)
func CRUD_affixHeader(p bool) opt {
	return func(o map[string]interface{}) {
		o["affixHeader"] = p
	}
}

// 是否开启查询区域，开启后会根据列元素的 `searchable` 属性值，自动生成查询条件表单
func CRUD_autoGenerateFilter(p bool) opt {
	return func(o map[string]interface{}) {
		o["autoGenerateFilter"] = p
	}
}

// 单条数据 ajax 操作后是否重置页码为第一页
func CRUD_resetPageAfterAjaxItemAction(p bool) opt {
	return func(o map[string]interface{}) {
		o["resetPageAfterAjaxItemAction"] = p
	}
}

// 内容区域自适应高度
func CRUD_autoFillHeight(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["autoFillHeight"] = p
	}
}

// 是否可排序
func CRUD_sortable(p bool) opt {
	return func(o map[string]interface{}) {
		o["sortable"] = p
	}
}

// 按钮文字
func CRUD_label(p string) opt {
	return func(o map[string]interface{}) {
		o["label"] = p
	}
}

// 按钮提示文字
func CRUD_tooltip(p string) opt {
	return func(o map[string]interface{}) {
		o["tooltip"] = p
	}
}

// 按钮禁用状态下的提示
func CRUD_disabledTip(p string) opt {
	return func(o map[string]interface{}) {
		o["disabledTip"] = p
	}
}

// 按钮样式，参考[按钮](./action)
func CRUD_level(p string) opt {
	return func(o map[string]interface{}) {
		o["level"] = p
	}
}

// 是否可通过拖拽排序
func CRUD_draggable(p bool) opt {
	return func(o map[string]interface{}) {
		o["draggable"] = p
	}
}

// 默认是否展开
func CRUD_defaultIsOpened(p bool) opt {
	return func(o map[string]interface{}) {
		o["defaultIsOpened"] = p
	}
}

// 是否隐藏展开的图标
func CRUD_hideExpandIcon(p bool) opt {
	return func(o map[string]interface{}) {
		o["hideExpandIcon"] = p
	}
}

// 是否显示遮罩层
func CRUD_overlay(p bool) opt {
	return func(o map[string]interface{}) {
		o["overlay"] = p
	}
}

// 点击外部是否关闭
func CRUD_closeOnOutside(p bool) opt {
	return func(o map[string]interface{}) {
		o["closeOnOutside"] = p
	}
}

// 点击内容是否关闭
func CRUD_closeOnClick(p bool) opt {
	return func(o map[string]interface{}) {
		o["closeOnClick"] = p
	}
}

// 是否只显示图标。
func CRUD_iconOnly(p bool) opt {
	return func(o map[string]interface{}) {
		o["iconOnly"] = p
	}
}

// 按钮的图标
func CRUD_icon(p string) opt {
	return func(o map[string]interface{}) {
		o["icon"] = p
	}
}

// 按钮的 CSS 类名
func CRUD_btnClassName(p string) opt {
	return func(o map[string]interface{}) {
		o["btnClassName"] = p
	}
}

func CRUD_quickSaveItemApi(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["quickSaveItemApi"] = p
	}
}
