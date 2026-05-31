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

	fmt.Println("\n--- First 3 Valid Database Entries ---")
	for i := 0; i < 3; i++ {
		fmt.Printf("ID: %d | Name: %s | TradeID: %s\n", db.Notables[i].ID, db.Notables[i].Name, db.Notables[i].TradeID)
	}
	fmt.Println("-------------------------------------\n")

	testKey1 := db.Notables[10].TradeID
	testKey2 := db.Notables[20].TradeID
	testKey3 := db.Notables[50].TradeID

	idx1, ok1 := db.TradeIDToID[testKey1]
	idx2, ok2 := db.TradeIDToID[testKey2]
	idx3, ok3 := db.TradeIDToID[testKey3]

	if !ok1 || !ok2 || !ok3 {
		log.Fatalf("Error: Failed to find valid map keys in database registration.")
	}

	matrixLocation := combinatorics.GetCombinationIndex(idx1, idx2, idx3)
	
	fmt.Printf("Real Dynamic Test Keys: [%s, %s, %s]\n", testKey1, testKey2, testKey3)
	fmt.Printf("Translated Match IDs:   [%d, %d, %d]\n", idx1, idx2, idx3)
	fmt.Printf("Mapped 1D Plotting Array Location: %d / 4455099\n", matrixLocation)
}