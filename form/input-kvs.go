package form

// 键值对象
func InputKVS(opts ...opt) map[string]interface{} {
	return newForm("input-kvs", opts...)
}

