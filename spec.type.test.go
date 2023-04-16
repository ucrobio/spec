package spec

type (
	Test interface {
		len() int
		run(handler Handler, hooks []Hook)
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
		return
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
		handler.Call(line.Call())
	}
}
