package orderedmap

type OrderedMap[T any] struct {
	Order []string
	Map   map[string]T
}

func New[T any](data []struct {
	Key   string
	Value T
}) OrderedMap[T] {

	om := OrderedMap[T]{
		Order: make([]string, 0),
		Map:   make(map[string]T),
	}

	if data == nil {
		return om
	}

	for _, item := range data {
		om.Add(item.Key, item.Value)
	}

	return om
}

func (om *OrderedMap[T]) Add(key string, value T) {
	if _, exists := om.Map[key]; exists {
		om.Del(key)
	}
	om.Order = append(om.Order, key)
	om.Map[key] = value
}

func (om *OrderedMap[T]) Get(key string) T {
	return om.Map[key]
}

func (om *OrderedMap[T]) GetWithExists(key string) (T, bool) {
	val, exists := om.Map[key]
	return val, exists
}

func (om *OrderedMap[T]) Del(key string) {
	if _, exists := om.Map[key]; !exists {
		return
	}

	delete(om.Map, key)

	for i, k := range om.Order {
		if k == key {
			om.Order = append(om.Order[:i], om.Order[i+1:]...)
			break
		}
	}
}

func (om *OrderedMap[T]) Keys() []string {
	return append([]string{}, om.Order...)
}

func (om *OrderedMap[T]) Values() []T {
	values := make([]T, len(om.Order))
	for i, key := range om.Order {
		values[i] = om.Map[key]
	}
	return values
}

func (om *OrderedMap[T]) Len() int {
	return len(om.Order)
}

func (om *OrderedMap[T]) Has(key string) bool {
	_, exists := om.Map[key]
	return exists
}

func (om *OrderedMap[T]) Clear() {
	om.Order = make([]string, 0)
	om.Map = make(map[string]T)
}

func (om *OrderedMap[T]) ForEach(fn func(key string, value T)) {
	for _, key := range om.Order {
		fn(key, om.Map[key])
	}
}

func (om *OrderedMap[T]) Clone() OrderedMap[T] {
	clone := OrderedMap[T]{
		Order: make([]string, len(om.Order)),
		Map:   make(map[string]T, len(om.Map)),
	}

	copy(clone.Order, om.Order)
	for k, v := range om.Map {
		clone.Map[k] = v
	}

	return clone
}
