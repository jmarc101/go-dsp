package creatives

type Creative struct {
	ID         string
	Advertiser string
	Width      int
	Height     int
	MimeTypes  []string
	CampaignID string
	PriceFloor float64
	Markup     string
}
