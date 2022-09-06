package main

import (
	"fmt"
	"testing"
)

/* 封装
字母大写开头的属性和方法可以由包外访问到
字母小写开头的属性和方法表示只能包内调用
*/

type Rect struct {
	Width  float64
	Height float64
}

func (r *Rect) area() float64 {
	return r.Width * r.Height
}

/* 继承
匿名组合
*/

type Base struct {
	Name string
}

func (b Base) Print() {
	fmt.Println(b.Name)
}

type Child struct {
	Base
	Age int64
}

// Print 方法重写
func (c Child) Print() {
	fmt.Println(c.Name, c.Age)
}

func TestInherit(t *testing.T) {
	foo := Child{
		Base: Base{
			Name: "tom",
		},
		Age: 10,
	}
	foo.Print()
	foo.Base.Print()
}

/* 多态

 */

type Interface interface {
	Greet()
}

type Foo struct {
	Name string
}

func (f *Foo) Greet() {
	fmt.Println(f.Name)
}

type Bar struct {
	Name string
}

func (b *Bar) Greet() {
	fmt.Println(b.Name)
}

func Print1(i Interface) {
	i.Greet()
}

func TestPolymorphism1(t *testing.T) {
	Print1(&Foo{Name: "tom"})
	Print1(&Bar{Name: "jerry"})
}

type IBird interface {
	ID() int
	Name() string
	Tweet() error
}

type IFlyableBird interface {
	IBird
	Fly() error
}

type IRunnableBird interface {
	IBird
	Run() error
}

type NormalBird struct {
	iID   int
	sName string
}

func NewNormalBird(id int, name string) *NormalBird {
	return &NormalBird{
		iID:   id,
		sName: name,
	}
}

func (nb *NormalBird) ID() int {
	return nb.iID
}

func (nb *NormalBird) Name() string {
	return nb.sName
}

func (nb *NormalBird) Tweet() error {
	fmt.Printf("NormalBird.Tweet, id=%d, name=%s\n", nb.ID(), nb.Name())
	return nil
}

type FlyableBird struct {
	NormalBird
}

func NewFlyableBird(id int, name string) IBird {
	return &FlyableBird{
		*NewNormalBird(id, name),
	}
}

func (fb *FlyableBird) Fly() error {
	fmt.Printf("FlyableBird.Fly, id=%d, name=%s\n", fb.ID(), fb.Name())
	return nil
}

type RunnableBird struct {
	NormalBird
}

func NewRunnableBird(id int, name string) IBird {
	return &RunnableBird{
		*NewNormalBird(id, name),
	}
}

func (rb *RunnableBird) Run() error {
	fmt.Printf("RunnableBird.Run, id=%d, name=%s\n", rb.ID(), rb.Name())
	return nil
}

func TestPolymorphism2(t *testing.T) {
	callAndLogFn := func(fn func() error) {
		e := fn()
		if e != nil {
			t.Logf("error = %s", e.Error())
		}
	}

	birdFn := func(b IBird) {
		callAndLogFn(b.Tweet)
		if fb, ok := b.(IFlyableBird); ok {
			callAndLogFn(fb.Fly)
		}

		if rb, ok := b.(IRunnableBird); ok {
			callAndLogFn(rb.Run)
		}
	}

	birdFn(NewFlyableBird(10, "飞鸟"))
	birdFn(NewRunnableBird(11, "鸵鸟"))
}
