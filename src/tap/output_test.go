package tap
import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"github.com/XANi/testvortex/src/testproto"
)

func TestTapOutput(t *testing.T) {
	var t testproto.Test;
	Convey("TestFail",t,func() {
		So (t.RenderTest(1), ShouldEqual, "ok 1")
	})
}
