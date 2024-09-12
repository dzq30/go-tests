package singleton

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// 上一版可能因为指令重排序导致返回一个未初始化的单例，这里我们通过atomic保证返回的单例一定是被初始化过的
// https://bytedance.larkoffice.com/wiki/wikcndUwBdob7ApxAaS9JAikvGe 原文档
type SingletonWithAtomic struct {
	name string
}

var (
	instanceWithAtomic *SingletonWithAtomic
	lockV2             sync.Mutex
	done               int32
)

func GetInstanceWithAtomic() *SingletonWithAtomic {
	if atomic.LoadInt32(&done) == 0 {
		lockV2.Lock()
		defer lockV2.Unlock()
		if atomic.LoadInt32(&done) == 0 {
			defer atomic.StoreInt32(&done, 1)
			fmt.Println("new instanceWithAtomic")
			instanceWithAtomic = &SingletonWithAtomic{
				name: "dzq",
			}
		}
	}
	return instanceWithAtomic
}
