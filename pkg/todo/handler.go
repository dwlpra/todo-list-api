package todo

type handler struct {
	useCase *UseCase
}

func NewHandler(useCase *UseCase) *handler {
	return &handler{
		useCase: useCase,
	}
}

func (h *handler) Route() {

}

// func (h *handler) GetAll(c echo.Context) error {
// 	return c.String(http.StatusOK, "/users/1/files/*")
// }
