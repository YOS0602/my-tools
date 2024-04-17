package notion

import (
	"encoding/json"
	"fmt"
)

func UnmarshalResponse(data []byte) (*APIResponse, error) {
	var response APIResponse
	if err := json.Unmarshal(data, &response); err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &response, nil
}
