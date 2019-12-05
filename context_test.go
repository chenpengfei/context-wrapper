package context

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestContext(t *testing.T) {
	assert := assert.New(t)

	t.Run("timeout parent context", func(t *testing.T) {
		parent, cancel := context.WithTimeout(context.Background(), time.Second)
		ctx := WithSignal(parent)

		<-ctx.Done()
		cancel()
		_, ok := ctx.Deadline()
		assert.True(ok)
	})
}
