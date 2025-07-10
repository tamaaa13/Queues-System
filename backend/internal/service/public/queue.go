package public

import (
	"errors"
	"fmt"
	"time"

	"github.com/tamaaa13/Queue/internal/model/public"
)

// CreateQueue membuat antrian baru berdasarkan kategori
func (s *service) CreateQueue(category string) (*public.QueueModel, error) {
	// Hitung jumlah antrian hari ini untuk kategori tersebut
	count, err := s.publicRepo.CountToday(category)
	if err != nil {
		return nil, err
	}

	// Generate nomor antrian (misalnya A001, B001, dst)
	number := fmt.Sprintf("%s%03d", category, count+1)
	now := time.Now()

	// Insert antrian ke database
	id, err := s.publicRepo.InsertQueue(category, number, now)
	if err != nil {
		return nil, err
	}

	// Langsung print setelah berhasil buat antrian
	if err := printQueueNumber(number); err != nil {
		fmt.Println("Gagal print:", err)
	}

	// Kembalikan data antrian yang baru
	return &public.QueueModel{
		ID:        int(id),
		Category:  category,
		Number:    number,
		Status:    "waiting",
		CreatedAt: now,
	}, nil
}

var adminAssignments = map[string]struct {
	Category string
	Loket    string
	Room     string
}{
	"adminA1": {Category: "A", Loket: "LOKET A", Room: "Ruang 1"},
	"adminA2": {Category: "A", Loket: "LOKET A", Room: "Ruang 4"},
	"adminB1": {Category: "B", Loket: "LOKET B", Room: "Ruang 2"},
	"adminB2": {Category: "B", Loket: "LOKET B", Room: "Ruang 3"},
}

func (s *service) GetQueuesForAdmins() ([]public.QueueResponse, error) {
	var responses []public.QueueResponse

	for admin := range adminAssignments {
		// Ambil antrian terakhir berdasarkan category saja,
		// karena room dan loket tidak ada di DB
		q, err := s.publicRepo.GetLastCalledQueueByAdmin(admin)
		if err != nil {
			responses = append(responses, public.QueueResponse{
				Admin:       admin,
				QueueNumber: "",
				Status:      "waiting",
			})
			continue
		}
		if q.ID == 0 {
			responses = append(responses, public.QueueResponse{
				Admin:       admin,
				QueueNumber: "",
				Status:      "waiting",
			})
			continue
		}

		// Di sini kamu bisa saja lakukan validasi tambahan kalau mau,
		// misal cocokkan dengan info Loket dan Room admin,
		// tapi karena data queue gak punya itu, skip saja

		responses = append(responses, public.QueueResponse{
			Admin:       admin,
			QueueNumber: q.Number,
			Status:      q.Status,
		})
	}

	if len(responses) == 0 {
		return nil, errors.New("no queue data found")
	}

	return responses, nil
}
