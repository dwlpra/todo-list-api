package activity

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"

	resp "github.com/dwlpra/todo-list-api/libs/response"
)

type handler struct {
	useCase *UseCase
}

func NewHandler(useCase *UseCase) *handler {
	return &handler{
		useCase: useCase,
	}
}

func (h *handler) Route() {
	http.HandleFunc("/activity-groups", h.CheckMethods)
	http.HandleFunc("/activity-groups/", h.GetOne)
}

func (h *handler) CheckMethods(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "GET" {
		h.GetAll(w, r)
		return
	} else if r.Method == "POST" {
		h.Create(w, r)
		return
	}

}

func (h *handler) GetAll(w http.ResponseWriter, r *http.Request) {
	uc := (*h.useCase)
	result := <-uc.GetAll()
	if result.Err != nil {
		http.Error(w, result.Err.Error(), http.StatusInternalServerError)
		return
	}
	response := resp.RespModel{
		Status:  "sucess",
		Message: "success",
		Data:    result.Data,
	}

	res, _ := json.Marshal(response)
	w.Write([]byte(res))

}

func (h *handler) GetOne(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != "GET" {
		if r.Method == "DELETE" {
			h.Delete(w, r)
			return
		} else if r.Method == "PATCH" {
			h.Update(w, r)
			return
		}
	} else {
		if r.URL.Path == "/activity-groups/" {
			if r.Method == "GET" {
				h.GetAll(w, r)
			} else if r.Method == "POST" {
				h.Create(w, r)
			}
		} else {
			rgx := regexp.MustCompile(`\/`)
			url := rgx.Split(r.URL.Path, -1)
			id := url[len(url)-1]

			uc := (*h.useCase)
			result := <-uc.GetOne(id)
			if result.Err != nil {
				response := resp.RespModel{
					Status:  "Not Found",
					Message: fmt.Sprintf("Activity with ID %s Not Found", id),
					Data:    resp.EmptyResp{},
				}
				res, _ := json.Marshal(response)
				http.Error(w, string(res), http.StatusNotFound)
				return
			}
			response := resp.RespModel{
				Status:  "Success",
				Message: "Success",
				Data:    result.Data,
			}

			res, _ := json.Marshal(response)
			w.Write([]byte(res))
			return
		}
	}
}

func (h *handler) Create(w http.ResponseWriter, r *http.Request) {

	var req RequestBody
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if req.Email == "" {
		response := resp.RespModel{
			Status:  "Bad Request",
			Message: "email cannot be null",
			Data:    resp.EmptyResp{},
		}
		res, _ := json.Marshal(response)
		http.Error(w, string(res), http.StatusBadRequest)
		return
	}
	if req.Title == "" {
		response := resp.RespModel{
			Status:  "Bad Request",
			Message: "title cannot be null",
			Data:    resp.EmptyResp{},
		}
		res, _ := json.Marshal(response)
		http.Error(w, string(res), http.StatusBadRequest)
		return
	}

	uc := (*h.useCase)
	result := <-uc.Create(req)
	if result.Err != nil {
		http.Error(w, result.Err.Error(), http.StatusInternalServerError)
		return
	}
	response := resp.RespModel{
		Status:  "sucess",
		Message: "success",
		Data:    result.Data,
	}

	res, _ := json.Marshal(response)
	w.Write([]byte(res))

}

func (h *handler) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	rgx := regexp.MustCompile(`\/`)
	url := rgx.Split(r.URL.Path, -1)
	id := url[len(url)-1]

	uc := (*h.useCase)
	result := <-uc.Delete(id)
	if result.Err != nil {
		response := resp.RespModel{
			Status:  "Not Found",
			Message: fmt.Sprintf("Activity with ID %s Not Found", id),
			Data:    resp.EmptyResp{},
		}
		res, _ := json.Marshal(response)
		http.Error(w, string(res), http.StatusNotFound)
		return
	}
	response := resp.RespModel{
		Status:  "Success",
		Message: "Success",
		Data:    result.Data,
	}

	res, _ := json.Marshal(response)
	w.Write([]byte(res))

}

func (h *handler) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var req RequestBody
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if req.Title == "" {
		response := resp.RespModel{
			Status:  "Bad Request",
			Message: "title cannot be null",
			Data:    resp.EmptyResp{},
		}
		res, _ := json.Marshal(response)
		http.Error(w, string(res), http.StatusBadRequest)
		return
	}

	rgx := regexp.MustCompile(`\/`)
	url := rgx.Split(r.URL.Path, -1)
	id := url[len(url)-1]

	uc := (*h.useCase)
	result := <-uc.Update(req, id)
	if result.Err != nil {
		response := resp.RespModel{
			Status:  "Not Found",
			Message: fmt.Sprintf("Activity with ID %s Not Found", id),
			Data:    resp.EmptyResp{},
		}
		res, _ := json.Marshal(response)
		http.Error(w, string(res), http.StatusNotFound)
		return
	}
	response := resp.RespModel{
		Status:  "Success",
		Message: "Success",
		Data:    result.Data,
	}

	res, _ := json.Marshal(response)
	w.Write([]byte(res))

}
