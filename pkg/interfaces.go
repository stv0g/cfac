package cfac

type Metrics map[string]Metric

type Metric interface {
	String() string
	Float() float64
}
