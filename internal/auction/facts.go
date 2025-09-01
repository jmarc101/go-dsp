package auction

type Facts struct {
	// Impression info
	ImpID        string   // Impression ID
	Width        int      // Requested width
	Height       int      // Requested height
	BidFloor     float64  // Publisher's floor for this impression
	AllowedMimes []string // Allowed MIME types from imp.banner.mimes

	// User / identity info
	UserID   string            // DSP cookie / ID
	BuyerUID string            // Synced Prebid/SSP ID
	EIDs     map[string]string // Map of external IDs, e.g. {"uid2": "UID2-xxx", "id5": "ID5-yyy"}
	Country  string            // Geo country
	Region   string            // Geo region/state
	City     string            // Geo city
	Lat      float64           // Geo latitude
	Lon      float64           // Geo longitude

	// Device info
	DeviceUA  string // User Agent
	DeviceIP  string // IP address
	DeviceIFA string // Advertising ID (IDFA/AAID)

	// Auction / timing
	TMax          int64   // Max time allowed (ms)
	TransportP95  float64 // 95th percentile latency from SSP / bid request transport
	TimeRemaining int64   // Remaining ms before auction deadline

	// Internal / derived
	CampaignsEligible []string // Campaign IDs eligible for this bid (based on size, floor, MIME)
	Score             float64  // Optional: precomputed score for ranking creatives
}

func BuildFacts(br *BidRequest) (facts []Facts, err error) {
	facts = make([]Facts, len(br.Imp))

	for _, imp := range br.Imp {
		facts = append(facts, &Facts{
			ImpID: imp.ID,
			Width: int(*imp.Banner.W),
		})
	}

}
