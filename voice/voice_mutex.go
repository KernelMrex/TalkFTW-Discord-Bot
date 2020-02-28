package voice

import "sync"

type ServersVoiceActivity struct {
	servers       map[string]*sync.Mutex
	checkingMutex *sync.Mutex
}

func NewServersVoiceActivity() *ServersVoiceActivity {
	return &ServersVoiceActivity{
		servers:       make(map[string]*sync.Mutex),
		checkingMutex: &sync.Mutex{},
	}
}

func (s *ServersVoiceActivity) ServerLock(guid string) {
	// Checking if mutex have been created
	s.checkingMutex.Lock()
	serverMutex, ok := s.servers[guid]
	if !ok {
		serverMutex = &sync.Mutex{}
		s.servers[guid] = serverMutex
	}
	s.checkingMutex.Unlock()

	// Locking server mutex
	serverMutex.Lock()
}

func (s *ServersVoiceActivity) ServerUnlock(guid string) {
	// Checking if mutex have been created
	s.checkingMutex.Lock()
	serverMutex, ok := s.servers[guid]
	s.checkingMutex.Unlock()
	if !ok {
		return
	}

	// Locking server mutex
	serverMutex.Unlock()
}
