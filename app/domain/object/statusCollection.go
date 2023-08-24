package object

// コレクションオブジェクトは不変になるようにメソッドを設計する

type (
	StatusCollection struct {
		Statuses []Status
	}
)

func NewStatusCollection(statuses []Status) *StatusCollection {
	if statuses == nil {
		panic("nilで初期化することはできません")
	}
	return &StatusCollection{statuses}
}

// AddStatus
// 要素を追加するたびに、新しいコレクションを生成して返す。
func (sc StatusCollection) AddStatus(status Status) *StatusCollection {
	result := sc.Statuses
	result = append(result, status)

	return NewStatusCollection(result)
}
