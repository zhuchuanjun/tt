package main

import (
	"sync"
)

type TaskFunc func() error

type Task struct {
	id     int
	fn     TaskFunc
	status string
}

type Scheduler struct {
	tasks       map[int]*Task
	taskChan    chan *Task
	cancelChan  chan int
	statusChan  chan int
	resultChan  chan string
	mu          sync.Mutex
	wg          sync.WaitGroup
	maxParallel int
}

func NewScheduler(maxParallel int) *Scheduler {
	return &Scheduler{
		tasks:       make(map[int]*Task),
		taskChan:    make(chan *Task),
		cancelChan:  make(chan int),
		statusChan:  make(chan int),
		resultChan:  make(chan string),
		maxParallel: maxParallel,
	}
}

func (s *Scheduler) AddTask(id int, fn TaskFunc) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.tasks[id] = &Task{id: id, fn: fn, status: "pending"}
}

func (s *Scheduler) CancelTask(id int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if task, exists := s.tasks[id]; exists && task.status == "pending" {
		task.status = "cancelled"
	}
}

func (s *Scheduler) GetTaskStatus(id int) string {
	s.statusChan <- id
	return <-s.resultChan
}

func (s *Scheduler) Run() {
	for i := 0; i < s.maxParallel; i++ {
		s.wg.Add(1)
		go s.worker()
	}

	go func() {
		for _, task := range s.tasks {
			if task.status == "pending" {
				s.taskChan <- task
			}
		}
	}()

	s.wg.Wait()
}

func (s *Scheduler) worker() {
	defer s.wg.Done()
	for {
		select {
		case task := <-s.taskChan:
			s.executeTask(task)
		case id := <-s.cancelChan:
			s.cancelTask(id)
		case id := <-s.statusChan:
			s.queryStatus(id)
		}
	}
}

func (s *Scheduler) executeTask(task *Task) {
	s.mu.Lock()
	if task.status == "cancelled" {
		s.mu.Unlock()
		return
	}
	task.status = "running"
	s.mu.Unlock()
	err := task.fn()
	s.mu.Lock()
	defer s.mu.Unlock()
	if err != nil {
		task.status = "failed"
	} else {
		task.status = "completed"
	}
}

func (s *Scheduler) cancelTask(id int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if task, exists := s.tasks[id]; exists {
		task.status = "cancelled"
	}
}

func (s *Scheduler) queryStatus(id int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if task, exists := s.tasks[id]; exists {
		s.resultChan <- task.status
	} else {
		s.resultChan <- "not found"
	}
}
