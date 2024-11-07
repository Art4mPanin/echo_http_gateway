package singleton

import "sync"

var globalSingleton Handler

type handler struct {
	mapper map[string]interface{}
	mu     *sync.RWMutex
}

type Handler interface {
	RegisterSingleton(name string, value interface{})
	GetSingleton(name string) (interface{}, bool)
}

type NotInitializedError struct{}

func (e NotInitializedError) Error() string {
	return "Singleton not initialized, Call singleton.InitializeSingletonHandler() first."
}

type AlreadyInitializedError struct{}

func (e AlreadyInitializedError) Error() string {
	return "Singleton already initialized, Call singleton.InitializeSingletonHandler() only once."
}

func InitializeSingletonHandler() Handler {
	if globalSingleton != nil {
		panic(&AlreadyInitializedError{})
	}
	globalSingleton = &handler{
		mapper: make(map[string]interface{}),
		mu:     &sync.RWMutex{},
	}
	return globalSingleton
}

func GetSingletonHandler() Handler {
	if globalSingleton == nil {
		panic(&NotInitializedError{})
	}
	return globalSingleton
}
