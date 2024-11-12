package parsing

import "encoding/json"

func convertFromStringToMap(s string) (map[string]string, error) {
	var result map[string]string

	err := json.Unmarshal([]byte(s), &result)
	if err != nil {
		return result, err
	}

	return result, nil
}
