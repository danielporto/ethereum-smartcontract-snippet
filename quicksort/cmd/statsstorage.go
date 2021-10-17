package cmd

import (
	"fmt"
	"sync"
	"sync/atomic"
)


type  StatsStorage struct {
	storage sync.Map
}


func (s *StatsStorage) AddLatencySample(key int64) bool {
	inc := int64(1)
	count, loaded := s.storage.LoadOrStore(key, &inc)
	if loaded {
		atomic.AddInt64(count.(*int64), inc)
	}
	return !loaded
}


func (s *StatsStorage) PrintStatsMap() {
	s.storage.Range(func(key, value interface{}) bool {
		fmt.Println(key,*value.(*int64))
		return true
	})
}