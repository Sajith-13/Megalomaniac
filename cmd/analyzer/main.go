package main

import (
	"fmt"
	"log"
	"runtime"

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

	fmt.Println("Allocating 4,455,100 slot market pricing matrix...")
	matrix := storage.NewMarketMatrix()
	
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("Matrix Allocated! System Memory currently in use: %v MiB\n", m.Alloc / 1024 / 1024)

	testKey1 := db.Notables[10].TradeID
	testKey2 := db.Notables[20].TradeID
	testKey3 := db.Notables[50].TradeID

	idx1 := db.TradeIDToID[testKey1]
	idx2 := db.TradeIDToID[testKey2]
	idx3 := db.TradeIDToID[testKey3]

	matrixLocation := combinatorics.GetCombinationIndex(idx1, idx2, idx3)

	fmt.Printf("\nSimulating market discovery for index %d...\n", matrixLocation)
	_ = matrix.InsertPrice(matrixLocation, 150.0) // Listing 1: 150 Chaos
	_ = matrix.InsertPrice(matrixLocation, 135.5) // Listing 2: 135.5 Chaos

	recordedPrices, _ := matrix.GetPrices(matrixLocation)
	fmt.Printf("Retrieved Matrix Prices at Index %d: %v Chaos Orbs\n", matrixLocation, recordedPrices)
}