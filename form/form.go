package form

func newForm(t string, opts ...opt) map[string]interface{} {
	var o = map[string]interface{}{"type": t}
	for _, op := range opts {
		op(o)
	}
	return o
}

type opt func(map[string]interface{})
