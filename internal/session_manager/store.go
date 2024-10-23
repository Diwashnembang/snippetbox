package sessionmanager

import (
	"errors"
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

func (s *mapStore) Commit(token string, value *Session)  {
	s.mu.Lock()
	s.Sessions[token] = value
	s.mu.Unlock()

}

func (s *mapStore) find(token string)(*Session,error){
	if value, exists :=s.Sessions[token]; exists{
		return value, nil
	}else{
		return nil ,errors.New("invaid session token")
	}


}


func mapStoreInit()mapStore{
	return mapStore{
		Sessions: make(map[string]*Session),
		mu: &sync.RWMutex{},
	}
}