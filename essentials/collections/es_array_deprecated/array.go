package es_array_deprecated

import (
	"github.com/watermint/toolbox/essentials/collections/es_number_deprecated"
	"github.com/watermint/toolbox/essentials/collections/es_value_deprecated"
	"os"
	"sort"
)

type Array interface {
	// Returns first entry of the array.
	// Returns Null es_value_deprecated.Value (not nil) when the array is empty.
	First() es_value_deprecated.Value

	// Returns last entry of the array.
	// Returns Null es_value_deprecated.Value (not nil) when the array is empty.
	Last() es_value_deprecated.Value

	// Returns first n entries.
	Left(n int) Array

	// Returns first elements while the f returns true.
	LeftWhile(f func(v es_value_deprecated.Value) bool) Array

	// Returns last n entries.
	Right(n int) Array

	// Returns last elements while the f returns true.
	RightWhile(f func(v es_value_deprecated.Value) bool) Array

	// Returns size of the array.
	Size() int

	// Reverse order
	Reverse() Array

	// Returns true if an array is empty.
	IsEmpty() bool

	// Counts the number of entries in the array which satisfy a condition.
	Count(f func(v es_value_deprecated.Value) bool) int

	// Returns an entry of given index.
	// Returns Null es_value_deprecated.Value (not nil) when the index is out of range.
	At(i int) es_value_deprecated.Value

	// Return entries in array.
	Entries() []es_value_deprecated.Value

	// Returns unique values
	Unique() Array

	// Returns joined array of other and this instance.
	Append(other Array) Array

	// Returns unique intersect entries of other and this array.
	Intersection(other Array) Array

	// Returns unique union entries of other and this array.
	Union(other Array) Array

	// Returns an array removing all occurrences of entries in other.
	Diff(other Array) Array

	// Returns sorted array
	Sort() Array

	// Returns array in []string
	AsStringArray() []string

	// Returns array in []Number
	AsNumberArray() []es_number_deprecated.Number

	// Returns array in []interface{}
	AsInterfaceArray() []interface{}

	// es_value_deprecated.Value#Hash and es_value_deprecated.Value map
	HashMap() map[string]es_value_deprecated.Value

	// Create a new array containing the values returned by the function
	Map(f func(v es_value_deprecated.Value) es_value_deprecated.Value) Array

	// For each values
	Each(f func(v es_value_deprecated.Value))
}

func Empty() Array {
	return &arrayImpl{
		entries: make([]interface{}, 0),
	}
}

func NewByString(entries ...string) Array {
	vals := make([]interface{}, len(entries))
	for i, entry := range entries {
		vals[i] = entry
	}
	return &arrayImpl{
		entries: vals,
	}
}

func NewByInterface(entries ...interface{}) Array {
	if entries == nil {
		return Empty()
	}
	return &arrayImpl{
		entries: entries,
	}
}

func NewByFileInfo(entries ...os.FileInfo) Array {
	vals := make([]interface{}, len(entries))
	for i, entry := range entries {
		vals[i] = entry
	}
	return &arrayImpl{
		entries: vals,
	}
}

func NewByValue(entries ...es_value_deprecated.Value) Array {
	vals := make([]interface{}, len(entries))
	for i, entry := range entries {
		vals[i] = entry
	}
	return &arrayImpl{
		entries: vals,
	}
}

func NewByHashValueMap(entries map[string]es_value_deprecated.Value) Array {
	vals := make([]interface{}, 0)
	for _, entry := range entries {
		vals = append(vals, entry)
	}
	return &arrayImpl{
		entries: vals,
	}
}

type arrayImpl struct {
	entries []interface{}
}

func (z arrayImpl) Each(f func(v es_value_deprecated.Value)) {
	for _, e := range z.entries {
		f(es_value_deprecated.New(e))
	}
}

func (z arrayImpl) Reverse() Array {
	entries := make([]es_value_deprecated.Value, 0)
	for i := len(z.entries) - 1; i >= 0; i-- {
		v := es_value_deprecated.New(z.entries[i])
		entries = append(entries, v)
	}
	return NewByValue(entries...)
}

func (z arrayImpl) LeftWhile(f func(v es_value_deprecated.Value) bool) Array {
	entries := make([]es_value_deprecated.Value, 0)
	for i := 0; i < len(z.entries); i++ {
		v := es_value_deprecated.New(z.entries[i])
		if !f(v) {
			break
		}
		entries = append(entries, v)
	}
	return NewByValue(entries...)
}

func (z arrayImpl) RightWhile(f func(v es_value_deprecated.Value) bool) Array {
	entries := make([]es_value_deprecated.Value, 0)
	for i := len(z.entries) - 1; i >= 0; i-- {
		v := es_value_deprecated.New(z.entries[i])
		if !f(v) {
			break
		}
		entries = append(entries, v)
	}
	return NewByValue(entries...).Reverse()
}

func (z arrayImpl) Left(n int) Array {
	ne := min(len(z.entries), n)
	entries := make([]es_value_deprecated.Value, 0)
	for i := 0; i < ne; i++ {
		entries = append(entries, es_value_deprecated.New(z.entries[i]))
	}
	return NewByValue(entries...)
}

func (z arrayImpl) Right(n int) Array {
	le := len(z.entries)
	ne := min(le, n)
	entries := make([]es_value_deprecated.Value, 0)
	for i := le - ne; i < le; i++ {
		entries = append(entries, es_value_deprecated.New(z.entries[i]))
	}
	return NewByValue(entries...)
}

func (z arrayImpl) IsEmpty() bool {
	return len(z.entries) < 1
}

func (z arrayImpl) Count(f func(v es_value_deprecated.Value) bool) int {
	count := 0
	for _, entry := range z.entries {
		if f(es_value_deprecated.New(entry)) {
			count++
		}
	}
	return count
}

func (z arrayImpl) First() es_value_deprecated.Value {
	if len(z.entries) < 1 {
		return es_value_deprecated.Null()
	}
	return es_value_deprecated.New(z.entries[0])
}

func (z arrayImpl) Last() es_value_deprecated.Value {
	n := len(z.entries)
	if n < 1 {
		return es_value_deprecated.Null()
	}
	return es_value_deprecated.New(z.entries[n-1])
}

func (z arrayImpl) Size() int {
	return len(z.entries)
}

func (z arrayImpl) At(i int) es_value_deprecated.Value {
	n := len(z.entries)
	if i < 0 || i <= n {
		return es_value_deprecated.Null()
	}
	return es_value_deprecated.New(z.entries[i])
}

func (z arrayImpl) Entries() []es_value_deprecated.Value {
	entries := make([]es_value_deprecated.Value, len(z.entries))
	for i, entry := range z.entries {
		entries[i] = es_value_deprecated.New(entry)
	}
	return entries
}

func (z arrayImpl) HashMap() map[string]es_value_deprecated.Value {
	em := make(map[string]es_value_deprecated.Value)
	for _, entry := range z.Entries() {
		em[entry.Hash()] = entry
	}
	return em
}

func (z arrayImpl) Unique() Array {
	em := z.HashMap()
	vals := make([]es_value_deprecated.Value, 0)
	for _, v := range em {
		vals = append(vals, v)
	}
	return NewByValue(vals...)
}

func (z arrayImpl) Append(other Array) Array {
	entries := z.Entries()
	entries = append(entries, other.Entries()...)
	return NewByValue(entries...)
}

func (z arrayImpl) Intersection(other Array) Array {
	em1 := z.HashMap()
	em2 := other.HashMap()
	ema := make([]es_value_deprecated.Value, 0)

	for k, e := range em1 {
		if _, ok := em2[k]; ok {
			ema = append(ema, e)
		}
	}
	return NewByValue(ema...)
}

func (z arrayImpl) Union(other Array) Array {
	em := z.HashMap()
	em2 := other.HashMap()
	for k, v := range em2 {
		em[k] = v
	}
	return NewByHashValueMap(em)
}

func (z arrayImpl) Diff(other Array) Array {
	em := z.HashMap()
	em2 := other.HashMap()
	for k := range em2 {
		delete(em, k)
	}
	return NewByHashValueMap(em)
}

func (z arrayImpl) Sort() Array {
	entries := z.Entries()
	sort.SliceStable(entries, func(i, j int) bool {
		return entries[i].Compare(entries[j]) < 0
	})
	return NewByValue(entries...)
}

func (z arrayImpl) AsStringArray() []string {
	entries := make([]string, len(z.entries))
	for i, entry := range z.entries {
		entries[i] = es_value_deprecated.New(entry).String()
	}
	return entries
}

func (z arrayImpl) AsNumberArray() []es_number_deprecated.Number {
	entries := make([]es_number_deprecated.Number, len(z.entries))
	for i, entry := range z.entries {
		entries[i] = es_value_deprecated.New(entry).AsNumber()
	}
	return entries
}

func (z arrayImpl) AsInterfaceArray() []interface{} {
	return z.entries
}

func (z arrayImpl) Map(f func(v es_value_deprecated.Value) es_value_deprecated.Value) Array {
	entries := make([]es_value_deprecated.Value, len(z.entries))
	for i, v := range z.entries {
		entries[i] = f(es_value_deprecated.New(v))
	}
	return NewByValue(entries...)
}
