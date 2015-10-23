package tap

import (
	"github.com/XANi/testvortex/src/testproto"
	"fmt"
)

func (t Test) RenderTest(id int) error {
	out := ""
	if t.Success {
		out += "ok "
	} else {
		out += "not ok "
	}
	if id > 0 {
		out += fmt.Sprintf("%d ", id)
	}
	// TODO sanitize it
	if t.Description {
		out += t.Description
	}
	return out + "\n"
}
