package filter

// Filter is the interface implemented by all filters.
type Filter interface {
	Allow(line []byte) bool
}

// CompositeFilter combines multiple filters with AND semantics:
// a line is allowed only if all filters allow it.
type CompositeFilter struct {
	filters []Filter
}

// NewCompositeFilter creates a CompositeFilter from the given filters.
// Nil filters are ignored.
func NewCompositeFilter(filters ...Filter) *CompositeFilter {
	active := make([]Filter, 0, len(filters))
	for _, f := range filters {
		if f != nil {
			active = append(active, f)
		}
	}
	return &CompositeFilter{filters: active}
}

// Allow returns true if every contained filter allows the line.
// If no filters are present, all lines are allowed.
func (c *CompositeFilter) Allow(line []byte) bool {
	for _, f := range c.filters {
		if !f.Allow(line) {
			return false
		}
	}
	return true
}

// Len returns the number of active filters in the composite.
func (c *CompositeFilter) Len() int {
	return len(c.filters)
}
