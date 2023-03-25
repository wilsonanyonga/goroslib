//autogenerated:yes
//nolint:revive,lll
package trajectory_msgs

import (
	"time"

	"github.com/bluenviron/goroslib/v2/pkg/msg"
	"github.com/bluenviron/goroslib/v2/pkg/msgs/geometry_msgs"
)

type MultiDOFJointTrajectoryPoint struct {
	msg.Package   `ros:"trajectory_msgs"`
	Transforms    []geometry_msgs.Transform
	Velocities    []geometry_msgs.Twist
	Accelerations []geometry_msgs.Twist
	TimeFromStart time.Duration
}
