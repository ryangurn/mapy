package models

import "time"

type GeoJSON struct {
	Type       string     `json:"type"`
	ID         string     `json:"id"`
	Properties Properties `json:"properties"`
	Features   []Features `json:"features"`
}

type Properties struct {
	ID           string    `json:"id"`
	Name         string    `json:"name"`
	UpdatedDate  time.Time `json:"updated_date"`
	TimeCreated  time.Time `json:"time_created"`
	Notes        string    `json:"notes"`
	CoverPhotoID any       `json:"cover_photo_id"`
}

type Geometry struct {
	Type        string    `json:"type"`
	Coordinates []float64 `json:"coordinates"`
}

type Features struct {
	Type       string            `json:"type"`
	Geometry   Geometry          `json:"geometry"`
	ID         string            `json:"id"`
	Properties FeatureProperties `json:"properties"`
}

type FeatureProperties struct {
	ID               string    `json:"id"`
	UpdatedDate      time.Time `json:"updated_date"`
	TimeCreated      time.Time `json:"time_created"`
	Deleted          bool      `json:"deleted"`
	Title            string    `json:"title"`
	Public           bool      `json:"public"`
	IsActive         bool      `json:"is_active"`
	Icon             string    `json:"icon"`
	Revision         int       `json:"revision"`
	Notes            string    `json:"notes"`
	Latitude         float64   `json:"latitude"`
	Longitude        float64   `json:"longitude"`
	Elevation        int       `json:"elevation"`
	Attr             any       `json:"attr"`
	TrackID          string    `json:"track_id"`
	Photos           []any     `json:"photos"`
	Order            any       `json:"order"`
	Folder           string    `json:"folder"`
	MarkerType       string    `json:"marker_type"`
	MarkerColor      string    `json:"marker_color"`
	MarkerDecoration string    `json:"marker_decoration"`
}
