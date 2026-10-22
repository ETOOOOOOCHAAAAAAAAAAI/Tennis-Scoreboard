package main

import (
	"Tennis-ScoreboardRe/internal/handler"
	"Tennis-ScoreboardRe/internal/repository"
	"Tennis-ScoreboardRe/internal/service"
	"Tennis-ScoreboardRe/internal/storage"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("file .env not found")
	}
	dbURL := os.Getenv("DATABASE_URL")
	db, err := storage.NewDB(dbURL)
	if err != nil {
		log.Fatal("connecting to db failed", err)
	}
	defer db.Close()
	playerRepo := repository.NewPlayerRepository(db)
	matchRepo := repository.NewMatchRepository(db)
	ongMatchRepo := repository.NewOnMatchRepository()

	scoreCalc := service.NewMatchScoreCalculationService()
	matchService := service.NewMatchService(playerRepo, ongMatchRepo, matchRepo, scoreCalc)

	appHandler := handler.NewHandler(matchService)

	http.HandleFunc("/", appHandler.IndexHandler)
	http.HandleFunc("/new-match", appHandler.NewMatchHandler)
	http.HandleFunc("/match-score", appHandler.NewMatchScoreHandler)
	http.HandleFunc("/matches", appHandler.NewMatchesListHandler)

	fileServer := http.FileServer(http.Dir("ui/static"))
	http.Handle("/css/", fileServer)
	http.Handle("/images/", fileServer)
	http.Handle("/js/", fileServer)
	http.Handle("/static/", http.StripPrefix("/static/", fileServer))

	log.Println("Server started in http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
