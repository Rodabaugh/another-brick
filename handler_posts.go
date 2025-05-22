package main

import (
	"encoding/json"
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
