package main

import "fmt"

type Builder struct {
	Tag string
}

func newbuilder() *Builder {
	return &Builder{}
}

func (builder *Builder) setTag() {
	builder.Tag = "Tag101"
}

func (builder *Builder) getA() A {
	return A{
		Tag: builder.Tag,
	}
}

func (builder *Builder) Build() A {
	builder.setTag()
	return builder.getA()
}

type A struct {
	Tag string
}

func main() {
	builder1 := &Builder{}
	A1 := builder1.Build()
	fmt.Println("A1 Tag: ", A1.Tag)
}
