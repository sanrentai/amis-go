package form

// InputExcel
func InputExcel(opts ...opt) map[string]interface{} {
	return newForm("input-excel", opts...)
}



// 是否解析所有 sheet
func InputExcel_allSheets(p bool) opt {
	return func(o map[string]interface{}) {
		o["allSheets"] = p
	}
} 

// 解析模式
func InputExcel_parseMode(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["parseMode"] = p
	}
} 

// 是否包含空值
func InputExcel_includeEmpty(p bool) opt {
	return func(o map[string]interface{}) {
		o["includeEmpty"] = p
	}
} 

// 是否解析为纯文本
func InputExcel_plainText(p bool) opt {
	return func(o map[string]interface{}) {
		o["plainText"] = p
	}
} 