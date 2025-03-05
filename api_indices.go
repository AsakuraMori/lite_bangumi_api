/**
 * @file 	api_indices.go
 * @brief 	关于indices的API接口
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
  - @brief 新建目录。
    API：/v0/persons/{person_id}
  - @param【client】：http.Client对象。
  - @return 返回一个[]byte和一个err。
  - @retval []byte是返回体，err表示错误。如果err为nil，则没有错误。
*/
func SetIndices(client *http.Client) ([]byte, error) {

	apiURL := "https://api.bgm.tv/v0/indices"

	jsonData, err := getJsonDataFromURL("POST", apiURL, "", client)
	if err != nil {
		return nil, err
	}
	return jsonData, nil
}

/*
*
  - @brief 通过搜索目录ID名获取ID信息。
    API：/v0/indices/{index_id}
  - @param 【idxID】：目录ID。
    【client】：http.Client对象。
  - @return 返回一个[]byte和一个err。
  - @retval []byte是返回体，err表示错误。如果err为nil，则没有错误。
*/
func GetIndicesByID(idxID string, client *http.Client) ([]byte, error) {
	baseURL := "https://api.bgm.tv/v0/indices/"
	apiURL := fmt.Sprintf("%s%s", baseURL, idxID)

	jsonData, err := getJsonDataFromURL("GET", apiURL, "", client)
	if err != nil {
		return nil, err
	}
	return jsonData, nil
}

/*
*
  - @brief 通过目录ID修改目录信息。
    API：/v0/indices/{index_id}
  - @param 【idxID】：目录ID。
    【requestBody】：请求体，格式如下：
    {
    "title": "string",
    "description": "string"
    }
    【client】：http.Client对象。
  - @return 返回一个[]byte和一个err。
  - @retval []byte是返回体，err表示错误。如果err为nil，则没有错误。
*/
func EditIndicesInformationByIDAndRequestBody(idxID, requestBody string, client *http.Client) ([]byte, error) {
	baseURL := "https://api.bgm.tv/v0/indices/"
	apiURL := fmt.Sprintf("%s%s", baseURL, idxID)

	jsonData, err := getJsonDataFromURL("PUT", apiURL, requestBody, client)
	if err != nil {
		return nil, err
	}
	return jsonData, nil
}

/*
*

  - @brief 通过目录ID获取内部条目。
    API：/v0/indices/{index_id}/subjects

  - @param 【idxID】：目录ID。
    【typeName】：类型名（只能是以下字符串：书籍、动漫、音乐、游戏、三次元。如果typeName不满足以上字符串，则将全局搜索）
    【limit】：当前页面显示条目最大数量。
    【offset】：开始的条目位置。
    【client】：http.Client对象。

  - @return 返回一个bool和一个err。

  - @retval  如果bool为true，err为nil。如果bool为false，err表示错误信息
*/
func GetIndicesSubjectByID(idxID, typeName, limit, offset string, client *http.Client) (bool, error) {
	sType := 0
	switch typeName {
	case "书籍":
		sType = 1
	case "动漫":
		sType = 2
	case "音乐":
		sType = 3
	case "游戏":
		sType = 4
	case "三次元":
		sType = 6
	default:
		sType = 0
	}
	params := url.Values{}
	params.Add("type", fmt.Sprintf("%d", sType))
	params.Add("limit", fmt.Sprintf("%s", limit))
	params.Add("offset", fmt.Sprintf("%s", offset))

	apiURL := fmt.Sprintf("https://api.bgm.tv/v0/indices/%s/subjects?%s", idxID, params.Encode())

	err := getBoolDataFromURL("GET", apiURL, "", client)
	if err != nil {
		return false, err
	}
	return true, nil
}

/*
*
  - @brief 将条目添加到目录。
    API：/v0/indices/{index_id}/subjects
  - @param 【idxID】：目录ID。
    【requestBody】：请求体，格式如下：
    {
    "subject_id": 0,
    "sort": 0,
    "comment": "string"
    }
    【client】：http.Client对象。
  - @return 返回一个bool和一个err。
  - @retval 如果bool为true，err为nil。如果bool为false，err表示错误信息
*/
func AddSubjectsToIndicesByIDAndRequestBody(idxID, requestBody string, client *http.Client) (bool, error) {
	apiURL := fmt.Sprintf("https://api.bgm.tv/v0/indices/%s/subjects", idxID)

	err := getBoolDataFromURL("POST", apiURL, requestBody, client)
	if err != nil {
		return false, err
	}
	return true, nil
}

/*
*
  - @brief 编辑在目录中的条目，如果不存在，则创建条目。
    API：/v0/indices/{index_id}/subjects/{subject_id}
  - @param 【idxID】：目录ID。
    【subID】：条目ID
    【requestBody】：请求体，格式如下：
    {
    "sort": 0,
    "comment": "string"
    }
    【client】：http.Client对象。
  - @return 返回一个bool和一个err。
  - @retval 如果bool为true，err为nil。如果bool为false，err表示错误信息
*/
func EditSubjectsInformationInIndiesByIDAndRequestBody(idxID, subID, requestBody string, client *http.Client) (bool, error) {
	apiURL := fmt.Sprintf("https://api.bgm.tv/v0/indices/%s/subjects/%s", idxID, subID)

	err := getBoolDataFromURL("PUT", apiURL, requestBody, client)
	if err != nil {
		return false, err
	}
	return true, nil
}

/*
*

  - @brief 从目录中删除条目。
    API：/v0/indices/{index_id}/subjects/{subject_id}

  - @param 【idxID】：目录ID。
    【subID】：条目ID
    【client】：http.Client对象。

  - @return 返回一个bool和一个err。

  - @retval  如果bool为true，err为nil。如果bool为false，err表示错误信息
*/
func DeleteSubjectsFromIndicesByID(idxID, subID string, client *http.Client) (bool, error) {
	apiURL := fmt.Sprintf("https://api.bgm.tv/v0/indices/%s/subjects/%s", idxID, subID)

	err := getBoolDataFromURL("DELETE", apiURL, "", client)
	if err != nil {
		return false, err
	}
	return true, nil
}

/*
*

  - @brief 收藏目录到当前用户。
    API：/v0/indices/{index_id}/collect

  - @param 【idxID】：目录ID。
    【client】：http.Client对象。

  - @return 返回一个bool和一个err。

  - @retval  如果bool为true，err为nil。如果bool为false，err表示错误信息
*/
func CollectIndicesForCurrentUserByID(idxID string, client *http.Client) (bool, error) {
	apiURL := fmt.Sprintf("https://api.bgm.tv/v0/indices/%s/collect", idxID)

	err := getBoolDataFromURL("POST", apiURL, "", client)
	if err != nil {
		return false, err
	}
	return true, nil
}

/*
*

  - @brief 删除当前用户里指定的目录。
    API：/v0/indices/{index_id}/collect

  - @param 【idxID】：目录ID。
    【client】：http.Client对象。

  - @return 返回一个bool和一个err。

  - @retval  如果bool为true，err为nil。如果bool为false，err表示错误信息
*/
func DeleteCollectIndicesForCurrentUserByID(idxID string, client *http.Client) (bool, error) {
	apiURL := fmt.Sprintf("https://api.bgm.tv/v0/indices/%s/collect", idxID)

	err := getBoolDataFromURL("DELETE", apiURL, "", client)
	if err != nil {
		return false, err
	}
	return true, nil
}
