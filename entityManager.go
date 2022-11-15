package ecs

//EntityManager manage Entity
type EntityManager struct {
	EntityMap        map[uint32]*Entity
	codeCounter      uint32
	sysManager       *SystemManager
	removeEntities   *SimpleContainer
	ObjPools         *ObjPools
	RemoveEntityFunc func(*Entity)
}

//NewEntityManager make instance
func NewEntityManager(sysManager *SystemManager, f NewInstanceFromTypeCodeFunc) *EntityManager {

	em := &EntityManager{
		EntityMap:        make(map[uint32]*Entity),
		codeCounter:      0,
		sysManager:       sysManager,
		ObjPools:         NewObjPools(),
		removeEntities:   NewSimpleContainer(5000),
		RemoveEntityFunc: nil,
	}
	em.ObjPools.SetNewInstanceFunc(f)
	return em
}

//NewEntity get entity from objpools, and initialize entity
func (em *EntityManager) NewEntity(initFunc func(*Entity)) *Entity {

	newEntity := em.ObjPools.Get((*Entity)(nil)).(*Entity)
	newEntity.ObjPools = em.ObjPools

	for {
		em.codeCounter++
		_, ex := em.EntityMap[em.codeCounter]
		if !ex {
			break
		}
	}
	newEntity.Id = em.codeCounter
	newEntity.Active = true
	em.EntityMap[newEntity.Id] = newEntity

	initFunc(newEntity)

	em.sysManager.AddEntity(newEntity)

	return newEntity

}

//RegistToRemoveEntities Add entity to RemoveReadycontainer
func (em *EntityManager) RegistToRemoveEntities(en *Entity) {
	if en.NeedRemove {
		return
	}
	en.NeedRemove = true
	en.Active = false
	em.removeEntities.Add(en)
}

func (em *EntityManager) Reset() {
	for _, e := range em.EntityMap {
		em.RegistToRemoveEntities(e)
	}
	em.RemoveEntities()
}

//RemoveEntities remove from RemoveReadycontainer
func (em *EntityManager) RemoveEntities() {
	em.removeEntities.Each(func(idx int, item interface{}) {

		en := item.(*Entity)

		if en.SysRefCount == 0 {
			if em.RemoveEntityFunc != nil {
				em.RemoveEntityFunc(en)
			}
			delete(em.EntityMap, en.Id)
			en.Reset()
			en.Active = false
			em.ObjPools.Put(en)
			em.removeEntities.RemoveIdx(idx)

		}

	})

}
