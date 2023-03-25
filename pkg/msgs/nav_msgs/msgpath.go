//autogenerated:yes
//nolint:revive,lll
package nav_msgs

import (
	"github.com/bluenviron/goroslib/v2/pkg/msg"
	"github.com/bluenviron/goroslib/v2/pkg/msgs/geometry_msgs"
	"github.com/bluenviron/goroslib/v2/pkg/msgs/std_msgs"
)

type Path struct {
	msg.Package `ros:"nav_msgs"`
	Header      std_msgs.Header
	Poses       []geometry_msgs.PoseStamped
}
