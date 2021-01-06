package utils

import (
	"crypto/md5" //nolint
	"encoding/hex"
	"net"
	"net/http"
	"strings"

	log "github.com/sirupsen/logrus"
)

// GenerateClientIDFromRequest 根据 Http 请求 生成一个 client id
// 这个 client id 由 ip 地址加 useragent md5 加密生成
func GenerateClientIDFromRequest(req *http.Request) string {
	ipAddress := GetIPAddressFromRequest(req)
	ipAddressComponents := strings.Split(ipAddress, ":")
	ipAddressComponents[len(ipAddressComponents)-1] = ""
	clientID := strings.Join(ipAddressComponents, ":") + req.UserAgent()

	// Create a MD5 hash of this ip + useragent
	b := md5.Sum([]byte(clientID)) // nolint
	return hex.EncodeToString(b[:])
}

// GetIPAddressFromRequest 从 Http 请求中返回 IP 地址
// 如果请求头设置了 X-FORWARDED-FOR 的话，从其中取值作为 IP 地址
func GetIPAddressFromRequest(req *http.Request) string {
	xForwardedFor := req.Header.Get("X-FORWARDED-FOR")
	if xForwardedFor != "" {
		return xForwardedFor
	}

	ip, _, err := net.SplitHostPort(req.RemoteAddr)
	if err != nil {
		log.Errorln(err)
		return ""
	}

	return ip
}
