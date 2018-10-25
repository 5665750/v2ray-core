package tcp

import (
	"v2ray.com/core/common"
	"v2ray.com/core/transport/internet"
)

const protocolName = "tcp"

func init() {
	common.Must(internet.RegisterProtocolConfigCreatorByName(protocolName, func() interface{} {
		return new(Config)
	}))
}
