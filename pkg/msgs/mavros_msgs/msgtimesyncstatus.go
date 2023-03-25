//autogenerated:yes
//nolint:revive,lll
package mavros_msgs

import (
	"github.com/bluenviron/goroslib/v2/pkg/msg"
	"github.com/bluenviron/goroslib/v2/pkg/msgs/std_msgs"
)

type TimesyncStatus struct {
	msg.Package       `ros:"mavros_msgs"`
	Header            std_msgs.Header
	RemoteTimestampNs uint64
	ObservedOffsetNs  int64
	EstimatedOffsetNs int64
	RoundTripTimeMs   float32
}
