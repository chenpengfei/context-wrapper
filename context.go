package context

import (
	"context"
	"os"
	"os/signal"
	"syscall"
)

// WithSignal create a context cancelled when SIGINT or SIGTERM are notified
func WithSignal(ctx context.Context) context.Context {
	newCtx, cancel := context.WithCancel(ctx)
	signals := make(chan os.Signal)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-signals
		cancel()
		close(signals)
	}()
	return newCtx
}
