package rabbitclient

import "github.com/google/wire"

var Set = wire.NewSet(NewRabbitClientConnection)
