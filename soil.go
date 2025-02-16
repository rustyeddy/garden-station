package main

// type soil struct {
// 	ID  int `json:"id"`
// 	Dry int `json:"dry"`
// 	Wet int `json:"wet"`
// }

// var (
// 	soilSensors [4]soil
// )

// func init() {
// 	for i := 0; i < stationCount; i++ {
// 		soilSensors[i] = soil{
// 			ID:  i,
// 			Dry: 20,
// 			Wet: 80,
// 		}
// 	}
// }

// func (s *soil) fetchData(readQ <-chan float64) {
// 	running := false
// 	for running {
// 		// select {
// 		// case val := <-readQ:
// 		// 	s.Publish(messanger.TopicData("soil/"+s.ID), val)

// 		// case <-done:
// 		// 	running = false
// 		// }
// 	}
// 	return
// }
