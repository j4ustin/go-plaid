package plaid

import "testing"

func TestAddUser(t *testing.T) {
	clnt := testConfiguration(t)
	auth := UseAuth(clnt)
	accts, atk, mfa, err := auth.AddUser("plaid_test", "plaid_good", "wells", "")
	if err != nil {
		t.Fatal(err)
	}
	if len(accts) != 4 {
		t.Errorf("Expected to get back 4 accounts, got %v", len(accts))
	}
	if atk != "test_wells" {
		t.Errorf("Expected to get back access token \"test_wells\" got %v", atk)
	}
	if mfa.Type != "" {
		t.Errorf("Expected to get back no mfa content")
	}
	_, _, _, err = auth.AddUser("plaid_test", "bad", "wells", "")
	if err == nil {
		t.Errorf("Expected to get back descriptive error. got nil")
	}
	_, atk, mfa, err = auth.AddUser("plaid_test", "plaid_good", "chase", "")
	if err != ErrMfaRequired {
		t.Fatalf("Expected to get back MFA required. Got %v", err)
	}
	if atk != "test_chase" {
		t.Errorf("Expected to get back valid access token for mfa usage")
	}
	if mfa.Type != "device" {
		t.Errorf("Expected to get mfa info back")
	}
}

func TestGetData(t *testing.T) {
	clnt := testConfiguration(t)
	auth := UseAuth(clnt)
	accts, mfa, err := auth.GetData("test_wells")
	if err != nil {
		t.Fatal(err)
	}
	if len(accts) != 4 {
		t.Errorf("Expected to get 4 accounts back")
	}
	if mfa.Type != "" {
		t.Errorf("Expected to get no mfa info back")
	}
}
