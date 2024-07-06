package middleware

import (
	"net"
	"net/http"
	"sync"

	"api.default.marincor.pt/config/constants"
	"api.default.marincor.pt/entity"
	"api.default.marincor.pt/pkg/helpers"
	"golang.org/x/time/rate"
)

// NewIPRateLimiter creates a new instance of IPRateLimiter with the given rate limit.
func newIPRateLimiter(r rate.Limit, burst int) *entity.IPRateLimiter {
	return &entity.IPRateLimiter{
		Ips:     make(map[string]*rate.Limiter),
		Mu:      &sync.RWMutex{},
		Limiter: rate.NewLimiter(r, burst),
	}
}

func Throttling() entity.Middleware {
	limiter := newIPRateLimiter(constants.MaxResquestLimit, constants.MaxResquestLimit*2)

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ip, _, _ := net.SplitHostPort(r.RemoteAddr)
			if !limiter.Allow(ip) {
				helpers.CreateResponse(w, "Calls Limit Reached", constants.HTTPStatusTooManyRequests)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
