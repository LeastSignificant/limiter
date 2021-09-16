package limiter

import "time"

func (l *Limiter) Start() bool {
	if !l.running {
		go l.interval()
		l.running = true
		return true
	}

	return false
}

func (l *Limiter) Stop() bool {
	if l.running {
		l.running = false
		return true
	}

	return false
}

func (l *Limiter) interval() {
	for l.running {
		time.Sleep(l.TimeFrame)
		l.mu.Lock()
		for key := range l.hits {
			delete(l.hits, key)
		}
		l.mu.Unlock()
	}
}
