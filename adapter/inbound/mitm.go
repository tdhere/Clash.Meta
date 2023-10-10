package inbound

import (
	"net"

	C "github.com/Dreamacro/clash/constant"
	"github.com/Dreamacro/clash/context"
	"github.com/Dreamacro/clash/transport/socks5"
)

// NewMitm receive mitm request and return MitmContext
func NewMitm(target socks5.Addr, source net.Addr, userAgent string, conn net.Conn) *context.ConnContext {
	metadata := parseSocksAddr(target)
	metadata.NetWork = C.TCP
	metadata.Type = C.MITM
	metadata.UserAgent = userAgent
	if addrPort, err := parseAddr(source); err == nil {
		metadata.SrcIP = addrPort.Addr()
		metadata.SrcPort = addrPort.Port()
	}
	return context.NewConnContext(conn, metadata)
}
