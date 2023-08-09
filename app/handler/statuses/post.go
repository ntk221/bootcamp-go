package status

import (
	"encoding/json"
	"log"
	"net/http"
	"yatter-backend-go/app/domain/object"

	"yatter-backend-go/app/handler/auth"
	"yatter-backend-go/app/handler/httperror"
)

// PostRequest
// Request body for `POST /v1/statuses`
type PostRequest struct {
	StatusContent string
	MediaIDs      []int
}

// PostResponse
// Response body for `POST /v1/statuses`
type PostResponse struct {
	ID              object.StatusID
	Account         object.Account
	Content         string
	CreateAt        object.DateTime
	MediaAttachment []string // 仮実装
}

// Post
// Handle request for `POST /v1/statuses`
func (h *handler) Post(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req PostRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		httperror.BadRequest(w, err)
		return
	}

	log.Printf("req: %+v", req)

	account := auth.AccountOf(r)
	accountID := account.ID
	content := req.StatusContent
	status, err := object.CreateStatus(accountID, content)
	if err != nil {
		httperror.InternalServerError(w, err)
	}

	statusRepo := h.app.Dao.Status()

	posted, err := statusRepo.PostStatus(ctx, status)
	if err != nil {
		httperror.InternalServerError(w, err)
	}

	response := PostResponse{
		ID:              posted.ID,
		Account:         *account,
		Content:         posted.Content,
		CreateAt:        posted.CreateAt,
		MediaAttachment: []string{}, // 仮実装
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		httperror.InternalServerError(w, err)
		return
	}
}
