package amis

// 卡片
func Card(opts ...opt) map[string]interface{} {
	return newCompent("card", opts...)
}

// 外层 Dom 的类名
func Card_className(p string) opt {
	return func(o map[string]interface{}) {
		o["className"] = p
	}
}

// 外部链接
func Card_href(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["href"] = p
	}
}

// Card 头部内容设置
func Card_header(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["header"] = p
	}
}

// // 头部类名
// func Card_header.className(p string) opt {
// 	return func(o map[string]interface{}) {
// 		o["header.className"] = p
// 	}
// }

// // 标题
// func Card_header.title(p interface{}) opt {
// 	return func(o map[string]interface{}) {
// 		o["header.title"] = p
// 	}
// }

// // 标题类名
// func Card_header.titleClassName(p string) opt {
// 	return func(o map[string]interface{}) {
// 		o["header.titleClassName"] = p
// 	}
// }

// // 副标题
// func Card_header.subTitle(p interface{}) opt {
// 	return func(o map[string]interface{}) {
// 		o["header.subTitle"] = p
// 	}
// }

// // 副标题类名
// func Card_header.subTitleClassName(p string) opt {
// 	return func(o map[string]interface{}) {
// 		o["header.subTitleClassName"] = p
// 	}
// }

// // 副标题占位
// func Card_header.subTitlePlaceholder(p string) opt {
// 	return func(o map[string]interface{}) {
// 		o["header.subTitlePlaceholder"] = p
// 	}
// }

// // 描述
// func Card_header.description(p interface{}) opt {
// 	return func(o map[string]interface{}) {
// 		o["header.description"] = p
// 	}
// }

// // 描述类名
// func Card_header.descriptionClassName(p string) opt {
// 	return func(o map[string]interface{}) {
// 		o["header.descriptionClassName"] = p
// 	}
// }

// // 描述占位
// func Card_header.descriptionPlaceholder(p string) opt {
// 	return func(o map[string]interface{}) {
// 		o["header.descriptionPlaceholder"] = p
// 	}
// }

// // 图片
// func Card_header.avatar(p interface{}) opt {
// 	return func(o map[string]interface{}) {
// 		o["header.avatar"] = p
// 	}
// }

// // 图片包括层类名
// func Card_header.avatarClassName(p string) opt {
// 	return func(o map[string]interface{}) {
// 		o["header.avatarClassName"] = p
// 	}
// }

// // 图片类名
// func Card_header.imageClassName(p string) opt {
// 	return func(o map[string]interface{}) {
// 		o["header.imageClassName"] = p
// 	}
// }

// // 如果不配置图片，则会在图片处显示该文本
// func Card_header.avatarText(p interface{}) opt {
// 	return func(o map[string]interface{}) {
// 		o["header.avatarText"] = p
// 	}
// }

// // 设置文本背景色，它会根据数据分配一个颜色
// func Card_header.avatarTextBackground(p interface{}) opt {
// 	return func(o map[string]interface{}) {
// 		o["header.avatarTextBackground"] = p
// 	}
// }

// // 图片文本类名
// func Card_header.avatarTextClassName(p string) opt {
// 	return func(o map[string]interface{}) {
// 		o["header.avatarTextClassName"] = p
// 	}
// }

// // 是否显示激活样式
// func Card_header.highlight(p bool) opt {
// 	return func(o map[string]interface{}) {
// 		o["header.highlight"] = p
// 	}
// }

// // 激活样式类名
// func Card_header.highlightClassName(p string) opt {
// 	return func(o map[string]interface{}) {
// 		o["header.highlightClassName"] = p
// 	}
// }

// // 点击卡片跳转的链接地址
// func Card_header.href(p interface{}) opt {
// 	return func(o map[string]interface{}) {
// 		o["header.href"] = p
// 	}
// }

// // 是否新窗口打开
// func Card_header.blank(p bool) opt {
// 	return func(o map[string]interface{}) {
// 		o["header.blank"] = p
// 	}
// }

// 内容容器，主要用来放置非表单项组件
func Card_body(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["body"] = p
	}
}

// 内容区域类名
func Card_bodyClassName(p string) opt {
	return func(o map[string]interface{}) {
		o["bodyClassName"] = p
	}
}

// 配置按钮集合
func Card_actions(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["actions"] = p
	}
}

// 按钮集合每行个数
func Card_actionsCount(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["actionsCount"] = p
	}
}

// 点击卡片的行为
func Card_itemAction(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["itemAction"] = p
	}
}

// Card 多媒体部内容设置
func Card_media(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["media"] = p
	}
}

// // 图片/视频链接
// func Card_media.url(p string) opt {
// 	return func(o map[string]interface{}) {
// 		o["media.url"] = p
// 	}
// }

// // 多媒体类名
// func Card_media.className(p string) opt {
// 	return func(o map[string]interface{}) {
// 		o["media.className"] = p
// 	}
// }

// // 视频是否为直播
// func Card_media.isLive(p bool) opt {
// 	return func(o map[string]interface{}) {
// 		o["media.isLive"] = p
// 	}
// }

// // 视频是否自动播放
// func Card_media.autoPlay(p bool) opt {
// 	return func(o map[string]interface{}) {
// 		o["media.autoPlay"] = p
// 	}
// }

// // 视频封面
// func Card_media.poster(p string) opt {
// 	return func(o map[string]interface{}) {
// 		o["media.poster"] = p
// 	}
// }

// 次要说明
func Card_secondary(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["secondary"] = p
	}
}

// 工具栏按钮
func Card_toolbar(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["toolbar"] = p
	}
}

// 是否显示拖拽图标
func Card_dragging(p bool) opt {
	return func(o map[string]interface{}) {
		o["dragging"] = p
	}
}

// 卡片是否可选
func Card_selectable(p bool) opt {
	return func(o map[string]interface{}) {
		o["selectable"] = p
	}
}

// 卡片选择按钮是否禁用
func Card_checkable(p bool) opt {
	return func(o map[string]interface{}) {
		o["checkable"] = p
	}
}

// 卡片选择按钮是否选中
func Card_selected(p bool) opt {
	return func(o map[string]interface{}) {
		o["selected"] = p
	}
}

// 卡片选择按钮是否隐藏
func Card_hideCheckToggler(p bool) opt {
	return func(o map[string]interface{}) {
		o["hideCheckToggler"] = p
	}
}

// 卡片是否为多选
func Card_multiple(p bool) opt {
	return func(o map[string]interface{}) {
		o["multiple"] = p
	}
}

// 卡片内容区的表单项 label 是否使用 Card 内部的样式
func Card_useCardLabel(p bool) opt {
	return func(o map[string]interface{}) {
		o["useCardLabel"] = p
	}
}
