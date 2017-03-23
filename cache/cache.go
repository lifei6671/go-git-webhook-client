package cache

import (
	"sync"
	"go-git-webhook-client/models"
	"errors"
)

var TokenCache = &MemoryCache{ members: make(map[string]models.Member,10)}

type MemoryCache struct {
	sync.RWMutex
	members map[string]models.Member
}

func (c *MemoryCache)Add(name string,member models.Member) {
	c.Lock()
	defer c.Unlock()

	c.members[name] = member
}

func (c *MemoryCache) Delete(name string) {
	c.Lock()
	defer c.Unlock()

	if _,ok := c.members[name];ok {
		delete(c.members,name)
	}
}

func (c *MemoryCache) GetMember(name string) (models.Member,error) {
	c.RLock()
	defer c.RUnlock()
	if item, ok := c.members[name]; ok {
		return item,nil
	}
	return models.Member{},errors.New("Token does not exist.")
}

func (c *MemoryCache) Contains(name string) bool {
	c.RLock()
	defer c.RUnlock()
	 _, ok := c.members[name];
	return ok
}