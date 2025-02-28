# WebhookSubscription

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type_** | **string** |  | [default to null]
**Uuid** | **string** | The webhook&#x27;s id | [optional] [default to null]
**Url** | **string** | The URL events get delivered to. | [optional] [default to null]
**Description** | **string** | A user-defined description of the webhook. | [optional] [default to null]
**SubjectType** | **string** | The type of entity. Set to either &#x60;repository&#x60; or &#x60;workspace&#x60; based on where the subscription is defined. | [optional] [default to null]
**Subject** | [***ModelError**](map.md) |  | [optional] [default to null]
**Active** | **bool** |  | [optional] [default to null]
**CreatedAt** | [**time.Time**](time.Time.md) |  | [optional] [default to null]
**Events** | **[]string** | The events this webhook is subscribed to. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

