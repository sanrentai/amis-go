package form

// 文件上传
func InputFile(opts ...opt) map[string]interface{} {
	return newForm("input-file", opts...)
}



// 上传文件接口
func InputFile_receiver(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["receiver"] = p
	}
} 

// 默认只支持纯文本，要支持其他类型，请配置此属性为文件后缀`.xxx`
func InputFile_accept(p string) opt {
	return func(o map[string]interface{}) {
		o["accept"] = p
	}
} 

// 将文件以`base64`的形式，赋值给当前组件
func InputFile_asBase64(p bool) opt {
	return func(o map[string]interface{}) {
		o["asBase64"] = p
	}
} 

// 将文件以二进制的形式，赋值给当前组件
func InputFile_asBlob(p bool) opt {
	return func(o map[string]interface{}) {
		o["asBlob"] = p
	}
} 

// 默认没有限制，当设置后，文件大小大于此值将不允许上传。单位为`B`
func InputFile_maxSize(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["maxSize"] = p
	}
} 

// 默认没有限制，当设置后，一次只允许上传指定数量文件。
func InputFile_maxLength(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["maxLength"] = p
	}
} 

// 是否多选。
func InputFile_multiple(p bool) opt {
	return func(o map[string]interface{}) {
		o["multiple"] = p
	}
} 

// 是否为拖拽上传
func InputFile_drag(p bool) opt {
	return func(o map[string]interface{}) {
		o["drag"] = p
	}
} 

// [拼接值](./options#%E6%8B%BC%E6%8E%A5%E5%80%BC-joinvalues)
func InputFile_joinValues(p bool) opt {
	return func(o map[string]interface{}) {
		o["joinValues"] = p
	}
} 

// [提取值](./options#%E6%8F%90%E5%8F%96%E5%A4%9A%E9%80%89%E5%80%BC-extractvalue)
func InputFile_extractValue(p bool) opt {
	return func(o map[string]interface{}) {
		o["extractValue"] = p
	}
} 

// [拼接符](./options#%E6%8B%BC%E6%8E%A5%E7%AC%A6-delimiter)
func InputFile_delimiter(p string) opt {
	return func(o map[string]interface{}) {
		o["delimiter"] = p
	}
} 

// 否选择完就自动开始上传
func InputFile_autoUpload(p bool) opt {
	return func(o map[string]interface{}) {
		o["autoUpload"] = p
	}
} 

// 隐藏上传按钮
func InputFile_hideUploadButton(p bool) opt {
	return func(o map[string]interface{}) {
		o["hideUploadButton"] = p
	}
} 

// 上传状态文案
func InputFile_stateTextMap(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["stateTextMap"] = p
	}
} 

// 如果你不想自己存储，则可以忽略此属性。
func InputFile_fileField(p string) opt {
	return func(o map[string]interface{}) {
		o["fileField"] = p
	}
} 

// 接口返回哪个字段用来标识文件名
func InputFile_nameField(p string) opt {
	return func(o map[string]interface{}) {
		o["nameField"] = p
	}
} 

// 文件的值用那个字段来标识。
func InputFile_valueField(p string) opt {
	return func(o map[string]interface{}) {
		o["valueField"] = p
	}
} 

// 文件下载地址的字段名。
func InputFile_urlField(p string) opt {
	return func(o map[string]interface{}) {
		o["urlField"] = p
	}
} 

// 上传按钮的文字
func InputFile_btnLabel(p string) opt {
	return func(o map[string]interface{}) {
		o["btnLabel"] = p
	}
} 

// 默认显示文件路径的时候会支持直接下载，可以支持加前缀如：`http://xx.dom/filename=` ，如果不希望这样，可以把当前配置项设置为 `false`。
func InputFile_downloadUrl(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["downloadUrl"] = p
	}
} 

// amis 所在服务器，限制了文件上传大小不得超出 10M，所以 amis 在用户选择大文件的时候，自动会改成分块上传模式。
func InputFile_useChunk(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["useChunk"] = p
	}
} 

// 分块大小
func InputFile_chunkSize(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["chunkSize"] = p
	}
} 

// startChunkApi
func InputFile_startChunkApi(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["startChunkApi"] = p
	}
} 

// chunkApi
func InputFile_chunkApi(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["chunkApi"] = p
	}
} 

// finishChunkApi
func InputFile_finishChunkApi(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["finishChunkApi"] = p
	}
} 

// 分块上传时并行个数
func InputFile_concurrency(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["concurrency"] = p
	}
} 

// 文档内容
func InputFile_documentation(p string) opt {
	return func(o map[string]interface{}) {
		o["documentation"] = p
	}
} 

// 文档链接
func InputFile_documentLink(p string) opt {
	return func(o map[string]interface{}) {
		o["documentLink"] = p
	}
} 

// 初表单反显时是否执行
func InputFile_initAutoFill(p bool) opt {
	return func(o map[string]interface{}) {
		o["initAutoFill"] = p
	}
} 

// 上传文件值变化时触发(上传失败同样会触发)
func InputFile_change(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["change"] = p
	}
} 

// 移除文件时触发
func InputFile_remove(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["remove"] = p
	}
} 

// 上传成功时触发
func InputFile_success(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["success"] = p
	}
} 

// 上传文件失败时触发
func InputFile_fail(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["fail"] = p
	}
} 