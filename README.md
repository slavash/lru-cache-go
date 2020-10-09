# LRU Cache

Implementation of LRU logic for cache in Go

Usage:

```go
cache := cache.New(2)

u1 := User{name: "slava", age:  40}
done := cache.Set("user1", u1)
if !done {
    fmt.Printf("failed to put user1 to the cache")
    return
}
...

user, found := cache.Get("user1")
if !found {
    fmt.Printf("User1 not found")
    return
}
    
```