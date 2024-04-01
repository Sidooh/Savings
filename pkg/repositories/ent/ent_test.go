package entRepo

import (
	"Savings/ent"
	"Savings/ent/enttest"
	"Savings/utils"
	"context"
	"strconv"
	"testing"
)

func createPersonalAccount(c *ent.Client) (*ent.PersonalAccount, error) {
	return c.PersonalAccount.Create().
		SetAccountID(1).
		SetType("TEST").
		Save(context.Background())
}

func createPersonalAccountRandom(c *ent.Client) (*ent.PersonalAccount, error) {
	return c.PersonalAccount.Create().
		SetAccountID(uint64(utils.RandomInt(1, 99999))).
		SetType(utils.RandomString(10)).
		Save(context.Background())
}

func fetchPersonalAccounts(c *ent.Client) ([]*ent.PersonalAccount, error) {
	return c.PersonalAccount.Query().All(context.Background())
}

func TestPersonalAccount(t *testing.T) {
	c := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1")
	defer c.Close()

	// Fetching empty table
	accs, err := fetchPersonalAccounts(c)
	if err != nil {
		t.Errorf("Failed to fetch account: %s", err)
	}
	if len(accs) != 0 {
		t.Errorf("Expected 1 account, fetched %v accounts", len(accs))
	}

	// Creating record
	_, err = createPersonalAccount(c)
	if err != nil {
		t.Errorf("Failed to create account: %s", err)
	}

	// Fetching created record
	accs, err = fetchPersonalAccounts(c)
	if err != nil {
		t.Errorf("Failed to fetch account: %s", err)
	}
	if len(accs) != 1 {
		t.Errorf("Expected 1 account, fetched %v accounts", len(accs))
	}

	// Testing unique index
	_, err = createPersonalAccount(c)
	if err == nil {
		t.Errorf("Expected unique constraint error")
	}

	_, err = createPersonalAccountRandom(c)
	if err != nil {
		t.Errorf("Failed to create account: %s", err)
	}
}

func BenchmarkAccountCreation(b *testing.B) {
	c := enttest.Open(b, "sqlite3", "file:ent?mode=memory&_fk=1")
	defer c.Close()

	for i := 0; i < b.N; i++ {
		_, err := createPersonalAccountRandom(c)
		if err != nil {
			b.Fatal(err)
		}
	}

}

func BenchmarkAccountsFetching(b *testing.B) {
	c := enttest.Open(b, "sqlite3", "file:ent?mode=memory&_fk=1")
	defer c.Close()

	fetchCases := []int{1, 1000, 100000, 1000000}

	for _, testCase := range fetchCases {
		for i := 0; i < testCase; i++ {
			_, err := createPersonalAccountRandom(c)
			if err != nil {
				b.Fatal(err)
			}
		}

		b.Run(strconv.Itoa(testCase), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_, err := fetchPersonalAccounts(c)
				if err != nil {
					b.Fatal(err)
				}
			}
		})
	}

}
