//nolint:golint
package mavros_msgs

import (
	"github.com/aler9/goroslib/pkg/msg"
	"github.com/aler9/goroslib/pkg/msgs/std_msgs"
)

type ESCStatus struct {
	msg.Package `ros:"mavros_msgs"`
	Header      std_msgs.Header
	EscStatus   []ESCStatusItem
}