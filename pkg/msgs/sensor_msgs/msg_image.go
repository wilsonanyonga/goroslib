//autogenerated:yes
//nolint:revive,lll
package sensor_msgs

import (
	"github.com/bluenviron/goroslib/v2/pkg/msg"
	"github.com/bluenviron/goroslib/v2/pkg/msgs/std_msgs"
)

type Image struct {
	msg.Package `ros:"sensor_msgs"`
	Header      std_msgs.Header
	Height      uint32
	Width       uint32
	Encoding    string
	IsBigendian uint8
	Step        uint32
	Data        []uint8
}