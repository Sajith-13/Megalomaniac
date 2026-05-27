package main

import (
	"fmt"
	"log"
	
	"megalo-analyzer/internal/combinatorics"
	"megalo-analyzer/internal/storage"
)

func main() {
	fmt.Println("--- Megalomaniac Combinatorial Engine ---")

	db, err := storage.LoadNotables("data/notables.json")
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	fmt.Printf("Successfully loaded %d nodes into memory storage.\n", len(db.Notables))

	testMods := []string{
		"passive_notable_acrobatics",
		"passive_notable_added_soar",
		"passive_notable_agent_of_destruction",
	}

	idx1 := db.TradeIDToID[testMods[0]]
	idx2 := db.TradeIDToID[testMods[1]]
	idx3 := db.TradeIDToID[testMods[2]]

	matrixLocation := combinatorics.GetCombinationIndex(idx1, idx2, idx3)
	
	fmt.Printf("Test Match IDs: [%d, %d, %d]\n", idx1, idx2, idx3)
	fmt.Printf("Mapped 1D Plotting Array Location: %d / 4455099\n", matrixLocation)
}