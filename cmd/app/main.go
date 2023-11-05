package main

import "context"

func main() {
	ctx := context.Background()
	_, err := NewApp(ctx)
	if err != nil {
		return
	}
}
