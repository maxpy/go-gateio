/*
 * Gate API v4
 *
 * APIv4 provides spot, margin and futures trading operations. There are public APIs to retrieve the real-time market statistics, and private APIs which needs authentication to trade on user's behalf.
 *
 * API version: 4.5.2
 * Contact: support@mail.gate.io
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package gateapi

type Position struct {
	// User ID
	User int64 `json:"user,omitempty"`
	// Futures contract
	Contract string `json:"contract,omitempty"`
	// Position size
	Size int64 `json:"size,omitempty"`
	// Position leverage
	Leverage string `json:"leverage,omitempty"`
	// Position risk limit
	RiskLimit string `json:"risk_limit,omitempty"`
	// Maximum leverage under current risk limit
	LeverageMax string `json:"leverage_max,omitempty"`
	// Maintenance rate under current risk limit
	MaintenanceRate string `json:"maintenance_rate,omitempty"`
	// Position value calculated in settlement currency
	Value string `json:"value,omitempty"`
	// Position margin
	Margin string `json:"margin,omitempty"`
	// Entry price
	EntryPrice string `json:"entry_price,omitempty"`
	// Liquidation price
	LiqPrice string `json:"liq_price,omitempty"`
	// Current mark price
	MarkPrice string `json:"mark_price,omitempty"`
	// Unrealized PNL
	UnrealisedPnl string `json:"unrealised_pnl,omitempty"`
	// Realized PNL
	RealisedPnl string `json:"realised_pnl,omitempty"`
	// History realized PNL
	HistoryPnl string `json:"history_pnl,omitempty"`
	// PNL of last position close
	LastClosePnl string `json:"last_close_pnl,omitempty"`
	// ADL ranking, range from 1 to 5
	AdlRanking int32 `json:"adl_ranking,omitempty"`
	// Current open orders
	PendingOrders int32 `json:"pending_orders,omitempty"`
	CloseOrder PositionCloseOrder `json:"close_order,omitempty"`
}
