package main

import (
	"log"

	_ "github.com/mattn/go-sqlite3"
	"github.com/mihirtunnel/cleanArchitecture/internal/app/book"
	"github.com/mihirtunnel/cleanArchitecture/internal/app/handler"
	"github.com/mihirtunnel/cleanArchitecture/internal/app/server"
	"github.com/mihirtunnel/cleanArchitecture/internal/database"
)

func main() {
	db, err := database.InitializeDB("./bookset.db")
	if err != nil {
		log.Fatal("Failed to initialize database:", err)
	}
	defer db.Close()

	bookRepository := database.NewBookRepository(db)
	bookUseCase := book.NewBookController(bookRepository)

	handler := handler.NewBookHandler(bookUseCase)

	server := server.NewServer(handler)
	server.Start("4000")
}
