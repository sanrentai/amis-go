package amis

// 视频
func Video(opts ...opt) map[string]interface{} {
	return newCompent("video", opts...)
}

// 外层 Dom 的类名
func Video_className(p string) opt {
	return func(o map[string]interface{}) {
		o["className"] = p
	}
}

// 视频地址
func Video_src(p string) opt {
	return func(o map[string]interface{}) {
		o["src"] = p
	}
}

// 是否为直播，视频为直播时需要添加上，支持`flv`和`hls`格式
func Video_isLive(p bool) opt {
	return func(o map[string]interface{}) {
		o["isLive"] = p
	}
}

// 指定直播视频格式
func Video_videoType(p string) opt {
	return func(o map[string]interface{}) {
		o["videoType"] = p
	}
}

// 视频封面地址
func Video_poster(p string) opt {
	return func(o map[string]interface{}) {
		o["poster"] = p
	}
}

// 是否静音
func Video_muted(p bool) opt {
	return func(o map[string]interface{}) {
		o["muted"] = p
	}
}

// 是否自动播放
func Video_autoPlay(p bool) opt {
	return func(o map[string]interface{}) {
		o["autoPlay"] = p
	}
}

// 倍数，格式为`[1.0, 1.5, 2.0]`
func Video_rates(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["rates"] = p
	}
}

// key 是时刻信息，value 可以可以为空，可有设置为图片地址，请看上方示例
func Video_frames(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["frames"] = p
	}
}

// 点击帧的时候默认是跳转到对应的时刻，如果想提前 3 秒钟，可以设置这个值为 3
func Video_jumpBufferDuration(p bool) opt {
	return func(o map[string]interface{}) {
		o["jumpBufferDuration"] = p
	}
}

// 到了下一帧默认是接着播放，配置这个会自动停止
func Video_stopOnNextFrame(p bool) opt {
	return func(o map[string]interface{}) {
		o["stopOnNextFrame"] = p
	}
}
