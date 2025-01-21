package pokeapi

import (
	"net/http"
	"time"

	"github.com/Aman-Shetty/pokedexcli/internal/cache"
)

const baseURL = "https://pokeapi.co/api/v2"

type Client struct {
	cache cache.Cache
	httpClient http.Client
}

// After a minute the http request will be closed 
func NewClient(cacheInterval time.Duration) Client {
	return Client{
		cache: cache.NewCache(cacheInterval),
		httpClient: http.Client{
			Timeout: time.Minute,
		},
	}
}

