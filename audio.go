package amis

// 音频
func Audio(opts ...opt) map[string]interface{} {
	return newCompent("audio", opts...)
}

// 外层 Dom 的类名
func Audio_className(p string) opt {
	return func(o map[string]interface{}) {
		o["className"] = p
	}
}

// 是否是内联模式
func Audio_inline(p bool) opt {
	return func(o map[string]interface{}) {
		o["inline"] = p
	}
}

// 音频地址
func Audio_src(p string) opt {
	return func(o map[string]interface{}) {
		o["src"] = p
	}
}

// 是否循环播放
func Audio_loop(p bool) opt {
	return func(o map[string]interface{}) {
		o["loop"] = p
	}
}

// 是否自动播放
func Audio_autoPlay(p bool) opt {
	return func(o map[string]interface{}) {
		o["autoPlay"] = p
	}
}

// 可配置音频播放倍速如：`[1.0, 1.5, 2.0]`
func Audio_rates(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["rates"] = p
	}
}

// 内部模块定制化
func Audio_controls(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["controls"] = p
	}
}
