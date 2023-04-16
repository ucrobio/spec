package spec

import "fmt"

type Fallible func() error

func (it Fallible) Call() (ret error) {
	defer func() {
		rec := recover()
		if rec != nil {
			ret = fmt.Errorf("recovered from: %v", rec)
		}
	}()

	return it()
}
