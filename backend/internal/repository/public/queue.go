package public

import (
	"time"

	"github.com/tamaaa13/Queue/internal/model/public"
)

// CountToday menghitung jumlah antrian berdasarkan kategori dan hari ini
func (r *repository) CountToday(category string) (int, error) {
	var count int
	err := r.db.QueryRow(`
		SELECT COUNT(*) FROM queue
		WHERE category = ? AND queue_date = CURDATE()`, category).Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}

// InsertQueue memasukkan antrian ke dalam database
func (r *repository) InsertQueue(category, number string, createdAt time.Time) (int64, error) {
	result, err := r.db.Exec(`
		INSERT INTO queue (category, number, status, created_at) 
		VALUES (?, ?, 'waiting', ?)`, category, number, createdAt)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

func (r *repository) GetLastCalledQueueByAdmin(admin string) (public.QueueModel, error) {
	var q public.QueueModel
	row := r.db.QueryRow(`
        SELECT id, category, number, status, called_by, created_at
        FROM queue
        WHERE called_by = ?
        ORDER BY created_at DESC
        LIMIT 1
    `, admin)

	err := row.Scan(&q.ID, &q.Category, &q.Number, &q.Status, &q.CalledBy, &q.CreatedAt)
	if err != nil {
		return public.QueueModel{}, err
	}
	return q, nil
}
