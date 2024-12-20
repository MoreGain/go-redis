package main

import (
	"fmt"
	"moregain/go-redis"
)

func main() {
    fmt.Println("start conn redis...")
    testConnRedis() 
}

func testConnRedis() {
    client := redis.Client{}
    fmt.Printf("%+v\n", client)
}
