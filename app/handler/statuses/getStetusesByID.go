package status

import (
	"database/sql"
	"encoding/json"
	"errors"
	"github.com/go-chi/chi"
	"log"
	"net/http"
	"strconv"
	"yatter-backend-go/app/domain/object"
	"yatter-backend-go/app/handler/httperror"
)

type GetResponse struct {
	ID              object.StatusID
	Account         object.Account
	Content         string
	CreateAt        object.DateTime
	MediaAttachment []string
}

func (h *handler) GetStatusByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	statusID, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		log.Println(err)
		httperror.BadRequest(w, errors.New("URLパラメータが適切な形式ではありません"))
		return
	}

	status := new(object.Status)
	status.ID = object.StatusID(statusID)

	statusRepo := h.app.Dao.Status()

	found, err := statusRepo.FindStatusByID(ctx, status.ID)
	if err != nil {
		log.Println(err)
		if errors.Is(err, sql.ErrNoRows) {
			httperror.BadRequest(w, errors.New("指定されたidのstatusは見つかりませんでした"))
			return
		}
		httperror.InternalServerError(w, errors.New("サーバー側でなんらかの問題が発生しました"))
		return
	}

	accountRepo := h.app.Dao.Account()
	account, err := accountRepo.FindByID(ctx, found.AccountID)
	if err != nil {
		log.Println(err)
		httperror.InternalServerError(w, errors.New("指定されたidのstatusにはユーザーが見つかりませんでした"))
		return
	}

	response := GetResponse{
		ID:              found.ID,
		Account:         *account,
		Content:         found.Content,
		CreateAt:        found.CreateAt,
		MediaAttachment: []string{},
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Println(err)
		httperror.InternalServerError(w, errors.New("サーバー側でなんらかの問題が発生しました。"))
		return
	}

}
