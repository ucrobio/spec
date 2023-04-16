package spec

type (
	Engine interface {
		run(handler Handler, hooks []Hook)
		len() int
	}

	engine struct {
		title   string
		hooks   []Hook
		tests   []Test
		engines []Engine
	}
)

func (it engine) init(title string, hooks []Hook, tests []Test, engines []Engine) engine {
	it.title = title
	it.hooks = hooks
	it.tests = tests
	it.engines = engines

	return it
}

func (it engine) run(handler Handler, hooks []Hook) {
	if it.len() <= 0 {
		return
	}

	for _, hook := range it.hooks {
		switch {
		case hook.isBeforeAll():
			hook.run(handler)
		case hook.isAfterAll():
			defer hook.run(handler)
		default:
			hooks = append(hooks, hook)
		}
	}

	for _, test := range it.tests {
		test.run(handler, hooks)
	}

	for _, engine := range it.engines {
		engine.run(handler, hooks)
	}
}

func (it engine) len() (ret int) {
	for _, engine := range it.engines {
		ret += engine.len()
	}

	for _, test := range it.tests {
		ret += test.len()
	}

	return ret
}
