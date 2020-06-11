package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"encoding/json"
	"log"

	"github.com/anjaneyulubatta505/gin-graphql-postgres/graph/generated"
	"github.com/anjaneyulubatta505/gin-graphql-postgres/graph/model"
)

func (r *queryResolver) Analyze(ctx context.Context, code *model.Code, settings *model.InputSettings) (*model.Analytics, error) {
	jsonData := []byte(`
    {
        "data": {
            "code": "5116",
            "unit": 1,
            "description": "Level 6 Musculoskeletal Procedures",
            "dos": "",
            "applyGAF": 0,
            "startYear": 2018,
            "endYear": 2020,
            "providerType": "Facility",
            "placeOfService": ""
        },
        "settings": {
            "minInflationyear": 2013,
            "maxInflationyear": 2020,
            "typeRequest": "future_needs",
            "typeClaim": "personal_injury",
            "typeReport": "",
            "state": "CO",
            "latitude": 39.566700,
            "longitude": -104.859962,
            "distance": 30,
            "radius": 30,
            "percentile": 80,
            "gaf": 0,
            "minPatient": 11,
            "multiplier": null,
            "billDatabase": "",
            "providerType": "Facility",
            "placeOfService": "",
            "typeMedicalClaim": "bill_category_outpatient",
            "billType": "",
            "billRepeated": null,
            "code": "5116",
            "unit": 1,
            "applyGAF": 0,
            "startYear": 2018,
            "endYear": 2020
        },
        "results": {
            "query": "SELECT hospital.hospitalID AS providerID, hospital.facilityName AS providerLastOrgName, 'O' AS placeOfService, '' AS providerFirstName, hospital.npi AS credentials, '' AS providerCity, '' AS providerZip, '' AS providerState, '' AS providerType, outpatient.capcCode AS hcpcsCode, '' AS hcpcsDescription, outpatient.totalCharges AS averageSubmittedCharge, outpatient.unitsOfService AS lineServicesCount, outpatient.numberOfPatients AS beneUniqueCount, hospital.latitude, hospital.longitude FROM outpatient INNER JOIN hospital ON outpatient.cmsCertificationNumber = hospital.cmsCertificationNumber WHERE outpatient.capcCode = '5116' AND outpatient.totalCharges > 0 ORDER BY averageSubmittedCharge DESC",
            "applyGAF": 1,
            "data": {
                "radius": "National",
                "totalProvider": 73,
                "totalPatient": 1235,
                "percentileCharge": 76212.9,
                "percentileChargeWithGAF": 0,
                "percentileChargeWithInflation": 79441.2992912,
                "percentileChargeWithInflationAndGAF": 0,
                "gafCharge": -79441.2992912,
                "totalCharge": 0,
                "totalChargePerUnit": 0,
                "applyMultiProcedureDiscount": 0,
                "data": [{
                    "providerID": "00004229",
                    "providerName": "Brookwood Baptist Medical Center",
                    "placeOfService": "O",
                    "credentials": "1003260480",
                    "totalPatient": 12,
                    "totalProvider": 1,
                    "distance": 0,
                    "charge": 177292.18,
                    "city": "",
                    "state": "",
                    "providerType": ""
                }, {
                    "providerID": "00004516",
                    "providerName": "Stanford Health Care",
                    "placeOfService": "O",
                    "credentials": "1871543215",
                    "totalPatient": 11,
                    "totalProvider": 2,
                    "distance": 0,
                    "charge": 165333.62,
                    "city": "",
                    "state": "",
                    "providerType": ""
                }]
            }
        }
    }
    `)

	var analytics model.Analytics
	err := json.Unmarshal(jsonData, &analytics)
	if err != nil {
		log.Println(err)
	}
	return &analytics, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
