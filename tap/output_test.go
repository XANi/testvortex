package tap

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestTapOutput(t *testing.T) {
	var te TestTAP
	Convey("TestFail", t, func() {
		So(te.RenderTest(1), ShouldEqual, "not ok 1")
	})
}
