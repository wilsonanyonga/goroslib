//autogenerated:yes
//nolint:revive,lll
package std_msgs

import (
	"time"

	"github.com/bluenviron/goroslib/v2/pkg/msg"
)

type Time struct {
	msg.Package `ros:"std_msgs"`
	Data        time.Time
}
