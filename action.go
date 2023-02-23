package amis

// 行为按钮
func Action(opts ...opt) map[string]interface{} {
	return newCompent("action", opts...)
}

// 【必填】这是 action 最核心的配置，来指定该 action 的作用类型，支持：`ajax`、`link`、`url`、`drawer`、`dialog`、`confirm`、`cancel`、`prev`、`next`、`copy`、`close`。
func Action_actionType(p string) opt {
	return func(o map[string]interface{}) {
		o["actionType"] = p
	}
}

// 按钮文本。可用 `${xxx}` 取值。
func Action_label(p string) opt {
	return func(o map[string]interface{}) {
		o["label"] = p
	}
}

// 按钮样式，支持：`link`、`primary`、`secondary`、`info`、`success`、`warning`、`danger`、`light`、`dark`、`default`。
func Action_level(p string) opt {
	return func(o map[string]interface{}) {
		o["level"] = p
	}
}

// 按钮大小，支持：`xs`、`sm`、`md`、`lg`。
func Action_size(p string) opt {
	return func(o map[string]interface{}) {
		o["size"] = p
	}
}

// 设置图标，例如`fa fa-plus`。
func Action_icon(p string) opt {
	return func(o map[string]interface{}) {
		o["icon"] = p
	}
}

// 给图标上添加类名。
func Action_iconClassName(p string) opt {
	return func(o map[string]interface{}) {
		o["iconClassName"] = p
	}
}

// 在按钮文本右侧设置图标，例如`fa fa-plus`。
func Action_rightIcon(p string) opt {
	return func(o map[string]interface{}) {
		o["rightIcon"] = p
	}
}

// 给右侧图标上添加类名。
func Action_rightIconClassName(p string) opt {
	return func(o map[string]interface{}) {
		o["rightIconClassName"] = p
	}
}

// 按钮是否高亮。
func Action_active(p bool) opt {
	return func(o map[string]interface{}) {
		o["active"] = p
	}
}

// 按钮高亮时的样式，配置支持同`level`。
func Action_activeLevel(p string) opt {
	return func(o map[string]interface{}) {
		o["activeLevel"] = p
	}
}

// 给按钮高亮添加类名。
func Action_activeClassName(p string) opt {
	return func(o map[string]interface{}) {
		o["activeClassName"] = p
	}
}

// 用`display:"block"`来显示按钮。
func Action_block(p bool) opt {
	return func(o map[string]interface{}) {
		o["block"] = p
	}
}

// 当设置后，操作在开始前会询问用户。可用 `${xxx}` 取值。
func Action_confirmText(p string) opt {
	return func(o map[string]interface{}) {
		o["confirmText"] = p
	}
}

// 指定此次操作完后，需要刷新的目标组件名字（组件的`name`值，自己配置的），多个请用 `,` 号隔开。
func Action_reload(p string) opt {
	return func(o map[string]interface{}) {
		o["reload"] = p
	}
}

// 鼠标停留时弹出该段文字，也可以配置对象类型：字段为`title`和`content`。可用 `${xxx}` 取值。                                                                                                                                     |
func Action_tooltip(p string) opt {
	return func(o map[string]interface{}) {
		o["tooltip"] = p
	}
}

// 被禁用后鼠标停留时弹出该段文字，也可以配置对象类型：字段为`title`和`content`。可用 `${xxx}` 取值。
func Action_disabledTip(p string) opt {
	return func(o map[string]interface{}) {
		o["disabledTip"] = p
	}
}

// 如果配置了`tooltip`或者`disabledTip`，指定提示信息位置，可配置`top`、`bottom`、`left`、`right`。
func Action_tooltipPlacement(p string) opt {
	return func(o map[string]interface{}) {
		o["tooltipPlacement"] = p
	}
}

// 当`action`配置在`dialog`或`drawer`的`actions`中时，配置为`true`指定此次操作完后关闭当前`dialog`或`drawer`。当值为字符串，并且是祖先层弹框的名字的时候，会把祖先弹框关闭掉。 |
func Action_close(p interface{}) opt {
	return func(o map[string]interface{}) {
		o[""] = p
	}
}

// 配置字符串数组，指定在`form`中进行操作之前，需要指定的字段名的表单项通过验证
func Action_required(p ...string) opt {
	return func(o map[string]interface{}) {
		o["required"] = p
	}
}
