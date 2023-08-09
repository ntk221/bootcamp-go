package object

type (
	Status struct {
		ID        StatusID  `json:"id"`
		AccountID AccountID `json:"account_id"`
		Content   string    `json:"content"`
		Posted    bool
		// The time the account was created
		CreateAt DateTime `json:"create_at,omitempty" db:"create_at"`
	}
)

func CreateStatus(accountID AccountID, content string) (Status, error) {
	// TODO: validation?

	return Status{
		AccountID: accountID,
		Content:   content,
		Posted:    false,
	}, nil
}
