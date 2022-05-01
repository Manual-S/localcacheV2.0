package fifo

import "container/list"

// fifo 是一个FIFO cache 不是并发安全的
type fifo struct {
	maxBytes  int        // fifo的最大容量
	usedBytes int        // 已经使用过的字节数
	ll        *list.List // 双向链表
	cache     map[string]*list.Element
}

//
type entry struct {
	key   string
	value interface{}
}

// Len 返回一个entry的字节数
func (e *entry) Len() int {
	return 1
}

// Set 向cache中添加一个元素
func (f *fifo) Set(key string, value interface{}) {
	value, ok := f.cache[key]
	if ok {
		// 说明元素存在 需要更新
	}

	// 直接插入元素
	e := entry{
		key:   key,
		value: value,
	}

	f.usedBytes = f.usedBytes + e.Len()
	element := f.ll.PushBack(e)
	f.cache[key] = element
	if f.usedBytes > f.maxBytes {
		// 删除队头元素
	}
}
