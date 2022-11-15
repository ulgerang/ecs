package ecs

import (
	"fmt"
	"time"
)

//SystemManager manage systems
type SystemManager struct {
	systems      []ISystem
	previousTime time.Time
}

//NewSystemManager new SystemManager
func NewSystemManager() *SystemManager {

	return &SystemManager{
		previousTime: time.Now(),
	}
}

//AddSystem add System
func (sm *SystemManager) AddSystem(sys ISystem) {

	sm.systems = append(sm.systems, sys)
}

//AddEntity add Entity
func (sm *SystemManager) AddEntity(entity *Entity) {

	for _, sys := range sm.systems {
		sys.GetBaseSystem().CheckAndAdd(entity)
	}
}

//Update
func (sm *SystemManager) Update() time.Duration {

	dt := time.Since(sm.previousTime)

	deltaSec := float32(dt.Seconds())

	if dt.Nanoseconds() == 0 {
		fmt.Println("dt", deltaSec)
	}
	sm.previousTime = time.Now()
	sm.UpdateWithDeltaTime(deltaSec)
	return dt
}

func (sm *SystemManager) UpdateWithDeltaTime(deltaSec float32) {
	for _, sys := range sm.systems {
		sys.Update(deltaSec)
	}
}
