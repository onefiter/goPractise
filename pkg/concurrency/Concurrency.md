# 并发编程
## 并发与并行的区别
> 并发不是并行，并发关乎结构，并行关乎执行。  --Rob Pike，Go语言之父

并发`Concurrency`和并行`Paralleism`
## goroutine的调度原理
### goroutine调度器
提到“调度”，我们首先会想到的是操作系统对进程、线程的调度。操作系统调度器会将系统中的多个线程按照一定算法调度到物理CPU上运行。正如上一条提到的，传统的编程语言（如C、C++等）的并发实现多是基于线程模型的，即应用程序负责创建线程（一般通过libpthread等库函数调用实现），操作系统负责调度线程。

Go采用用户层轻量级线程来解决这些问题，并将之称为goroutine。

>由于一个goroutine占用资源很少，一个Go程序中可以创建成千上万个并发的goroutine。

将这些goroutine按照一定算法放到CPU上执行的程序就称为goroutine调度器（`goroutine scheduler`）

goroutine是由Go运行时管理的用户轻量级线程

> 一个Go程序对于操作系统来说只是一个用户层程序，操作系统眼中只有线程，goroutine的调度全要靠Go自己完成。

### goroutine调度模型与演进过程
#### G-M模型
#### G-P-M模型
#### 抢占式调度
#### NUMA调度模型


## Go并发模型和常见并发模式
## channel的妙用
## sync包的正确用法
## 使用atomic包实现伸缩性更好的并发读取
