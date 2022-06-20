package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()

	ctx1 := context.WithValue(ctx, "Nombre", "Juan")
	fmt.Println("Usando contexto con valor", ctx1.Value("Nombre"))

	deadline := time.Now().Add(time.Second * 5)

	ctx, cancel := context.WithDeadline(ctx, deadline)
	defer cancel()

	<-ctx.Done()
	fmt.Println(ctx.Err().Error()) // context deadline exceeded
}
