package spec

type Handler func(err error)

func (it Handler) Call(err error) {
	if err != nil {
		it(err)
	}
}
