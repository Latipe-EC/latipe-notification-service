package mongodb

import "github.com/google/wire"

var Set = wire.NewSet(OpenMongoDBConnection)
