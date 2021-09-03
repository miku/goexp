package main

import (
	"context"
	"log"
	"time"
)

func Query(ctx context.Context) (string, error) {
	// this select block until one case is available
	select {
	case <-time.After(1 * time.Second):
		return "ok", nil
	case <-ctx.Done():
		log.Println("cancel requested ...")
		time.Sleep(200 * time.Millisecond)
		return "", ctx.Err()
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()
	// Query would normally take a second, but we're cancelling it.
	result, err := Query(ctx)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(result)
}
