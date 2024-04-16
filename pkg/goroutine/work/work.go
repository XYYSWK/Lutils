package work

import "sync"

/*
	工作池，它使用一组固定数量的工作线程来执行任务队列中的工作单元。
	可以处理：“一组耗时的任务需要执行，我们希望并发执行它们，但同时限制开发度。“的问题
*/

type Worker struct {
	config   Config
	taskChan chan func()      // 任务通道，用于接收待执行的任务函数
	errChan  chan interface{} // 错误通道，用于传递任务执行过程中的错误信息
	wg       sync.WaitGroup   // WaitGroup，用于等待所有任务执行完成
}

type Config struct {
	TaskChanCapacity int // 任务 channel 容量
	WorkerNum        int // 协程工人数
	ErrChanCapacity  int // 错误 channel 容量
}

func Init(config *Config) *Worker {
	w := &Worker{
		config:   *config,
		taskChan: make(chan func(), config.TaskChanCapacity),
		errChan:  make(chan interface{}, config.ErrChanCapacity),
		wg:       sync.WaitGroup{},
	}
	w.run()
	return w
}

func (w *Worker) run() {
	w.wg.Add(int(w.config.WorkerNum))
	for i := 0; i < w.config.WorkerNum; i++ {
		go w.work()
	}
}

func (w *Worker) work() {
	defer func() {
		err := recover()
		if err == nil {
			w.wg.Done()
			return
		}
		select {
		case w.errChan <- err:
		default:
		}
		go w.work() // 重新启动
	}()
	for task := range w.taskChan {
		task()
	}
}

func (w *Worker) SendTask(task func()) {
	w.taskChan <- task
}

func (w *Worker) Err() <-chan interface{} {
	return w.errChan
}

func (w *Worker) Stop() {
	close(w.taskChan)
	w.wg.Wait()
}

func (w *Worker) Restart(config *Config) {
	w.Stop()
	if config != nil {
		w.config = *config
	}
	w.run()
}
