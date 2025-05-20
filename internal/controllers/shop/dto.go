package shop

type ShopData struct {
	Name  string     `json:"name"`
	Items []ShopItem `json:"items""`
}

type ShopItem struct {
	Name              string `json:"name"`
	Description       string `json:"description"`
	Encumbrance       int32  `json:"encumbrance"`
	Cost              int32  `json:"cost"`
	Unit              string `json:"unit"`
	QuantityAvailable int32  `json:"quantityAvailable"`
	Category          string `json:"category"`
}
