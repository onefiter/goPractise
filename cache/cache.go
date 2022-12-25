package cache

type Cache interface {
	Set(key string, value interface{})
	Get(key string) interface{}
	Del(key string)
	DelOldest()
	Len() int
}

type Value interface {
	Len() int
}
