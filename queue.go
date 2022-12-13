package unicycle

type element[VALUE_TYPE any] struct {
	value VALUE_TYPE
	next  *element[VALUE_TYPE]
}

// used to ensure that Queue can be passed by reference or value and always behave the same
type innerQueue[VALUE_TYPE any] struct {
	first  *element[VALUE_TYPE]
	last   *element[VALUE_TYPE]
	length int
}

func (queue *innerQueue[VALUE_TYPE]) Push(value VALUE_TYPE) int {
	node := &element[VALUE_TYPE]{
		value: value,
	}
	if queue.first == nil {
		queue.first = node
		queue.length = 1
	} else {
		queue.last.next = node
		queue.length++
	}
	queue.last = node
	return queue.length
}

func (queue *innerQueue[VALUE_TYPE]) More() bool {
	return queue.first != nil
}

func (queue *innerQueue[VALUE_TYPE]) Len() int {
	return queue.length
}

func (queue *innerQueue[VALUE_TYPE]) Pop() VALUE_TYPE {
	if queue.first == nil {
		panic("runtime error: Pop() called on queue with length 0. Make sure to call queue.More() while iterating.")
	} else {
		value := queue.first.value
		if queue.first == queue.last {
			queue.last = nil
		}
		queue.first = queue.first.next
		queue.length--
		return value
	}
}

func (queue *innerQueue[VALUE_TYPE]) PopAll() []VALUE_TYPE {
	output := make([]VALUE_TYPE, queue.length)
	index := 0
	for queue.More() {
		output[index] = queue.Pop()
		index++
	}
	return output
}

// A simple FIFO queue, implemented as a monodirectional linked list
type Queue[VALUE_TYPE any] struct {
	inner *innerQueue[VALUE_TYPE]
}

func (queue *Queue[VALUE_TYPE]) Init() {
	if queue.inner == nil {
		queue.inner = &innerQueue[VALUE_TYPE]{}
	}
}

func (queue *Queue[VALUE_TYPE]) Push(value VALUE_TYPE) int {
	queue.Init()
	return queue.inner.Push(value)
}

func (queue *Queue[VALUE_TYPE]) More() bool {
	queue.Init()
	return queue.inner.More()
}

func (queue *Queue[VALUE_TYPE]) Len() int {
	queue.Init()
	return queue.inner.Len()
}

func (queue *Queue[VALUE_TYPE]) Pop() VALUE_TYPE {
	queue.Init()
	return queue.inner.Pop()
}

func (queue *Queue[VALUE_TYPE]) PopAll() []VALUE_TYPE {
	queue.Init()
	return queue.inner.PopAll()
}
