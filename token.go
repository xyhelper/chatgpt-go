package chatgpt

import "github.com/gogf/gf/v2/errors/gerror"

// getAccessToken will return accessToken, if expired than fetch a new one
func (c *Client) getAccessToken() (string, error) {
	if c.opts.AccessToken != "" {
		return c.opts.AccessToken, nil
	}
	return "", gerror.New("accessToken is empty")

	// // fetch new accessToken
	// res, err := c.authSession()
	// if err != nil {
	// 	return "", fmt.Errorf("fetch new accessToken failed: %v", err)
	// }

	// accessToken := res.Get("accessToken").String()
	// if accessToken == "" {
	// 	return "", fmt.Errorf("invalid session data: %s", accessToken)
	// }

	// return accessToken, nil
}
