package automaxprocs

import (
	"go.uber.org/automaxprocs/maxprocs"
)

func init() {
	maxprocs.Set()
}
