package retry

import (
	"context"
	"time"
)

type RetryableFunc func() error

type Retryer interface {
	ShouldRetry(ctx context.Context, err error) bool
	Do(ctx context.Context, f RetryableFunc) error
}

func NewRetryer(opts ...Option) Retryer {
	o := &options{}
	for _, opt := range opts {
		opt(o)
	}
	return &retryer{
		o: o,
	}
}

type retryer struct {
	o *options
}

func (r *retryer) ShouldRetry(ctx context.Context, err error) bool {
	return true
}

func (r *retryer) Do(ctx context.Context, f RetryableFunc) (err error) {
	maxDuration := r.o.MaxDuration
	maxRetryTimes := r.o.MaxRetryTimes

	var retryTimes int
	startTime := time.Now()
	for i := 0; i <= maxRetryTimes; i++ {
		//var retryStart time.Time
		if i == 0 {
			//retryStart = startTime
		} else if i > 0 {
			if maxDuration > 0 && time.Since(startTime) > maxDuration {
				break
			}
			if ok := r.ShouldRetry(ctx, err); !ok {
				break
			}
			//retryStart = time.Now()
		}
		retryTimes++
		err = f()
		if err == nil {
			break
		} else {
			if i == maxRetryTimes {
				//todo wrap error

			} else if !r.isRetryErr(err) {
				break
			}
		}
	}
	return
}

func (r *retryer) isRetryErr(err error) bool {
	if err == nil {
		return false
	}
	if r.o.ErrorRetry != nil && r.o.ErrorRetry(err) {
		return true
	}
	return false
}
