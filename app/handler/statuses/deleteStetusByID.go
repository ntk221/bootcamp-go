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

func (h *handler) DeleteStatusByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		log.Println(err)
		httperror.BadRequest(w, errors.New("idの形式が正しくありません"))
		return
	}

	statusID := object.StatusID(id)

	statusRepo := h.app.Dao.Status()
	err = statusRepo.DeleteStatusByID(ctx, statusID)
	if err != nil {
		log.Println(err)
		if errors.Is(err, sql.ErrNoRows) {
			httperror.BadRequest(w, errors.New("指定されたidのstatusが見つかりませんでした"))
			return
		}
		httperror.InternalServerError(w, errors.New("サーバー側でなんらかの不具合が発生しました"))
		return
	}

	// 成功メッセージを含むJSONレスポンスを返す
	response := map[string]string{"message": "ステータスの削除が成功しました"}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)

	return
}
