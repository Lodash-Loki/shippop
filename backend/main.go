package main

import (
	"context"

	"github.com/Lodash-Loki/shippop/bootstrap"
)

func main() {
	ctx := context.Background()
	bootstrap.Dispatch(ctx)
}