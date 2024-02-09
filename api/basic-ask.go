package api

import (
	"encoding/json"
	"net/http"

	"github.com/projet-x/metier"
	"github.com/projet-x/repository"
)

func HandleBasicAsk(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		handleBasicAskPost(rw, r, getNewServices())
	default:
		rw.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func handleBasicAskPost(rw http.ResponseWriter, r *http.Request, service metier.Service) {
	var question metier.Question
	if err := json.NewDecoder(r.Body).Decode(&question); err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	resp := service.Ask(question)
	respJSON, err := json.Marshal(resp)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	rw.Write(respJSON)
}

func getNewServices() metier.Service {
	repo := repository.NewQuestion()
	return metier.NewService(repo)
}
