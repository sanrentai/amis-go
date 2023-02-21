package form

// 密码输入框
func InputPassword(opts ...opt) map[string]interface{} {
	return newForm("input-password", opts...)
}



// 是否展示密码显/隐按钮
func InputPassword_revealPassword(p bool) opt {
	return func(o map[string]interface{}) {
		o["revealPassword"] = p
	}
} 