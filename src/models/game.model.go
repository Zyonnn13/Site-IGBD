package models

type Game struct {
	ID               int          `json:"id"`
	Name             string       `json:"name"`
	Summary          string       `json:"summary"`
	Storyline        string       `json:"storyline"`
	FirstReleaseDate int64        `json:"first_release_date"`
	Rating           float64      `json:"rating"`
	Slug             string       `json:"slug"`
	Category         int          `json:"category"`
	Cover            Cover        `json:"cover"`
	Screenshots      []Screenshot `json:"screenshots"`
	Genres           []Genre      `json:"genres"`
	Platforms        []Platform   `json:"platforms"`
}

type Cover struct {
	URL string `json:"url"`
}

type Screenshot struct {
	URL string `json:"url"`
}

type Genre struct {
	Name string `json:"name"`
}

type Platform struct {
	Name string `json:"name"`
}

type Company struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Logo        Logo   `json:"logo"`
	StartDate   int64  `json:"start_date"`
}

type Logo struct {
	URL string `json:"url"`
}
