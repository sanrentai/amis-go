package amis

// 表格
func Table2(opts ...opt) map[string]interface{} {
	return newCompent("table2", opts...)
}

// 标题
func Table2_title(p string) opt {
	return func(o map[string]interface{}) {
		o["title"] = p
	}
}

// 数据源, 绑定当前环境变量
func Table2_source(p string) opt {
	return func(o map[string]interface{}) {
		o["source"] = p
	}
}

// 是否粘性头部
func Table2_sticky(p bool) opt {
	return func(o map[string]interface{}) {
		o["sticky"] = p
	}
}

// 表格是否加载中
func Table2_loading(p bool) opt {
	return func(o map[string]interface{}) {
		o["loading"] = p
	}
}

// 展示列显示开关, 自动即：列数量大于或等于 5 个时自动开启
func Table2_columnsTogglable(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["columnsTogglable"] = p
	}
}

// 行相关配置
func Table2_rowSelection(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["rowSelection"] = p
	}
}

// 展开行配置
func Table2_expandable(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["expandable"] = p
	}
}

// 底部外层 CSS 类名
func Table2_footerClassName(p string) opt {
	return func(o map[string]interface{}) {
		o["footerClassName"] = p
	}
}

// 工具栏 CSS 类名
func Table2_toolbarClassName(p string) opt {
	return func(o map[string]interface{}) {
		o["toolbarClassName"] = p
	}
}

// 用来设置列信息
func Table2_columns(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["columns"] = p
	}
}

// 自动合并单元格
func Table2_combineNum(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["combineNum"] = p
	}
}

// 悬浮行操作按钮组
func Table2_itemActions(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["itemActions"] = p
	}
}

// 配置当前行是否可勾选的条件，要用 [表达式](../../docs/concepts/expression)
func Table2_itemCheckableOn(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["itemCheckableOn"] = p
	}
}

// 配置当前行是否可拖拽的条件，要用 [表达式](../../docs/concepts/expression)
func Table2_itemDraggableOn(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["itemDraggableOn"] = p
	}
}

// 点击数据行是否可以勾选当前行
func Table2_checkOnItemClick(p bool) opt {
	return func(o map[string]interface{}) {
		o["checkOnItemClick"] = p
	}
}

// 给行添加 CSS 类名
func Table2_rowClassName(p string) opt {
	return func(o map[string]interface{}) {
		o["rowClassName"] = p
	}
}

// 通过模板给行添加 CSS 类名
func Table2_rowClassNameExpr(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["rowClassNameExpr"] = p
	}
}

// 顶部总结行
func Table2_prefixRow(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["prefixRow"] = p
	}
}

// 底部总结行
func Table2_affixRow(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["affixRow"] = p
	}
}

// 行角标配置
func Table2_itemBadge(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["itemBadge"] = p
	}
}

// 内容区域自适应高度
func Table2_autoFillHeight(p bool) opt {
	return func(o map[string]interface{}) {
		o["autoFillHeight"] = p
	}
}

// 选择列是否固定，只能左侧固定
func Table2_fixed(p bool) opt {
	return func(o map[string]interface{}) {
		o["fixed"] = p
	}
}

// 当前行是否可选择条件，要用 [表达式](../../docs/concepts/expression)
func Table2_disableOn(p string) opt {
	return func(o map[string]interface{}) {
		o["disableOn"] = p
	}
}

// 自定义筛选菜单，内置`all`（全选）、`invert`（反选）、`none`（取消选择）、`odd`（选择奇数项）、`even`（选择偶数项）
func Table2_selections(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["selections"] = p
	}
}

// 已选择项正则表达式
func Table2_selectedRowKeysExpr(p string) opt {
	return func(o map[string]interface{}) {
		o["selectedRowKeysExpr"] = p
	}
}

// 自定义选择列列宽
func Table2_columnWidth(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["columnWidth"] = p
	}
}

// 单条任意区域选中
func Table2_rowClick(p bool) opt {
	return func(o map[string]interface{}) {
		o["rowClick"] = p
	}
}

// 自定义菜单项文本
func Table2_text(p string) opt {
	return func(o map[string]interface{}) {
		o["text"] = p
	}
}

// 说明
func Table2_属性名(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["属性名"] = p
	}
}

// 指定可展开的行，要用 [表达式](../../docs/concepts/expression)
func Table2_expandableOn(p string) opt {
	return func(o map[string]interface{}) {
		o["expandableOn"] = p
	}
}

// 对应数据源的 key 值，默认是`key`，可指定为`id`、`shortId`等
func Table2_keyField(p string) opt {
	return func(o map[string]interface{}) {
		o["keyField"] = p
	}
}

// 表头文本内容
func Table2_label(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["label"] = p
	}
}

// 通过名称关联数据
func Table2_name(p string) opt {
	return func(o map[string]interface{}) {
		o["name"] = p
	}
}

// 弹出框
func Table2_popOver(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["popOver"] = p
	}
}

// 快速编辑
func Table2_quickEdit(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["quickEdit"] = p
	}
}

// 是否可复制
func Table2_copyable(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["copyable"] = p
	}
}

// 是否可排序
func Table2_sortable(p bool) opt {
	return func(o map[string]interface{}) {
		o["sortable"] = p
	}
}

// 列宽
func Table2_width(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["width"] = p
	}
}

// 提示信息
func Table2_remark(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["remark"] = p
	}
}
