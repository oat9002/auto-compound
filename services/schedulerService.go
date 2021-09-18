package services

import "github.com/robfig/cron/v3"

type schedulerService struct {
	cron *cron.Cron
}

func NewSchedulerService(cron *cron.Cron) *schedulerService {
	service := &schedulerService{cron: cron}

	return service
}

func (s *schedulerService) AddFunc(cron string, action func()) (cron.EntryID, error) {
	return s.cron.AddFunc(cron, action)
}

func (s *schedulerService) RunAsync() {
	s.cron.Start()
}

func (s *schedulerService) Run() {
	s.cron.Run()
}

func (s *schedulerService) Stop() {
	s.cron.Stop()
}

func (s *schedulerService) RemoveFunc(id cron.EntryID) {
	s.cron.Remove(id)
}
