package services

import (
	"encoding/json"
	"os"
)

type Favorite struct {
	ID    string `json:"id"`
	Type  string `json:"type"`
	Name  string `json:"name"`
	Image string `json:"image"`
}

const favoritesFile = "favorites.json"

func GetAllFavorites() ([]Favorite, error) {

	if _, err := os.Stat(favoritesFile); os.IsNotExist(err) {
		return []Favorite{}, nil
	}

	data, err := os.ReadFile(favoritesFile)
	if err != nil {
		return nil, err
	}

	var favorites []Favorite
	err = json.Unmarshal(data, &favorites)
	if err != nil {
		return nil, err
	}

	return favorites, nil
}

func AddFavorite(id, favType, name, image string) error {

	favorites, err := GetAllFavorites()
	if err != nil {
		return err
	}

	for _, fav := range favorites {
		if fav.ID == id && fav.Type == favType {
			return nil
		}
	}

	newFav := Favorite{
		ID:    id,
		Type:  favType,
		Name:  name,
		Image: image,
	}
	favorites = append(favorites, newFav)

	return saveFavorites(favorites)
}

func RemoveFavorite(id, favType string) error {
	favorites, err := GetAllFavorites()
	if err != nil {
		return err
	}

	var newFavorites []Favorite
	for _, fav := range favorites {
		if !(fav.ID == id && fav.Type == favType) {
			newFavorites = append(newFavorites, fav)
		}
	}

	return saveFavorites(newFavorites)
}

func saveFavorites(favorites []Favorite) error {
	data, err := json.MarshalIndent(favorites, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(favoritesFile, data, 0644)
}

func IsFavorite(id, favType string) bool {
	favorites, err := GetAllFavorites()
	if err != nil {
		return false
	}

	for _, fav := range favorites {
		if fav.ID == id && fav.Type == favType {
			return true
		}
	}
	return false
}
