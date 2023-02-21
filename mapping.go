package amis

// 映射
func Mapping(opts ...opt) map[string]interface{} {
	return newCompent("mapping", opts...)
}

// 外层 CSS 类名
func Mapping_className(p string) opt {
	return func(o map[string]interface{}) {
		o["className"] = p
	}
}

// 占位文本
func Mapping_placeholder(p string) opt {
	return func(o map[string]interface{}) {
		o["placeholder"] = p
	}
}

// 映射配置
func Mapping_map(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["map"] = p
	}
}

// [API](../../../docs/types/api) 或 [数据映射](../../../docs/concepts/data-mapping)
func Mapping_source(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["source"] = p
	}
}

// `2.5.2` map或source为`Array<object>`时，用来匹配映射的字段名
func Mapping_valueField(p string) opt {
	return func(o map[string]interface{}) {
		o["valueField"] = p
	}
}

// `2.5.2` map或source为`Array<object>`时，用来展示的字段名<br />注：配置后映射值无法作为`schema`组件渲染
func Mapping_labelField(p string) opt {
	return func(o map[string]interface{}) {
		o["labelField"] = p
	}
}

// `2.5.2` 自定义渲染模板，支持`html`或`schemaNode`；<br /> 当映射值是`非object`时，可使用`${item}`获取映射值；<br />当映射值是`object`时，可使用映射语法: `${xxx}`获取`object`的值；<br /> 也可使用数据映射语法：`${xxx}`获取数据域中变量值。
func Mapping_itemSchema(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["itemSchema"] = p
	}
}
