package main

import (
	"context"
	"github.com/gin-contrib/cors"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/go-co-op/gocron"
	"go.uber.org/zap"
	"net/http"
	"os"
	"os/signal"
	newsHttpDelivery "sports_news/app/news/delivery/http"
	"sports_news/app/news/repository"
	"sports_news/app/news/usecase"
	"sports_news/app/newsconsumer"
	"sports_news/app/newsconsumer/hullcity"
	"sports_news/config"
	"sports_news/domain"
	"sports_news/infrastructure/datastore"
	"syscall"
	"time"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// Load config
	configApp := config.LoadConfig()

	// Setup logger
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	httpClient := http.Client{}

	// Setup infra
	mongoClient, err := datastore.NewDatabase(ctx, configApp.DatabaseURL)
	if err != nil {
		logger.Error("connection with mongodb failed", zap.Error(err))
		os.Exit(1)
	}
	// Setup repos.
	newsRepo := repository.NewMongoNewsRepository(mongoClient, configApp.DatabaseName)

	// Setup usecase.
	newsUc := usecase.NewNewsUsecase(newsRepo, logger)

	// Setup consumers.
	hullCityConsumer := hullcity.NewHullCityConsumer(httpClient, logger, newsRepo, configApp.HullCityConfig)

	consumers := []domain.NewsConsumer{
		hullCityConsumer,
	}

	consumerTask := newsconsumer.NewConsumerTask(consumers, configApp.QuantityToFetch)

	// Setup cron.
	s := gocron.NewScheduler(time.UTC)
	s.Every(configApp.CronFrequencyMinutes).Minute().Do(consumerTask.ConsumeNews, ctx)
	s.StartAsync()

	// Setup server.
	srv := startAndSetupServer(logger, configApp, newsUc)

	<-ctx.Done()

	// Restore default behavior on the interrupt signal and notify user of shutdown.
	stop()
	s.Stop()
	logger.Info("shutting down gracefully, press Ctrl+C again to force")

	// The context is used to inform the server it has 5 seconds to finish.
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logger.Fatal("Server forced to shutdown: ", zap.Error(err))
	}

	logger.Info("Server exiting")
}

func startAndSetupServer(logger *zap.Logger, configApp *config.Config, newsUc domain.NewsUsecase) *http.Server {
	router := gin.Default()
	router.Use(ginzap.Ginzap(logger, time.RFC3339, true))
	router.Use(ginzap.RecoveryWithZap(logger, true))
	router.Use(cors.Default())
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "I am Alive!"})
	})

	v1 := router.Group("/v1")
	newsHttpDelivery.NewNewsHandler(v1, newsUc, logger)
	srv := &http.Server{
		Addr:    configApp.HttpPort,
		Handler: router,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatal("listen:", zap.Error(err))
		}
	}()

	return srv
}
