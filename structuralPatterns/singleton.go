package structuralpatterns

import "sync"

// import "sync"

// var lock = &sync.Mutex{}

// type single struct {
// }

// var singleInstance *single

// func getInstance() *single {
// 	if singleInstance == nil {
// 		lock.Lock()
// 		defer lock.Unlock()
// 		if singleInstance == nil {
// 			return &single{}
// 		}
// 	}
// 	return singleInstance
// }

// another example
var once sync.Once

type single struct {
}

var singleInstance *single

func GetInstance() *single {
	if singleInstance == nil {
		once.Do(
			func() {
				singleInstance = &single{}
			})
	}
	return singleInstance
}
