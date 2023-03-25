//autogenerated:yes
//nolint:revive,lll
package mavros_msgs

import (
	"github.com/bluenviron/goroslib/v2/pkg/msg"
)

type WaypointClearReq struct {
	msg.Package `ros:"mavros_msgs"`
}

type WaypointClearRes struct {
	msg.Package `ros:"mavros_msgs"`
	Success     bool
}

type WaypointClear struct {
	msg.Package `ros:"mavros_msgs"`
	WaypointClearReq
	WaypointClearRes
}
