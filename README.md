# gosudoblock
Block users from running your app with root permissions

### How to install
```go
go get github.com/shamsher31/gosudoblock
```

### How to use
```go
package main

import (
  "fmt"
  "github.com/shamsher31/gosudoblock"
)

func main() {
   sudoblock.Is()
   // You are not allowed to run this app with root permissions
}
```

### Related
[goisroot](https://github.com/shamsher31/goisroot)<br>
[ansicolor](https://github.com/shiena/ansicolor)<br>

### Why
This package is inspired by [sudo-block](https://www.npmjs.com/package/sudo-block) npm module.

### License
MIT © [Shamsher Ansari](https://github.com/shamsher31)
