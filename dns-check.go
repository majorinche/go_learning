/*
	20210120 - defer用不用，对接口还是有影响，但是这个影响是否有什么实际价值，不是很确定, 必须运用这些基本概念，避免go沦为python

*/

package main

import (
	"context"
	"flag"
	"fmt"
	"net"
	"sync/atomic"
	"time"
)

var host string
var connections int
var duration int64
var limit int64
var timeoutCount int64

func main() {
	// os.Args = append(os.Args, "-host", "www.baidu.com", "-c", "200", "-d", "30", "-l", "5000")

	flag.StringVar(&host, "host", "", "Resolve host")
	flag.IntVar(&connections, "c", 100, "Connections")
	flag.Int64Var(&duration, "d", 0, "Duration(s)")
	flag.Int64Var(&limit, "l", 0, "Limit(ms)")
	flag.Parse()

	var count int64 = 0
	var errCount int64 = 0
	pool := make(chan interface{}, connections)
	exit := make(chan bool)
	var (
		min int64 = 0
		max int64 = 0
		sum int64 = 0
	)

	go func() { //这里的go表示以并发方式，调用匿名函数func(), 我理解是会一直调用，而不是侧重于并发，而是侧重于连续性，从某个角度看，也的确是并发的
		time.Sleep(time.Second * time.Duration(duration))
		exit <- true
	}()
endD:
	for {
		select {
		case pool <- nil: //这里最终会赋值10次nil到pool中，因为pool的长度是10
			//			fmt.Println(len(pool)) //有意思，我们不一定知道通道里面是什么，但是至少知道，用了一次，就会减少一次，通道就是go版的kafka
			go func() {
				defer func() {
					//fmt.Println(len(pool))
					<-pool //这里就是把通道里面的东西用了一次，不过为什么用defer没有想通
				}()
				//<-pool
				resolver := &net.Resolver{}
				now := time.Now()
				_, err := resolver.LookupIPAddr(context.Background(), host)
				use := time.Since(now).Nanoseconds() / int64(time.Millisecond)
				if min == 0 || use < min {
					min = use
				}
				if use > max {
					max = use
				}
				sum += use
				if limit > 0 && use >= limit {
					timeoutCount++
				}
				atomic.AddInt64(&count, 1) //这里居然用到原子性，后期再看，有意思
				if err != nil {
					fmt.Println(err.Error())
					atomic.AddInt64(&errCount, 1)
				}
			}()
		case <-exit:
			break endD
		}
	}

	fmt.Printf("request count：%d\nerror count：%d\n", count, errCount)
	fmt.Printf("request time：min(%dms) max(%dms) avg(%dms) timeout(%dn)\n", min, max, sum/count, timeoutCount)
}
