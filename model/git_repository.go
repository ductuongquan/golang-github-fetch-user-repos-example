package model

import (
	"encoding/json"
	"io"
)

// GitRepository we only list some required fields in this exam and omit others
type GitRepository struct {
	Id              int64               `json:"id"`
	Name            string              `json:"name"`
	Description     string              `json:"description"`
	StargazersCount int64               `json:"stargazers_count"`
	Owner           *GitRepositoryOwner `json:"owner"`
}

// GitRepositoryOwner we only list some required fields in this exam and omit others
type GitRepositoryOwner struct {
	Id    int64  `json:"id"`
	Login string `json:"login"`
}

func GitRepositoryFromJson(data io.Reader) *GitRepository {
	var content *GitRepository
	json.NewDecoder(data).Decode(&content)
	return content
}

// NOTE: API schema required has some differences from Github API response
// So we need a type for our API response
type GitRepositoryResponse struct {
	Id              int64  `json:"id"`
	Name            string `json:"name"`
	Description     string `json:"description"`
	StargazersCount int64  `json:"stargazersCount"`
	OwnerId         int64  `json:"ownerId"`
	OwnerLogin      string `json:"ownerLogin"`
}

// ToClientResponseItem A method to convert from Github response item to our client response item
func (o *GitRepository) ToClientResponseItem() *GitRepositoryResponse {
	return &GitRepositoryResponse{
		o.Id,
		o.Name,
		o.Description,
		o.StargazersCount,
		o.Owner.Id,
		o.Owner.Login,
	}
}

func GitRepositoryResponseListToJson(t []*GitRepositoryResponse) string {
	b, _ := json.Marshal(t)
	return string(b)
}
