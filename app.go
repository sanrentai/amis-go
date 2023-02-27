package amis

// 多页应用
func App(opts ...opt) map[string]interface{} {
	return newCompent("app", opts...)
}

// 页面配置接口，如果你想远程拉取页面配置请配置。返回配置路径 `json>data>pages`，具体格式请参考 `pages` 属性。
func App_api(p string) opt {
	return func(o map[string]interface{}) {
		o["api"] = p
	}
}

// 应用名称。
func App_brandName(p string) opt {
	return func(o map[string]interface{}) {
		o["brandName"] = p
	}
}

func App_showBreadcrumb(p bool) opt {
	return func(o map[string]interface{}) {
		o["showBreadcrumb"] = p
	}
}

// 支持图片地址，或者 svg。
func App_logo(p string) opt {
	return func(o map[string]interface{}) {
		o["logo"] = p
	}
}

// css 类名。
func App_className(p string) opt {
	return func(o map[string]interface{}) {
		o["className"] = p
	}
}

// 顶部区域。
func App_header(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["header"] = p
	}
}

// 页面菜单上前面区域。
func App_asideBefore(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["asideBefore"] = p
	}
}

// 页面菜单下前面区域。
func App_asideAfter(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["asideAfter"] = p
	}
}

// 页面。
func App_footer(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["footer"] = p
	}
}

// `Array<页面配置>`具体的页面配置。
//  通常为数组，数组第一层为分组，一般只需要配置 label 集合，如果你不想分组，直接不配置，真正的页面请在第二层开始配置，即第一层的 children 中。
func App_pages(p ...interface{}) opt {
	return func(o map[string]interface{}) {
		o["pages"] = p
	}
}

//   具体的页面配置也支持属性结构，每层有以下配置。

func Schema(opts ...opt) map[string]interface{} {
	var o = map[string]interface{}{}
	for _, op := range opts {
		op(o)
	}
	return o
}

// 菜单名称。
func Schema_label(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["label"] = p
	}
}

// 菜单图标，比如：fa fa-file.
func Schema_icon(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["icon"] = p
	}
}

// 页面路由路径，当路由命中该路径时，启用当前页面。当路径不是 `/` 打头时，会连接父级路径。比如：父级的路径为 `folder`，而此时配置 `pageA`, 那么当页面地址为 `/folder/pageA` 时才会命中此页面。当路径是 `/` 开头如： `/crud/list` 时，则不会拼接父级路径。另外还支持 `/crud/view/:id` 这类带参数的路由，页面中可以通过 `${params.id}` 取到此值。
func Schema_url(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["url"] = p
	}
}

//`schema` 页面的配置，具体配置请前往 [Page 页面说明](./page)
func Schema_schema(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["schema"] = p
	}
}

// 如果想通过接口拉取，请配置。返回路径为 `json>data`。schema 和 schemaApi 只能二选一。
func Schema_schemaApi(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["schemaApi"] = p
	}
}

// 如果想配置个外部链接菜单，只需要配置 link 即可。
func Schema_link(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["link"] = p
	}
}

// 跳转，当命中当前页面时，跳转到目标页面。
func Schema_redirect(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["redirect"] = p
	}
}

// 改成渲染其他路径的页面，这个方式页面地址不会发生修改。
func Schema_rewrite(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["rewrite"] = p
	}
}

// 当你需要自定义 404 页面的时候有用，不要出现多个这样的页面，因为只有第一个才会有用。
func Schema_isDefaultPage(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["isDefaultPage"] = p
	}
}

// 有些页面可能不想出现在菜单中，可以配置成 `false`，另外带参数的路由无需配置，直接就是不可见的。
func Schema_visible(p bool) opt {
	return func(o map[string]interface{}) {
		o["visible"] = p
	}
}

// 菜单类名。
func Schema_className(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["className"] = p
	}
}

func Schema_children(p ...interface{}) opt {
	return func(o map[string]interface{}) {
		o["children"] = p
	}
}
