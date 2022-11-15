package ecs

type ISystem interface {
	GetBaseSystem() *BaseSystem
	Update(float32)
}

type BaseSystem struct {
	Filter          *ComponentFilter
	Priority        int
	Entities        *SimpleContainer
	WhenCheckAndAdd func(*Entity)
}

func NewBaseSystem() *BaseSystem {

	return &BaseSystem{
		Filter:          NewComponentFilter(),
		Priority:        0,
		Entities:        NewSimpleContainer(1000),
		WhenCheckAndAdd: nil,
	}
}

func (sys *BaseSystem) CheckAndAdd(en *Entity) {

	if sys.Filter.IsFiltered2(en.ComponentFilter) {
		sys.Entities.Add(en)
		en.SysRefCount++
	}
	if sys.WhenCheckAndAdd != nil {
		sys.WhenCheckAndAdd(en)
	}

}

//Each iterate valid entity, it remove entity which need to be remove.
func (sys *BaseSystem) Each(f func(en *Entity)) {

	sys.Entities.Each(func(idx int, item interface{}) {

		en := item.(*Entity)
		if sys.CheckAndRemove(en, idx) || en.Active == false {
			return
		}

		f(en)

	})
}

func (sys *BaseSystem) CheckAndRemove(en *Entity, idx int) bool {

	if en.NeedRemove == false {
		return false
	}
	sys.Entities.RemoveIdx(idx)
	en.SysRefCount--

	return true

}
