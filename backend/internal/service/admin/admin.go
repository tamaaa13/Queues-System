package admin

import (
	"database/sql"
	"fmt"

	"github.com/tamaaa13/Queue/internal/model/admin"
	"github.com/tamaaa13/Queue/internal/model/public"
)

// Get All User
func (s *service) GetAllUser() ([]*admin.User, error) {
	return s.adminRepo.GetAllUsers()
}

// CreateUser untuk membuat user baru
func (s *service) CreateUser(user *admin.User) error {
	// Cek apakah username sudah ada
	_, err := s.adminRepo.FindUserByUsername(user.Username)
	if err == nil {
		return fmt.Errorf("username already exists")
	}

	// Simpan user baru
	return s.adminRepo.CreateUser(user)
}

func (s *service) UpdateAdmin(username, newUsername, newPassword, role, category string) error {
	user, err := s.adminRepo.FindUserByUsername(username)
	if err != nil {
		return fmt.Errorf("user not found")
	}

	if newUsername != "" {
		user.Username = newUsername
	}

	if newPassword != "" {
		user.PasswordHash = newPassword
	}

	if role != "" {
		user.Role = role
	}

	if category != "" {
		user.Category = sql.NullString{String: category, Valid: true}
	} else {
		user.Category = sql.NullString{Valid: false} // agar validasi kosong juga tertangani
	}

	return s.adminRepo.UpdateUser(user)
}

func (s *service) DeleteUser(userID string) error {
	return s.adminRepo.DeleteUser(userID)
}

func (s *service) FindUserByUsername(username string) (*admin.User, error) {
	return s.adminRepo.FindUserByUsername(username)
}

// GetTodayQueues mengambil semua antrian hari ini
func (s *service) GetTodayQueues() ([]*public.QueueModel, error) {
	return s.adminRepo.GetTodayQueues()
}

func (s *service) GetTodayQueuesFiltered(category string) ([]*public.QueueModel, error) {
	return s.adminRepo.GetTodayQueuesFiltered(category)
}

// GetNextQueue mendapatkan antrian berikutnya berdasarkan kategori
func (s *service) GetNextQueue(category string, calledBy string) (*public.QueueModel, error) {
	return s.adminRepo.GetNextQueue(category, calledBy)
}
