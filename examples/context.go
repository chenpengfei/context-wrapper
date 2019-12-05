package main

import (
	"context"
	"fmt"
	cw "github.com/chenpengfei/context-wrapper"
	"os"
)

func main() {
	ctx := cw.WithSignal(context.Background())

	p, err := os.FindProcess(os.Getpid())
	if err == nil {
		p.Signal(os.Interrupt)
	}

	select {
	case <-ctx.Done():
		fmt.Println("I have to go")
	}
}
