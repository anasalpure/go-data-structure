package main

type SimpleStack struct {
	stack []string
}

func (this *SimpleStack) push(el string) {
	this.stack = append(this.stack, el)
}

func (this *SimpleStack) pop() (string, bool) {
	if len(this.stack) == 0 {
		return "", false
	}
	size := len(this.stack)
	lastEl := this.stack[size-1]

	this.stack = this.stack[:size-1]
	return lastEl, true
}

func (this SimpleStack) isEmpty() bool {
	return len(this.stack) == 0
}
