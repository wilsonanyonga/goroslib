//autogenerated:yes
//nolint:revive,lll
package geographic_msgs

import (
	"github.com/bluenviron/goroslib/v2/pkg/msg"
	"github.com/bluenviron/goroslib/v2/pkg/msgs/uuid_msgs"
)

type MapFeature struct {
	msg.Package `ros:"geographic_msgs"`
	Id          uuid_msgs.UniqueID
	Components  []uuid_msgs.UniqueID
	Props       []KeyValue
}
