package main

import "fmt"

const (
	DEFAULT_MAX_TOTAL = 10
	DEFAULT_MAX_IDLE  = 8
	DEFAULT_MIN_IDLE  = 2
)

// 产品类
type ResourcePoolConfig struct {
	name     string
	maxTotal int
	maxIdle  int
	minIdle  int
}

func (r *ResourcePoolConfig) Info() string {
	return fmt.Sprintf("name: %s\nmaxTotal: %d\nmaxIdle: %d\nminIdle: %d",
		r.name, r.maxTotal, r.maxIdle, r.minIdle)
}

// 抽象建造者
type IBuilder interface {
	SetMaxTotal(maxTotal int) IBuilder
	SetMaxIdle(maxIdle int) IBuilder
	SetMinIdle(minIdle int) IBuilder
	Build() *ResourcePoolConfig
}

// 建造者
type ResourcePoolConfigBuilder struct {
	builder *ResourcePoolConfig
}

func NewResourcePoolConfigBuilder(name string) *ResourcePoolConfigBuilder {
	return &ResourcePoolConfigBuilder{builder: &ResourcePoolConfig{name: name}}
}

func (b *ResourcePoolConfigBuilder) SetMaxTotal(maxTotal int) IBuilder {
	if maxTotal <= 0 {
		panic("maxTotal is less than or equal to zero")
	}
	b.builder.maxTotal = maxTotal
	return b
}

func (b *ResourcePoolConfigBuilder) SetMaxIdle(maxIdle int) IBuilder {
	if maxIdle < 0 {
		panic("maxIdle is less than zero")
	}
	b.builder.maxIdle = maxIdle
	return b
}

func (b *ResourcePoolConfigBuilder) SetMinIdle(minIdle int) IBuilder {
	if minIdle < 0 {
		panic("minIdle is less than zero")
	}
	b.builder.minIdle = minIdle
	return b
}

func (b *ResourcePoolConfigBuilder) Build() *ResourcePoolConfig {
	if len(b.builder.name) == 0 {
		panic("invalid name")
	}
	if b.builder.maxTotal == 0 {
		b.builder.maxTotal = DEFAULT_MAX_TOTAL
	}
	if b.builder.maxIdle == 0 {
		b.builder.maxIdle = DEFAULT_MAX_IDLE
	}
	if b.builder.minIdle == 0 {
		b.builder.minIdle = DEFAULT_MIN_IDLE
	}
	if b.builder.maxIdle > b.builder.maxTotal {
		panic("maxIdle is more than maxTotal")
	}
	if b.builder.minIdle > b.builder.maxTotal || b.builder.minIdle > b.builder.maxIdle {
		panic("minIdle is more than maxTotal or minIdle is more than maxIdle")
	}

	return &ResourcePoolConfig{
		name:     b.builder.name,
		maxTotal: b.builder.maxTotal,
		maxIdle:  b.builder.maxIdle,
		minIdle:  b.builder.minIdle,
	}
}

func main() {
	info := NewResourcePoolConfigBuilder("dbconnectionpool").
		SetMaxTotal(15).
		SetMaxIdle(10).
		SetMinIdle(5).
		Build().
		Info()
	fmt.Println(info)
}
