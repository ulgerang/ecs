package ecs

import bitset "github.com/bits-and-blooms/bitset"

type IGetTypeCode interface {
	GetTypeCode() uint8
	ReadyToUse()
}

// ComponentFilter has bitset  using typecode from typeInfo
type ComponentFilter struct {
	componentBits *bitset.BitSet
}

//NewComponentFilter make instance
func NewComponentFilter() *ComponentFilter {

	return &ComponentFilter{
		componentBits: bitset.New(32),
	}
}

//UsingComponent it mark to componentBits using typeInfo code, params consist of  (*type)(nil)
func (c *ComponentFilter) UsingComponent(oArr ...IGetTypeCode) {

	for _, o := range oArr {
		c.componentBits.Set(uint(o.GetTypeCode()))
	}
}

//IsFiltered check param bitset, it returns true when the intersection of componentBits between param equals param
func (c *ComponentFilter) IsFiltered(bitset *bitset.BitSet) bool {

	return c.componentBits.Intersection(bitset).Equal(c.componentBits)
}

//IsFiltered2 check param ComponentFilter, it returns true when the intersection of componentBits between param equals param
func (c *ComponentFilter) IsFiltered2(f *ComponentFilter) bool {

	return c.IsFiltered(f.componentBits)
}

//Reset clear all componentBits
func (c *ComponentFilter) Reset() {
	c.componentBits.ClearAll()
}
