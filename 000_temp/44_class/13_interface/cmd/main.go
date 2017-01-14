package cmd

import (
	"fmt"
	"github.com/GoesToEleven/golang-web-dev/000_temp/44_class/13_interface/cache"
)

func CacheUser(cache cache.Cache, id, def string) {
	val := cache.Get(id)
	if val == nil {
		val = def
		cache.Set(id, val)
	}
	fmt.Println(val)
}
