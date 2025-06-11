package middleware

import (
	"net/http"
	"sync"
	"time"
	"go_health_sentiment/utils"
)

type RateLimiter struct {
	visitors map[string]*visitor
	mu       sync.RWMutex
	rate     int
	burst    int
}

type visitor struct {
	limiter  *time.Ticker
	lastSeen time.Time
	count    int
}

func NewRateLimiter(rate, burst int) *RateLimiter {
	rl := &RateLimiter{
		visitors: make(map[string]*visitor),
		rate:     rate,
		burst:    burst,
	}
	
	// Clean up old visitors every minute
	go rl.cleanupVisitors()
	
	return rl
}

func (rl *RateLimiter) cleanupVisitors() {
	for {
		time.Sleep(time.Minute)
		rl.mu.Lock()
		for ip, v := range rl.visitors {
			if time.Since(v.lastSeen) > 3*time.Minute {
				delete(rl.visitors, ip)
			}
		}
		rl.mu.Unlock()
	}
}

func (rl *RateLimiter) getVisitor(ip string) *visitor {
	rl.mu.Lock()
	defer rl.mu.Unlock()
	
	v, exists := rl.visitors[ip]
	if !exists {
		v = &visitor{
			lastSeen: time.Now(),
			count:    0,
		}
		rl.visitors[ip] = v
	}
	
	v.lastSeen = time.Now()
	return v
}

func (rl *RateLimiter) RateLimitMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip := r.RemoteAddr
		visitor := rl.getVisitor(ip)
		
		// Simple rate limiting: allow burst requests, then limit
		if visitor.count >= rl.burst {
			utils.WriteError(w, http.StatusTooManyRequests, "Rate limit exceeded")
			return
		}
		
		visitor.count++
		
		// Reset count after rate interval
		go func() {
			time.Sleep(time.Duration(60/rl.rate) * time.Second)
			rl.mu.Lock()
			if v, exists := rl.visitors[ip]; exists {
				v.count--
				if v.count < 0 {
					v.count = 0
				}
			}
			rl.mu.Unlock()
		}()
		
		next.ServeHTTP(w, r)
	})
}