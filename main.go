package redis

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

func main() {
  fmt.Print("hello world")
  client := redis.NewClient(&redis.Options{
    Addr: "127.0.0.1:6379",
    Password: "",
    DB: 0,
  })

  ctx := context.Background()
}
