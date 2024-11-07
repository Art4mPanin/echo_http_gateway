package singleton

func (s *handler) RegisterSingleton(name string, value interface{}) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.mapper[name] = value
}

func (s *handler) GetSingleton(name string) (interface{}, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	val, ok := s.mapper[name]
	return val, ok
}

func GetAndConvertSingleton[T any](name string) (result T, ok bool) {
	singleton := GetSingletonHandler()
	val, ok := singleton.GetSingleton(name)
	if !ok {
		return result, false
	}
	result, ok = val.(T)
	return result, ok
}
