package admin

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/tamaaa13/Queue/internal/model/admin"
	"github.com/tamaaa13/Queue/internal/model/public"
)

// GetAllUsers mengambil semua data user dari tabel users
func (r *repository) GetAllUsers() ([]*admin.User, error) {
	query := `SELECT id, username, password_hash, role, category, created_at FROM users`
	rows, err := r.db.Query(query)
	if err != nil {
		log.Println("Error saat query GetAllUsers:", err)
		return nil, err
	}
	defer rows.Close()

	var users []*admin.User
	for rows.Next() {
		user := &admin.User{}
		err := rows.Scan(&user.ID, &user.Username, &user.PasswordHash, &user.Role, &user.Category, &user.CreatedAt)
		if err != nil {
			log.Println("Error saat scan user:", err)
			continue // Lewati user ini dan lanjutkan
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		log.Println("Error saat iterasi rows:", err)
		return nil, err
	}

	return users, nil
}

// FindUserByUsername mencari user berdasarkan username
func (r *repository) FindUserByUsername(username string) (*admin.User, error) {
	user := &admin.User{}
	query := `SELECT id, username, password_hash, role, category, created_at 
	          FROM users WHERE username = ?`
	err := r.db.QueryRow(query, username).Scan(&user.ID, &user.Username, &user.PasswordHash, &user.Role, &user.Category, &user.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}
		log.Println("Error: ", err)
		return nil, err
	}
	return user, nil
}

// CreateUser menambahkan user baru
func (r *repository) CreateUser(user *admin.User) error {
	_, err := r.db.Exec(`INSERT INTO users (username, password_hash, role, category) VALUES (?, ?, ?, ?)`,
		user.Username, user.PasswordHash, user.Role, user.Category)
	if err != nil {
		return err
	}
	return nil
}

// UpdateUser memperbarui data user di database
func (r *repository) UpdateUser(user *admin.User) error {
	_, err := r.db.Exec(`UPDATE users SET username = ?, password_hash = ?, role = ?, category = ? WHERE id = ?`,
		user.Username, user.PasswordHash, user.Role, user.Category, user.ID)
	if err != nil {
		log.Println("Error saat update user:", err)
		return err
	}

	return nil
}

// DeleteUser menghapus data user di database
func (r *repository) DeleteUser(userID string) error {
	_, err := r.db.Exec(`DELETE FROM users WHERE id = ?`, userID)
	if err != nil {
		log.Println("Error saat delete user:", err)
		return err
	}
	return nil
}

// GetTodayQueues mengambil semua antrian yang ada hari ini
func (r *repository) GetTodayQueues() ([]*public.QueueModel, error) {
	rows, err := r.db.Query(`
		SELECT id, category, number, status, created_at 
		FROM queue 
		WHERE DATE(created_at) = CURDATE()
		ORDER BY created_at ASC`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var queues []*public.QueueModel
	for rows.Next() {
		var q public.QueueModel
		if err := rows.Scan(&q.ID, &q.Category, &q.Number, &q.Status, &q.CreatedAt); err != nil {
			return nil, err
		}
		queues = append(queues, &q)
	}
	return queues, nil
}

func (r *repository) GetTodayQueuesFiltered(category string) ([]*public.QueueModel, error) {
	var rows *sql.Rows
	var err error

	if category != "" {
		rows, err = r.db.Query(`
			SELECT id, category, number, status, created_at 
			FROM queue 
			WHERE DATE(created_at) = CURDATE() AND category = ?
			ORDER BY created_at ASC`, category)
	} else {
		rows, err = r.db.Query(`
			SELECT id, category, number, status, created_at 
			FROM queue 
			WHERE DATE(created_at) = CURDATE()
			ORDER BY created_at ASC`)
	}

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var queues []*public.QueueModel
	for rows.Next() {
		var q public.QueueModel
		if err := rows.Scan(&q.ID, &q.Category, &q.Number, &q.Status, &q.CreatedAt); err != nil {
			return nil, err
		}
		queues = append(queues, &q)
	}
	return queues, nil
}

// GetNextQueue mendapatkan antrian selanjutnya berdasarkan kategori
func (r *repository) GetNextQueue(category string, calledBy string) (*public.QueueModel, error) {
	row := r.db.QueryRow(`
		SELECT id, category, number, status, created_at
		FROM queue
		WHERE category = ? AND status = 'waiting' AND queue_date = CURDATE()
		ORDER BY created_at ASC LIMIT 1`, category)

	var q public.QueueModel
	if err := row.Scan(&q.ID, &q.Category, &q.Number, &q.Status, &q.CreatedAt); err != nil {
		return nil, err
	}

	// Simpan admin yang memanggil antrian
	_, err := r.db.Exec(`UPDATE queue SET status = 'called', called_by = ? WHERE id = ?`, calledBy, q.ID)
	if err != nil {
		return nil, err
	}

	q.Status = "called"
	q.CalledBy = calledBy
	return &q, nil
}
