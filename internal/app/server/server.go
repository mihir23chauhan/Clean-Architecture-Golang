package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/mihirtunnel/cleanArchitecture/internal/app/handler"
)

type Server struct {
	handler *handler.BookHandler
}

func NewServer(handler *handler.BookHandler) *Server {
	return &Server{handler: handler}
}

func (s *Server) Start(port string) {

	router := gin.Default()
	router.GET("/books", s.handler.GetAllBooks)
	router.GET("/books/:id", s.handler.GetBookByID)
	router.POST("/books", s.handler.AddBook)
	router.PUT("/books/:id", s.handler.UpdateBook)
	router.DELETE("/books/:id", s.handler.DeleteBook)

	fmt.Printf("Server is running on http://localhost%s\n", port)
	router.Run(":" + port)
}
