# \AllianceApi

All URIs are relative to *https://esi.tech.ccp.is/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAlliancesAllianceId**](AllianceApi.md#GetAlliancesAllianceId) | **Get** /alliances/{alliance_id}/ | Get alliance information


# **GetAlliancesAllianceId**
> GetAlliancesAllianceIdOk GetAlliancesAllianceId(allianceId, optional)
Get alliance information

Public information about an alliance  ---  Alternate route: `/latest/alliances/{alliance_id}/`   ---  This route is cached for up to 3600 seconds

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
  **allianceId** | **int32**| An Eve alliance ID | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **allianceId** | **int32**| An Eve alliance ID | 
 **datasource** | **string**| The server name you would like data from | [default to tranquility]

### Return type

[**GetAlliancesAllianceIdOk**](get_alliances_alliance_id_ok.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

