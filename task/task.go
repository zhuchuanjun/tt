package main

import "sync"

type Scheduler struct {
	Tasks []Task
	Mu    sync.Mutex
}

type Task struct {
	Id    int
	Func  func()
	State int
	Quit  chan struct{}
}

func (s *Scheduler) Add(t Task) {
	s.Mu.Lock()
	defer s.Mu.Unlock()
	s.Tasks = append(s.Tasks, t)
}

func (s *Scheduler) Cancel(id int) {
	for _, task := range s.Tasks {
		if task.Id == id {
			task.Quit <- struct{}{}
		}
	}
}

func (s *Scheduler) Run(num int) {

}

func (s *Scheduler) GetState(id int) int {
	for _, task := range s.Tasks {
		if task.Id == id {
			return task.State
		}
	}
	return 0
}
