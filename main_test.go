package main

import (
	"testing"
)

func TestFrequentFlierAccount(t *testing.T) {
	history := []interface{}{
		FrequentFlierAccountCreated{AccountId: "1234567", OpeningMiles: 10000, OpeningTierPoints: 0},
		StatusMatched{NewStatus: StatusSilver},
		FlightTaken{MilesAdded: 2525, TierPointsAdded: 5},
		FlightTaken{MilesAdded: 2512, TierPointsAdded: 5},
		FlightTaken{MilesAdded: 5600, TierPointsAdded: 5},
		FlightTaken{MilesAdded: 3000, TierPointsAdded: 3},
	}

	aggregate := NewFrequentFlierAccountFromHistory(history)

	if aggregate.miles != 23637 {
		t.Errorf("Expected miles to be 23637, got %d", aggregate.miles)
	}

	if aggregate.tierPoints != 18 {
		t.Errorf("Expected tier points to be 18, got %d", aggregate.tierPoints)
	}

	if aggregate.status != StatusSilver {
		t.Errorf("Expected status to be Silver, got %s", aggregate.status)
	}

	aggregate.RecordFlightTaken(1000, 3)

	if aggregate.miles != 24637 {
		t.Errorf("Expected miles to be 24637, got %d", aggregate.miles)
	}

	if aggregate.tierPoints != 21 {
		t.Errorf("Expected tier points to be 21, got %d", aggregate.tierPoints)
	}

	if aggregate.status != StatusGold {
		t.Errorf("Expected status to be Gold, got %s", aggregate.status)
	}
}
