package main

import (
	"blog-platform/internal/handler"
	"blog-platform/internal/models"
	"blog-platform/internal/service"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize database (Singleton pattern)
	db := models.GetDatabaseInstance()

	// Initialize repositories
	postRepo := models.NewPostRepository(db.DB)

	// Initialize services
	commandService := service.NewCommandService(postRepo)
	queryService := service.NewQueryService(postRepo)
	contentFactory := &service.ContentFactory{}
	postService := service.NewPostService()

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
	}

	// Start the server
	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
