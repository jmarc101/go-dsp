package auction

import openrtb2 "github.com/prebid/openrtb/v20/openrtb2"

type BidRequest = openrtb2.BidRequest
type BidResponse = openrtb2.BidResponse
type Imp = openrtb2.Imp
type Device = openrtb2.Device
type User = openrtb2.User
type SeatBid = openrtb2.SeatBid

/*
     BID REQUEST EXAMPLE
{
  "id": "req-1234567890",              // Unique request ID
  "imp": [
    {
      "id": "1",                       // Impression ID
      "banner": {
        "format": [
          { "w": 300, "h": 250 }       // Banner size (300x250 ad slot)
        ]
      },
      "bidfloor": 0.50,                // Publisher's price floor (USD)
      "bidfloorcur": "USD"
    },
    {
      "id": "2",                       // Another impression
      "banner": {
        "format": [
          { "w": 728, "h": 90 }        // Leaderboard ad slot
        ]
      },
      "bidfloor": 1.20,                // Higher floor for premium slot
      "bidfloorcur": "USD"
    }
  ],
  "site": {
    "domain": "example.com",           // Publisher domain
    "page": "https://example.com/news" // Page where ad will render
  },
  "device": {
    "ua": "Mozilla/5.0 (Windows NT 10.0; Win64; x64)...", // User agent
    "ip": "203.0.113.42",             // Device IP
    "ifa": "AEBE52E7-03EE-455A-B3C4-E57283966239",        // ID for advertisers
    "geo": {
      "lat": 40.7128,                  // Approx latitude
      "lon": -74.0060,                 // Approx longitude
      "country": "USA",                // Country code
      "region": "NY",                  // Region/state
      "city": "New York",              // City
      "zip": "10001"                   // ZIP code
    }
  },
  "user": {
    "id": "user-xyz",                  // DSP-side cookie or ID
    "buyeruid": "buyer-789",           // Synced buyer ID from Prebid
    "eids": [                          // External identity sources
      {
        "source": "uidapi.com",
        "uids": [
          {
            "id": "UID2-abcdef123456", // UID2 identity
            "atype": 1
          }
        ]
      },
      {
        "source": "id5-sync.com",
        "uids": [
          {
            "id": "ID5-xyz-98765",     // ID5 identity
            "atype": 1
          }
        ]
      },
    ]
  },
  "regs": {
    "ext": {
      "gdpr": 1,                       // GDPR applies
      "us_privacy": "1YNN"             // CCPA string
    }
  },
  "tmax": 120,                         // Time in ms allowed for auction
  "cur": ["USD"]                       // Currencies accepted
}

*/
