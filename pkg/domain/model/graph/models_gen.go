// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package graph

type Mutation struct {
}

type NewPost struct {
	Text   string `json:"text"`
	UserID int    `json:"userId"`
}

type NewUser struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Query struct {
}
