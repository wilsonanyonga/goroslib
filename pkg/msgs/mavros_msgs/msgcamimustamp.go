//autogenerated:yes
//nolint:revive,lll
package mavros_msgs

import (
	"time"

	"github.com/bluenviron/goroslib/v2/pkg/msg"
)

type CamIMUStamp struct {
	msg.Package `ros:"mavros_msgs"`
	FrameStamp  time.Time
	FrameSeqId  int32
}
