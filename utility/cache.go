package utility
import (
	cache "github.com/patrickmn/go-cache"
	"time"
)
type Cache struct{
	d cache.Cache
}
func (l *Cache) New(default_expiration, purge_expired time.Duration){
	cache.New(default_expiration , purge_expired )
}
func (l *Cache) Get(key string) (interface{}, bool){
	return l.d.Get(key)
}
func (l *Cache) Set(k string, x interface{}, d time.Duration){
	l.d.Set(k,x,d)
}