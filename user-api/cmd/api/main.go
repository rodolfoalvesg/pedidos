package main

import (
	"context"
	"log"
	"time"
	"user-api/app/domain/entities"
	"user-api/app/gateways/http"
	"user-api/app/gateways/repository/db/postgres/migration"
	elastic "user-api/app/gateways/repository/search/elasticsearch"
	"user-api/config"

	"github.com/redis/go-redis/v9"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/avast/retry-go/v4"

	_ "user-api/docs/swagger"
)

func main() {
	ctx := context.Background()

	// Load the configuration.
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load configuration: %v", err)
	}

	// Connect to the database.
	var db *gorm.DB

	// Retry to connect to the database.
	err = retry.Do(
		func() error {
			var err error
			db, err = gorm.Open(postgres.Open(cfg.DB.DSN()), &gorm.Config{})
			if err != nil {
				log.Printf("failed to connect to the database: %v", err)
				return err
			}

			return nil
		},
		retry.Attempts(10),
		retry.Delay(2*time.Second),
	)

	if err != nil {
		log.Fatalf("failed to connect to the database: %v", err)
	}

	if err := migration.EnableUUIDPostgres(db); err != nil {
		log.Printf("failed to enable the UUID extension: %v", err)
	} else {
		log.Printf("extension UUID enabled")
	}

	// Auto-migrate the database.
	if err := db.AutoMigrate(&entities.User{}); err != nil {
		log.Fatalf("failed to auto-migrate the database: %v", err)
	}

	// Verify the connection to the database is alive.
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("failed to get the database connection: %v", err)
	}

	log.Printf("connected to the database: %s", cfg.DB.Host)

	defer sqlDB.Close()

	// Connect to the Redis database.
	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.Redis.Host + ":" + cfg.Redis.Port,
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.DB,
	})

	pong, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("failed to connect to the Redis database: %v", err)
	}

	log.Printf("connected to the Redis database: %s", pong)

	// Create a new client elastic search
	es, err := elastic.NewClient(&cfg)
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}

	err = retry.Do(
		func() error {
			_, err := es.Info()
			if err != nil {
				return err
			}

			return nil
		},
		retry.Attempts(10),
		retry.Delay(6*time.Second),
	)

	if err != nil {
		log.Fatalf("failed to connect to the elasticsearch: %v", err)
	}

	exists, err := elastic.IndexExists(ctx, cfg.ES.NameIndex, es)
	if err != nil {
		log.Fatalf("Error checking if the index exists: %s", err)
	}

	if !exists {
		err = elastic.CreateIndex(ctx, cfg.ES.NameIndex, es)
		if err != nil {
			log.Fatalf("Error creating the index: %s", err)
		}
	}

	// Start the API server.
	srv, err := http.NewServer(ctx, &cfg, db, rdb, es)
	if err != nil {
		log.Fatalf("failed to create the server: %v", err)
	}

	log.Printf("server is running on %s", srv.Addr)

	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("server failed to start: %v", err)
	}
}
