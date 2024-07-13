package stream

type Stream struct {
	list []string
}

func List(data []string) Stream {
	return Stream{list: data}
}

func (r Stream) Map(mapper func(interface{}) string) Stream {

	mapped := make([]string, len(r.list))

	for i, item := range r.list {
		mapped[i] = mapper(item)
	}

	r.list = mapped

	return r
}

func (r Stream) Filter(predicate func(interface{}) bool) Stream {

	filtered := make([]string, 0)

	for _, item := range r.list {
		if predicate(item) {
			filtered = append(filtered, item)
		}
	}

	r.list = filtered

	return r
}

func (r Stream) Collect() []string {
	return r.list
}
