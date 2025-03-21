package scheduler

import (
	"context"

	"auto-messenger/internal/domain"
	"auto-messenger/pkg/logger"

	"github.com/robfig/cron/v3"
)

type Scheduler struct {
	service domain.MessageService
	logger  logger.Logger
	cron    *cron.Cron
}

func NewScheduler(service domain.MessageService, logger logger.Logger) *Scheduler {
	return &Scheduler{
		service: service,
		logger:  logger,
		cron:    cron.New(),
	}
}

func (s *Scheduler) Start(ctx context.Context) {
	_, err := s.cron.AddFunc("0 19 * * *", func() {
		//_, err := s.cron.AddFunc("@every 30s", func() {
		s.logger.Info("(scheduler.Start): Running scheduled message job")
		if err := s.service.SendScheduledMessage(ctx); err != nil {
			s.logger.Error("(scheduler.Start):" + err.Error())
		}
	})

	if err != nil {
		s.logger.Error("(scheduler.Start):" + err.Error())
		return
	}

	s.cron.Start()
	s.logger.Info("(scheduler.Start)")
}
