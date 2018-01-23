package hyper

import (
	"encoding/json"
)

func JSONString(v interface{}) string {
	bs, _ := json.Marshal(v)
	return string(bs)
}
