package storage

import (
	"encoding/json"
	"os"
)

type Notable struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	TradeID string `json:"trade_id"`
}

type MemoryDB struct {
	Notables    []Notable
	TradeIDToID map[string]int
}

func LoadNotables(filePath string) (*MemoryDB, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var list []Notable
	if err := json.NewDecoder(file).Decode(&list); err != nil {
		return nil, err
	}

	tradeMap := make(map[string]int)
	for _, n := range list {
		tradeMap[n.TradeID] = n.ID
	}

	return &MemoryDB{
		Notables:    list,
		TradeIDToID: tradeMap,
	}, nil
}