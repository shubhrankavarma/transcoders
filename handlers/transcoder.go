package handlers

import (
	"sync/atomic"

	"github.com/amagimedia/transcoders/config"

	"github.com/amagimedia/transcoders/dbiface"
)

type TranscoderHandler struct {
	Col     dbiface.MongoCollectionAPI
	Cfg     config.Properties
	IsReady *atomic.Value
}
