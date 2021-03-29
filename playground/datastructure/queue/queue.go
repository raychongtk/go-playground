package queue

import "container/list"

var queue = list.New()

func Add(item string) {
	queue.PushBack(item)
}

func Poll() string {
	f := queue.Front()
	queue.Remove(f)
	return f.Value.(string)
}

func GetQueue() list.List {
	return *queue
}
