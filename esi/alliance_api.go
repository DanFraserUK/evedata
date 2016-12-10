/* 
 * EVE Swagger Interface
 *
 * An OpenAPI for EVE Online
 *
 * OpenAPI spec version: 0.3.2.dev3
 * 
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package esi

import (
	"net/url"
	"strings"
	"time"
	"errors"
	"golang.org/x/net/context"
	"encoding/json"
	"fmt"
)

var _ context.Context

type AllianceApiService service


/**
 * List all alliances
 * List all active player alliances  ---  Alternate route: &#x60;/v1/alliances/&#x60;  Alternate route: &#x60;/legacy/alliances/&#x60;  Alternate route: &#x60;/dev/alliances/&#x60;   ---  This route is cached for up to 3600 seconds
 *
 * @param datasource(string) The server name you would like data from 
 * @return []int32
 */
func (a AllianceApiService) GetAlliances(datasource interface{}) ([]int32,  time.Time, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	)

	// create path and map variables
	localVarPath := "https://esi.tech.ccp.is/latest/alliances/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := a.client.typeCheckParameter(datasource, "string", "datasource"); err != nil {
		return nil, time.Now(), err
	}
	if datasource != nil {
		localVarQueryParams.Add("datasource", a.client.parameterToString(datasource, ""))
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.client.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	 var successPayload = new([]int32)

	 r, err := a.client.prepareRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes, "application/json")
	 if err != nil {
		  return *successPayload, time.Now(), err
	 }

	 localVarHttpResponse, err := a.client.callAPI(r)
	 if err != nil || localVarHttpResponse == nil {
		  return *successPayload, time.Now(), err
	 }
	 defer localVarHttpResponse.Body.Close()
	 if localVarHttpResponse.StatusCode >= 300 {
		return *successPayload, time.Now(), errors.New(localVarHttpResponse.Status)
	 }
	
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
	 	return *successPayload, time.Now(), err
	}

	expires := cacheExpires(localVarHttpResponse)
	return *successPayload, expires, err
}

/**
 * Get alliance information
 * Public information about an alliance  ---  Alternate route: &#x60;/v2/alliances/{alliance_id}/&#x60;   ---  This route is cached for up to 3600 seconds
 *
 * @param allianceId An Eve alliance ID 
 * @param datasource(string) The server name you would like data from 
 * @return *GetAlliancesAllianceIdOk
 */
func (a AllianceApiService) GetAlliancesAllianceId(allianceId int32, datasource interface{}) (*GetAlliancesAllianceIdOk,  time.Time, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	)

	// create path and map variables
	localVarPath := "https://esi.tech.ccp.is/latest/alliances/{alliance_id}/"
	localVarPath = strings.Replace(localVarPath, "{"+"alliance_id"+"}", fmt.Sprintf("%v", allianceId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := a.client.typeCheckParameter(datasource, "string", "datasource"); err != nil {
		return nil, time.Now(), err
	}
	if datasource != nil {
		localVarQueryParams.Add("datasource", a.client.parameterToString(datasource, ""))
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.client.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	 var successPayload = new(GetAlliancesAllianceIdOk)

	 r, err := a.client.prepareRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes, "application/json")
	 if err != nil {
		  return successPayload, time.Now(), err
	 }

	 localVarHttpResponse, err := a.client.callAPI(r)
	 if err != nil || localVarHttpResponse == nil {
		  return successPayload, time.Now(), err
	 }
	 defer localVarHttpResponse.Body.Close()
	 if localVarHttpResponse.StatusCode >= 300 {
		return successPayload, time.Now(), errors.New(localVarHttpResponse.Status)
	 }
	
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
	 	return successPayload, time.Now(), err
	}

	expires := cacheExpires(localVarHttpResponse)
	return successPayload, expires, err
}

/**
 * List alliance&#39;s corporations
 * List all current member corporations of an alliance  ---  Alternate route: &#x60;/v1/alliances/{alliance_id}/corporations/&#x60;  Alternate route: &#x60;/legacy/alliances/{alliance_id}/corporations/&#x60;  Alternate route: &#x60;/dev/alliances/{alliance_id}/corporations/&#x60;   ---  This route is cached for up to 3600 seconds
 *
 * @param allianceId An EVE alliance ID 
 * @param datasource(string) The server name you would like data from 
 * @return []int32
 */
func (a AllianceApiService) GetAlliancesAllianceIdCorporations(allianceId int32, datasource interface{}) ([]int32,  time.Time, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	)

	// create path and map variables
	localVarPath := "https://esi.tech.ccp.is/latest/alliances/{alliance_id}/corporations/"
	localVarPath = strings.Replace(localVarPath, "{"+"alliance_id"+"}", fmt.Sprintf("%v", allianceId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := a.client.typeCheckParameter(datasource, "string", "datasource"); err != nil {
		return nil, time.Now(), err
	}
	if datasource != nil {
		localVarQueryParams.Add("datasource", a.client.parameterToString(datasource, ""))
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.client.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	 var successPayload = new([]int32)

	 r, err := a.client.prepareRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes, "application/json")
	 if err != nil {
		  return *successPayload, time.Now(), err
	 }

	 localVarHttpResponse, err := a.client.callAPI(r)
	 if err != nil || localVarHttpResponse == nil {
		  return *successPayload, time.Now(), err
	 }
	 defer localVarHttpResponse.Body.Close()
	 if localVarHttpResponse.StatusCode >= 300 {
		return *successPayload, time.Now(), errors.New(localVarHttpResponse.Status)
	 }
	
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
	 	return *successPayload, time.Now(), err
	}

	expires := cacheExpires(localVarHttpResponse)
	return *successPayload, expires, err
}

/**
 * Get alliance icon
 * Get the icon urls for a alliance  ---  Alternate route: &#x60;/v1/alliances/{alliance_id}/icons/&#x60;  Alternate route: &#x60;/legacy/alliances/{alliance_id}/icons/&#x60;  Alternate route: &#x60;/dev/alliances/{alliance_id}/icons/&#x60;   ---  This route is cached for up to 3600 seconds
 *
 * @param allianceId An EVE alliance ID 
 * @param datasource(string) The server name you would like data from 
 * @return *GetAlliancesAllianceIdIconsOk
 */
func (a AllianceApiService) GetAlliancesAllianceIdIcons(allianceId int32, datasource interface{}) (*GetAlliancesAllianceIdIconsOk,  time.Time, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	)

	// create path and map variables
	localVarPath := "https://esi.tech.ccp.is/latest/alliances/{alliance_id}/icons/"
	localVarPath = strings.Replace(localVarPath, "{"+"alliance_id"+"}", fmt.Sprintf("%v", allianceId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := a.client.typeCheckParameter(datasource, "string", "datasource"); err != nil {
		return nil, time.Now(), err
	}
	if datasource != nil {
		localVarQueryParams.Add("datasource", a.client.parameterToString(datasource, ""))
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.client.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	 var successPayload = new(GetAlliancesAllianceIdIconsOk)

	 r, err := a.client.prepareRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes, "application/json")
	 if err != nil {
		  return successPayload, time.Now(), err
	 }

	 localVarHttpResponse, err := a.client.callAPI(r)
	 if err != nil || localVarHttpResponse == nil {
		  return successPayload, time.Now(), err
	 }
	 defer localVarHttpResponse.Body.Close()
	 if localVarHttpResponse.StatusCode >= 300 {
		return successPayload, time.Now(), errors.New(localVarHttpResponse.Status)
	 }
	
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
	 	return successPayload, time.Now(), err
	}

	expires := cacheExpires(localVarHttpResponse)
	return successPayload, expires, err
}

/**
 * Get alliance names
 * Resolve a set of alliance IDs to alliance names  ---  Alternate route: &#x60;/v1/alliances/names/&#x60;  Alternate route: &#x60;/legacy/alliances/names/&#x60;  Alternate route: &#x60;/dev/alliances/names/&#x60;   ---  This route is cached for up to 3600 seconds
 *
 * @param allianceIds A comma separated list of alliance IDs 
 * @param datasource(string) The server name you would like data from 
 * @return []GetAlliancesNames200Ok
 */
func (a AllianceApiService) GetAlliancesNames(allianceIds []int64, datasource interface{}) ([]GetAlliancesNames200Ok,  time.Time, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	)

	// create path and map variables
	localVarPath := "https://esi.tech.ccp.is/latest/alliances/names/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := a.client.typeCheckParameter(datasource, "string", "datasource"); err != nil {
		return nil, time.Now(), err
	}
		localVarQueryParams.Add("alliance_ids", a.client.parameterToString(allianceIds, "csv"))
	if datasource != nil {
		localVarQueryParams.Add("datasource", a.client.parameterToString(datasource, ""))
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.client.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	 var successPayload = new([]GetAlliancesNames200Ok)

	 r, err := a.client.prepareRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes, "application/json")
	 if err != nil {
		  return *successPayload, time.Now(), err
	 }

	 localVarHttpResponse, err := a.client.callAPI(r)
	 if err != nil || localVarHttpResponse == nil {
		  return *successPayload, time.Now(), err
	 }
	 defer localVarHttpResponse.Body.Close()
	 if localVarHttpResponse.StatusCode >= 300 {
		return *successPayload, time.Now(), errors.New(localVarHttpResponse.Status)
	 }
	
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
	 	return *successPayload, time.Now(), err
	}

	expires := cacheExpires(localVarHttpResponse)
	return *successPayload, expires, err
}

