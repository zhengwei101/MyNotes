package main

import "fmt"

// func redisPool() {
// 	// 初始化10个连接
// 	pool := make(chan redis.Conn, 10)
// 	for i := 0; i < 10; i++ {
// 		pool <- redis.Dail()
// 	}

// 	// 10个连接在1000个协程中共享
// 	for i := 0; i < 1000; i++ {
// 		go func() {
// 			// 取出连接， 如果连接池中没有，则等待
// 			conn <- pool
// 			// conn op ...
// 			// 使用完结放回，继续被其它线程使用
// 			pool <- conn
// 		}()
// 	}
// }

//消费者
func consumer(data chan int, done chan bool) {
	for x := range data { //接收数据，直到通道被关闭
		fmt.Println("recv: ", x)
	}
	done <- true //通知main,消费结束
}

//生产者
func producer(data chan int) {
	for i := 0; i < 4; i++ {
		data <- i //发送数据
	}
	close(data) //生产结束，关闭通道
}

func main() {
	done := make(chan bool) //用于接收消费结束信号
	data := make(chan int)  //数据管道

	go consumer(data, done) //启动消费者
	go producer(data)       // 启动生产者
	<-done                  //阻塞，直到消费者发回结束信息

	//指定启动10个工作任务
	const taskNum int = 10
	// 定义一个可存储一定数量的int数据的channel类型
	chs := make(chan int, taskNum)

	for i := 0; i < taskNum; i++ {
		// 启动指定数量的协程数
		go func() {
			// do something
			// ...
			// 此异步逻辑执行完成后，写入标记到chan变量中
			chs <- 1
		}()
	}
	// 这里等待所有的异步逻辑执行完成
	for i := 0; i < taskNum; i++ {
		<-chs
	}

	// 继续其它主逻辑
	// TODO
}
