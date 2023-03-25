//autogenerated:yes
//nolint:revive,lll
package geometry_msgs

import (
	"github.com/bluenviron/goroslib/v2/pkg/msg"
	"github.com/bluenviron/goroslib/v2/pkg/msgs/std_msgs"
)

type PoseStamped struct {
	msg.Package `ros:"geometry_msgs"`
	Header      std_msgs.Header
	Pose        Pose
}
