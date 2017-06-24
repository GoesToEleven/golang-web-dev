package cache

type Cache interface {
	Set(k string, val interface{})
	Get(k string) interface{}
}
