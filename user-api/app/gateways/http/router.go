package http

import (
	"context"
	"net/http"
	"user-api/app/domain/usecases"
	"user-api/app/gateways/http/handlers"
	"user-api/app/gateways/http/rest"
	cache "user-api/app/gateways/repository/cache/redis"
	database "user-api/app/gateways/repository/db/postgres"
	elastic "user-api/app/gateways/repository/search/elasticsearch"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/go-chi/chi"
	"github.com/redis/go-redis/v9"
	httpSwagger "github.com/swaggo/http-swagger"
	"gorm.io/gorm"
)

// newHandler creates a new HTTP handler.
// @title User API
// @version 2.0
// @BasePath /
// @schemes http
// @host localhost:3000
// @produce json
// @consumes json

// @summary This is a medical catalog API.
// @description This is a medical catalog API.
func newHandler(ctx context.Context, db *gorm.DB, rdb *redis.Client, es *elasticsearch.Client) (http.Handler, error) {
	r := chi.NewRouter()

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	r.Route("/docs/v1/user-api", func(r chi.Router) {
		r.Get("/swagger", func(w http.ResponseWriter, r *http.Request) {
			http.Redirect(w, r, "swagger/index.html", http.StatusMovedPermanently)
		})
		r.Get("/swagger/*", httpSwagger.Handler())
	})

	repoPostgres := database.NewRepository(db)
	repoRedis := cache.NewRepository(rdb)
	repoElastic := elastic.NewRepository(es)
	usecase := usecases.NewUsecase(repoPostgres, repoRedis, repoElastic)
	handler := handlers.NewHandler(usecase)

	r.Route("/api/v1/user-api/users", func(r chi.Router) {
		r.Post("/", rest.Handle(handler.CreateUser))
		r.Get("/{user-id}", rest.Handle(handler.GetUserByID))
		r.Get("/", rest.Handle(handler.GetAllUsers))
		r.Put("/{user-id}", rest.Handle(handler.UpdateUser))
		r.Delete("/{user-id}", rest.Handle(handler.DeleteUser))
	})

	return r, nil
}
