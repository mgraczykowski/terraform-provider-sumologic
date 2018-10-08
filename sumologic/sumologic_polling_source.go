package sumologic

import (
	"encoding/json"
	"fmt"
)

type PollingSource struct {
	Source
	ContentType   string               `json:"contentType"`
	ScanInterval  int                  `json:"scanInterval"`
	Paused        bool                 `json:"paused"`
	URL           bool                 `json:"url"`
	ThirdPartyRef PollingThirdPartyRef `json:"thirdPartyRef,omitempty"`
}

type PollingThirdPartyRef struct {
	Resources []PollingResource `json:"resources"`
}

type PollingResource struct {
	ServiceType    string                `json:"serviceType"`
	Authentication PollingAuthentication `json:"authentication"`
	Path           PollingPath           `json:"path"`
}

type PollingAuthentication struct {
	Type    string `json:"type"`
	AwsID   string `json:"awsId"`
	AwsKey  string `json:"awsKey"`
	RoleARN string `json:"roleARN"`
}

type PollingPath struct {
	Type           string `json:"type"`
	BucketName     string `json:"bucketName"`
	PathExpression string `json:"pathExpression"`
}

func (s *Client) CreatePollingSource(source PollingSource, collectorID int) (int, error) {

	type PollingSourceMessage struct {
		Source PollingSource `json:"source"`
	}

	request := PollingSourceMessage{
		Source: source,
	}

	urlPath := fmt.Sprintf("collectors/%d/sources", collectorID)

	body, err := s.Post(urlPath, request)

	if err != nil {
		return -1, err
	}

	var response PollingSourceMessage
	err = json.Unmarshal(body, &response)

	if err != nil {
		return -1, err
	}

	return response.Source.ID, nil
}

func (s *Client) GetPollingSource(collectorID, sourceID int) (*PollingSource, error) {
	urlPath := fmt.Sprintf("collectors/%d/sources/%d", collectorID, sourceID)
	body, _, err := s.Get(urlPath)

	if err != nil {
		return nil, err
	}

	type PollingSourceResponse struct {
		Source PollingSource `json:"source"`
	}

	var response PollingSourceResponse
	err = json.Unmarshal(body, &response)

	if err != nil {
		return nil, err
	}

	return &response.Source, nil
}

func (s *Client) UpdatePollingSource(source PollingSource, collectorID int) error {
	url := fmt.Sprintf("collectors/%d/sources/%d", collectorID, source.ID)

	type PollingSourceMessage struct {
		Source PollingSource `json:"source"`
	}

	request := PollingSourceMessage{
		Source: source,
	}

	_, err := s.Put(url, request)

	return err
}
