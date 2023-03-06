# sync包

## sync和channel的运行对比

```shell
go test -bench  . go-sync-package-1_test.go 

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i7-8700 CPU @ 3.20GHz
BenchmarkCriticalSectionSyncByMutex-12          100000000               11.92 ns/op
BenchmarkCriticalSectionSyncByChan-12           32173220                36.39 ns/op
PASS
ok      command-line-arguments  2.421s

```

可以看到`sync.Mutex`实现的同步机制的性能要比channel实现的高出两倍多


不想转移结构体对象所有权，但又要保证结构体内部状态数据的同步访问的场景基于channel的并发设计的一个特点是，在goroutine间通过channel转移数据对象的所有权。只有拥有数据对象所有权（从channel接收到该数据）的goroutine才可以对该数据对象进行状态变更。如果你的设计中没有转移结构体对象所有权，但又要保证结构体内部状态数据能在多个goroutine之间同步访问，那么你可以使用sync包提供的低级同步原语来实现，比如最常用的sync.Mutex。
