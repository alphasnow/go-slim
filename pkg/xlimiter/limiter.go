package xlimiter

import (
	libgin "github.com/gin-gonic/gin"
	libredis "github.com/go-redis/redis/v8"
	"github.com/ulule/limiter/v3"
	mgin "github.com/ulule/limiter/v3/drivers/middleware/gin"
	sredis "github.com/ulule/limiter/v3/drivers/store/redis"
	"strings"
)

type Limiter struct {
	Redis          *libredis.Client
	Prefix         string
	Rate           string
	ReachedHandler mgin.LimitReachedHandler
	Config         *Config
}

func (l *Limiter) SetPrefix(prefix string) *Limiter {
	l.Prefix = l.Config.JoinPrefix(prefix)
	return l
}
func (l *Limiter) SetRate(rate string) *Limiter {
	l.Rate = rate
	return l
}

func (l *Limiter) newRate() (limiter.Rate, error) {
	return limiter.NewRateFromFormatted(l.Rate)
}
func (l *Limiter) newStore() (limiter.Store, error) {
	return sredis.NewStoreWithOptions(l.Redis, limiter.StoreOptions{
		Prefix: l.Prefix,
	})
}

func (l *Limiter) NewMiddleware() libgin.HandlerFunc {
	// * 5 reqs/second: "5-S"
	// * 10 reqs/minute: "10-M"
	// * 1000 reqs/hour: "1000-H"
	// * 2000 reqs/day: "2000-D"
	rate, err := l.newRate()
	if err != nil {
		panic(err)
	}

	// Create a store with the redis client.
	store, err := l.newStore()
	if err != nil {
		panic(err)
	}

	// error
	opt := mgin.WithLimitReachedHandler(l.ReachedHandler)

	// middleware
	return mgin.NewMiddleware(limiter.New(store, rate), opt)
}

func (l *Limiter) NewMiddlewares(rules []string) map[string]libgin.HandlerFunc {
	res := map[string]libgin.HandlerFunc{}
	// auth-login:6-M
	for _, rule := range rules {
		splits := strings.Split(rule, ":")
		prefix := splits[0]
		rate := splits[1]
		res[rule] = l.SetRate(rate).SetPrefix(prefix).NewMiddleware()
	}
	return res
}

func NewLimiter(client *libredis.Client, handler mgin.LimitReachedHandler, config *Config) *Limiter {
	return &Limiter{Redis: client, Prefix: config.JoinPrefix("default"), Rate: config.DefaultRate, Config: config, ReachedHandler: handler}
}
