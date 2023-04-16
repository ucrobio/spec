package spec

import "fmt"

func Run(engine Engine, handler Handler) {
	engine.run(handler, []Hook{})
}

func Define(title string, body ...Element) Engine {
	return Describe(title, body...).Engine()
}

func Describe(title string, body ...Element) Element {
	hooks := []Hook{}
	tests := []Test{}
	engines := []Engine{}

	for _, child := range body {
		switch {
		case child.IsHook():
			hooks = append(hooks, child.Hook())
		case child.IsTest():
			tests = append(tests, child.Test())
		case child.IsEngine():
			engines = append(engines, child.Engine())
		}
	}

	element := new(element).init()
	element.engine = new(engine).init(title, hooks, tests, engines)

	return element
}

func BeforeAll(lines ...Evaluable) Element {
	element := new(element).init()
	element.hook = new(hook).init(true, false, false, false, lines)

	return element
}

func BeforeEach(lines ...Evaluable) Element {
	element := new(element).init()
	element.hook = new(hook).init(false, true, false, false, lines)

	return element
}

func AfterAll(lines ...Evaluable) Element {
	element := new(element).init()
	element.hook = new(hook).init(false, false, true, false, lines)

	return element
}

func AfterEach(lines ...Evaluable) Element {
	element := new(element).init()
	element.hook = new(hook).init(false, false, false, true, lines)

	return element
}

func It(title string, lines ...Fallible) Element {
	element := new(element).init()
	element.test = new(test).init(title, lines)

	return element
}

func enriches(title string, err error) error {
	if err == nil {
		return nil
	}

	return fmt.Errorf("%s: %w", title, err)
}
