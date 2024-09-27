package repository

import (
	"encoding/json"
	"log"
	"microservices/core/constant"
	"microservices/src/entity"
	"microservices/src/input"
	"io"
	"bytes"
	"fmt"
	"net/http"

	supa "github.com/nedpals/supabase-go"
)

type TargetRepository interface {
	GetForcast(forecastData input.ForecastRequest) (entity.ForecastData, error)
	GetAnalytics() ([]entity.ForecastData, error)
}

type targetRepository struct {
	supabaseClient *supa.Client
}

func NewTargetRepository(supabaseClient *supa.Client) *targetRepository {
	return &targetRepository{supabaseClient}
}

func (r *targetRepository) GetForcast(forecastData input.ForecastRequest) (entity.ForecastData, error) {
	 // Prepare the payload
	 value, err := json.Marshal(forecastData)
	 if err != nil {
		 return entity.ForecastData{}, err
	 }
 
	 requestURL := constant.ForcastURL()
	 sessionJWT := constant.GetToken()
	 return r.postForecast(requestURL, value,sessionJWT)

}


func (r *targetRepository) GetAnalytics() ([]entity.ForecastData, error) {
	var result []map[string]interface{}
	err := r.supabaseClient.DB.From(constant.TableFiles).Select("*").Execute(&result)

	if err != nil {
		return nil, err
	}


	s := []entity.ForecastData{}

	if len(result) > 0 {
		value, _ := json.Marshal(result)
		_ = json.Unmarshal(value, &s)
	}

	return s, nil
}

func (r *targetRepository) postForecast(requestURL string, payload []byte, jwtToken string) (entity.ForecastData, error) {
    var result entity.ForecastData

    // Create a new request
    req, err := http.NewRequest(http.MethodPost, requestURL, bytes.NewBuffer(payload))
    if err != nil {
        return result, err
    }

    // Set the Content-Type header
    req.Header.Set("Content-Type", "application/json")

    // Set the Authorization header with the JWT token
    req.Header.Set("Authorization", "Bearer "+jwtToken)

    // Perform the request
    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return result, err
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return result, fmt.Errorf("failed to get valid response from API: %s", resp.Status)
    }

    // Read and log the raw response body
    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return result, err
    }
    log.Printf("Raw response body: %s", body)

    // Unmarshal the response
    if err := json.Unmarshal(body, &result); err != nil {
        return result, err
    }

    return result, nil
}
