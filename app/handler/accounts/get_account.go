package accounts

import (
	"encoding/json"
	"net/http"

	"yatter-backend-go/app/domain/object"
	"yatter-backend-go/app/handler/httperror"

	"github.com/go-chi/chi"
)

// GetAccount
// Handle request for `POST /v1/accounts/username`
func (h *handler) GetAccount(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	req := chi.URLParam(r, "username")

	account := new(object.Account)
	account.Username = req

	accountRepo := h.app.Dao.Account() // domain/repository の取得

	account, err := accountRepo.FindByUsername(ctx, account.Username)
	if err != nil {
		httperror.InternalServerError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(account); err != nil {
		httperror.InternalServerError(w, err)
		return
	}
}
