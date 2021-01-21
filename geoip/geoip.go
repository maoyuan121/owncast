// This package utilizes the MaxMind GeoLite2 GeoIP database https://dev.maxmind.com/geoip/geoip2/geolite2/.
// You must provide your own copy of this database for it to work.
// Read more about how this works at http://owncast.online/docs/geoip

package geoip

import (
	"net"

	"github.com/oschwald/geoip2-golang"
	"github.com/owncast/owncast/config"
	log "github.com/sirupsen/logrus"
)

var _geoIPCache = map[string]GeoDetails{}
var _enabled = true // Try to use GeoIP support it by default.

// GeoDetails stores details about a location.
type GeoDetails struct {
	CountryCode string `json:"countryCode"` // 国家代码
	RegionName  string `json:"regionName"`  // 区域
	TimeZone    string `json:"timeZone"`    // 时区
}

// 从缓存中获取 ip 对应的 geo 信息
func GetGeoFromIP(ip string) *GeoDetails {
	if cachedGeoDetails, ok := _geoIPCache[ip]; ok {
		return &cachedGeoDetails
	}

	if ip == "::1" || ip == "127.0.0.1" {
		return &GeoDetails{
			CountryCode: "N/A",
			RegionName:  "Localhost",
			TimeZone:    "",
		}
	}

	return nil
}

// 获取 ip 的 geo 信息写到缓存
// 先从缓存中找有的话直接返回，找不到才从数据库中找，并写回到缓存
func FetchGeoForIP(ip string) {
	// If GeoIP has been disabled then don't try to access it.
	if !_enabled {
		return
	}

	// Don't re-fetch if we already have it.
	if _, ok := _geoIPCache[ip]; ok {
		return
	}

	go func() {
		db, err := geoip2.Open(config.GeoIPDatabasePath)
		if err != nil {
			log.Traceln("GeoIP support is disabled. visit http://owncast.online/docs/geoip to learn how to enable.", err)
			_enabled = false
			return
		}

		defer db.Close()

		ipObject := net.ParseIP(ip)

		record, err := db.City(ipObject)
		if err != nil {
			log.Warnln(err)
			return
		}

		// If no country is available then exit
		if record.Country.IsoCode == "" {
			return
		}

		// If we believe this IP to be anonymous then no reason to report it
		if record.Traits.IsAnonymousProxy {
			return
		}

		var regionName = "Unknown"
		if len(record.Subdivisions) > 0 {
			if region, ok := record.Subdivisions[0].Names["en"]; ok {
				regionName = region
			}
		}

		response := GeoDetails{
			CountryCode: record.Country.IsoCode,
			RegionName:  regionName,
			TimeZone:    record.Location.TimeZone,
		}

		_geoIPCache[ip] = response
	}()
}
