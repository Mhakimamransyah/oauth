package entities

import "time"

const (
	ManualRegister = iota
	GithubAccount
	GoogleAccount
)

type User struct {
	Id           int
	Name         string
	Email        string
	Image        string
	Token        string
	CountLogin   int
	Account      int
	RegisteredAt time.Time
}
