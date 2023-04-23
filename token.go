package chatgpt

import "github.com/gogf/gf/v2/errors/gerror"

// getAccessToken will return accessToken, if expired than fetch a new one
func (c *Client) getAccessToken() (string, error) {
	if c.opts.AccessToken != "" {
		return c.opts.AccessToken, nil
	}
	return "", gerror.New("accessToken is empty")
}
