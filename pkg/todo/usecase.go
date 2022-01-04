package todo

type UseCase interface {
}

type useCase struct {
	repository *Repository
}

func NewUseCase(repository *Repository) UseCase {
	return &useCase{
		repository: repository,
	}
}
