package retry

import "time"

type Option func(o *options)

type options struct {
	MaxRetryTimes int
	MaxDuration   time.Duration
	ErrorRetry    func(err error) bool
}

func WithMaxRetryTimes(maxRetryTimes int) Option {
	return func(o *options) {
		o.MaxRetryTimes = maxRetryTimes
	}
}

func WithMaxDuration(maxDuration time.Duration) Option {
	return func(o *options) {
		o.MaxDuration = maxDuration
	}
}

func WithErrorRetry(errorRetry func(err error) bool) Option {
	return func(o *options) {
		o.ErrorRetry = errorRetry
	}
}
