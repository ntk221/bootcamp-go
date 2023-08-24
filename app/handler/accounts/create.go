package accounts

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"yatter-backend-go/app/domain/object"
	"yatter-backend-go/app/handler/httperror"
)

// Request body for `POST /v1/accounts`
type AddRequest struct {
	Username string
	Password string
}

// Handle request for `POST /v1/accounts`
func (h *handler) Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req AddRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Println(err)
		httperror.BadRequest(w, err)
		return
	}

	account := new(object.Account)
	account.Username = req.Username
	if err := account.SetPassword(req.Password); err != nil {
		log.Println(err)
		httperror.InternalServerError(w, err)
		return
	}

	accountRepo := h.app.Dao.Account() // domain/repository の取得

	if err := accountRepo.CreateUser(ctx, account); err != nil {
		log.Println(err)
		// panic("TODO: エラーハンドリング")
		httperror.InternalServerError(w, errors.New("サーバー側で問題が発生しました"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(account); err != nil {
		log.Println(err)
		httperror.InternalServerError(w, err)
		return
	}
}
