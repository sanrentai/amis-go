package amis

import "encoding/json"

func newCompent(t string, opts ...opt) map[string]interface{} {
	var o = map[string]interface{}{"type": t}
	for _, op := range opts {
		op(o)
	}
	return o
}

type opt = func(map[string]interface{})

func JSON(o map[string]interface{}, pretty ...bool) string {
	if len(pretty) == 1 && pretty[0] {
		bts, _ := json.MarshalIndent(o, "", "	")
		return string(bts)
	} else {
		bts, _ := json.Marshal(o)
		return string(bts)
	}

}
