package cmd

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"sort"
	"sync"
	"sync/atomic"
)

type StatsStorage struct {
	storage       sync.Map
	max_latency   int64
	total_samples int64
	total_keys	  int64
	latencies 	  []int64
}

func (s *StatsStorage) StoreLatencySample(latency int64) bool {
	inc := int64(1)
	s.total_samples++
	if latency > s.max_latency {
		s.max_latency = latency
	}
	count, loaded := s.storage.LoadOrStore(latency, &inc)
	if loaded {
		// latency already exists, increment the counter
		atomic.AddInt64(count.(*int64), inc)
	} else {
		// new key (latency sample) added
		s.total_keys++
	}

	return !loaded
}

// GetLatencyValuesFlattened
// this version uses more memory and is not required as the compute methods are aware of repetitions
func (s *StatsStorage) GetLatencyValuesFlattened() *[]int64 {
	if s.latencies != nil {
		//sorted previously and still valid
		log.Infof("Previous latency slice is still valid. return")
		return &s.latencies
	}

	s.latencies = []int64{}
	//put the elements there
	s.storage.Range(func(key, value interface{}) bool {
		repetitions := *value.(*int64)
		var i int64
		for i = 0; i < repetitions; i++ {
			_ = append(s.latencies, *key.(*int64))
		}
		return true
	})
	//sort it
	sort.Slice(s.latencies, func(i, j int) bool { return s.latencies[i] < s.latencies[j] })
	return &s.latencies
}

func (s *StatsStorage) GetLatencyValues() *[]int64 {

	if s.latencies != nil {
		//sorted previously and still valid
		log.Infof("Previous latency slice is still valid. return")
		return &s.latencies
	}

	// create the slice
	s.latencies = []int64{}
	//iterate over the map, read the latencies and store them into the slice
	s.storage.Range(func(key, value interface{}) bool {
		//log.Infof("Appending new value to the slice: %v\n", key.(int64))
		s.latencies = append(s.latencies, key.(int64))
		return true
	})

	//sort it
	sort.Slice(s.latencies, func(i, j int) bool { return s.latencies[i] < s.latencies[j] })
	return &s.latencies
}

func (s *StatsStorage) GetAverage(percent bool) int64 {
	latencies := *s.GetLatencyValues()
	latency := int64(0)
	consumedSamples := int64(0)
	limit := int64(0)

	if percent {
		//default is take 10%, and thus compute the 90%
		limit = s.total_samples / 10
	}

	pos := s.total_samples - limit
	//find the position in the flat array
	for i := 0 ; i < len(latencies); i++ {
		latencyValue := latencies[i]
		latencySampleQty, _ := s.storage.Load(latencyValue)
		latencySamples := *latencySampleQty.(*int64)

		if consumedSamples + latencySamples >= pos {
			qty := consumedSamples + latencySamples - pos
			latency += (latencySamples -qty)*latencyValue
			break
		} else {
			consumedSamples += latencySamples
			latency += latencySamples * latencyValue
		}
	} //for
	return latency / pos
}

func (s *StatsStorage) GetMedian(limit bool) int64 {
	return 0

}

func (s *StatsStorage) GetSTD(limit bool) int64 {
	return 0

}

func (s *StatsStorage) GetMax(limit bool) int64 {
	return 0

}

func (s *StatsStorage) GetPercentile(limit bool) int64 {
	return 0
}

func (s *StatsStorage) GetLatencies(limit bool) int64 {
	return 0
}

func (s *StatsStorage) PrintStatsMap() {
	s.storage.Range(func(key, value interface{}) bool {
		fmt.Println(key, *value.(*int64))
		return true
	})


	a := s.GetLatencyValues()
	fmt.Println("Latencies:", a)

	fmt.Printf("Max Latency: %v, Total samples: %v, Total keys: %v, average: %v  \n", s.max_latency, s.total_samples, s.total_keys, s.GetAverage(false))
}
