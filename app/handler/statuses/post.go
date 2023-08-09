package status

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"yatter-backend-go/app/handler/auth"
	"yatter-backend-go/app/handler/httperror"
)

// AddRequest
// Request body for `POST /v1/accounts`
type AddRequest struct {
	Status   string
	MediaIDs []int
}

// Post
// Handle request for `POST /v1/statuses`
func (h *handler) Post(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req AddRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		httperror.BadRequest(w, err)
		return
	}

	log.Printf("req: %+v", req)

	account := auth.AccountOf(r)
	err := account.SetStatus(req.Status)
	if err != nil {
		httperror.InternalServerError(w, errors.New("failed to set status"))
		return
	}

	log.Printf("account: %+v", account)

	accountRepo := h.app.Dao.Account() // domain/repository の取得

	err = accountRepo.PostStatus(ctx, account)
	if err != nil {
		httperror.InternalServerError(w, errors.New("failed to post status"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(account); err != nil {
		httperror.InternalServerError(w, err)
		return
	}
}
