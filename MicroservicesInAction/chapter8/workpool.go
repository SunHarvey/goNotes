package main

import (
	"fmt"
	"reflect"
	"time"
)

type Task struct {
	Num int
}

type Job struct {
	Task Task
}
var (
	MaxWorker = 5
	JobQueue chan Job
)

type Worker struct {
	id int
	WorkerPool chan chan Job
	JobChannel chan Job
	exit chan bool
}

func NewWorker(workerPool chan chan Job, id int) Worker {
	fmt.Printf("new a worker(%d)\n", id)
	return Worker {
		id: id,
		WorkerPool: workerPool,
		JobChannel: make(chan Job),
		exit: make(chan bool),
	}
}


func (w Worker) Start() {
	go func() {
		for {
			w.WorkerPool <- w.JobChannel
			fmt.Println("register private JobChannel to public WorkerPool",w)
			select {
				case job := <- w.JobChannel:
					fmt.Println("get a job from private w.JobChannel")
					fmt.Println(job)
					time.Sleep(5 * time.Second)
				case <- w.exit:
					fmt.Println("worker exit",w)
					return
			}
		}
	}()
}


func (w Worker) Stop() {
	go func() {
		w.exit <- true
	}()
}


type Scheduler struct {
	WorkerPool chan chan Job
	MaxWorkers int
	Workers []*Worker
}


func NewScheduler(maxWorkers int) *Scheduler {
	pool := make(chan chan Job, maxWorkers)
	return &Scheduler{WorkerPool: pool, MaxWorkers: maxWorkers}
}


func (s *Scheduler) Create() {
	workers := make([]*Worker, s.MaxWorkers)
	for i := 0; i<s.MaxWorkers; i++ {
		worker := NewWorker(s.WorkerPool, i)
		worker.Start()
		workers[i] = &worker
	}
	s.Workers  = workers
	go s.schedule()
}


func (s *Scheduler) Shutdown() {
	workers := s.Workers
	for _,w := range workers {
		w.Stop()
	}
	time.Sleep(time.Second)
	close(s.WorkerPool)
}


func (s *Scheduler) schedule() {
	for {
		select {
			case job := <- JobQueue:
				fmt.Println("get a job from JobQueue")
				go func(job Job) {
					//
					jobChannel := <- s.WorkerPool
					fmt.Println("get a private jobChannel from public s.WorkerPool", reflect.TypeOf(jobChannel))
					jobChannel <- job
					fmt.Println("worker's private jobChannel and job")
				}(job)
		}
	}
}

func main() {
	JobQueue = make(chan Job, 5)
	scheduler := NewScheduler(MaxWorker)
	scheduler.Create()
	time.Sleep(1 * time.Second)
	go createJobQueue()
	time.Sleep(100 * time.Second)
	scheduler.Shutdown()
	time.Sleep(10 * time.Second)
}


func createJobQueue() {
	for i := 0; i<30; i++ {
		task := Task{Num: 1}
		job := Job{Task: task}
		JobQueue <- job
		fmt.Println("JobQueue add one job")
		time.Sleep(1 * time.Second)
	}
}
