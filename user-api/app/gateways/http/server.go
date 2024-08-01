package http

import (
	"context"
	"net/http"
	"user-api/config"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func NewServer(
	ctx context.Context,
	cfg *config.Config,
	db *gorm.DB,
	rdb *redis.Client,
	es *elasticsearch.Client,
) (*http.Server, error) {
	handler, err := newHandler(ctx, db, rdb, es)
	if err != nil {
		return nil, err
	}

	s := http.Server{
		Addr:    cfg.HTTP.Host + ":" + cfg.HTTP.Port,
		Handler: handler,
	}

	return &s, nil
}
