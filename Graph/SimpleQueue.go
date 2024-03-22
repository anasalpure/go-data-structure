package main

type SimpleQueue struct {
	queue []string
}

func (this *SimpleQueue) enqueue(el string) {
	this.queue = append(this.queue, el)
}

func (this *SimpleQueue) dequeue() (string, bool) {
	if len(this.queue) == 0 {
		return "", false
	}
	firstEl := this.queue[0]

	this.queue = this.queue[1:]
	return firstEl, true
}

func (this SimpleQueue) isEmpty() bool {
	return len(this.queue) == 0
}
