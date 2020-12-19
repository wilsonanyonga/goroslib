//nolint:golint
package nav_msgs

import (
	"github.com/aler9/goroslib/pkg/msg"
	"github.com/aler9/goroslib/pkg/msgs/geometry_msgs"
)

type SetMapReq struct {
	msg.Package `ros:"nav_msgs"`
	Map         OccupancyGrid
	InitialPose geometry_msgs.PoseWithCovarianceStamped
}

type SetMapRes struct {
	msg.Package `ros:"nav_msgs"`
	Success     bool
}