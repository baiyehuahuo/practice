package session

import (
	"geeorm/log"
	"testing"
)

type Account struct {
	ID       int `geeorm:"PRIMARY KEY"`
	Password string
}

func (a *Account) BeforeInsert(s *Session) error {
	log.Info("before insert ", a)
	a.ID += 1000
	return nil
}

func (a *Account) AfterQuery(s *Session) error {
	log.Info("after query ", a)
	a.Password = "*****"
	return nil
}

func TestSession_CallMethod(t *testing.T) {
	s := NewSession().Model(&Account{})
	s.DropTable()
	s.CreateTable()
	s.Insert(&Account{1, "123456"}, &Account{2, "qwerty"})
	u := &Account{}
	err := s.First(u)
	if err != nil || u.ID != 1001 || u.Password != "*****" {
		t.Fatal("Failed to call hooks after query, got", u)
	}
}
