package amis

// 返回成功
func Ok(data interface{}) map[string]interface{} {
	return map[string]interface{}{"status": 0, "data": data}
}

// 返回错误
func Err(err error) map[string]interface{} {
	return Fail(-1, err.Error())
}

// 返回失败
func Fail(status int, msg string) map[string]interface{} {
	return map[string]interface{}{"status": status, "msg": msg}
}
