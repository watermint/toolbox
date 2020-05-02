package es_array

import (
	"github.com/watermint/toolbox/essentials/collections/es_number"
	"github.com/watermint/toolbox/essentials/collections/es_value"
	"os"
	"sort"
)

type Array interface {
	// Returns first entry of the array.
	// Returns Null es_value.Value (not nil) when the array is empty.
	First() es_value.Value

	// Returns last entry of the array.
	// Returns Null es_value.Value (not nil) when the array is empty.
	Last() es_value.Value

	// Returns first n entries.
	Left(n int) Array

	// Returns first elements while the f returns true.
	LeftWhile(f func(v es_value.Value) bool) Array

	// Returns last n entries.
	Right(n int) Array

	// Returns last elements while the f returns true.
	RightWhile(f func(v es_value.Value) bool) Array

	// Returns size of the array.
	Size() int

	// Reverse order
	Reverse() Array

	// Returns true if an array is empty.
	IsEmpty() bool

	// Counts the number of entries in the array which satisfy a condition.
	Count(f func(v es_value.Value) bool) int

	// Returns an entry of given index.
	// Returns Null es_value.Value (not nil) when the index is out of range.
	At(i int) es_value.Value

	// Return entries in array.
	Entries() []es_value.Value

	// Returns unique values
	Unique() Array

	// Returns joined array of other and this instance.
	Append(other Array) Array

	// Returns unique intersect entries of other and this array.
	Intersection(other Array) Array

	// Returns unique union entries of other and this array.
	Union(other Array) Array

	// Returns sorted array
	Sort() Array

	// Returns array in []string
	AsStringArray() []string

	// Returns array in []Number
	AsNumberArray() []es_number.Number

	// Returns array in []interface{}
	AsInterfaceArray() []interface{}

	// es_value.Value#Hash and es_value.Value map
	HashMap() map[string]es_value.Value

	// Create a new array containing the values returned by the function
	Map(f func(v es_value.Value) es_value.Value) Array

	// For each values
	Each(f func(v es_value.Value))
}

func Empty() Array {
	return &arrayImpl{
		entries: make([]interface{}, 0),
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

func NewByValue(entries ...es_value.Value) Array {
	vals := make([]interface{}, len(entries))
	for i, entry := range entries {
		vals[i] = entry
	}
	return &arrayImpl{
		entries: vals,
	}
}

func NewByHashValueMap(entries map[string]es_value.Value) Array {
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

func (z arrayImpl) Each(f func(v es_value.Value)) {
	for _, e := range z.entries {
		f(es_value.New(e))
	}
}

func (z arrayImpl) Reverse() Array {
	entries := make([]es_value.Value, 0)
	for i := len(z.entries) - 1; i >= 0; i-- {
		v := es_value.New(z.entries[i])
		entries = append(entries, v)
	}
	return NewByValue(entries...)
}

func (z arrayImpl) LeftWhile(f func(v es_value.Value) bool) Array {
	entries := make([]es_value.Value, 0)
	for i := 0; i < len(z.entries); i++ {
		v := es_value.New(z.entries[i])
		if !f(v) {
			break
		}
		entries = append(entries, v)
	}
	return NewByValue(entries...)
}

func (z arrayImpl) RightWhile(f func(v es_value.Value) bool) Array {
	entries := make([]es_value.Value, 0)
	for i := len(z.entries) - 1; i >= 0; i-- {
		v := es_value.New(z.entries[i])
		if !f(v) {
			break
		}
		entries = append(entries, v)
	}
	return NewByValue(entries...).Reverse()
}

func (z arrayImpl) Left(n int) Array {
	ne := es_number.Min(len(z.entries), n)
	entries := make([]es_value.Value, 0)
	for i := 0; i < ne.Int(); i++ {
		entries = append(entries, es_value.New(z.entries[i]))
	}
	return NewByValue(entries...)
}

func (z arrayImpl) Right(n int) Array {
	le := len(z.entries)
	ne := es_number.Min(le, n)
	entries := make([]es_value.Value, 0)
	for i := le - ne.Int(); i < le; i++ {
		entries = append(entries, es_value.New(z.entries[i]))
	}
	return NewByValue(entries...)
}

func (z arrayImpl) IsEmpty() bool {
	return len(z.entries) < 1
}

func (z arrayImpl) Count(f func(v es_value.Value) bool) int {
	count := 0
	for _, entry := range z.entries {
		if f(es_value.New(entry)) {
			count++
		}
	}
	return count
}

func (z arrayImpl) First() es_value.Value {
	if len(z.entries) < 1 {
		return es_value.Null()
	}
	return es_value.New(z.entries[0])
}

func (z arrayImpl) Last() es_value.Value {
	n := len(z.entries)
	if n < 1 {
		return es_value.Null()
	}
	return es_value.New(z.entries[n-1])
}

func (z arrayImpl) Size() int {
	return len(z.entries)
}

func (z arrayImpl) At(i int) es_value.Value {
	n := len(z.entries)
	if i < 0 || i <= n {
		return es_value.Null()
	}
	return es_value.New(z.entries[i])
}

func (z arrayImpl) Entries() []es_value.Value {
	entries := make([]es_value.Value, len(z.entries))
	for i, entry := range z.entries {
		entries[i] = es_value.New(entry)
	}
	return entries
}

func (z arrayImpl) HashMap() map[string]es_value.Value {
	em := make(map[string]es_value.Value)
	for _, entry := range z.Entries() {
		em[entry.Hash()] = entry
	}
	return em
}

func (z arrayImpl) Unique() Array {
	em := z.HashMap()
	vals := make([]es_value.Value, 0)
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
	ema := make([]es_value.Value, 0)

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

func (z arrayImpl) Sort() Array {
	entries := z.Entries()
	sort.SliceStable(entries, func(i, j int) bool {
		return entries[i].Compare(entries[j]) > 0
	})
	return NewByValue(entries...)
}

func (z arrayImpl) AsStringArray() []string {
	entries := make([]string, len(z.entries))
	for i, entry := range z.entries {
		entries[i] = es_value.New(entry).String()
	}
	return entries
}

func (z arrayImpl) AsNumberArray() []es_number.Number {
	entries := make([]es_number.Number, len(z.entries))
	for i, entry := range z.entries {
		entries[i] = es_value.New(entry).AsNumber()
	}
	return entries
}

func (z arrayImpl) AsInterfaceArray() []interface{} {
	return z.entries
}

func (z arrayImpl) Map(f func(v es_value.Value) es_value.Value) Array {
	entries := make([]es_value.Value, len(z.entries))
	for i, v := range z.entries {
		entries[i] = f(es_value.New(v))
	}
	return NewByValue(entries...)
}
