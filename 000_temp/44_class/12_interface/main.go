package main

import "fmt"

// package cache
type Cache interface {
	Set(k string, val interface{})
	Get(k string) interface{}
}

// package memcache
type memCache struct {
	m map[string]interface{}
}

func (m *memCache) Set(k string, val interface{}) {
	m.m[k] = val
}

func (m *memCache) Get(k string) interface{} {
	return m.m[k]
}

// package cmd
// imports cache
func cacheUser(cache Cache, id, def string) {
	val := cache.Get(id)
	if val == nil {
		val = def
		cache.Set(id, val)
	}
	fmt.Println(val)
}

// package main
// imports memcache
func main() {
	c := &memCache{
		m: map[string]interface{}{},
	}
	cacheUser(c, "Bob", "Hello")
	cacheUser(c, "Bob", "Goodbye")
}
