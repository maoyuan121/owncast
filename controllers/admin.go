package controllers

import (
	"net/http"

	"github.com/owncast/owncast/core/rtmp"
)

// DisconnectInboundConnection 将强制关闭 inbound stream 的连接
func DisconnectInboundConnection(w http.ResponseWriter, r *http.Request) {
	rtmp.Disconnect()
	w.WriteHeader(http.StatusOK)
}
