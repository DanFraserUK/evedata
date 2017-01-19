# \UniverseApi

All URIs are relative to *https://esi.tech.ccp.is/latest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetUniverseCategories**](UniverseApi.md#GetUniverseCategories) | **Get** /universe/categories/ | Get item categories
[**GetUniverseCategoriesCategoryId**](UniverseApi.md#GetUniverseCategoriesCategoryId) | **Get** /universe/categories/{category_id}/ | Get item category information
[**GetUniverseGroups**](UniverseApi.md#GetUniverseGroups) | **Get** /universe/groups/ | Get item groups
[**GetUniverseGroupsGroupId**](UniverseApi.md#GetUniverseGroupsGroupId) | **Get** /universe/groups/{group_id}/ | Get item group information
[**GetUniverseStationsStationId**](UniverseApi.md#GetUniverseStationsStationId) | **Get** /universe/stations/{station_id}/ | Get station information
[**GetUniverseStructures**](UniverseApi.md#GetUniverseStructures) | **Get** /universe/structures/ | List all public structures
[**GetUniverseStructuresStructureId**](UniverseApi.md#GetUniverseStructuresStructureId) | **Get** /universe/structures/{structure_id}/ | Get structure information
[**GetUniverseSystemsSystemId**](UniverseApi.md#GetUniverseSystemsSystemId) | **Get** /universe/systems/{system_id}/ | Get solar system information
[**GetUniverseTypes**](UniverseApi.md#GetUniverseTypes) | **Get** /universe/types/ | Get types
[**GetUniverseTypesTypeId**](UniverseApi.md#GetUniverseTypesTypeId) | **Get** /universe/types/{type_id}/ | Get type information
[**PostUniverseNames**](UniverseApi.md#PostUniverseNames) | **Post** /universe/names/ | Get names and categories for a set of ID&#39;s


# **GetUniverseCategories**
> []int32 GetUniverseCategories(optional)
Get item categories

Get a list of item categories

---

Alternate route: `/v1/universe/categories/`

Alternate route: `/legacy/universe/categories/`

Alternate route: `/dev/universe/categories/`


---

This route is cached for up to 3600 seconds

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **datasource** | **string**| The server name you would like data from | [default to tranquility]

### Return type

**[]int32**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUniverseCategoriesCategoryId**
> GetUniverseCategoriesCategoryIdOk GetUniverseCategoriesCategoryId(categoryId, optional)
Get item category information

Get information of an item category

---

Alternate route: `/v1/universe/categories/{category_id}/`

Alternate route: `/legacy/universe/categories/{category_id}/`

Alternate route: `/dev/universe/categories/{category_id}/`


---

This route is cached for up to 3600 seconds

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
  **categoryId** | **int32**| An Eve item category ID | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **categoryId** | **int32**| An Eve item category ID | 
 **language** | **string**| Language to use in the response | [default to en-us]
 **datasource** | **string**| The server name you would like data from | [default to tranquility]

### Return type

[**GetUniverseCategoriesCategoryIdOk**](get_universe_categories_category_id_ok.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUniverseGroups**
> []int32 GetUniverseGroups(optional)
Get item groups

Get a list of item groups

---

Alternate route: `/v1/universe/groups/`

Alternate route: `/legacy/universe/groups/`

Alternate route: `/dev/universe/groups/`


---

This route is cached for up to 3600 seconds

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32**| Which page to query | 
 **datasource** | **string**| The server name you would like data from | [default to tranquility]

### Return type

**[]int32**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUniverseGroupsGroupId**
> GetUniverseGroupsGroupIdOk GetUniverseGroupsGroupId(groupId, optional)
Get item group information

Get information on an item group

---

Alternate route: `/v1/universe/groups/{group_id}/`

Alternate route: `/legacy/universe/groups/{group_id}/`

Alternate route: `/dev/universe/groups/{group_id}/`


---

This route is cached for up to 3600 seconds

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
  **groupId** | **int32**| An Eve item group ID | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **groupId** | **int32**| An Eve item group ID | 
 **language** | **string**| Language to use in the response | [default to en-us]
 **datasource** | **string**| The server name you would like data from | [default to tranquility]

### Return type

[**GetUniverseGroupsGroupIdOk**](get_universe_groups_group_id_ok.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUniverseStationsStationId**
> GetUniverseStationsStationIdOk GetUniverseStationsStationId(stationId, optional)
Get station information

Public information on stations

---

Alternate route: `/v1/universe/stations/{station_id}/`

Alternate route: `/legacy/universe/stations/{station_id}/`

Alternate route: `/dev/universe/stations/{station_id}/`


---

This route is cached for up to 3600 seconds

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
  **stationId** | **int32**| An Eve station ID | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **stationId** | **int32**| An Eve station ID | 
 **datasource** | **string**| The server name you would like data from | [default to tranquility]

### Return type

[**GetUniverseStationsStationIdOk**](get_universe_stations_station_id_ok.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUniverseStructures**
> []int64 GetUniverseStructures(optional)
List all public structures

List all public structures

---

Alternate route: `/v1/universe/structures/`

Alternate route: `/legacy/universe/structures/`

Alternate route: `/dev/universe/structures/`


---

This route is cached for up to 3600 seconds

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **datasource** | **string**| The server name you would like data from | [default to tranquility]

### Return type

**[]int64**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUniverseStructuresStructureId**
> GetUniverseStructuresStructureIdOk GetUniverseStructuresStructureId(ctx, structureId, optional)
Get structure information

Returns information on requested structure, if you are on the ACL. Otherwise, returns "Forbidden" for all inputs.

---

Alternate route: `/v1/universe/structures/{structure_id}/`

Alternate route: `/legacy/universe/structures/{structure_id}/`

Alternate route: `/dev/universe/structures/{structure_id}/`


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **structureId** | **int64**| An Eve structure ID | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **structureId** | **int64**| An Eve structure ID | 
 **datasource** | **string**| The server name you would like data from | [default to tranquility]

### Return type

[**GetUniverseStructuresStructureIdOk**](get_universe_structures_structure_id_ok.md)

### Authorization

[evesso](../README.md#evesso)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUniverseSystemsSystemId**
> GetUniverseSystemsSystemIdOk GetUniverseSystemsSystemId(systemId, optional)
Get solar system information

Information on solar systems

---

Alternate route: `/v1/universe/systems/{system_id}/`

Alternate route: `/legacy/universe/systems/{system_id}/`

Alternate route: `/dev/universe/systems/{system_id}/`


---

This route is cached for up to 3600 seconds

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
  **systemId** | **int32**| An Eve solar system ID | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **systemId** | **int32**| An Eve solar system ID | 
 **datasource** | **string**| The server name you would like data from | [default to tranquility]

### Return type

[**GetUniverseSystemsSystemIdOk**](get_universe_systems_system_id_ok.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUniverseTypes**
> []int32 GetUniverseTypes(optional)
Get types

Get a list of type ids

---

Alternate route: `/v1/universe/types/`

Alternate route: `/legacy/universe/types/`

Alternate route: `/dev/universe/types/`


---

This route is cached for up to 3600 seconds

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32**| Which page to query | 
 **datasource** | **string**| The server name you would like data from | [default to tranquility]

### Return type

**[]int32**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUniverseTypesTypeId**
> GetUniverseTypesTypeIdOk GetUniverseTypesTypeId(typeId, optional)
Get type information

Get information on a type

---

Alternate route: `/v2/universe/types/{type_id}/`

Alternate route: `/dev/universe/types/{type_id}/`


---

This route is cached for up to 3600 seconds

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
  **typeId** | **int32**| An Eve item type ID | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **typeId** | **int32**| An Eve item type ID | 
 **language** | **string**| Language to use in the response | [default to en-us]
 **datasource** | **string**| The server name you would like data from | [default to tranquility]

### Return type

[**GetUniverseTypesTypeIdOk**](get_universe_types_type_id_ok.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostUniverseNames**
> []PostUniverseNames200Ok PostUniverseNames(ids, optional)
Get names and categories for a set of ID's

Resolve a set of IDs to names and categories. Supported ID's for resolving are: Characters, Corporations, Alliances, Stations, Solar Systems, Constellations, Regions, Types.

---

Alternate route: `/v1/universe/names/`

Alternate route: `/legacy/universe/names/`


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
  **ids** | [**PostUniverseNamesIds**](PostUniverseNamesIds.md)| The ids to resolve | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ids** | [**PostUniverseNamesIds**](PostUniverseNamesIds.md)| The ids to resolve | 
 **datasource** | **string**| The server name you would like data from | [default to tranquility]

### Return type

[**[]PostUniverseNames200Ok**](post_universe_names_200_ok.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

