# Demand Side Platform

This is a toy implementation of a DSP (Demand Side Platform) written in Go.
It’s not production-ready — the goal is to explore how real-time bidding (RTB) flows work end-to-end in a simplified way.

## What It Does

- Bid Request Handling: Listens for incoming bid requests (simulating an exchange).
- Fact Extraction: Pulls relevant information such as user ID, country, and device.
- Creative Matching: Checks available creatives and picks one that fits the request.
- Pricing Logic:
- Base price determined by presence of a known ID.
- Adjusted by country factors.
- Identity graph lookup if the user ID is found.
- Latency & Timeout Aware:
- Tracks p95 latency from the transport layer.
- Considers remaining time before the bid request would be refused.
- Decides early whether to launch a bid or skip.

##  Picture Flow

- Receive bid request
- Evaluate latency/time budget
- Choose appropriate flow based on budget
- Read own identity graph
- Extract facts (ID, country, latency budget)
- Check creatives (filter by targeting rules)
- Apply pricing logic (ID, country, identity graph)


