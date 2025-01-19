package main

import (
	"testing"
)

func TestRecordFlightTaken(t *testing.T) {
	account := &FrequentFlierAccount{
		id:         "1234567",
		miles:      10000,
		tierPoints: 18,
		status:     StatusSilver,
	}

	account.RecordFlightTaken(1000, 3)

	if account.miles != 11000 {
		t.Errorf("Expected miles to be 11000, got %d", account.miles)
	}

	if account.tierPoints != 21 {
		t.Errorf("Expected tier points to be 21, got %d", account.tierPoints)
	}

	if account.status != StatusGold {
		t.Errorf("Expected status to be Gold, got %s", account.status)
	}
}

func TestNewFrequentFlierAccountFromHistory(t *testing.T) {
	history := []interface{}{
		FrequentFlierAccountCreated{AccountId: "1234567", OpeningMiles: 10000, OpeningTierPoints: 0},
		StatusMatched{NewStatus: StatusSilver},
		FlightTaken{MilesAdded: 2525, TierPointsAdded: 5},
	}

	account := NewFrequentFlierAccountFromHistory(history)

	if account.miles != 12525 {
		t.Errorf("Expected miles to be 12525, got %d", account.miles)
	}

	if account.tierPoints != 5 {
		t.Errorf("Expected tier points to be 5, got %d", account.tierPoints)
	}

	if account.status != StatusSilver {
		t.Errorf("Expected status to be Silver, got %s", account.status)
	}
}
