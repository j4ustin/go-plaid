package plaid

import "testing"

const (
	testID     = "test_id"
	testSecret = "test_secret"
)

func TestConfigure(t *testing.T) {
	clnt, err := Configure("", "")
	if err != nil {
		t.Fatal(err)
	}
	if clnt.envURL != Development {
		t.Errorf("Expected the default route to point to UAT")
	}
	if _, err = Configure("", "", SetEnvironment("bad")); err == nil {
		t.Error("Expected to get error for bad environment")
	}
	clnt, err = Configure("", "", SetEnvironment(Production))
	if err != nil {
		t.Fatal(err)
	}
	if clnt.envURL != Production {
		t.Errorf("Expected to set route to Production")
	}
}

func testConfiguration(t *testing.T) Client {
	clnt, err := Configure(testID, testSecret, SetEnvironment(Sandbox))
	if err != nil {
		t.Fatal(err)
	}
	return clnt
}
