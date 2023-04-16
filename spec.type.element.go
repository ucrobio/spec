package spec

type (
	Element interface {
		Test() Test
		Hook() Hook
		Engine() Engine
		IsTest() bool
		IsHook() bool
		IsEngine() bool
	}

	element struct {
		test   Test
		hook   Hook
		engine Engine
	}
)

func (it element) init() element { return it }

func (it element) Test() Test     { return it.test }
func (it element) Hook() Hook     { return it.hook }
func (it element) Engine() Engine { return it.engine }
func (it element) IsTest() bool   { return it.test != nil }
func (it element) IsHook() bool   { return it.hook != nil }
func (it element) IsEngine() bool { return it.engine != nil }
