package metier

type Question struct {
	Text string `json:"text" validate:"required`
	Tag  string `json:"tag,omitempty""`
}

type Response struct {
	Text string `json:"text"`
}

type QuestionRepository interface {
	Ask(question Question) []Response
}

type Service struct {
	repository QuestionRepository
}

func NewService(repository QuestionRepository) Service {
	return Service{repository: repository}
}

func (s Service) Ask(question Question) []Response {
	return s.repository.Ask(question)
}
