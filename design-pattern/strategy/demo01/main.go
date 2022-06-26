package main

import (
	"fmt"
)

// IStrategy 抽象策略
type IStrategy interface {
	algorithm()
}

// concreteStrategyA 具体策略
type concreteStrategyA struct{}

func (*concreteStrategyA) algorithm() {
	fmt.Println("Strategy A")
}

// concreteStrategyB 具体策略
type concreteStrategyB struct{}

func (*concreteStrategyB) algorithm() {
	fmt.Println("Strategy B")
}

// Context 上下文环境
type Context struct {
	strategy IStrategy
}

func NewContext(strategy IStrategy) *Context {
	return &Context{
		strategy: strategy,
	}
}

// 调用策略中的方法
func (context *Context) do() {
	context.strategy.algorithm()
}

// 策略工厂
type StrategyFactory struct {
	strategys map[string]IStrategy
}

func NewStrategyFactory() *StrategyFactory {
	factory := new(StrategyFactory)
	strategys := map[string]IStrategy{
		"A": new(concreteStrategyA),
		"B": new(concreteStrategyB),
	}

	factory.strategys = strategys
	return factory
}

func (factory *StrategyFactory) GetStrategy(name string) (strategy IStrategy, ok bool) {
	if v, ok := factory.strategys[name]; ok {
		return v, true
	}
	return nil, false
}

func main() {
	// 1. 使用策略模式
	// 选择一个具体策略
	strategy := new(concreteStrategyA)
	// 创建一个上下文环境
	context := NewContext(strategy)
	// 客户端直接让上下文角色执行算法
	context.do()

	// 2. 使用策略工厂来管理策略类
	factory := NewStrategyFactory()
	if strategy, ok := factory.GetStrategy("B"); ok {
		strategy.algorithm()
	}
}
