package form

// Editor
func JSONSchemaEditor(opts ...opt) map[string]interface{} {
	return newForm("json-schema-editor", opts...)
}

// 顶级类型是否可配置
func JSONSchemaEditorrootTypeMutable(p bool) opt {
	return func(o map[string]interface{}) {
		o["rootTypeMutable"] = p
	}
}

// 是否显示顶级类型信息
func JSONSchemaEditorshowRootInfo(p bool) opt {
	return func(o map[string]interface{}) {
		o["showRootInfo"] = p
	}
}

// 用来禁用默认数据类型，默认类型有：string、number、interger、object、number、array、boolean、null
func JSONSchemaEditordisabledTypes(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["disabledTypes"] = p
	}
}

// 用来配置预设类型
func JSONSchemaEditordefinitions(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["definitions"] = p
	}
}
