package medium

import (
	"testing"
)

func TestConstructor(t *testing.T) {
	maxQueue := Constructor()
	maxQueue.Push_back(2)
	maxQueue.Push_back(1)
	maxQueue.Push_back(3)
	t.Log(maxQueue.Max_value())
	maxQueue.Push_back(5)
	t.Log(maxQueue.Max_value())
	maxQueue.Push_back(6)
	t.Log(maxQueue.Max_value())
	maxQueue.Pop_front()
	t.Log(maxQueue.Max_value())

	maxQueue = Constructor()
	t.Log(maxQueue.Max_value())
	maxQueue.Pop_front()
	maxQueue.Pop_front()
	maxQueue.Push_back(94)
	maxQueue.Push_back(16)
	maxQueue.Push_back(89)
	maxQueue.Pop_front()
	maxQueue.Push_back(22)
	maxQueue.Pop_front()
}

