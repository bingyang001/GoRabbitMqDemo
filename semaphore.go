package goRabbitMq

//import (
//	"sync"
//)
//type i func()
////等待
//func (s *Semaphore) Wait() {
//	wg.Add(s.max)
//	defer s.w.Done()
//	s.c <- 1 // 向 sem 发送数据，阻塞或者成功。
//	s.w.Wait()
//}

//// 释放信号，使得其他阻塞 goroutine 可以发送数据。
//func (s *Semaphore) Release() {
//	<-s.c
//}
