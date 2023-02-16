package dialer

import (
	"github.com/Dreamacro/clash/common/cache"
	"github.com/Dreamacro/clash/log"
	"net"
	"time"
)

const ttl = 600 * time.Second

var lruCache = cache.New(cache.WithSize(4096), cache.WithStale(true))

func getIP(host string) (net.IP, bool) {
	ip, ok := lruCache.Get(host)
	if !ok {
		return nil, ok
	}
	return ip.(net.IP), ok
}

func setIP(host string, ip net.IP) {
	log.Debugln("[DNS cache] set %s --> %s", host, ip)
	lruCache.SetWithExpire(host, ip, time.Now().Add(ttl))
}

func deleteIP(host string) {
	log.Debugln("[DNS cache] delete %s", host)
	lruCache.Delete(host)
}
