//autogenerated:yes
//nolint:revive,lll
package actionlib_msgs

import (
	"github.com/bluenviron/goroslib/v2/pkg/msg"
	"github.com/bluenviron/goroslib/v2/pkg/msgs/std_msgs"
)

type GoalStatusArray struct {
	msg.Package `ros:"actionlib_msgs"`
	Header      std_msgs.Header
	StatusList  []GoalStatus
}
