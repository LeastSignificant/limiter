package limiter

import (
	"fmt"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

func (l *Limiter) CheckUpdate(b *gotgbot.Bot, u *gotgbot.Update) bool {
	return true
}

func (l *Limiter) HandleUpdate(b *gotgbot.Bot, ctx *ext.Context) error {
	if ctx.EffectiveUser == nil {
		return ext.ContinueGroups
	}

	key := l.KeyGenerator(ctx)
	if key == "" {
		return ext.ContinueGroups
	}

	l.mu.Lock()
	l.hits[key]++
	hits := l.hits[key]
	l.mu.Unlock()

	if hits > l.Limit {
		if l.OnLimitExceeded != nil {
			go l.OnLimitExceeded(b, ctx)
		}

		return ext.EndGroups
	}

	return ext.ContinueGroups
}

func (l Limiter) Name() string {
	return fmt.Sprintf("limiter_%p", l.HandleUpdate)
}
