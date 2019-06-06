package sql

import (
	"errors"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// ErrSessionNil means session is nil
var ErrSessionNil = errors.New("session is nil")

// Store is a connecton to MongoDB
type Store struct {
	session  *gorm.DB
	Address  string
	Username string
	Password string
	Database string
}

// GetSession returns sesssion
func (s *Store) GetSession() (*gorm.DB, error) {
	if s.session == nil {
		return nil, ErrSessionNil
	}
	return s.session, nil
}

// Connect connects the db
func (s *Store) Connect() error {
	url := fmt.Sprintf("%s:%s@/%s", s.Username, s.Password, s.Database)
	db, err := gorm.Open("mysql", url)
	if err != nil {
		return err
	}

	s.session = db

	return nil
}

// Close the database connection
func (s *Store) Close() {
	if s.session != nil {
		s.session.Close()
	}
}
