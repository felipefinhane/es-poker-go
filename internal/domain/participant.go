package domain

import "time"

type Participant struct {
	ID       string
	Name     string
	Email    string
	JoinedAt time.Time
}
