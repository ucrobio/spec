package spec

import "fmt"

type Evaluable func()

func (it Evaluable) Call() (ret error) {
	defer func() {
		rec := recover()
		if rec != nil {
			ret = fmt.Errorf("recovered from: %v", rec)
		}
	}()

	it()

	return ret
}
