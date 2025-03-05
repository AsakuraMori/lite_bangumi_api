/**
 * @file 	api_users.go
 * @brief 	关于users的API接口
 * @author 	AsakuraMori
 * @version 0.3.0
 * @date 	2025-03-25
 */

package lite_bangumi_api

import (
	"fmt"
	"net/http"
)

/*
*
  - @brief 通过搜索用户名获取用户信息。
    API：/v0/users/{username}
  - @param 【userName】：用户名。
    【client】：http.Client对象。
  - @return 返回一个[]byte和一个err。
  - @retval []byte是返回体，err表示错误。如果err为nil，则没有错误。
*/
func SearchUserNameByName(userName string, client *http.Client) ([]byte, error) {
	baseURL := "https://api.bgm.tv/v0/users/"
	apiURL := fmt.Sprintf("%s%s", baseURL, userName)

	jsonData, err := getJsonDataFromURL("GET", apiURL, "", client)
	if err != nil {
		return nil, err
	}
	return jsonData, nil
}

/*
*
  - @brief 获取当前用户信息。
    API：/v0/me
  - @param 无
  - @return 返回一个[]byte和一个err。
  - @retval []byte是返回体，err表示错误。如果err为nil，则没有错误。
*/
func GetMe(client *http.Client) ([]byte, error) {
	apiURL := "https://api.bgm.tv/v0/me"

	jsonData, err := getJsonDataFromURL("GET", apiURL, "", client)
	if err != nil {
		return nil, err
	}
	return jsonData, nil
}
