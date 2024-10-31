package sessionmanager

import (
	"bytes"
	"errors"
	"os"
	"sync"
	"time"
)

// Store is the interface for session stores.
type Store interface {
	// Delete should remove the session token and corresponding data from the
	// session store. If the token does not exist then Delete should be a no-op
	// and return nil (not an error).
	Delete(token string) (err error)

	// Find should return the data for a session token from the store. If the
	// session token is not found or is expired, the found return value should
	// be false (and the err return value should be nil). Similarly, tampered
	// or malformed tokens should result in a found return value of false and a
	// nil err value. The err return value should be used for system errors only.
	Find(token string) (b []byte, found bool, err error)

	// Commit should add the session token and data to the store, with the given
	// expiry time. If the session token already exists, then the data and
	// expiry time should be overwritten.
	Commit(token string, b []byte, expiry time.Time) (err error)
}

type Data struct {
	data   []byte
	expiry time.Time
}

// type memoSave struct {
// 	data Data
// }

// func NewMemoSave() *memoSave {
// 	return &memoSave{}
// }

// func (s *memoSave) Delete(token string) (err error) {
// 	if _, exists := s.session[token]; exists {

// 		delete(s.session, token)
// 	} else {
// 		return fmt.Errorf("token doesn't exists")
// 	}
// 	return
// }

// func (s *memoSave) Find(token string) (b []byte, found bool, err error) {
// 	if value, exists := s.session.Value[token]; exists {
// 		b = []byte(value)
// 	} else {
// 		err = fmt.Errorf("cannot find data")
// 		return
// 	}

// 	return b, true, nil

// }

// func (s *memoSave) Commit(token string, b []byte, expiry time.Time) (err error) {

// }

type mapStore struct {
	Sessions map[string]*Session
	mu       *sync.RWMutex
}

func (s *mapStore) Put(token string, session *Session) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.Sessions[token] = session

}

func (s *mapStore) Get(token string) (*Session, error) {
	if value, exists := s.Sessions[token]; exists {
		return value, nil
	} else {
		return nil, errors.New("invaid session token")
	}

}

func (s *mapStore) FindAll() (map[string]*Session, error) {
	return s.Sessions, nil
}

func mapStoreInit() mapStore {
	return mapStore{
		Sessions: make(map[string]*Session),
		mu:       &sync.RWMutex{},
	}
}

func (s *mapStore) AddSessionValue(token string, key string, value any) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.Sessions[token].Value[key] = value
}

func (s *mapStore) GetSessionValue(token string, key string) (any, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	if value, exists := s.Sessions[token].Value[key]; exists {
		return value, nil
	}
	return nil, errors.New("invalid key")

}

func (s *mapStore) Commit(token string, b []byte, expiry time.Time) (err error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	file, err := os.OpenFile("sessins.txt", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0664)
	if err != nil {
		return err
	}
	defer file.Close()
	buffer := bytes.NewBuffer(make([]byte, 0, 512))
	_, err = buffer.Write(b)
	if err != nil {
		return err
	}
	buffer.WriteTo(file)
	return nil

}
