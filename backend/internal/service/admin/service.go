package admin

import (
	"github.com/tamaaa13/Queue/internal/configs"
	"github.com/tamaaa13/Queue/internal/model/admin"
	"github.com/tamaaa13/Queue/internal/model/public"
)

type adminRepo interface {
	GetAllUsers() ([]*admin.User, error)
	FindUserByUsername(username string) (*admin.User, error)
	CreateUser(user *admin.User) error
	UpdateUser(user *admin.User) error
	DeleteUser(userID string) error
	GetTodayQueues() ([]*public.QueueModel, error)
	GetTodayQueuesFiltered(category string) ([]*public.QueueModel, error)
	GetNextQueue(category string, calledBy string) (*public.QueueModel, error)
}

type service struct {
	cfg       *configs.Config
	adminRepo adminRepo
}

func NewService(cfg *configs.Config, adminRepo adminRepo) *service {
	return &service{
		cfg:       cfg,
		adminRepo: adminRepo,
	}
}
