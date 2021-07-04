//nolint:golint
package tf

import (
	"github.com/aler9/goroslib/pkg/msg"
	"github.com/aler9/goroslib/pkg/msgs/geometry_msgs"
)

type TfMessage struct {
	msg.Package `ros:"tf"`
	Transforms  []geometry_msgs.TransformStamped
}