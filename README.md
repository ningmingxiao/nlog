# a log lib can easy log and print stack


## install
```
go get -u github.com/ningmingxiao/nlog@v0.0.1
```

### example
```
package main

import (
	"github.com/ningmingxiao/nlog"
)

func main() {
	nlog.PrintStack()
	nlog.LogStack("/var/log/test.log")
	nlog.LogFile("/var/log/test.log", "hello")
}
```
