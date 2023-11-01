package qualysapi

import (
	"net/http"
	"time"

	"github.com/derrynEdwards/qualys-toolbox/internal/qualyscache"
)

type Client struct {
	cache      qualyscache.Cache
	httpClient http.Client
}

func NewClient(timeout, cacheInterval time.Duration) Client {
	return Client{
		cache: qualyscache.NewCache(cacheInterval),
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}
