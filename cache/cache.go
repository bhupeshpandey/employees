package cache

import (
	model "github.com/bhupeshpandey/employees/model"
)

func New(cacheConfig *model.CacheConfig) model.Cache {
	return NewLRU(cacheConfig.Size)
}


//func New()
