package main

import (
	"fmt"

	cache "github.com/slavash/lru-cache-go"
)

type User struct {
	name string
	age  int
}

func main() {

	cache := cache.New(2)

	u1 := User{name: "slava", age:  40}
	u2 := User{name: "james", age:  39}
	u3 := User{name: "wyatt", age:  42}



	cache.Set("user1", u1)
	fmt.Printf("Cache: %+v\n", cache)

	cache.Set("user2", u2)
	fmt.Printf("Cache: %+v\n", cache)

	user, found := cache.Get("user1")
	if !found {
		fmt.Printf("User1 not found")
		return
	}
	
	fmt.Printf("Cache: %+v\n", cache)
	fmt.Printf("User1 from cache: %+v is now least recently used\n", user)

	cache.Set("user3", u3) // Replace user2
	fmt.Printf("Cache: %+v\n", cache)
}
