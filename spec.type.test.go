package spec

import "fmt"

type (
	Test interface {
		len() int
		run(handler Handler, hooks []Hook)
		context(title string) Test
	}

	test struct {
		title string
		lines []Fallible
	}
)

func (it test) init(title string, lines []Fallible) test {
	it.title = title
	it.lines = lines

	return it
}

func (it test) len() int {
	return len(it.lines)
}

func (it test) run(handler Handler, hooks []Hook) {
	if it.len() <= 0 {
		handler.Call(enriches(it.title, fmt.Errorf("empty test")))
	}

	for _, hook := range hooks {
		switch {
		case hook.isBeforeEach():
			hook.run(handler)
		case hook.isAfterEach():
			defer hook.run(handler)
		}
	}

	for _, line := range it.lines {
		handler.Call(enriches(it.title, line.Call()))
	}
}

func (it test) context(title string) Test {
	it.title = fmt.Sprintf("%s: %s", title, it.title)

	return it
}
