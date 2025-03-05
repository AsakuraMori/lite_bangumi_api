/**
 * @file 	api_persons.go
 * @brief 	关于persons的API接口
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
*

  - @brief 通过搜索人物名获取人物信息。
    API：/v0/search/persons

  - @param 【limit】：当前页最大数量
    【offset】：起始位置
    【requestBody】：请求体，结构如下：
    {
    "keyword": "string",
    "filter": {
    "career": [
    "artist",
    "director"
    ]
    }
    }
    【client】：http.Client对象。

  - @return 返回一个[]byte和一个err。

  - @retval []byte是返回体，err表示错误。如果err为nil，则没有错误。
*/
func SearchPersonsByName(limit, offset string, requestBody string, client *http.Client) ([]byte, error) {

	baseURL := "https://api.bgm.tv/v0/search/persons"

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
*
  - @brief 通过搜索人物ID名获取ID信息。
    API：/v0/persons/{person_id}
  - @param 【perID】：人物ID。
    【client】：http.Client对象。
  - @return 返回一个[]byte和一个err。
  - @retval []byte是返回体，err表示错误。如果err为nil，则没有错误。
*/
func SearchPersonsById(perID string, client *http.Client) ([]byte, error) {
	baseURL := "https://api.bgm.tv/v0/persons/"
	apiURL := fmt.Sprintf("%s%s", baseURL, perID)

	jsonData, err := getJsonDataFromURL("GET", apiURL, "", client)
	if err != nil {
		return nil, err
	}
	return jsonData, nil
}

/*
*
  - @brief 为当前用户收藏人物。
    API：/v0/persons/{person_id}/collect
  - @param 【perID】：人物ID。
    【client】：http.Client对象。
  - @return 返回一个bool和一个err。
  - @retval  如果bool为true，err为nil。如果bool为false，err表示错误信息
*/
func SetCollectPersonsById(perID string, client *http.Client) (bool, error) {

	apiURL := fmt.Sprintf("https://api.bgm.tv/v0/persons/%s/collect", perID)

	err := getBoolDataFromURL("POST", apiURL, "", client)
	if err != nil {
		return false, err
	}
	return true, nil
}

/*
*
  - @brief 为当前用户取消收藏人物。
    API：/v0/persons/{person_id}/collect
  - @param 【perID】：人物ID。
    【client】：http.Client对象。
  - @return 返回一个bool和一个err。
  - @retval  如果bool为true，err为nil。如果bool为false，err表示错误信息
*/
func DeleteCollectPersonsById(perID string, client *http.Client) (bool, error) {

	apiURL := fmt.Sprintf("https://api.bgm.tv/v0/persons/%s/collect/", perID)

	err := getBoolDataFromURL("DELETE", apiURL, "", client)
	if err != nil {
		return false, err
	}
	return true, nil
}
