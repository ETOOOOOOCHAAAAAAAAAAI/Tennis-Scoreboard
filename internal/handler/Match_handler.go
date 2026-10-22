package handler

import (
	"Tennis-ScoreboardRe/internal/service"
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/google/uuid"
)

type Handler struct {
	matchService *service.MatchService
}

func NewHandler(ms *service.MatchService) *Handler {
	return &Handler{matchService: ms}
}

func (h *Handler) NewMatchHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		temp, err := template.ParseFiles("ui/templates/new-match.html")
		if err != nil {
			http.Error(w, "Failed to load template", http.StatusInternalServerError)
			return
		}
		temp.Execute(w, nil)
		return
	}
	if r.Method == http.MethodPost {
		player1Name := r.FormValue("player1")
		player2Name := r.FormValue("player2")
		if player1Name == "" || player2Name == "" {
			http.Error(w, "Player names cannot be empty", http.StatusBadRequest)
			return
		}
		matchID, err := h.matchService.CreateMatch(player1Name, player2Name)
		if err != nil {
			http.Error(w, "Failed to create match", http.StatusInternalServerError)
			return
		}
		redirectURL := fmt.Sprintf("/match-score?uuid=%s", matchID.String())
		http.Redirect(w, r, redirectURL, http.StatusSeeOther)
		return
	}
	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)

}

func (h *Handler) NewMatchScoreHandler(w http.ResponseWriter, r *http.Request) {
	uuidStr := r.URL.Query().Get("uuid")
	if uuidStr == "" {
		http.Error(w, "Missing match UUID", http.StatusBadRequest)
		return
	}
	matchID, err := uuid.Parse(uuidStr)
	if err != nil {
		http.Error(w, "Invalid UUID format", http.StatusBadRequest)
		return
	}
	if r.Method == http.MethodGet {
		temp, err := template.ParseFiles("ui/templates/match-score.html")
		if err != nil {
			http.Error(w, "Failed to load template", http.StatusInternalServerError)
			return
		}
		match, err := h.matchService.GetMatch(matchID)
		if err != nil {
			http.Error(w, "Match not found", http.StatusNotFound)
			return
		}
		temp.Execute(w, match)
		return
	}
	if r.Method == http.MethodPost {
		winnerID := r.FormValue("winner")
		if winnerID == "" {
			http.Error(w, "Winner ID is missing", http.StatusBadRequest)
			return
		}
		integerWinnerID, err := strconv.ParseInt(winnerID, 10, 64)
		if err != nil {
			http.Error(w, "Invalid Winner ID format", http.StatusBadRequest)
			return
		}
		err = h.matchService.ProcessPoint(matchID, integerWinnerID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		_, err = h.matchService.GetMatch(matchID)
		if err != nil {
			http.Redirect(w, r, "/matches", http.StatusSeeOther)
			return
		}

		redirectURL := fmt.Sprintf("/match-score?uuid=%s", matchID.String())
		http.Redirect(w, r, redirectURL, http.StatusSeeOther)
		return
	}
	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
}

func (h *Handler) NewMatchesListHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		pageStr := r.URL.Query().Get("page")
		page := 1
		if pageStr != "" {
			var err error
			page, err = strconv.Atoi(pageStr)
			if err != nil || page < 1 {
				page = 1
			}
		}
		matches, err := h.matchService.GetFinishedMatch(page)
		if err != nil {
			http.Error(w, "Failed to retrieve matches", http.StatusInternalServerError)
			return
		}
		temp, err := template.ParseFiles("ui/templates/matches.html")
		if err != nil {
			http.Error(w, "Failed to load template", http.StatusInternalServerError)
			return
		}
		temp.Execute(w, matches)
		return
	}
	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)

}

func (h *Handler) IndexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	if r.Method == http.MethodGet {
		temp, err := template.ParseFiles("ui/templates/index.html")
		if err != nil {
			http.Error(w, "Ошибка загрузки главной страницы", http.StatusInternalServerError)
			return
		}
		temp.Execute(w, nil)
		return
	}
	http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
}
