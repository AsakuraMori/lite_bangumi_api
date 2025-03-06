/**
 * @file 	api_characters.go
 * @brief 	关于characters的API接口
 * @author 	AsakuraMori
 * @version 0.3.0
 * @date 	2025-03-25
 */

package lite_bangumi_api

import (
	"fmt"
	"net/http"
	"net/url"
)

/*
SearchCharactersByName

  - @brief 通过搜索角色名获取任务信息。

    API：/v0/search/characters

  - @param

    【limit】：当前页最大数量

    【offset】：起始位置

    【requestBody】：请求体，结构如下：

    {
    "keyword": "string",
    "filter": {
    "nsfw": true
    }
    }

    【client】：http.Client对象。

  - @return 返回一个[]byte和一个err。

  - @retval []byte是返回体，err表示错误。如果err为nil，则没有错误。
*/
func SearchCharactersByName(limit, offset string, requestBody string, client *http.Client) ([]byte, error) {

	baseURL := "https://api.bgm.tv/v0/search/characters"

	params := url.Values{}
	params.Add("limit", fmt.Sprintf("%s", limit))
	params.Add("offset", fmt.Sprintf("%s", offset))

	apiURL := fmt.Sprintf("%s?%s", baseURL, params.Encode())
	jsonData, err := getJsonDataFromURL("POST", apiURL, requestBody, client)
	if err != nil {
		return nil, err
	}
	return jsonData, nil
}

/*
SearchCharactersById

  - @brief 通过搜索角色ID名获取ID信息。

    API：/v0/characters/{character_id}

  - @param

    【chrID】：角色ID。

    【client】：http.Client对象。

  - @return 返回一个[]byte和一个err。

  - @retval []byte是返回体，err表示错误。如果err为nil，则没有错误。
*/
func SearchCharactersById(chrID string, client *http.Client) ([]byte, error) {
	baseURL := "https://api.bgm.tv/v0/characters/"
	apiURL := fmt.Sprintf("%s%s", baseURL, chrID)

	jsonData, err := getJsonDataFromURL("GET", apiURL, "", client)
	if err != nil {
		return nil, err
	}
	return jsonData, nil
}

/*
SetCollectCharactersById

  - @brief 为当前用户收藏角色。

    API：/v0/characters/{character_id}/collect

  - @param

    【chrID】：角色ID。

    【client】：http.Client对象。

  - @return 返回一个bool和一个err。

  - @retval  如果bool为true，err为nil。如果bool为false，err表示错误信息
*/
func SetCollectCharactersById(chrID string, client *http.Client) (bool, error) {

	apiURL := fmt.Sprintf("https://api.bgm.tv/v0/characters/%s/collect", chrID)
	err := getBoolDataFromURL("POST", apiURL, "", client)
	if err != nil {
		return false, err
	}
	return true, nil
}

/*
DeleteCollectCharactersById

  - @brief 为当前用户取消收藏角色。

    API：/v0/characters/{character_id}/collect

  - @param

    【chrID】：角色ID。

    【client】：http.Client对象。

  - @return 返回一个bool和一个err。

  - @retval  如果bool为true，err为nil。如果bool为false，err表示错误信息
*/
func DeleteCollectCharactersById(chrID string, client *http.Client) (bool, error) {

	apiURL := fmt.Sprintf("https://api.bgm.tv/v0/characters/%s/collect", chrID)

	err := getBoolDataFromURL("DELETE", apiURL, "", client)
	if err != nil {
		return false, err
	}
	return true, nil
}
