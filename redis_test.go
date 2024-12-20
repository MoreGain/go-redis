package redis

import (
	"fmt"
	"testing"
)

func TestClient(t *testing.T) {
	client := NewClient(&Options{
		Network: "tcp",
		Addr:    "127.0.0.1:6379",
	})
	fmt.Print(client)
}
