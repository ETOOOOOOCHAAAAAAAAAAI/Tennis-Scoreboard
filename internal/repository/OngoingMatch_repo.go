package repository

import (
	"Tennis-ScoreboardRe/internal/model"
	"sync"

	"github.com/google/uuid"
)

type OngMatchRepository struct {
	matches map[uuid.UUID]model.Matches
	mu      sync.RWMutex
}

func NewOnMatchRepository() *OngMatchRepository {
	return &OngMatchRepository{matches: make(map[uuid.UUID]model.Matches)}
}

func (o *OngMatchRepository) Save(id uuid.UUID, match model.Matches) {
	o.mu.Lock()
	defer o.mu.Unlock()
	o.matches[id] = match
}

func (o *OngMatchRepository) Get(id uuid.UUID) (model.Matches, bool) {
	o.mu.RLock()
	defer o.mu.RUnlock()
	if _, exists := o.matches[id]; !exists {
		return model.Matches{}, false
	}
	return o.matches[id], true
}

func (o *OngMatchRepository) Delete(id uuid.UUID) {
	o.mu.Lock()
	defer o.mu.Unlock()
	delete(o.matches, id)
}
