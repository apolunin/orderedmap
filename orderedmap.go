package orderedmap

// OrderedMap represents a hash map which maintains key insertion order.
//   K - type of keys, should satisfy 'comparable' constraint
//   V - value type, has no restrictions
//
// NOTE: This type is NOT thread-safe.
type OrderedMap[K comparable, V any] struct {
	data  map[K]*element[K, V]
	items *list[K]
}

// New creates a new instance of OrderedMap and returns a pointer to it.
func New[K comparable, V any]() *OrderedMap[K, V] {
	return &OrderedMap[K, V]{
		data:  make(map[K]*element[K, V]),
		items: &list[K]{},
	}
}

// Get retrieves a value corresponding to `key`.
//
// Parameters:
//   - `key` - a key in the map.
//
// Returns:
//   - (value, true) if corresponding key->value pair is present in a map;
//   - (<zero>, false) is returned otherwise, where <zero> represents a default value for type V.
func (om *OrderedMap[K, V]) Get(key K) (V, bool) {
	if elem, ok := om.data[key]; ok {
		return elem.value, true
	}

	var def V
	return def, false
}

// Set adds a key->value entry to a map.
//
// If `key` is already present in a map, corresponding entry is updated with a new value.
// In this case insertion order of keys is not changed (i.e. `key` is not moved to the end of the keys list).
//
// Parameters:
//   - `key` - map entry key.
//   - `value` - map entry value.
//
// Returns:
//   - (old, true) if `key` already existed in a map, where `old` is a previous value of the entry;
//   - (<zero>, false) if `key` didn't exist before where <zero> represents a default value for type V.
func (om *OrderedMap[K, V]) Set(key K, value V) (V, bool) {
	if old, ok := om.data[key]; ok {
		om.data[key].value = value
		return old.value, true
	}

	item := &node[K]{value: key}
	om.items.push(item)
	om.data[key] = &element[K, V]{value, item}

	var def V
	return def, false
}

// Delete removes a key->value entry from a map.
//
// Parameters:
//   - `key` - map entry key.
//
// Returns:
//   - (value, true) if key->value entry was present in a map;
//   - (<zero>, false) is returned otherwise where <zero> represents a default value for type V.
func (om *OrderedMap[K, V]) Delete(key K) (V, bool) {
	if val, ok := om.data[key]; ok {
		om.items.remove(val.item)
		delete(om.data, key)
		return val.value, true
	}

	var val V
	return val, false
}

// Len returns total number of elements in a map.
func (om *OrderedMap[K, V]) Len() int {
	return len(om.data)
}

// Iterator returns a function which can be used to iterate over key->value pairs of a map
// in keys insertion order.
//
// Example:
//  om := orderedmap.New()
//  // ... insert entries in the map ...
//  next := om.Iterator()
//  for k, v, ok := next(); ok; k, v, ok = next() {
//    fmt.Printf("key: %v, value: %v", k, v)
//  }
//
// Function next() returns 3 values: key, value and a bool flag which indicates
// if there are any unvisited elements left.
//
// NOTE: if a map is modified when iteration is in progress,
// the result of a subsequent call to next() is undefined.
func (om *OrderedMap[K, V]) Iterator() func() (K, V, bool) {
	curr := om.items.head
	return func() (K, V, bool) {
		if curr == nil {
			var key K
			var val V
			return key, val, false
		}

		key := curr.value
		val := om.data[key].value
		curr = curr.next

		return key, val, true
	}
}

type node[T any] struct {
	value      T
	prev, next *node[T]
}

type element[K comparable, V any] struct {
	value V
	item  *node[K]
}

type list[T any] struct {
	head, tail *node[T]
}

func (lst *list[T]) push(n *node[T]) {
	if lst.head == nil {
		lst.head = n
		lst.tail = n
	} else {
		lst.tail.next = n
		n.prev = lst.tail
		lst.tail = n
	}
}

func (lst *list[T]) remove(n *node[T]) {
	if n.prev != nil {
		n.prev.next = n.next
	}

	if n.next != nil {
		n.next.prev = n.prev
	}

	if n == lst.head {
		lst.head = n.next
	}

	if n == lst.tail {
		lst.tail = n.prev
	}
}
