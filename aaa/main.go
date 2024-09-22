package main

import (
	"fmt"
	"sync"
	"time"
)

const (
	statePending = iota
	stateRunning
	stateCompleted
	stateCanceled
)

type Task struct {
	Id     int
	Action func()
	State  int
}

func NewSchedule(maxWorker int) *Schedule {
	return &Schedule{
		Mu:          sync.Mutex{},
		Tasks:       make(map[int]*Task),
		TaskQueue:   make(chan int, 100),
		MaxWorker:   maxWorker,
		WorkerQueue: make(chan struct{}, maxWorker),
		CancelQueue: make(chan int, 100),
	}
}

type Schedule struct {
	Mu          sync.Mutex
	Tasks       map[int]*Task
	TaskQueue   chan int
	MaxWorker   int
	WorkerQueue chan struct{}
	CancelQueue chan int
}

func (s *Schedule) AddTask(id int, action func()) {
	s.Mu.Lock()
	defer s.Mu.Unlock()
	fmt.Printf("task add %d\n", id)

	task := &Task{
		Id:     id,
		Action: action,
		State:  statePending,
	}
	s.Tasks[id] = task
	s.TaskQueue <- id
}

func (s *Schedule) Run() {
	for {
		select {
		case id := <-s.TaskQueue:
			s.WorkerQueue <- struct{}{}
			if task, ok := s.Tasks[id]; ok {
				go s.executeWorker(task)
			}
		case id := <-s.CancelQueue:
			s.Mu.Lock()
			if task, ok := s.Tasks[id]; ok && task.State == statePending {
				task.State = stateCanceled
			}
			s.Mu.Unlock()
		}
	}
}

func (s *Schedule) executeWorker(task *Task) {
	defer func() {
		<-s.WorkerQueue
	}()
	s.Mu.Lock()
	if task.State == stateCanceled {
		s.Mu.Unlock()
		return
	}
	if task.State == statePending {
		task.State = stateRunning
	}
	s.Mu.Unlock()

	task.Action()

	s.Mu.Lock()
	if task.State == stateRunning {
		task.State = stateCompleted
	}
	s.Mu.Unlock()
}

func (s *Schedule) getState(id int) int {
	defer s.Mu.Unlock()
	s.Mu.Lock()
	if task, ok := s.Tasks[id]; ok {
		return task.State
	}
	return stateCanceled
}

func (s *Schedule) cancelTask(id int) {
	s.CancelQueue <- id
}

func main() {
	sch := NewSchedule(3)
	go sch.Run()

	for i := 1; i <= 5; i++ {
		id := i
		sch.AddTask(i, func() {
			fmt.Printf("task run %d\n", id)
			time.Sleep(time.Second)
		})
	}

	sch.cancelTask(2)
	sch.cancelTask(4)

	time.Sleep(2 * time.Second)

	for i := 1; i <= 5; i++ {
		state := sch.getState(i)
		fmt.Printf("task %d state %d\n", i, state)
	}

	time.Sleep(5 * time.Second)
}
