// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package client

import (
	"github.com/chanxuehong/wechat/menu"
)

// 创建自定义菜单.
//  NOTE: 创建自定义菜单后，由于微信客户端缓存，需要24小时微信客户端才会展现出来。
//  建议测试时可以尝试取消关注公众账号后再次关注，则可以看到创建后的效果。
func (c *Client) MenuCreate(_menu menu.Menu) (err error) {
	token, err := c.Token()
	if err != nil {
		return
	}
	_url := menuCreateURL(token)

	var result Error
	if err = c.postJSON(_url, _menu, &result); err != nil {
		return
	}

	if result.ErrCode != 0 {
		return result
	}
	return
}

// 删除自定义菜单
func (c *Client) MenuDelete() (err error) {
	token, err := c.Token()
	if err != nil {
		return
	}
	_url := menuDeleteURL(token)

	var result Error
	if err = c.getJSON(_url, &result); err != nil {
		return
	}

	if result.ErrCode != 0 {
		return result
	}
	return
}

// 获取自定义菜单
func (c *Client) MenuGet() (_menu menu.Menu, err error) {
	token, err := c.Token()
	if err != nil {
		return
	}
	_url := menuGetURL(token)

	var result struct {
		Menu menu.Menu `json:"menu"`
		Error
	}
	if err = c.getJSON(_url, &result); err != nil {
		return
	}

	if result.ErrCode != 0 {
		err = result.Error
		return
	}

	_menu = result.Menu
	return
}
