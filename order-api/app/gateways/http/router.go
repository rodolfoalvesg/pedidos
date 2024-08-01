package http

import (
	"context"
	"net/http"
	"order-api/app/domain/usecases"
	"order-api/app/gateways/http/client"
	"order-api/app/gateways/http/handlers"
	"order-api/app/gateways/http/rest"
	cache "order-api/app/gateways/repository/cache/redis"
	database "order-api/app/gateways/repository/db/postgres"
	elastic "order-api/app/gateways/repository/search/elasticsearch"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/go-chi/chi"
	"github.com/redis/go-redis/v9"
	httpSwagger "github.com/swaggo/http-swagger"
	"gorm.io/gorm"
)

// newHandler creates a new HTTP handler.
// @title Order API
// @version 2.0
// @BasePath /
// @schemes http
// @host localhost:3001
// @produce json
// @consumes json

// @summary This is a order API.
// @description This is a order API.
func newHandler(
	ctx context.Context,
	db *gorm.DB,
	rdb *redis.Client,
	es *elasticsearch.Client,
	client *client.Client,
) (http.Handler, error) {
	r := chi.NewRouter()

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	r.Route("/docs/v1/order-api", func(r chi.Router) {
		r.Get("/swagger", func(w http.ResponseWriter, r *http.Request) {
			http.Redirect(w, r, "swagger/index.html", http.StatusMovedPermanently)
		})
		r.Get("/swagger/*", httpSwagger.Handler())
	})

	repoPostgres := database.NewRepository(db)
	repoRedis := cache.NewRepository(rdb)
	repoElastic := elastic.NewRepository(es)
	usecase := usecases.NewUsecase(repoPostgres, repoRedis, repoElastic, client)
	handler := handlers.NewHandler(usecase)

	r.Route("/api/v1/order-api/orders", func(r chi.Router) {
		r.Post("/", rest.Handle(handler.CreateOrder))
		r.Get("/{order-id}", rest.Handle(handler.GetOrderByID))
		r.Get("/", rest.Handle(handler.GetAllOrders))
		r.Put("/{order-id}", rest.Handle(handler.UpdateOrder))
		r.Delete("/{order-id}", rest.Handle(handler.DeleteOrder))
	})

	return r, nil
}
