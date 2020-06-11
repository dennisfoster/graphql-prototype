// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Analytics struct {
	Data     *Data     `json:"data"`
	Settings *Settings `json:"settings"`
	Results  *Results  `json:"results"`
}

type Code struct {
	Code        string  `json:"code"`
	Unit        int     `json:"unit"`
	Description string  `json:"description"`
	Dos         *string `json:"dos"`
}

type Data struct {
	Code           *string `json:"code"`
	Unit           *int    `json:"unit"`
	Description    *string `json:"description"`
	Dos            *string `json:"dos"`
	ApplyGaf       *int    `json:"applyGAF"`
	StartYear      *int    `json:"startYear"`
	EndYear        *int    `json:"endYear"`
	ProviderType   *string `json:"providerType"`
	PlaceOfService *string `json:"placeOfService"`
}

type InputSettings struct {
	Requesttype      int     `json:"requesttype"`
	Claimtype        int     `json:"claimtype"`
	Address          string  `json:"address"`
	State            string  `json:"state"`
	Latitude         float64 `json:"latitude"`
	Longitude        float64 `json:"longitude"`
	Percentile       int     `json:"percentile"`
	Distance         float64 `json:"distance"`
	Gaf              float64 `json:"gaf"`
	TypeMedicalClaim string  `json:"typeMedicalClaim"`
	ProviderGroup    string  `json:"providerGroup"`
	ListDatabaseID   int     `json:"listDatabaseID"`
	ProviderType     int     `json:"providerType"`
	PlaceOfService   string  `json:"placeOfService"`
	BillType         string  `json:"billType"`
}

type Provider struct {
	ProviderID     *string  `json:"providerID"`
	ProviderName   *string  `json:"providerName"`
	PlaceOfService *string  `json:"placeOfService"`
	Credentials    *string  `json:"credentials"`
	TotalPatient   *int     `json:"totalPatient"`
	TotalProvider  *int     `json:"totalProvider"`
	Distance       *float64 `json:"distance"`
	Charge         *float64 `json:"charge"`
	City           *string  `json:"city"`
	State          *string  `json:"state"`
	ProviderType   *string  `json:"providerType"`
}

type Results struct {
	Query    *string      `json:"query"`
	ApplyGaf *int         `json:"applyGAF"`
	Data     *ResultsData `json:"data"`
}

type ResultsData struct {
	Radius                              *string     `json:"radius"`
	TotalProvider                       *int        `json:"totalProvider"`
	TotalPatient                        *int        `json:"totalPatient"`
	PercentileCharge                    *float64    `json:"percentileCharge"`
	PercentileChargeWithGaf             *float64    `json:"percentileChargeWithGAF"`
	PercentileChargeWithInflation       *float64    `json:"percentileChargeWithInflation"`
	PercentileChargeWithInflationAndGaf *float64    `json:"percentileChargeWithInflationAndGAF"`
	GafCharge                           *float64    `json:"gafCharge"`
	TotalCharge                         *float64    `json:"totalCharge"`
	TotalChargePerUnit                  *float64    `json:"totalChargePerUnit"`
	ApplyMultiProcedureDiscount         *int        `json:"applyMultiProcedureDiscount"`
	Data                                []*Provider `json:"data"`
}

type Settings struct {
	MinInflationyear *int     `json:"minInflationyear"`
	MaxInflationyear *int     `json:"maxInflationyear"`
	TypeRequest      *string  `json:"typeRequest"`
	TypeClaim        *string  `json:"typeClaim"`
	TypeReport       *string  `json:"typeReport"`
	State            *string  `json:"state"`
	Latitude         *float64 `json:"latitude"`
	Longitude        *float64 `json:"longitude"`
	Distance         *int     `json:"distance"`
	Radius           *int     `json:"radius"`
	Percentile       *int     `json:"percentile"`
	Gaf              *float64 `json:"gaf"`
	MinPatient       *int     `json:"minPatient"`
	Multiplier       *float64 `json:"multiplier"`
	BillDatabase     *string  `json:"billDatabase"`
	ProviderType     *string  `json:"providerType"`
	PlaceOfService   *string  `json:"placeOfService"`
	TypeMedicalClaim *string  `json:"typeMedicalClaim"`
	BillType         *string  `json:"billType"`
	BillRepeated     *int     `json:"billRepeated"`
	Code             *string  `json:"code"`
	Unit             *int     `json:"unit"`
	ApplyGaf         *int     `json:"applyGAF"`
	StartYear        *int     `json:"startYear"`
	EndYear          *int     `json:"endYear"`
}