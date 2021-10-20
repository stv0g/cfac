package cfac

type Measurement interface{}

type Measurable interface {
	Fetch()
}
