package public

import "time"

type QueueModel struct {
	ID        int       `db:"id"`
	Category  string    `db:"category"`
	Number    string    `db:"number"`
	Status    string    `db:"status"`
	CalledBy  string    `db:"called_by"` // <- ini baru
	CreatedAt time.Time `db:"created_at"`
}

type QueueResponse struct {
	Admin       string `json:"admin"`
	QueueNumber string `json:"queue_number"`
	Status      string `json:"status"`
}
