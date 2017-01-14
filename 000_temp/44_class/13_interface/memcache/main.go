package memcache

type MemCache struct {
	M map[string]interface{}
}

func (m *MemCache) Set(k string, val interface{}) {
	m.M[k] = val
}

func (m *MemCache) Get(k string) interface{} {
	return m.M[k]
}
