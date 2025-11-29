package entity

// Fruit represents a single fruit document exposed to the web layer.
// Field names are capitalized for JSON so they are easy to use in JS.
type Fruit struct {
	Nama     string  `json:"Nama"`
	Kategori string  `json:"Kategori"`
	Harga    float64 `json:"Harga"`
}
