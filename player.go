package tcpx

import (
	"sync"
)

type PlayerManager struct {
	sync.RWMutex
	data map[int]map[string]interface{}
}

var playerManager *PlayerManager

func init() {
	playerManager = &PlayerManager{
		RWMutex:sync.RWMutex{},
		data:make(map[int]map[string]interface{}),
	}
}

//PlayerManager 玩家管理对象.
func GetPlayerManager() *PlayerManager {
	return playerManager
}

//Get 获取玩家一个属性.
func (p *PlayerManager) Get(playerid int, keys ...string) interface{} {
	defer p.RUnlock()
	p.RLock()
	_, ok := p.data[playerid]
	if !ok {
		p.data[playerid] = make(map[string]interface{})
	}

	pdata := p.data[playerid]
	if len(keys) > 0 {
		val,ok := pdata[keys[0]]
		if !ok {
			return nil
		}
		return val
	}
	return pdata
}

//Set 设置玩家属性.
func (p *PlayerManager) Set(playerid int, key string, value interface{}) {
	defer p.Unlock()
	p.Lock()
	_, ok := p.data[playerid]
	if !ok {
		p.data[playerid] = make(map[string]interface{})
	}
	p.data[playerid][key] = value
}

//Remove 移除一个玩家.
func (p *PlayerManager) Remove(playerid int) {
	defer p.Unlock()
	p.Lock()
	delete(p.data, playerid)
}