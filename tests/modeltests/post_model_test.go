package modeltests

import (
	"github.com/thaibuixuanDEV/forum/api/models"
	"gopkg.in/go-playground/assert.v1"
	"log"
	"testing"
	"time"
)

func TestFindAllPosts(t *testing.T) {
	err := refreshUserAndPostTable()
	if err != nil {
		log.Fatalf("")
	}

	_, _, err = seedUserAndPosts()
	if err != nil {
		log.Fatalf("can not seed data: %v", err)
	}

	posts, err := postInstance.FindAllPosts(server.DB)
	if err != nil {
		t.Errorf("can not get all posts: %v", err)
		return
	}
	assert.Equal(t, len(*posts), 2)
}

func TestSavePost(t *testing.T) {
	err := refreshUserAndPostTable()
	if err != nil {
		log.Fatalf("err when refesh User and Post table: %v", err)
	}

	user, err := seedOneUser()
	if err != nil {
		log.Fatalf("can not seed user: %v", err)
	}
	post := models.Post{
		ID:       1,
		Title:    "Title",
		Content:  ":Content",
		AuthorID: user.ID,
	}
	savedPost, err := post.SavePost(server.DB)
	if err != nil {
		t.Errorf("error when save post: %v", err)
		return
	}
	assert.Equal(t, savedPost.ID, post.ID)
	assert.Equal(t, savedPost.Title, post.Title)
	assert.Equal(t, savedPost.Content, post.Content)
	assert.Equal(t, savedPost.AuthorID, post.AuthorID)
}

func TestFindPostByID(t *testing.T) {
	err := refreshUserAndPostTable()
	if err != nil {
		log.Fatalf("can not refresh user and post table: %v", err)
	}
	post, err := seedOneUserAndPost()
	if err != nil {
		log.Fatalf("can not seed user and post: %v", err)
	}

	foundPost, err := post.FindPostByID(server.DB, post.ID)
	if err != nil {
		t.Errorf("get err when findPost by ID: %v", err)
		return
	}
	assert.Equal(t, foundPost.ID, post.ID)
	assert.Equal(t, foundPost.Title, post.Title)
	assert.Equal(t, foundPost.Content, post.Content)
	assert.Equal(t, foundPost.AuthorID, post.AuthorID)
}

func TestUpdateAPost(t *testing.T) {
	err := refreshUserAndPostTable()
	if err != nil {
		log.Fatalf("can not refresh user and post table: %v", err)
	}
	post, err := seedOneUserAndPost()
	if err != nil {
		log.Fatalf("can not seed user and post")
	}

	postUpdate := models.Post{
		ID:        post.ID,
		Title:     "modiTitle",
		Content:   "modiContent",
		AuthorID:  post.AuthorID,
		UpdatedAt: time.Now(),
	}

	updatedPost, err := postUpdate.UpdateAPost(server.DB)
	if err != nil {
		t.Errorf("can not update post: %v", err)
		return
	}
	assert.Equal(t, updatedPost.ID, postUpdate.ID)
	assert.Equal(t, updatedPost.Title, postUpdate.Title)
	assert.Equal(t, updatedPost.Content, postUpdate.Content)
	assert.Equal(t, updatedPost.AuthorID, postUpdate.AuthorID)
}

func TestDeleteAPost(t *testing.T) {
	err := refreshUserAndPostTable()
	if err != nil {
		log.Fatalf("can not refresh user and post table: %v", err)
	}

	post, err := seedOneUserAndPost()
	if err != nil {
		log.Fatalf("cannot seed user and post")
	}

	isDeleted, err := post.DeleteAPost(server.DB, post.ID, post.AuthorID)
	if err != nil {
		t.Errorf("can not delete a post: %v", err)
		return
	}
	assert.Equal(t, isDeleted, int64(1))
}
