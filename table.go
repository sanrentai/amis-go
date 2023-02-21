package amis

// 表格
func Table(opts ...opt) map[string]interface{} {
	return newCompent("table", opts...)
}

// 标题
func Table_title(p string) opt {
	return func(o map[string]interface{}) {
		o["title"] = p
	}
}

// 数据源, 绑定当前环境变量
func Table_source(p string) opt {
	return func(o map[string]interface{}) {
		o["source"] = p
	}
}

// 是否固定表头
func Table_affixHeader(p bool) opt {
	return func(o map[string]interface{}) {
		o["affixHeader"] = p
	}
}

// 展示列显示开关, 自动即：列数量大于或等于 5 个时自动开启
func Table_columnsTogglable(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["columnsTogglable"] = p
	}
}

// 当没数据的时候的文字提示
func Table_placeholder(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["placeholder"] = p
	}
}

// 外层 CSS 类名
func Table_className(p string) opt {
	return func(o map[string]interface{}) {
		o["className"] = p
	}
}

// 表格 CSS 类名
func Table_tableClassName(p string) opt {
	return func(o map[string]interface{}) {
		o["tableClassName"] = p
	}
}

// 顶部外层 CSS 类名
func Table_headerClassName(p string) opt {
	return func(o map[string]interface{}) {
		o["headerClassName"] = p
	}
}

// 底部外层 CSS 类名
func Table_footerClassName(p string) opt {
	return func(o map[string]interface{}) {
		o["footerClassName"] = p
	}
}

// 工具栏 CSS 类名
func Table_toolbarClassName(p string) opt {
	return func(o map[string]interface{}) {
		o["toolbarClassName"] = p
	}
}

// 用来设置列信息
func Table_columns(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["columns"] = p
	}
}

// 自动合并单元格
func Table_combineNum(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["combineNum"] = p
	}
}

// 悬浮行操作按钮组
func Table_itemActions(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["itemActions"] = p
	}
}

// 配置当前行是否可勾选的条件，要用 [表达式](../../docs/concepts/expression)
func Table_itemCheckableOn(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["itemCheckableOn"] = p
	}
}

// 配置当前行是否可拖拽的条件，要用 [表达式](../../docs/concepts/expression)
func Table_itemDraggableOn(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["itemDraggableOn"] = p
	}
}

// 点击数据行是否可以勾选当前行
func Table_checkOnItemClick(p bool) opt {
	return func(o map[string]interface{}) {
		o["checkOnItemClick"] = p
	}
}

// 给行添加 CSS 类名
func Table_rowClassName(p string) opt {
	return func(o map[string]interface{}) {
		o["rowClassName"] = p
	}
}

// 通过模板给行添加 CSS 类名
func Table_rowClassNameExpr(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["rowClassNameExpr"] = p
	}
}

// 顶部总结行
func Table_prefixRow(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["prefixRow"] = p
	}
}

// 底部总结行
func Table_affixRow(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["affixRow"] = p
	}
}

// 行角标配置
func Table_itemBadge(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["itemBadge"] = p
	}
}

// 内容区域自适应高度
func Table_autoFillHeight(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["autoFillHeight"] = p
	}
}

// 列宽度是否支持调整
func Table_resizable(p bool) opt {
	return func(o map[string]interface{}) {
		o["resizable"] = p
	}
}

// 支持勾选
func Table_selectable(p bool) opt {
	return func(o map[string]interface{}) {
		o["selectable"] = p
	}
}

// 勾选 icon 是否为多选样式`checkbox`， 默认为`radio`
func Table_multiple(p bool) opt {
	return func(o map[string]interface{}) {
		o["multiple"] = p
	}
}

// 说明
func Table_属性名(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["属性名"] = p
	}
}

// 表头文本内容
func Table_label(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["label"] = p
	}
}

// 通过名称关联数据
func Table_name(p string) opt {
	return func(o map[string]interface{}) {
		o["name"] = p
	}
}

// 提示信息
func Table_remark(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["remark"] = p
	}
}

// 弹出框
func Table_popOver(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["popOver"] = p
	}
}

// 是否可复制
func Table_copyable(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["copyable"] = p
	}
}
