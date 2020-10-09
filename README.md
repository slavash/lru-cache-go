# LRU Cache

Implementation of LRU logic for cache in Go

Usage:

```go
cache := cache.New(2)

u1 := User{name: "slava", age:  40}
...

user, found := cache.Get("user1")
if !found {
    fmt.Printf("User1 not found")
    return
}
    
```