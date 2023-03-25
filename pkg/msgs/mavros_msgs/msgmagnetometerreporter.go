//autogenerated:yes
//nolint:revive,lll
package mavros_msgs

import (
	"github.com/bluenviron/goroslib/v2/pkg/msg"
	"github.com/bluenviron/goroslib/v2/pkg/msgs/std_msgs"
)

type MagnetometerReporter struct {
	msg.Package `ros:"mavros_msgs"`
	Header      std_msgs.Header
	Report      uint8
	Confidence  float32
}
