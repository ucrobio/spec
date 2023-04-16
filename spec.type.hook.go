package spec

type (
	Hook interface {
		run(handler Handler)

		isBeforeAll() bool
		isBeforeEach() bool
		isAfterAll() bool
		isAfterEach() bool
	}

	hook struct {
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

	it.lines = lines

	return it
}

func (it hook) run(handler Handler) {
	for _, line := range it.lines {
		handler.Call(line.Call())
	}
}

func (it hook) isBeforeAll() bool  { return it.before.all }
func (it hook) isBeforeEach() bool { return it.before.each }
func (it hook) isAfterAll() bool   { return it.after.all }
func (it hook) isAfterEach() bool  { return it.after.each }
