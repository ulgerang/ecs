package ecs

//Entity has id and components, and  component filter
type Entity struct {
	Id              uint32
	components      map[uint8]IGetTypeCode
	ComponentFilter *ComponentFilter
	Active          bool
	NeedRemove      bool
	SysRefCount     int
	ObjPools        *ObjPools
}

//NewEntity make new instance of Entity
func NewEntity() *Entity {

	return &Entity{
		components:      make(map[uint8]IGetTypeCode),
		ComponentFilter: NewComponentFilter(),
		Active:          true,
		NeedRemove:      false,
	}
}

//FillEntity fill entity
func FillEntity(en *Entity) {

	en.components = make(map[uint8]IGetTypeCode)
	en.ComponentFilter = NewComponentFilter()
	en.Active = true
	en.NeedRemove = false

}

func (en *Entity) GetTypeCode() uint8 {
	return 0
}

//GetComponent add component
func (en *Entity) GetComponent(p IGetTypeCode) IGetTypeCode {

	tCode := p.GetTypeCode()
	c, ok := en.components[tCode]

	if ok == false {
		panic("No componenet")
	}
	return c
}

func (en *Entity) GetComponentFromCode(tCode uint8) IGetTypeCode {

	c, ok := en.components[tCode]

	if ok == false {
		panic("No componenet")
	}
	return c
}

//HasComponent check being of component
func (en *Entity) HasComponent(p IGetTypeCode) bool {

	tCode := p.GetTypeCode()
	_, ok := en.components[tCode]

	return ok
}

func (en *Entity) HasComponentByTypeCode(tCode uint8) bool {
	_, ok := en.components[tCode]

	return ok
}

//AddComponent add component
func (en *Entity) AddComponent(p IGetTypeCode) IGetTypeCode {

	tCode := p.GetTypeCode()
	c := en.ObjPools.Get(p)
	c.ReadyToUse()
	en.components[tCode] = c
	en.ComponentFilter.componentBits.Set(uint(tCode))

	return c
}

func (en *Entity) AddComponentFromCode(tCode uint8) IGetTypeCode {

	c := en.ObjPools.GetFromTypeCode(tCode)
	c.ReadyToUse()
	en.components[tCode] = c
	en.ComponentFilter.componentBits.Set(uint(tCode))

	return c
}



func (en *Entity) ReadyToUse() {

}

//Reset reset component and componentFilter
func (en *Entity) Reset() {

	for k, v := range en.components {
		en.ObjPools.Put(v)
		delete(en.components, k)
	}
	en.ComponentFilter.Reset()
	en.NeedRemove = false
	en.Active = true
}
