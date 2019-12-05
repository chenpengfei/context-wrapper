# context-wrapper
[![Build Status](https://travis-ci.com/chenpengfei/context-wrapper.svg)](https://travis-ci.com/chenpengfei/context-wrapper)
[![Coverage Status](https://coveralls.io/repos/github/chenpengfei/context-wrapper/badge.svg)](https://coveralls.io/github/chenpengfei/context-wrapper)

> context wrapper

## Usage
graceful shutdown

```
ctx := cw.WithSignal(context.Background())

// do something

<-ctx.Done()
fmt.Println("I have to go")
```
