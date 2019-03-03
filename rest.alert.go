package sdk

import (
	"encoding/json"
	"fmt"
)

// CreateAlertNotification creates a new alert notification.
// It reflects POST /api/alert-notifications API call.
func (r *Client) CreateAlertNotification(an AlertNotification) (AlertNotification, error) {
	var (
		raw  []byte
		resp AlertNotification
		err  error
	)
	if raw, err = json.Marshal(an); err != nil {
		return AlertNotification{}, err
	} else if raw, _, err = r.post("api/alert-notifications", nil, raw); err != nil {
		return AlertNotification{}, err
	}
	return resp, json.Unmarshal(raw, &resp)
}

// GetAlertNotifications loads all alert notifications.
// It reflects GET /api/alert-notifications API call.
func (r *Client) GetAlertNotifications() ([]AlertNotification, error) {
	var (
		raw  []byte
		ds   []AlertNotification
		code int
		err  error
	)
	if raw, code, err = r.get("api/alert-notifications", nil); err != nil {
		return nil, err
	} else if code != 200 {
		return nil, fmt.Errorf("HTTP error %d: returns %s", code, raw)
	}
	return ds, json.Unmarshal(raw, &ds)
}
