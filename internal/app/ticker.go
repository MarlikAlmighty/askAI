package app

import (
	"github.com/MarlikAlmighty/kickHisAss/internal/data"
	"time"
)

// CleanUserData function clearing map with user data
func CleanUserData(users *data.UserData) {

	ticker := time.NewTicker(10 * time.Minute)
	defer ticker.Stop()

	go func() {
		for {
			select {
			case <-ticker.C:
				users.Clear()
			}
		}
	}()
}
