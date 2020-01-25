package storage

import (
	"strings"
	"testing"
)

func TestCheckQuotaNotifiesUser(t *testing.T) {
	// Save and restore original notifyUser func.
	savedNotifyUser := notifyUser
	defer func() {
		notifyUser = savedNotifyUser
	}()

	// Install a fake test notifiedUser func.
	var notifiedUser, notifiedMsg string
	notifyUser = func(user, msg string) {
		notifiedUser, notifiedMsg = user, msg
	}

	const user = "joe@example.org"
	usage[user] = 980000000 // Simulate 98% used condition

	CheckQuota(user)
	if notifiedUser == "" && notifiedMsg == "" {
		t.Fatal("notifyUser not called")
	}

	if notifiedUser != user {
		t.Fatalf("wrong user (%s) notified, want %s", notifiedUser, user)
	}

	const wantSubstring = "98% of your quota"
	if !strings.Contains(notifiedMsg, wantSubstring) {
		t.Errorf("unexpected notification message <<%s>>, want substring %q", notifiedMsg, wantSubstring)
	}
}
