/**
 * @file 	api_revisions.go
 * @brief 	关于revisions的API接口
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
SearchPersonsRevisionsById

  - @brief 获取人物ID的编辑历史。

    API：/v0/revisions/persons

  - @param

    【perID】：人物ID

    【limit】：当前页最大数

    【offset】：起始位置

    【client】：http.Client对象。

  - @return 返回一个[]byte和一个err。

  - @retval []byte是返回体，err表示错误。如果err为nil，则没有错误。
*/
func SearchPersonsRevisionsById(perID, limit, offset string, client *http.Client) ([]byte, error) {
	baseURL := "https://api.bgm.tv/v0/revisions/persons"
	params := url.Values{}
	params.Add("person_id", fmt.Sprintf("%s", perID))
	params.Add("limit", fmt.Sprintf("%s", limit))
	params.Add("offset", fmt.Sprintf("%s", offset))
	apiURL := fmt.Sprintf("%s?%s", baseURL, params.Encode())
	jsonData, err := getJsonDataFromURL("GET", apiURL, "", client)
	if err != nil {
		return nil, err
	}
	return jsonData, nil
}

/*
SearchPersonsRevisionsByRevisionsId

  - @brief 获取人物编辑历史ID内详细信息。

    API：/v0/revisions/persons/{revision_id}

  - @param

    【revID】：历史ID

    【client】：http.Client对象。

  - @return 返回一个[]byte和一个err。

  - @retval []byte是返回体，err表示错误。如果err为nil，则没有错误。
*/
func SearchPersonsRevisionsByRevisionsId(revID string, client *http.Client) ([]byte, error) {
	baseURL := "https://api.bgm.tv/v0/revisions/persons/"
	apiURL := fmt.Sprintf("%s%s", baseURL, revID)

	jsonData, err := getJsonDataFromURL("GET", apiURL, "", client)
	if err != nil {
		return nil, err
	}
	return jsonData, nil
}

/*
SearchCharactersRevisionsById

  - @brief 获取角色ID的编辑历史。

    API：/v0/revisions/characters

  - @param

    【chrID】：角色ID

    【limit】：当前页最大数

    【offset】：起始位置

    【client】：http.Client对象。

  - @return 返回一个[]byte和一个err。

  - @retval []byte是返回体，err表示错误。如果err为nil，则没有错误。
*/
func SearchCharactersRevisionsById(chrID, limit, offset string, client *http.Client) ([]byte, error) {
	baseURL := "https://api.bgm.tv/v0/revisions/characters"
	params := url.Values{}
	params.Add("character_id", fmt.Sprintf("%s", chrID))
	params.Add("limit", fmt.Sprintf("%s", limit))
	params.Add("offset", fmt.Sprintf("%s", offset))
	apiURL := fmt.Sprintf("%s?%s", baseURL, params.Encode())

	jsonData, err := getJsonDataFromURL("GET", apiURL, "", client)
	if err != nil {
		return nil, err
	}
	return jsonData, nil
}

/*
SearchCharactersRevisionsByRevisionsId

  - @brief 获取角色编辑历史ID内详细信息。

    API：/v0/revisions/characters/{revision_id}

  - @param

    【revID】：历史ID

    【client】：http.Client对象。

  - @return 返回一个[]byte和一个err。

  - @retval []byte是返回体，err表示错误。如果err为nil，则没有错误。
*/
func SearchCharactersRevisionsByRevisionsId(revID string, client *http.Client) ([]byte, error) {
	baseURL := "https://api.bgm.tv/v0/revisions/characters/"
	apiURL := fmt.Sprintf("%s%s", baseURL, revID)

	jsonData, err := getJsonDataFromURL("GET", apiURL, "", client)
	if err != nil {
		return nil, err
	}
	return jsonData, nil
}

/*
SearchSubjectsRevisionsById

  - @brief 获取条目ID的编辑历史。

    API：/v0/revisions/subjects

  - @param

    【subID】：条目ID

    【limit】：当前页最大数

    【offset】：起始位置

    【client】：http.Client对象。

  - @return 返回一个[]byte和一个err。

  - @retval []byte是返回体，err表示错误。如果err为nil，则没有错误。
*/
func SearchSubjectsRevisionsById(subID, limit, offset string, client *http.Client) ([]byte, error) {
	baseURL := "https://api.bgm.tv/v0/revisions/subjects"
	params := url.Values{}
	params.Add("subject_id", fmt.Sprintf("%s", subID))
	params.Add("limit", fmt.Sprintf("%s", limit))
	params.Add("offset", fmt.Sprintf("%s", offset))
	apiURL := fmt.Sprintf("%s?%s", baseURL, params.Encode())

	jsonData, err := getJsonDataFromURL("GET", apiURL, "", client)
	if err != nil {
		return nil, err
	}
	return jsonData, nil
}

/*
SearchSubjectsRevisionsByRevisionsId

  - @brief 获取条目编辑历史ID内详细信息。

    API：/v0/revisions/subjects/{revision_id}

  - @param

    【revID】：历史ID

    【client】：http.Client对象。

  - @return 返回一个[]byte和一个err。

  - @retval []byte是返回体，err表示错误。如果err为nil，则没有错误。
*/
func SearchSubjectsRevisionsByRevisionsId(revID string, client *http.Client) ([]byte, error) {
	baseURL := "https://api.bgm.tv/v0/revisions/subjects/"
	apiURL := fmt.Sprintf("%s%s", baseURL, revID)

	jsonData, err := getJsonDataFromURL("GET", apiURL, "", client)
	if err != nil {
		return nil, err
	}
	return jsonData, nil
}

/*
SearchEpisodesRevisionsById

  - @brief 获取章节ID的编辑历史。

    API：/v0/revisions/episodes

  - @param

    【epiID】：章节ID

    【limit】：当前页最大数

    【offset】：起始位置

    【client】：http.Client对象。

  - @return 返回一个[]byte和一个err。

  - @retval []byte是返回体，err表示错误。如果err为nil，则没有错误。
*/
func SearchEpisodesRevisionsById(epiID, limit, offset string, client *http.Client) ([]byte, error) {
	baseURL := "https://api.bgm.tv/v0/revisions/episodes"
	params := url.Values{}
	params.Add("episode_id", fmt.Sprintf("%s", epiID))
	params.Add("limit", fmt.Sprintf("%s", limit))
	params.Add("offset", fmt.Sprintf("%s", offset))
	apiURL := fmt.Sprintf("%s?%s", baseURL, params.Encode())

	jsonData, err := getJsonDataFromURL("GET", apiURL, "", client)
	if err != nil {
		return nil, err
	}
	return jsonData, nil
}

/*
SearchEpisodesRevisionsByRevisionsId

  - @brief 获取章节编辑历史ID内详细信息。

    API：/v0/revisions/episodes/{revision_id}

  - @param

    【revID】：历史ID

    【client】：http.Client对象。

  - @return 返回一个[]byte和一个err。

  - @retval []byte是返回体，err表示错误。如果err为nil，则没有错误。
*/
func SearchEpisodesRevisionsByRevisionsId(revID string, client *http.Client) ([]byte, error) {
	baseURL := "https://api.bgm.tv/v0/revisions/episodes/"
	apiURL := fmt.Sprintf("%s%s", baseURL, revID)

	jsonData, err := getJsonDataFromURL("GET", apiURL, "", client)
	if err != nil {
		return nil, err
	}
	return jsonData, nil
}
