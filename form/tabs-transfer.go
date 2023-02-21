package form

// 组合穿梭器
func TabsTransfer(opts ...opt) map[string]interface{} {
	return newForm("tabs-transfer", opts...)
}



// 更新数据，开启`multiple`支持设置多项，开启`joinValues`时，多值用`,`分隔，否则多值用数组
func TabsTransfer_setValue(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["setValue"] = p
	}
} 