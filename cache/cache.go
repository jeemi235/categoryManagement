package cache

import (
	"categorymanagement/models"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/patrickmn/go-cache"
)

const (
	defaultExpiration = 10 * time.Minute
	purgeTime         = 10 * time.Minute
)

type allCache struct {
	products *cache.Cache
}

var c = newCache()

func CheckCache(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		id:=name
		res, ok := c.read(id)
		if ok {
			w.Header().Set("Content-Type", "application/json")
			w.Write(res)
			return
		}
		log.Println("Got Data From Controller")
		handler.ServeHTTP(w, r)
	})
}

func newCache() *allCache {
	Cache := cache.New(defaultExpiration, purgeTime)
	return &allCache{
		products: Cache,
	}
}

func (c *allCache) read(id string) (item []byte, ok bool) {
	data, ok := c.products.Get(id)
	if ok {
		log.Println("Got Data From Cache")
		res, err := json.Marshal(data)
		if err != nil {
			log.Fatal("Error")
		}
		return res, true
	}
	return nil, false
}

func UpdateCache(id string, data []models.Types) {
	c.Update(id, data)
}

func (c *allCache) Update(id string, data []models.Types) {
	c.products.Set(id, data, cache.DefaultExpiration)
}
