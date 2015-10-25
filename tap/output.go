package tap

import (
	"fmt"
)


func (t TestTAP) RenderTest(id int) string {
	out := ""
	if t.Success {
		out += "ok"
	} else {
		out += "not ok"
	}
	if id > 0 {
		out += fmt.Sprintf(" %d", id)
	}
	// TODO sanitize it
	if len(t.Description) > 0 {
		out += t.Description
	}
	return out
}
