package unicycle

import (
	"reflect"
	"testing"
)

func assertQueueLength(t *testing.T, queue Queue[string], length int) {
	if queue.Len() != length {
		t.Errorf("queue.Len() = %d; want %d", queue.Len(), length)
	}
}

func TestQueue(t *testing.T) {
	queue := NewQueue[string]()
	assertQueueLength(t, queue, 0)
	queue.Push("hello")
	queue.Push("world")
	queue.Push("lorem ipsum dolor sit amet")
	assertQueueLength(t, queue, 3)
	popped := queue.Pop()
	assertQueueLength(t, queue, 2)
	if popped != "hello" {
		t.Errorf("queue.Pop() returned %s out of order", popped)
	}
	queue.Push("foo")
	queue.Push("bar")
	assertQueueLength(t, queue, 4)
	slice := queue.PopAll()
	if !reflect.DeepEqual(slice, []string{"world", "lorem ipsum dolor sit amet", "foo", "bar"}) {
		t.Errorf("queue.PopAll() returned unexpected %s", slice)
	}
	assertQueueLength(t, queue, 0)
}
