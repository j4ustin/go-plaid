package plaid

import "testing"

func TestAddUser(t *testing.T) {
	auth := UseAuth(testConfiguration(t))
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
	auth := UseAuth(testConfiguration(t))
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

func TestMfaStep(t *testing.T) {
	auth := UseAuth(testConfiguration(t))
	accts, _, _, err := auth.MfaStep("test_chase", "1234")
	if err != nil {
		t.Fatal(err)
	}
	if len(accts) != 4 {
		t.Errorf("Expected 4 accounts got %v", len(accts))
	}
}

func TestUpdate(t *testing.T) {
	auth := UseAuth(testConfiguration(t))
	accts, _, mfa, err := auth.UpdateUser("plaid_test", "plaid_good", "test_wells")
	if err != nil {
		t.Fatal(err)
	}
	if len(accts) != 4 {
		t.Errorf("Expected 4 accounts got %v", len(accts))
	}
	if mfa.Type != "" {
		t.Errorf("Expected to get no mfa info back")
	}
}

func TestDelete(t *testing.T) {
	auth := UseAuth(testConfiguration(t))
	if err := auth.DeleteUser("test_wells"); err != nil {
		t.Fatal(err)
	}
}
