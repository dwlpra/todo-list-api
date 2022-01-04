package activity

import resp "github.com/dwlpra/todo-list-api/libs/response"

type UseCase interface {
	GetAll() chan resp.Result
	GetOne(id string) chan resp.Result
	Create(req RequestBody) chan resp.Result
	Delete(id string) chan resp.Result
	Update(req RequestBody, id string) chan resp.Result
}

type useCase struct {
	repository *Repository
}

func NewUseCase(repository *Repository) UseCase {
	return &useCase{
		repository: repository,
	}
}

func (uc *useCase) GetAll() chan resp.Result {
	output := make(chan resp.Result)
	go func() {
		repo := (*uc.repository)
		result := <-repo.GetAll()
		if result.Err != nil {
			output <- resp.Result{Err: result.Err}
			return
		}
		output <- resp.Result{Data: result.Data}
	}()

	return output
}

func (uc *useCase) GetOne(id string) chan resp.Result {
	output := make(chan resp.Result)
	go func() {
		repo := (*uc.repository)
		result := <-repo.GetOne(id)
		if result.Err != nil {
			output <- resp.Result{Err: result.Err}
			return
		}
		output <- resp.Result{Data: result.Data}
	}()

	return output
}

func (uc *useCase) Create(req RequestBody) chan resp.Result {
	output := make(chan resp.Result)
	go func() {
		repo := (*uc.repository)
		result := <-repo.Create(req)
		if result.Err != nil {
			output <- resp.Result{Err: result.Err}
			return
		}
		output <- resp.Result{Data: result.Data}
	}()

	return output
}

func (uc *useCase) Delete(id string) chan resp.Result {
	output := make(chan resp.Result)
	go func() {
		repo := (*uc.repository)
		result := <-repo.Delete(id)
		if result.Err != nil {
			output <- resp.Result{Err: result.Err}
			return
		}
		output <- resp.Result{Data: result.Data}
	}()

	return output
}

func (uc *useCase) Update(req RequestBody, id string) chan resp.Result {
	output := make(chan resp.Result)
	go func() {
		repo := (*uc.repository)
		result := <-repo.Update(req, id)
		if result.Err != nil {
			output <- resp.Result{Err: result.Err}
			return
		}
		output <- resp.Result{Data: result.Data}
	}()

	return output
}
