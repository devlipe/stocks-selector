package model

type Stock struct {
	Company_id          int     `json:"companyId,omitempty"`
	Company_name        string  `json:"companyName,omitempty"`
	Ticker              string  `json:"ticker,omitempty"`
	Price               float64 `json:"price,omitempty"`
	P_L                 float64 `json:"p_L,omitempty"`
	Dy                  float64 `json:"dy,omitempty"`
	P_Vp                float64 `json:"p_VP,omitempty"`
	P_Ebit              float64 `json:"p_Ebit,omitempty"`
	EV_Ebit             float64 `json:"eV_Ebit,omitempty"`
	MargemEbit          float64 `json:"margemEbit,omitempty"`
	MargemBruta         float64 `json:"margemBruta,omitempty"`
	Roe                 float64 `json:"roe,omitempty"`
	Roa                 float64 `json:"roa,omitempty"`
	Roic                float64 `json:"roic,omitempty"`
	Vpa                 float64 `json:"vpa,omitempty"`
	ValorMercado        float64 `json:"valorMercado,omitempty"`
	LiquidezMediaDiaria float64 `json:"liquidezMediaDiaria,omitempty"`
	RankPl              int
	RankRoa             int
	RankEvEbit          int
	RankGeral           int
	ExcludedReason      string
}
