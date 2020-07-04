package modeltests

import (
	"github.com/thaibuixuanDEV/forum/api/models"
	"gopkg.in/go-playground/assert.v1"
	"log"
	"testing"
	"time"
)

func TestFindAllUsers(t *testing.T) {
	err := refreshUserTable()
	if err != nil {
		log.Fatal(err)
	}
	err = seedUsers()
	if err != nil {
		log.Fatal(err)
	}
	users, err := userInstance.FindAllUsers(server.DB)
	if err != nil {
		t.Errorf("err when getting the users %v", err)
		return
	}
	assert.Equal(t, len(*users), 2)
}

func TestSaveUser(t *testing.T) {
	err := refreshUserTable()
	if err != nil {
		log.Fatal(err)
	}
	newUser := models.User{
		ID:       1,
		Nickname: "nickname1",
		Email:    "email1@gmail.com",
		Password: "thaibuixuan",
	}

	savedUser, err := newUser.SaveUser(server.DB)
	if err != nil {
		t.Errorf("err when get user: %v", err)
		return
	}
	assert.Equal(t, newUser.ID, savedUser.ID)
	assert.Equal(t, newUser.Email, savedUser.Email)
	assert.Equal(t, newUser.Nickname, savedUser.Nickname)
}

func TestFindUserByID(t *testing.T) {
	user, err := seedOneUser()
	if err != nil {
		log.Fatalf("can not seed user: %v", err)
	}

	newUser, err := user.FindUserByID(server.DB, user.ID)
	if err != nil {
		t.Errorf("err when getting user: %v", err)
		return
	}
	assert.Equal(t, user.ID, newUser.ID)
	assert.Equal(t, user.Nickname, newUser.Nickname)
	assert.Equal(t, user.Email, newUser.Email)
}

func TestUpdateAUser(t *testing.T) {

	err := refreshUserTable()
	if err != nil {
		log.Fatal(err)
	}

	user, err := seedOneUser()
	if err != nil {
		log.Fatalf("Cannot seed user: %v\n", err)
	}

	userUpdate := models.User{
		ID:        1,
		Nickname:  "modiUpdate",
		Email:     "modiupdate@gmail.com",
		Password:  "password",
		UpdatedAt: time.Now(),
	}
	updatedUser, err := userUpdate.UpdateAUser(server.DB, user.ID)
	if err != nil {
		t.Errorf("this is the error updating the user: %v\n", err)
		return
	}
	assert.Equal(t, updatedUser.ID, userUpdate.ID)
	assert.Equal(t, updatedUser.Email, userUpdate.Email)
	assert.Equal(t, updatedUser.Nickname, userUpdate.Nickname)
}

func TestDeleteUser(t *testing.T) {
	err := refreshUserTable()
	if err != nil {
		log.Fatalf(err.Error())
	}
	user, err := seedOneUser()
	if err != nil {
		log.Fatal("can not seed user")
	}

	isDelete, err := user.DeleteAUser(server.DB, user.ID)
	if err != nil {
		t.Errorf("error when delete user: %v", err)
		return
	}

	assert.Equal(t, isDelete, int64(1))

}
