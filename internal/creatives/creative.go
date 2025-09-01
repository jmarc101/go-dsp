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

var MemCreatives = []Creative{
	{
		ID:         "cr-001",
		Advertiser: "Nike",
		Width:      300,
		Height:     250,
		MimeTypes:  []string{"image/jpeg", "image/png"},
		CampaignID: "cmp-101",
		PriceFloor: 0.60, // DSP won’t bid lower than this
		Markup:     `<img src="https://ads.cdn.example.com/nike-300x250.jpg" width="300" height="250" />`,
	},
	{
		ID:         "cr-002",
		Advertiser: "Apple",
		Width:      728,
		Height:     90,
		MimeTypes:  []string{"image/jpeg"},
		CampaignID: "cmp-102",
		PriceFloor: 1.10,
		Markup:     `<img src="https://ads.cdn.example.com/apple-728x90.jpg" width="728" height="90" />`,
	},
	{
		ID:         "cr-003",
		Advertiser: "CocaCola",
		Width:      300,
		Height:     600,
		MimeTypes:  []string{"image/gif", "image/png"},
		CampaignID: "cmp-103",
		PriceFloor: 0.85,
		Markup:     `<img src="https://ads.cdn.example.com/coke-300x600.gif" width="300" height="600" />`,
	},
	{
		ID:         "cr-004",
		Advertiser: "Amazon",
		Width:      160,
		Height:     600,
		MimeTypes:  []string{"image/png"},
		CampaignID: "cmp-104",
		PriceFloor: 0.40,
		Markup:     `<img src="https://ads.cdn.example.com/amazon-160x600.png" width="160" height="600" />`,
	},
	{
		ID:         "cr-005",
		Advertiser: "Spotify",
		Width:      300,
		Height:     250,
		MimeTypes:  []string{"video/mp4"},
		CampaignID: "cmp-105",
		PriceFloor: 1.50,
		Markup:     `<video width="300" height="250" autoplay muted><source src="https://ads.cdn.example.com/spotify-ad.mp4" type="video/mp4"></video>`,
	},
}
