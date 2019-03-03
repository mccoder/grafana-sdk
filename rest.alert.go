package sdk

import "encoding/json"

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
