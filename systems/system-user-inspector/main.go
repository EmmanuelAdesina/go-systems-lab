package main

import (
	"fmt"
	"time"
)

type SystemUser struct {
	Username string
	UID int
	GID int
	HomeDirectory string
	Shell string
	Groups []string
	LastLogin time.Time
}

// HasHomeDirectory reports whether the user has a home directory configured.
func (u SystemUser) HasHomeDirectory() bool {
	return u.HomeDirectory != ""
}

// IsAdmin reports whether the user belongs to an administrative group.
func (u SystemUser) IsAdmin() bool {
	for _, group := range u.Groups {
		if group == "sudo" || group == "wheel" {
			return true
		}
	}
	return false
}

func SampleUsers() []SystemUser {
	return []SystemUser{
		
		{
			Username: "alice",
			UID: 1001,
			GID: 1001,
			HomeDirectory: "/home/alice",
			Shell: "/bin/bash",
			Groups: []string{"staff"},
			LastLogin: time.Date(2026, time.July, 1, 9, 30, 0, 0, time.UTC),
		},

		{
			Username: "roothelper",
			UID: 1002,
			GID: 1002,
			HomeDirectory: "/home/roothelper",
			Shell: "/bin/bash",
			Groups: []string{"staff", "sudo"},
			LastLogin: time.Date(2026, time.July, 3, 18, 10, 0, 0, time.UTC),
		},

		{
			Username: "systemsvc",
			UID: 1003,
			GID: 1003,
			Shell: "/usr/sbin/nologin",
			Groups: []string{"services"},
			LastLogin: time.Date(2026, time.July, 2, 6, 0, 0, 0, time.UTC),
		},

	}
}

func main() {
	users := SampleUsers()
	for _, user := range users {
		fmt.Println("Username:", user.Username)
		fmt.Println("Is admin:", user.IsAdmin())
		fmt.Println("Has home directory:", user.HasHomeDirectory())
		fmt.Println()
	}
}
