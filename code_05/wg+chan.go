package main

import "sync"

type Pool struct {
	queue chan int
	wg    *sync.WaitGroup
}

func New(size int) *Pool {
	if size <= 0 {
		size = 1
	}
	return &Pool{
		queue: make(chan int, size),
		wg:    &sync.WaitGroup{},
	}
}

// Add 方法将指定数量的工作项（delta）添加到 Pool 中。
// 如果 delta 为正数，它会向队列中添加 delta 个单位的工作量，
// 并相应地更新工作组（wg）的计数，以跟踪额外的并发任务。
// 如果 delta 为负数，它会从队列中移除 -delta 个单位的工作量，
// 这意味着这些任务已完成，等待中的工作者可以处理更多工作。
// 注意，这个方法没有实际检查 delta 的值是否大于0，因此在调用时需要确保传入正确的参数。
func (p *Pool) Add(delta int) {
	// 添加工作项到队列
	for i := 0; i < delta; i++ {
		p.queue <- 1
	}

	// 移除已完成的工作项，此循环可能永远不会执行，除非 delta 为负数
	for i := 0; i > delta; i-- {
		<-p.queue
	}

	// 更新工作小组的等待计数器
	p.wg.Add(delta)
}

func (p *Pool) Done() {
	<-p.queue
	p.wg.Done()
}

func (p *Pool) Wait() {
	p.wg.Wait()
}

func main() {

}
