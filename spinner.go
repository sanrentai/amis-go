package amis

// 加载中
func Spinner(opts ...opt) map[string]interface{} {
	return newCompent("spinner", opts...)
}

// 是否显示 spinner 组件
func Spinner_show(p bool) opt {
	return func(o map[string]interface{}) {
		o["show"] = p
	}
}

// spinner 图标父级标签的自定义 class
func Spinner_className(p string) opt {
	return func(o map[string]interface{}) {
		o["className"] = p
	}
}

// 组件中 icon 所在标签的自定义 class
func Spinner_spinnerClassName(p string) opt {
	return func(o map[string]interface{}) {
		o["spinnerClassName"] = p
	}
}

// 作为容器使用时组件最外层标签的自定义 class
func Spinner_spinnerWrapClassName(p string) opt {
	return func(o map[string]interface{}) {
		o["spinnerWrapClassName"] = p
	}
}

// 组件大小 `sm` `lg`
func Spinner_size(p string) opt {
	return func(o map[string]interface{}) {
		o["size"] = p
	}
}

// 组件图标，可以是`amis`内置图标，也可以是字体图标或者网络图片链接，作为 ui 库使用时也可以是自定义组件
func Spinner_icon(p string) opt {
	return func(o map[string]interface{}) {
		o["icon"] = p
	}
}

// 配置组件文案，例如`加载中...`
func Spinner_tip(p string) opt {
	return func(o map[string]interface{}) {
		o["tip"] = p
	}
}

// 配置组件 `tip` 相对于 `icon` 的位置
func Spinner_tipPlacement(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["tipPlacement"] = p
	}
}

// 配置组件显示延迟的时间（毫秒）
func Spinner_delay(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["delay"] = p
	}
}

// 配置组件显示 spinner 时是否显示遮罩层
func Spinner_overlay(p bool) opt {
	return func(o map[string]interface{}) {
		o["overlay"] = p
	}
}

// 作为容器使用时，被包裹的内容
func Spinner_body(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["body"] = p
	}
}

// 为 `Spinner` 指定挂载的容器, `root` 是一个selector，在拥有`Spinner`的组件上都可以通过传递`loadingConfig`改变Spinner的挂载位置，开启后，会强制开启属性`overlay=true`
func Spinner_loadingConfig(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["loadingConfig"] = p
	}
}
