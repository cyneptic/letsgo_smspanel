package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/url"

	"github.com/google/uuid"
)

var (
	DisableUserEndpoint = "disable-user"
	IsAdminEndpoint     = "is-admin"
)

type AdminActionClient struct {
	address string
}

type DisableUserRequest struct {
	TargetID string `json:"target_id"`
	Toggle   bool   `json:"toggle"`
}

func NewAdminActionClient() *AdminActionClient {
	return &AdminActionClient{
		address: "http://localhost:8001/",
	}
}

func (c *AdminActionClient) DisableUserAccount(targetId uuid.UUID, toggle bool) error {
	cl := &http.Client{
		Transport: &http.Transport{},
	}

	payload, err := json.Marshal(DisableUserRequest{
		TargetID: targetId.String(),
		Toggle:   toggle,
	})

	if err != nil {
		return err
	}

	url, err := url.Parse(c.address + DisableUserEndpoint)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", url.String(), bytes.NewBuffer(payload))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	resp, err := cl.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errors.New("Internal Server Error: Bad Request")
	}

	return nil
}

func (c *AdminActionClient) IsAdmin(userid uuid.UUID) (bool, error) {
	cl := &http.Client{
		Transport: &http.Transport{},
	}

	url, err := url.Parse(c.address + IsAdminEndpoint)
	if err != nil {
		return false, err
	}

	q := url.Query()
	q.Set("user_id", userid.String())
	url.RawQuery = q.Encode()

	req, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		return false, err
	}

	resp, err := cl.Do(req)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	var result bool
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	return result, nil
}
