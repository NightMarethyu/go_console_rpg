package game

import (
	"encoding/json"
	"fmt"
	"os"
)

func loadJSON[T any](path string) ([]T, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var result []T
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func LoadItems() (map[string]Item, error) {
	items, err := loadJSON[Item]("../data/items.json")
	if err != nil {
		return nil, err
	}
	result := make(map[string]Item)
	for _, item := range items {
		result[item.ID] = item
	}
	return result, nil
}

func LoadBackgrounds() ([]Background, error) {
	return loadJSON[Background]("../data/backgrounds.json")
}

func LoadEnemies() ([]Enemy, error) {
	return loadJSON[Enemy]("../data/enemies.json")
}

func ResolveStartingItems(bg Background, itemMap map[string]Item) ([]Item, error) {
	items := make([]Item, 0, len(bg.StartingItemIDs))
	for _, id := range bg.StartingItemIDs {
		item, ok := itemMap[id]
		if !ok {
			return nil, fmt.Errorf("item ID %s not found in item map", id)
		}
		items = append(items, item)
	}
	return items, nil
}
