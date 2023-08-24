package timelines

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"
	"yatter-backend-go/app/domain/object"
	"yatter-backend-go/app/handler/httperror"
)

type TimelineItem struct {
	ID              object.StatusID
	Account         object.Account
	Content         string
	CreateAt        object.DateTime
	MediaAttachment []string // 仮実装
}

func (h *handler) Public(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	_ = r.URL.Query().Get("only_media")
	maxID, _ := strconv.Atoi(r.URL.Query().Get("max_id"))
	sinceID, _ := strconv.Atoi(r.URL.Query().Get("since_id"))
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))

	statusRepo := h.app.Dao.Status()
	sc, err := statusRepo.GetStatusesByParams(ctx, maxID, sinceID, limit)
	if err != nil {
		httperror.InternalServerError(w, errors.New("タイムラインの取得に失敗しました"))
	}

	// Account 情報を追加してtimelineを構築する
	var timeline []TimelineItem
	accountRepo := h.app.Dao.Account()
	for _, status := range sc.Statuses {
		account, err := accountRepo.FindByID(ctx, status.AccountID)
		if err != nil {
			log.Println(err)
			httperror.InternalServerError(w, errors.New("タイムラインの取得中にサーバーでエラーが発生しました"))
		}
		timelineItem := TimelineItem{
			ID:              status.ID,
			Account:         *account,
			Content:         status.Content,
			CreateAt:        status.CreateAt,
			MediaAttachment: []string{},
		}
		timeline = append(timeline, timelineItem)
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(timeline); err != nil {
		log.Println(err)
		httperror.InternalServerError(w, errors.New("サーバー側でなんらかの問題が発生しました。"))
		return
	}
	return
}
