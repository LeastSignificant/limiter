package limiter

import (
	"strconv"
	"sync"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

type Limiter struct {
	Limit           int
	TimeFrame       time.Duration
	OnLimitExceeded func(b *gotgbot.Bot, ctx *ext.Context)
	KeyGenerator    func(ctx *ext.Context) string
	mu              *sync.Mutex
	running         bool
	hits            map[string]int
}

func New(start bool) *Limiter {
	limiter := &Limiter{
		Limit:     1,
		TimeFrame: 5 * time.Second,
		KeyGenerator: func(ctx *ext.Context) string {
			return string(strconv.AppendInt([]byte(""), ctx.EffectiveUser.Id, 10))
		},
		mu:   &sync.Mutex{},
		hits: make(map[string]int),
	}

	if start {
		limiter.Start()
	}

	return limiter
}
