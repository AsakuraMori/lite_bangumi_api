/**
 * @file 	api_collections.go
 * @brief 	关于collections的API接口
 * @author 	AsakuraMori
 * @version 0.3.0
 * @date 	2025-03-25
 */

package lite_bangumi_api

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
)

/*
*
  - @brief 获取用户收藏。
    API：/v0/users/{username}/collections
  - @param 【userName】：用户名。
    【subjectTypeName】：条目类型（只能是以下字符串：书籍、动漫、音乐、游戏、三次元。如果不满足以上字符串，则会返回错误）
    【typeName】：收藏类型（只能是以下字符串：想看、看过、在看、搁置、抛弃。如果不满足以上字符串，则将全局搜索）
    【limit】：当前页面显示条目最大数量。
    【offset】：开始的条目位置。
    【client】：http.Client对象。
  - @return 返回一个[]byte和一个err。
  - @retval []byte是返回体，err表示错误。如果err为nil，则没有错误。
*/
func SearchCollectionsByUserName(userName, subjectTypeName, typeName, limit, offset string, client *http.Client) ([]byte, error) {

	params := url.Values{}
	sType := 0
	switch subjectTypeName {
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
		errMsg := errors.New("不匹配的subjectTypeName")
		return nil, errMsg
	}
	tType := 0
	switch typeName {
	case "想看":
		tType = 1
	case "看过":
		tType = 2
	case "在看":
		tType = 3
	case "搁置":
		tType = 4
	case "抛弃":
		tType = 5
	default:
		tType = 0
	}
	params.Add("subject_type", fmt.Sprintf("%d", sType))
	params.Add("type", fmt.Sprintf("%d", tType))
	params.Add("limit", fmt.Sprintf("%s", limit))
	params.Add("offset", fmt.Sprintf("%s", offset))

	apiURL := fmt.Sprintf("https://api.bgm.tv/v0/users/%s/collections?%s", userName, params.Encode())
	jsonData, err := getJsonDataFromURL("GET", apiURL, "", client)
	if err != nil {
		return nil, err
	}
	return jsonData, nil
}

/*
*
  - @brief 获取对应用户的收藏，查看私有收藏需要 access token。
    API：/v0/users/{username}/collections/{subject_id}
  - @param 【userName】：用户名。
    【subID】：条目类ID
    【client】：http.Client对象。
  - @return 返回一个[]byte和一个err。
  - @retval []byte是返回体，err表示错误。如果err为nil，则没有错误。
*/
func SearchCollectionsByID(userName, subID string, client *http.Client) ([]byte, error) {
	apiURL := fmt.Sprintf("https://api.bgm.tv/v0/users/%s/collections/%s", userName, subID)
	jsonData, err := getJsonDataFromURL("GET", apiURL, "", client)
	if err != nil {
		return nil, err
	}
	return jsonData, nil
}

/*
*
  - @brief 修改条目收藏状态, 如果不存在则创建，如果存在则修改。
    由于直接修改剧集条目的完成度可能会引起意料之外效果，只能用于修改书籍类条目的完成度。
    方法的所有请求体字段均可选。
    API：/v0/users/-/collections/{subject_id}
  - @param 【subID】：条目ID。
    【requestBody】：请求体，格式如下：
    {
    "type": 3,
    "rate": 10,
    "ep_status": 0,
    "vol_status": 0,
    "comment": "string",
    "private": true,
    "tags": [
    "string"
    ]
    }
    【client】：http.Client对象。
  - @return 返回一个bool和一个err。
  - @retval 如果bool为true，err为nil。如果bool为false，err表示错误信息
*/
func AddOrEditCollectionsSubjectsInUsersByID(subID, requestBody string, client *http.Client) (bool, error) {
	apiURL := fmt.Sprintf("https://api.bgm.tv/v0/users/-/collections/%s", subID)

	err := getBoolDataFromURL("POST", apiURL, requestBody, client)
	if err != nil {
		return false, err
	}
	return true, nil
}

/*
*
  - @brief 修改条目收藏状态。
    由于直接修改剧集条目的完成度可能会引起意料之外效果，只能用于修改书籍类条目的完成度。
    PATCH 方法的所有请求体字段均可选。
    API：/v0/users/-/collections/{subject_id}
  - @param 【subID】：条目ID。
    【requestBody】：请求体，格式如下：
    {
    "type": 3,
    "rate": 10,
    "ep_status": 0,
    "vol_status": 0,
    "comment": "string",
    "private": true,
    "tags": [
    "string"
    ]
    }
    【client】：http.Client对象。
  - @return 返回一个bool和一个err。
  - @retval 如果bool为true，err为nil。如果bool为false，err表示错误信息
*/
func EditCollectionsSubjectsInUsersByID(subID, requestBody string, client *http.Client) (bool, error) {
	apiURL := fmt.Sprintf("https://api.bgm.tv/v0/users/-/collections/%s", subID)
	err := getBoolDataFromURL("PATCH", apiURL, requestBody, client)
	if err != nil {
		return false, err
	}
	return true, nil
}

/*
*
  - @brief 搜索章节收藏信息。
    API：/v0/users/-/collections/{subject_id}/episodes
  - @param 【subID】：条目ID。
    【limit】：当前页面显示条目最大数量。
    【offset】：开始的条目位置。
    【episodesType】：章节类型（只能是以下字符串：本篇、特别篇、OP、ED、预告/宣传/广告、MAD、其他。如果不满足以上字符串，则将全局搜索）
    【client】：http.Client对象。
  - @return 返回一个[]byte和一个err。
  - @retval []byte是返回体，err表示错误。如果err为nil，则没有错误。
*/
func SearchUsersCollectionsEpisodesBySubjectsID(subID, offset, limit, episodesType string, client *http.Client) ([]byte, error) {
	params := url.Values{}
	sType := 0
	switch episodesType {
	case "本篇":
		sType = 0
	case "特别篇":
		sType = 1
	case "OP":
		sType = 2
	case "ED":
		sType = 3
	case "预告/宣传/广告":
		sType = 4
	case "MAD":
		sType = 5
	case "其他":
		sType = 6
	default:
		sType = 0
	}
	params.Add("episode_type", fmt.Sprintf("%d", sType))
	params.Add("limit", fmt.Sprintf("%s", limit))
	params.Add("offset", fmt.Sprintf("%s", offset))

	apiURL := fmt.Sprintf("https://api.bgm.tv/v0/users/-/collections/%s/episodes?", subID, params.Encode())
	jsonData, err := getJsonDataFromURL("GET", apiURL, "", client)
	if err != nil {
		return nil, err
	}
	return jsonData, nil
}

/*
*
  - @brief 章节收藏信息。同时会重新计算条目的完成度。
    API：/v0/users/-/collections/{subject_id}/episodes
  - @param 【subID】：条目ID。
    【requestBody】：请求体，格式如下：
    {
    "episode_id": [
    1,
    2,
    8
    ],
    "type": 2
    ]
    }
    【client】：http.Client对象。
  - @return 返回一个bool和一个err。
  - @retval 如果bool为true，err为nil。如果bool为false，err表示错误信息
*/
func GetCollectionsSubjectsEpisodesInfo(subID, requestBody string, client *http.Client) (bool, error) {
	apiURL := fmt.Sprintf("https://api.bgm.tv/v0/users/-/collections/%s", subID)
	err := getBoolDataFromURL("PATCH", apiURL, requestBody, client)
	if err != nil {
		return false, err
	}
	return true, nil
}

/*
*
  - @brief 获取章节收藏信息。
    API：/v0/users/-/collections/-/episodes/{episode_id}
  - @param 【epiID】：章节ID。
    【client】：http.Client对象。
  - @return 返回一个[]byte和一个err。
  - @retval []byte是返回体，err表示错误。如果err为nil，则没有错误。
*/
func SearchCollectionsEpisodesInfo(epiID string, client *http.Client) ([]byte, error) {
	apiURL := fmt.Sprintf("https://api.bgm.tv/v0/users/-/collections/-/episodes/%s", epiID)
	jsonData, err := getJsonDataFromURL("GET", apiURL, "", client)
	if err != nil {
		return nil, err
	}
	return jsonData, nil
}

/*
*
  - @brief 更新章节收藏信息。
    API：/v0/users/-/collections/-/episodes/{episode_id}
  - @param 【epiID】：章节ID。
    【requestBody】：请求体，格式如下：
    {
    "type": 2
    }
    【client】：http.Client对象。
  - @return 返回一个bool和一个err。
  - @retval 如果bool为true，err为nil。如果bool为false，err表示错误信息
*/
func UpdateCollectionEpisodesInfo(epiID, requestBody string, client *http.Client) (bool, error) {
	apiURL := fmt.Sprintf("https://api.bgm.tv/v0/users/-/collections/-/episodes/%s", epiID)
	err := getBoolDataFromURL("PUT", apiURL, requestBody, client)
	if err != nil {
		return false, err
	}
	return true, nil
}

/*
*
  - @brief 获取用户角色收藏列表。
    API：/v0/users/{username}/collections/-/characters
  - @param 【userName】：用户名。
    【client】：http.Client对象。
  - @return 返回一个[]byte和一个err。
  - @retval []byte是返回体，err表示错误。如果err为nil，则没有错误。
*/
func SearchCharactersCollectionsByUserName(userName string, client *http.Client) ([]byte, error) {
	apiURL := fmt.Sprintf("https://api.bgm.tv/v0/users/%s/collections/-/characters", userName)
	jsonData, err := getJsonDataFromURL("GET", apiURL, "", client)
	if err != nil {
		return nil, err
	}
	return jsonData, nil
}

/*
*
  - @brief 获取用户单个角色收藏信息。
    API：/v0/users/{username}/collections/-/characters/{character_id}
  - @param 【userName】：用户名。
    【client】：http.Client对象。
  - @return 返回一个[]byte和一个err。
  - @retval []byte是返回体，err表示错误。如果err为nil，则没有错误。
*/
func SearchCharactersCollectionsByUserNameAndID(userName, chrID string, client *http.Client) ([]byte, error) {
	apiURL := fmt.Sprintf("https://api.bgm.tv/v0/users/%s/collections/-/characters/%s", userName, chrID)
	jsonData, err := getJsonDataFromURL("GET", apiURL, "", client)
	if err != nil {
		return nil, err
	}
	return jsonData, nil
}

/*
*
  - @brief 获取用户人物收藏列表。
    API：/v0/users/{username}/collections/-/persons
  - @param 【userName】：用户名。
    【client】：http.Client对象。
  - @return 返回一个[]byte和一个err。
  - @retval []byte是返回体，err表示错误。如果err为nil，则没有错误。
*/
func SearchPersonsCollectionsByUserName(userName string, client *http.Client) ([]byte, error) {
	apiURL := fmt.Sprintf("https://api.bgm.tv/v0/users/%s/collections/-/persons", userName)
	jsonData, err := getJsonDataFromURL("GET", apiURL, "", client)
	if err != nil {
		return nil, err
	}
	return jsonData, nil
}

/*
*
  - @brief 获取用户单个人物收藏信息。
    API：/v0/users/{username}/collections/-/persons/{person_id}
  - @param 【userName】：用户名。
    【client】：http.Client对象。
  - @return 返回一个[]byte和一个err。
  - @retval []byte是返回体，err表示错误。如果err为nil，则没有错误。
*/
func SearchPersonsCollectionsByUserNameAndID(userName, perID string, client *http.Client) ([]byte, error) {
	apiURL := fmt.Sprintf("https://api.bgm.tv/v0/users/%s/collections/-/persons/%s", userName, perID)
	jsonData, err := getJsonDataFromURL("GET", apiURL, "", client)
	if err != nil {
		return nil, err
	}
	return jsonData, nil
}
