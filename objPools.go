package ecs

import (
	"sync"
)

//ObjPools have sync.Pool per type , and init function per type.
type ObjPools struct {
	pools           map[uint8]*sync.Pool
	initFunc        map[uint8]func(IGetTypeCode)
	newInstanceFunc NewInstanceFromTypeCodeFunc
}

type NewInstanceFromTypeCodeFunc func(uint8) IGetTypeCode

var objPool *ObjPools

func NewObjPools() *ObjPools {

	return &ObjPools{
		pools:    make(map[uint8]*sync.Pool),
		initFunc: make(map[uint8]func(IGetTypeCode)),
	}

}

func (g *ObjPools) SetNewInstanceFunc(f NewInstanceFromTypeCodeFunc) {

	g.newInstanceFunc = f
}

//SetInitFunc set InitFunc Per type, the type is from (*type)(nil)  as
func (g *ObjPools) SetInitFunc(p IGetTypeCode, init func(IGetTypeCode)) {

	g.initFunc[p.GetTypeCode()] = init
}

//GetPool get Pool from type,(*type)(nil) as p
func (g *ObjPools) GetPool(p IGetTypeCode) *sync.Pool {

	typeCode := p.GetTypeCode()
	return g.GetPoolFromTypeCode( typeCode)

}

func (g *ObjPools) GetPoolFromTypeCode( typeCode uint8) *sync.Pool {
	var pool *sync.Pool
	var ok bool
	if pool, ok = g.pools[typeCode]; ok == false {

		pool = &sync.Pool{
			New: func() interface{} {

				newOne := g.newInstanceFunc(typeCode)

				if init, ok := g.initFunc[typeCode]; ok {
					init(newOne)
				}

				return newOne
			},
		}
		g.pools[typeCode] = pool
	}

	return pool
}

//Get get istance using GetPool from the type , (*type)(nil) as p
func (g *ObjPools) GetFromTypeCode(typeCode uint8) IGetTypeCode {

	return g.GetPoolFromTypeCode(typeCode).Get().(IGetTypeCode)
}

//Get get istance using GetPool from the type , (*type)(nil) as p
func (g *ObjPools) Get(p IGetTypeCode) IGetTypeCode {

	return g.GetPool(p).Get().(IGetTypeCode)
}

//Put put instance using GetPool from the type , (*type)(nil) as p
func (g *ObjPools) Put(p IGetTypeCode) {

	g.GetPool(p).Put(p)
}
