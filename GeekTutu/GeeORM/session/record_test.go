package session

import (
	"geeorm/log"
	"testing"
)

var (
	user1 = &User{"Tom", 18}
	user2 = &User{"Sam", 25}
	user3 = &User{"Jack", 25}
)

func testRecordInit(t *testing.T) *Session {
	t.Helper()
	s := NewSession().Model(&User{})
	err1 := s.DropTable()
	err2 := s.CreateTable()
	_, err3 := s.Insert(user1, user2)
	if err1 != nil || err2 != nil || err3 != nil {
		t.Fatal("failed init test records", err1, err2, err3)
	}
	return s
}

func TestSession_Insert(t *testing.T) {
	s := testRecordInit(t)
	affected, err := s.Insert(user3)
	if err != nil || affected != 1 {
		t.Fatal("failed to create record")
	}
}

func TestSession_Find(t *testing.T) {
	s := testRecordInit(t)
	var users []User
	if err := s.Find(&users); err != nil || len(users) != 2 {
		t.Fatal("failed to query all")
	}
	log.Infof("Get users %v", users)
}

func TestLimit(t *testing.T) {
	s := testRecordInit(t)
	var users []User
	if err := s.Limit(1).Find(&users); err != nil || len(users) != 1 {
		t.Fatal("failed to query with limit")
	}
}

func TestSession_Update(t *testing.T) {
	s := testRecordInit(t)
	affected, err := s.Where("Name = ?", "Tom").Update("Age", "30")
	if err != nil {
		t.Fatal("failed to update", err)
	}
	u := &User{}
	if err = s.OrderBy("Age DESC").First(u); affected != 1 || err != nil || u.Age != 30 {
		t.Fatal("failed to update", affected, err, u.Age)
	}
}

func TestSession_DeleteAndCount(t *testing.T) {
	s := testRecordInit(t)
	affected, err := s.Where("Name = ?", "Tom").Delete()
	if err != nil {
		t.Fatal("failed to delete", err)
	}
	count, err := s.Count()
	if err != nil {
		t.Fatal("failed to count")
	}
	if affected != 1 || count != 1 {
		t.Fatal("failed to delete or count")
	}
}
