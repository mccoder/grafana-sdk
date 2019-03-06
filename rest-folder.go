package sdk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"
)

func (r *Client) GetAllFolder(limit string) ([]Folder, error) {
	var (
		raw    []byte
		result = make([]Folder, 0)
		code   int
		err    error
	)
	params := url.Values{}
	if "" != limit {
		params.Set("limit", limit)
	}
	if raw, code, err = r.get("/api/folders", params); err != nil {
		return nil, err
	}
	if code != 200 {
		return nil, fmt.Errorf("HTTP error %d: returns %s", code, raw)
	}
	dec := json.NewDecoder(bytes.NewReader(raw))
	dec.UseNumber()
	if err := dec.Decode(&result); err != nil {
		return nil, fmt.Errorf("unmarshal folder error: %s\n%s", err, raw)
	}
	return result, nil
}

func (r *Client) GetFolderByUID(uid string) (Folder, error) {
	var (
		raw    []byte
		result Folder
		code   int
		err    error
	)
	if raw, code, err = r.get("/api/folders/"+uid, nil); err != nil {
		return Folder{}, err
	}
	if code != 200 {
		return Folder{}, fmt.Errorf("HTTP error %d: returns %s", code, raw)
	}
	dec := json.NewDecoder(bytes.NewReader(raw))
	dec.UseNumber()
	if err := dec.Decode(&result); err != nil {
		return Folder{}, fmt.Errorf("unmarshal folder error: %s\n%s", err, raw)
	}
	return result, nil
}

func (r *Client) SetFolderPermissionForUser(uid string, userID, perm uint) (StatusMessage, error) {
	var (
		raw    []byte
		result StatusMessage
		code   int
		err    error
	)
	params := make(map[string][]map[string]uint)
	params["items"] = []map[string]uint{
		map[string]uint{"userID": userID, "permission": perm},
	}
	paramsBytes, err := json.Marshal(params)
	if err != nil {
		return StatusMessage{}, err
	}
	if raw, code, err = r.post("/api/folders/"+uid+"/permissions", nil, paramsBytes); err != nil {
		return StatusMessage{}, err
	}
	if code != 200 {
		return StatusMessage{}, fmt.Errorf("HTTP error %d: returns %s", code, raw)
	}
	dec := json.NewDecoder(bytes.NewReader(raw))
	dec.UseNumber()
	if err := dec.Decode(&result); err != nil {
		return StatusMessage{}, fmt.Errorf("unmarshal folder error: %s\n%s", err, raw)
	}
	return result, nil
}

func (r *Client) CreateFolder(uid, title string) (Folder, error) {
	var (
		raw    []byte
		code   int
		result Folder
		err    error
	)
	params := make(map[string]string)
	params["uid"] = uid
	params["title"] = title
	paramsBytes, err := json.Marshal(params)
	if err != nil {
		return Folder{}, err
	}
	if raw, code, err = r.post("/api/folders", nil, paramsBytes); err != nil {
		return Folder{}, err
	}
	if code != 200 {
		return Folder{}, fmt.Errorf("HTTP error %d: returns %s", code, raw)
	}
	if err := json.Unmarshal(raw, &result); err != nil {
		return Folder{}, err
	}
	return result, nil
}

func (r *Client) DeleteFolder(uid string) (StatusMessage, error) {
	var resp StatusMessage
	if raw, code, err := r.delete(fmt.Sprintf("api/folders/%s", uid)); err != nil {
		return StatusMessage{}, err
	} else if code != 200 {
		return StatusMessage{}, fmt.Errorf("HTTP error %d: returns %s", code, raw)
	} else if err := json.Unmarshal(raw, &resp); err != nil {
		return StatusMessage{}, err
	}
	return resp, nil
}
