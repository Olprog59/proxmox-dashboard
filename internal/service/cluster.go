package service

// func (c *Client) HttpGetClient[T any](url string, t T) ([]T, error) {
// 	body, err := c.doRequest("GET", url, nil)
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	var result struct {
// 		Data []T `json:"data"`
// 	}
//
// 	if err := json.Unmarshal(body, &result); err != nil {
// 		return nil, err
// 	}
// 	return result.Data, nil
// }
