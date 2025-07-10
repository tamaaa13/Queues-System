package public

import (
	"time"

	"github.com/tamaaa13/Queue/internal/configs"
	"github.com/tamaaa13/Queue/internal/model/public"
)

type publicRepo interface {
	CountToday(category string) (int, error)
	InsertQueue(category, number string, createdAt time.Time) (int64, error)
	GetLastCalledQueueByAdmin(admin string) (public.QueueModel, error)
}

type service struct {
	cfg        *configs.Config
	publicRepo publicRepo
}

func NewService(cfg *configs.Config, publicRepo publicRepo) *service {
	return &service{
		cfg:        cfg,
		publicRepo: publicRepo,
	}
}
