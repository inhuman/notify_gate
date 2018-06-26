package workerpool

import (
	"sync"
)

// Task interface used for executing by worker pool
type Task interface {
	Execute()
}

// Pool is used for manage workers
type Pool struct {
	mu    sync.Mutex
	size  int
	tasks chan Task
	kill  chan struct{}
	wg    sync.WaitGroup
}

// NewPool is used for create worker pool by giving size, and initialize it
func NewPool(size int) *Pool {
	pool := &Pool{
		tasks: make(chan Task, 128),
		kill:  make(chan struct{}),
	}
	pool.Resize(size)
	return pool
}

func (p *Pool) worker() {
	defer p.wg.Done()
	for {
		select {
		case task, ok := <-p.tasks:
			if !ok {
				return
			}
			task.Execute()
		case <-p.kill:
			return
		}
	}
}

// Resize is used for resize worker pool
func (p *Pool) Resize(n int) {
	p.mu.Lock()
	defer p.mu.Unlock()
	for p.size < n {
		p.size++
		p.wg.Add(1)
		go p.worker()
	}
	for p.size > n {
		p.size--
		p.kill <- struct{}{}
	}
}

// Close is used for closing task channel
func (p *Pool) Close() {
	close(p.tasks)
}

// Wait is used for waiting until all workers done
func (p *Pool) Wait() {
	p.wg.Wait()
}

// Exec is uses for executing Task
func (p *Pool) Exec(task Task) {
	p.tasks <- task
}
