package common

import (
	"context"
	"errors"
	"github.com/qiniu/qmgo"
	"sync"
	"time"
)

func ensureProbeDeadlineFromContext(ctx context.Context) time.Time {
	ctxDeadline, hasDeadline := ctx.Deadline()
	if !hasDeadline {
		ctxDeadline = time.Now().Add(1 * time.Second)
	}

	return ctxDeadline
}

// WaitTimeout waits for the waitgroup for the specified max timeout.
// Returns nil on completion or ErrWaitTimeout if waiting timed out.
// See https://stackoverflow.com/questions/32840687/timeout-for-waitgroup-wait
// Note that the spawned goroutine to wg.Wait() gets leaked and will continue running detached
func WaitTimeout(wg *sync.WaitGroup, timeout time.Duration) error {
	c := make(chan struct{})
	go func() {
		defer close(c)
		wg.Wait()
	}()
	select {
	case <-c:
		return nil // completed normally
	case <-time.After(timeout):
		return errors.New("WaitGroup has timed out") // timed out
	}
}

func ProbeReadiness(ctx context.Context, dbClient *qmgo.Client) error {
	ctxDeadline := ensureProbeDeadlineFromContext(ctx)

	var dbPingWg sync.WaitGroup
	var dbErr error

	dbPingWg.Add(1)

	go func() {
		dbErr = dbClient.Ping(5)
		dbPingWg.Done()
	}()

	if err := WaitTimeout(&dbPingWg, time.Until(ctxDeadline)); err != nil {
		return err
	}

	if dbErr != nil {
		return dbErr
	}
	return nil
}

func ProbeLiveness(ctx context.Context, dbClient *qmgo.Client) error {
	err := ProbeReadiness(ctx, dbClient)
	if err != nil {
		return err
	}

	return nil
}
