package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type Post struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Content   string    `json:"content"`
}

func (cfg *apiConfig) handlerPostsCreate(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Content string `json:"content"`
	}

	type response struct {
		Post
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Was unable to decode parameters", err)
		return
	}

	post, err := cfg.db.CreatePost(r.Context(), params.Content)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Was unable to create post", err)
		return
	}

	respondWithJSON(w, http.StatusCreated, response{
		Post: Post{
			ID:        post.ID,
			CreatedAt: post.CreatedAt,
			UpdatedAt: post.UpdatedAt,
			Content:   post.Content,
		},
	})
}

func (cfg *apiConfig) Posts() ([]Post, error) {
	dbPosts, err := cfg.db.GetPosts(context.Background())
	if err != nil {
		return nil, fmt.Errorf("unable to get posts from database", err)
	}

	posts := []Post{}

	for _, dbPost := range dbPosts {
		posts = append(posts, Post{
			ID:        dbPost.ID,
			CreatedAt: dbPost.CreatedAt,
			UpdatedAt: dbPost.UpdatedAt,
			Content:   dbPost.Content,
		})
	}

	return posts, nil
}

func (cfg *apiConfig) handlerPostsGet(w http.ResponseWriter, r *http.Request) {
	posts, err := cfg.Posts()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Unable to get posts", err)
		return
	}

	respondWithJSON(w, http.StatusOK, posts)
}
