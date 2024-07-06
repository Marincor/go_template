package entity

import (
	"sync"

	"golang.org/x/time/rate"
)

type IPRateLimiter struct {
	Ips     map[string]*rate.Limiter
	Mu      *sync.RWMutex
	Limiter *rate.Limiter
}

// Allow checks if the request from the given IP is allowed.
func (lim *IPRateLimiter) Allow(ip string) bool {
	lim.Mu.RLock()
	rl, exists := lim.Ips[ip]
	lim.Mu.RUnlock()

	if !exists {
		lim.Mu.Lock()
		rl, exists = lim.Ips[ip]
		if !exists {
			rl = rate.NewLimiter(lim.Limiter.Limit(), lim.Limiter.Burst())
			lim.Ips[ip] = rl
		}
		lim.Mu.Unlock()
	}

	return rl.Allow()
}
