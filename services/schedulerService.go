package services

import "github.com/robfig/cron/v3"

type SchedulerService struct {
	cron *cron.Cron
}

func NewSchedulerService(cron *cron.Cron) *SchedulerService {
	service := &SchedulerService{cron: cron}

	return service
}

func (s *SchedulerService) AddFunc(cron string, action func()) (cron.EntryID, error) {
	return s.cron.AddFunc(cron, action)
}

func (s *SchedulerService) RunAsync() {
	s.cron.Start()
}

func (s *SchedulerService) Run() {
	s.cron.Run()
}

func (s *SchedulerService) Stop() {
	s.cron.Stop()
}

func (s *SchedulerService) RemoveFunc(id cron.EntryID) {
	s.cron.Remove(id)
}
