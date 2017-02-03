# \CharacterApi

All URIs are relative to *https://esi.tech.ccp.is/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCharactersCharacterIdPortrait**](CharacterApi.md#GetCharactersCharacterIdPortrait) | **Get** /characters/{character_id}/portrait/ | Get character portraits


# **GetCharactersCharacterIdPortrait**
> GetCharactersCharacterIdPortraitOk GetCharactersCharacterIdPortrait(characterId, optional)
Get character portraits

Get portrait urls for a character  ---  Alternate route: `/latest/characters/{character_id}/portrait/`  Alternate route: `/dev/characters/{character_id}/portrait/`   ---  This route is cached for up to 3600 seconds

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
  **characterId** | **int32**| An EVE character ID | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **characterId** | **int32**| An EVE character ID | 
 **datasource** | **string**| The server name you would like data from | [default to tranquility]

### Return type

[**GetCharactersCharacterIdPortraitOk**](get_characters_character_id_portrait_ok.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

