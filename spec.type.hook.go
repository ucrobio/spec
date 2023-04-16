package spec

import "fmt"

type (
	Hook interface {
		run(handler Handler)

		isBeforeAll() bool
		isBeforeEach() bool
		isAfterAll() bool
		isAfterEach() bool

		context(title string) Hook
	}

	hook struct {
		title  string
		before struct {
			all  bool
			each bool
		}

		after struct {
			all  bool
			each bool
		}

		lines []Evaluable
	}
)

func (it hook) init(beforeAll, beforeEach, afterAll, afterEach bool, lines []Evaluable) hook {
	it.before.all = beforeAll
	it.before.each = beforeEach
	it.after.all = afterAll
	it.after.each = afterEach

	switch {
	case it.before.all:
		it.title = "before all"
	case it.before.each:
		it.title = "before each"
	case it.after.all:
		it.title = "after all"
	case it.after.each:
		it.title = "after each"
	}

	it.lines = lines

	return it
}

func (it hook) run(handler Handler) {
	for _, line := range it.lines {
		handler.Call(enriches(it.title, line.Call()))
	}
}

func (it hook) isBeforeAll() bool  { return it.before.all }
func (it hook) isBeforeEach() bool { return it.before.each }
func (it hook) isAfterAll() bool   { return it.after.all }
func (it hook) isAfterEach() bool  { return it.after.each }

func (it hook) context(title string) Hook {
	it.title = fmt.Sprintf("%s: %s", title, it.title)

	return it
}
