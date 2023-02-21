package amis

// Html
func Html(opts ...opt) map[string]interface{} {
	return newCompent("html", opts...)
}
