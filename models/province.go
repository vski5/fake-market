package models

import (
	"sync"
	"sync/atomic"

	"github.com/gin-gonic/gin"
)

// 定义结构体，包含一个锁和一个用于存储访问数量的map
type visitInfo struct {
	sync.RWMutex                   //读写锁，防止多个goroutine同时写入map导致的数据冲突
	PageVisit    map[string]*int64 //map结构，用于存储访问的页面和相应的访问数量
}

// 创建 visitInfo 结构的新实例。
// 这个函数会返回一个指向新创建的 visitInfo 结构体的指针。
func NewVisitInfo() *visitInfo {
	return &visitInfo{
		PageVisit: make(map[string]*int64),
	}
}

// gin获取访问者/客户端的ip
func (v *visitInfo) GetIP(c *gin.Context) string {
	return c.ClientIP()
}

// 当有新的访问发生时，调用这个函数来增加对应页面的访问次数
// page 表示要访问的页面，addr 表示访问者的地址（IP地址或者省份）
func (v *visitInfo) AddVisit(c *gin.Context) {

	key := c.ClientIP()         //c.ClientIP()是访问者的ip，现在只能获取访问者的ip地址
	v.Lock()                    //获取写锁
	defer v.Unlock()            // 函数结束时释放写锁
	val, ok := v.PageVisit[key] //检查是否已经有该页面的访问记录
	if !ok {                    //如果没有，初始化该页面的访问次数为1
		var i int64 = 1
		v.PageVisit[key] = &i
	} else { //已经有该页面的访问记录，将访问次数加1
		atomic.AddInt64(val, 1) //使用原子操作确保并发时数据的一致性
	}
}

// 获取指定页面的访问次数
// page 表示要获取访问次数的页面，province 表示省份
func (v *visitInfo) GetVisit(page string, province string) int64 {
	key := page + "-" + province
	v.RLock()                   //获取读锁
	defer v.RUnlock()           //函数结束时释放读锁
	val, ok := v.PageVisit[key] //获取该页面的访问次数
	if !ok {
		return 0 //如果没有访问记录，返回0
	}
	return atomic.LoadInt64(val) //使用原子操作获取访问次数
}
