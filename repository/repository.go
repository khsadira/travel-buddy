package repository

import "github.com/projet-x/metier"

type Question struct{}

func NewQuestion() Question {
	return Question{}
}

func (q Question) Ask(metier.Question) []metier.Response {
	// do relational stuff with chat gpt so that we can return a list of responses
	return []metier.Response{{Text: "I don't know"}}
}
