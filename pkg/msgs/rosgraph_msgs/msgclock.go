//autogenerated:yes
//nolint:revive,lll
package rosgraph_msgs

import (
	"time"

	"github.com/bluenviron/goroslib/v2/pkg/msg"
)

type Clock struct {
	msg.Package `ros:"rosgraph_msgs"`
	Clock       time.Time
}
