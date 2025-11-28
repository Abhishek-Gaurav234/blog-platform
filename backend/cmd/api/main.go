package main

import (
	"blog-platform/internal/handler"
	"blog-platform/internal/models"
	"blog-platform/internal/service"
	"blog-platform/pkg/proxy"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize database (Singleton pattern)
	db := models.GetDatabaseInstance()

	// Initialize repositories
	realPostRepo := models.NewPostRepository(db.DB)
	
	// Wrap repository with Caching Proxy (Proxy pattern)
	// Cache up to 100 posts with 5-minute TTL
	postRepo := proxy.NewPostRepositoryCachingProxy(realPostRepo, 100, 5*time.Minute)
	log.Println("✅ Caching Proxy enabled: Max 100 posts, 5min TTL")

	// Initialize services
	commandService := service.NewCommandService(postRepo)
	queryService := service.NewQueryService(postRepo)
	contentFactory := &service.ContentFactory{}
	postService := service.NewPostService()
	searchService := service.NewSearchService(postRepo)

	// Register observers
	postService.Subscribe(&service.SearchIndexObserver{})
	postService.Subscribe(&service.NotificationObserver{})
	postService.Subscribe(&service.AnalyticsObserver{})

	// Initialize handlers
	postHandler := handler.NewPostHandler(
		commandService,
		queryService,
		contentFactory,
		postService,
		searchService,
	)

	// Set up Gin router
	router := gin.Default()

	// CORS middleware for React frontend
	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	// API routes
	api := router.Group("/api/v1")
	{
		// Post routes
		posts := api.Group("/posts")
		{
			posts.POST("", postHandler.CreatePost)
			posts.GET("", postHandler.ListPosts)
			posts.GET("/:id", postHandler.GetPost)
			posts.PUT("/:id", postHandler.UpdatePost)
			posts.DELETE("/:id", postHandler.DeletePost)
			posts.GET("/search", postHandler.SearchPosts)
		}

		// Cache statistics endpoint (demonstrates Proxy pattern benefits)
		api.GET("/cache/stats", func(c *gin.Context) {
			stats := postRepo.GetStatistics()
			c.JSON(200, gin.H{
				"cache_statistics": stats,
				"description":      "Proxy Pattern: Transparent caching layer for post repository",
			})
		})
	}

	// Start the server
	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
