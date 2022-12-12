package unicycle

import (
	"reflect"
	"testing"
)

func assertQueueLength(t *testing.T, queue *Queue[string], length int) {
	if queue.length != length {
		t.Errorf("queue.length = %d; want %d", queue.length, length)
	}
}

func assertQueueIsEmpty(t *testing.T, queue *Queue[string]) {
	assertQueueLength(t, queue, 0)
	if queue.first != nil {
		t.Error("empty queue.first is not nil")
	}
	if queue.last != nil {
		t.Error("empty queue.first is not nil")
	}
}

func assertQueueStartsWith(t *testing.T, queue *Queue[string], startValue string) {
	if queue.first == nil {
		t.Error("filled queue.first is nil")
	} else if queue.first.value != startValue {
		t.Errorf("queue.first has the wrong value '%s' instead of expected '%s'", queue.first.value, startValue)
	}
}

func assertQueueEndsWith(t *testing.T, queue *Queue[string], endValue string) {
	if queue.last == nil {
		t.Error("filled queue.last is nil")
	} else if queue.last.value != endValue {
		t.Errorf("queue.last has the wrong value '%s' instead of expected '%s'", queue.last.value, endValue)
	}
}

func TestQueue(t *testing.T) {
	var queue Queue[string]
	assertQueueIsEmpty(t, &queue)
	queue.Push("hello")
	queue.Push("world")
	queue.Push("lorem ipsum dolor sit amet")
	assertQueueLength(t, &queue, 3)
	assertQueueStartsWith(t, &queue, "hello")
	assertQueueEndsWith(t, &queue, "lorem ipsum dolor sit amet")
	popped := queue.Pop()
	assertQueueLength(t, &queue, 2)
	if popped != "hello" {
		t.Errorf("queue.Pop() returned %s out of order", popped)
	}
	assertQueueStartsWith(t, &queue, "world")
	queue.Push("foo")
	queue.Push("bar")
	assertQueueLength(t, &queue, 4)
	assertQueueStartsWith(t, &queue, "world")
	assertQueueEndsWith(t, &queue, "bar")
	slice := queue.PopAll()
	if !reflect.DeepEqual(slice, []string{"world", "lorem ipsum dolor sit amet", "foo", "bar"}) {
		t.Errorf("queue.PopAll() returned unexpected %s", slice)
	}
	assertQueueIsEmpty(t, &queue)
}
