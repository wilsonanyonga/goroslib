//autogenerated:yes
//nolint:revive,lll
package geometry_msgs

import (
	"github.com/bluenviron/goroslib/v2/pkg/msg"
)

type Quaternion struct {
	msg.Package `ros:"geometry_msgs"`
	X           float64
	Y           float64
	Z           float64
	W           float64
}