package main

import "fmt"

// Adaptee is
type Adaptee struct{}

// ExistingMethod is
func (a *Adaptee) ExistingMethod() {
	fmt.Println("using existing method")
}

// Adapter is
type Adapter struct {
	adaptee *Adaptee
}

// NewAdapter returns a pointer of Adapter
func NewAdapter() *Adapter {
	return &Adapter{new(Adaptee)}
}

// ExpectedMethod is
func (a *Adapter) ExpectedMethod() {
	fmt.Println("doning some work")
	a.adaptee.ExistingMethod()
}

func main() {
	adaptor := NewAdapter()
	adaptor.ExpectedMethod()
}
