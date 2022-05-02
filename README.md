# automaxprocs

Automatically set `GOMAXPROCS` to match Linux container CPU quota.

> This package is silent version of original Uber package.
> 
> Please read original documentation: [![GoDoc][doc-img]][doc]

## Installation

`go get -u github.com/gonevo/automaxprocs`

## Quick Start

```go
import _ "github.com/gonevo/automaxprocs"

func main() {
  // Your application logic here.
}
```

<hr>

Released under the [MIT License](LICENSE).

[doc-img]: https://godoc.org/go.uber.org/automaxprocs?status.svg
[doc]: https://godoc.org/go.uber.org/automaxprocs
