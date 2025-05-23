/*
Conversations Inbox & Messages

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package conversations_inbox_&amp;_messages

import (
	"encoding/json"
	"fmt"
	"gopkg.in/validator.v2"
)

// CollectionResponsePublicMessageForwardPagingResultsInner - struct for CollectionResponsePublicMessageForwardPagingResultsInner
type CollectionResponsePublicMessageForwardPagingResultsInner struct {
	PublicAssignmentMessage *PublicAssignmentMessage
	PublicComment *PublicComment
	PublicConversationsMessage *PublicConversationsMessage
	PublicThreadInboxChange *PublicThreadInboxChange
	PublicThreadStatusChange *PublicThreadStatusChange
	PublicWelcomeMessage *PublicWelcomeMessage
}

// PublicAssignmentMessageAsCollectionResponsePublicMessageForwardPagingResultsInner is a convenience function that returns PublicAssignmentMessage wrapped in CollectionResponsePublicMessageForwardPagingResultsInner
func PublicAssignmentMessageAsCollectionResponsePublicMessageForwardPagingResultsInner(v *PublicAssignmentMessage) CollectionResponsePublicMessageForwardPagingResultsInner {
	return CollectionResponsePublicMessageForwardPagingResultsInner{
		PublicAssignmentMessage: v,
	}
}

// PublicCommentAsCollectionResponsePublicMessageForwardPagingResultsInner is a convenience function that returns PublicComment wrapped in CollectionResponsePublicMessageForwardPagingResultsInner
func PublicCommentAsCollectionResponsePublicMessageForwardPagingResultsInner(v *PublicComment) CollectionResponsePublicMessageForwardPagingResultsInner {
	return CollectionResponsePublicMessageForwardPagingResultsInner{
		PublicComment: v,
	}
}

// PublicConversationsMessageAsCollectionResponsePublicMessageForwardPagingResultsInner is a convenience function that returns PublicConversationsMessage wrapped in CollectionResponsePublicMessageForwardPagingResultsInner
func PublicConversationsMessageAsCollectionResponsePublicMessageForwardPagingResultsInner(v *PublicConversationsMessage) CollectionResponsePublicMessageForwardPagingResultsInner {
	return CollectionResponsePublicMessageForwardPagingResultsInner{
		PublicConversationsMessage: v,
	}
}

// PublicThreadInboxChangeAsCollectionResponsePublicMessageForwardPagingResultsInner is a convenience function that returns PublicThreadInboxChange wrapped in CollectionResponsePublicMessageForwardPagingResultsInner
func PublicThreadInboxChangeAsCollectionResponsePublicMessageForwardPagingResultsInner(v *PublicThreadInboxChange) CollectionResponsePublicMessageForwardPagingResultsInner {
	return CollectionResponsePublicMessageForwardPagingResultsInner{
		PublicThreadInboxChange: v,
	}
}

// PublicThreadStatusChangeAsCollectionResponsePublicMessageForwardPagingResultsInner is a convenience function that returns PublicThreadStatusChange wrapped in CollectionResponsePublicMessageForwardPagingResultsInner
func PublicThreadStatusChangeAsCollectionResponsePublicMessageForwardPagingResultsInner(v *PublicThreadStatusChange) CollectionResponsePublicMessageForwardPagingResultsInner {
	return CollectionResponsePublicMessageForwardPagingResultsInner{
		PublicThreadStatusChange: v,
	}
}

// PublicWelcomeMessageAsCollectionResponsePublicMessageForwardPagingResultsInner is a convenience function that returns PublicWelcomeMessage wrapped in CollectionResponsePublicMessageForwardPagingResultsInner
func PublicWelcomeMessageAsCollectionResponsePublicMessageForwardPagingResultsInner(v *PublicWelcomeMessage) CollectionResponsePublicMessageForwardPagingResultsInner {
	return CollectionResponsePublicMessageForwardPagingResultsInner{
		PublicWelcomeMessage: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *CollectionResponsePublicMessageForwardPagingResultsInner) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into PublicAssignmentMessage
	err = newStrictDecoder(data).Decode(&dst.PublicAssignmentMessage)
	if err == nil {
		jsonPublicAssignmentMessage, _ := json.Marshal(dst.PublicAssignmentMessage)
		if string(jsonPublicAssignmentMessage) == "{}" { // empty struct
			dst.PublicAssignmentMessage = nil
		} else {
			if err = validator.Validate(dst.PublicAssignmentMessage); err != nil {
				dst.PublicAssignmentMessage = nil
			} else {
				match++
			}
		}
	} else {
		dst.PublicAssignmentMessage = nil
	}

	// try to unmarshal data into PublicComment
	err = newStrictDecoder(data).Decode(&dst.PublicComment)
	if err == nil {
		jsonPublicComment, _ := json.Marshal(dst.PublicComment)
		if string(jsonPublicComment) == "{}" { // empty struct
			dst.PublicComment = nil
		} else {
			if err = validator.Validate(dst.PublicComment); err != nil {
				dst.PublicComment = nil
			} else {
				match++
			}
		}
	} else {
		dst.PublicComment = nil
	}

	// try to unmarshal data into PublicConversationsMessage
	err = newStrictDecoder(data).Decode(&dst.PublicConversationsMessage)
	if err == nil {
		jsonPublicConversationsMessage, _ := json.Marshal(dst.PublicConversationsMessage)
		if string(jsonPublicConversationsMessage) == "{}" { // empty struct
			dst.PublicConversationsMessage = nil
		} else {
			if err = validator.Validate(dst.PublicConversationsMessage); err != nil {
				dst.PublicConversationsMessage = nil
			} else {
				match++
			}
		}
	} else {
		dst.PublicConversationsMessage = nil
	}

	// try to unmarshal data into PublicThreadInboxChange
	err = newStrictDecoder(data).Decode(&dst.PublicThreadInboxChange)
	if err == nil {
		jsonPublicThreadInboxChange, _ := json.Marshal(dst.PublicThreadInboxChange)
		if string(jsonPublicThreadInboxChange) == "{}" { // empty struct
			dst.PublicThreadInboxChange = nil
		} else {
			if err = validator.Validate(dst.PublicThreadInboxChange); err != nil {
				dst.PublicThreadInboxChange = nil
			} else {
				match++
			}
		}
	} else {
		dst.PublicThreadInboxChange = nil
	}

	// try to unmarshal data into PublicThreadStatusChange
	err = newStrictDecoder(data).Decode(&dst.PublicThreadStatusChange)
	if err == nil {
		jsonPublicThreadStatusChange, _ := json.Marshal(dst.PublicThreadStatusChange)
		if string(jsonPublicThreadStatusChange) == "{}" { // empty struct
			dst.PublicThreadStatusChange = nil
		} else {
			if err = validator.Validate(dst.PublicThreadStatusChange); err != nil {
				dst.PublicThreadStatusChange = nil
			} else {
				match++
			}
		}
	} else {
		dst.PublicThreadStatusChange = nil
	}

	// try to unmarshal data into PublicWelcomeMessage
	err = newStrictDecoder(data).Decode(&dst.PublicWelcomeMessage)
	if err == nil {
		jsonPublicWelcomeMessage, _ := json.Marshal(dst.PublicWelcomeMessage)
		if string(jsonPublicWelcomeMessage) == "{}" { // empty struct
			dst.PublicWelcomeMessage = nil
		} else {
			if err = validator.Validate(dst.PublicWelcomeMessage); err != nil {
				dst.PublicWelcomeMessage = nil
			} else {
				match++
			}
		}
	} else {
		dst.PublicWelcomeMessage = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.PublicAssignmentMessage = nil
		dst.PublicComment = nil
		dst.PublicConversationsMessage = nil
		dst.PublicThreadInboxChange = nil
		dst.PublicThreadStatusChange = nil
		dst.PublicWelcomeMessage = nil

		return fmt.Errorf("data matches more than one schema in oneOf(CollectionResponsePublicMessageForwardPagingResultsInner)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(CollectionResponsePublicMessageForwardPagingResultsInner)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src CollectionResponsePublicMessageForwardPagingResultsInner) MarshalJSON() ([]byte, error) {
	if src.PublicAssignmentMessage != nil {
		return json.Marshal(&src.PublicAssignmentMessage)
	}

	if src.PublicComment != nil {
		return json.Marshal(&src.PublicComment)
	}

	if src.PublicConversationsMessage != nil {
		return json.Marshal(&src.PublicConversationsMessage)
	}

	if src.PublicThreadInboxChange != nil {
		return json.Marshal(&src.PublicThreadInboxChange)
	}

	if src.PublicThreadStatusChange != nil {
		return json.Marshal(&src.PublicThreadStatusChange)
	}

	if src.PublicWelcomeMessage != nil {
		return json.Marshal(&src.PublicWelcomeMessage)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *CollectionResponsePublicMessageForwardPagingResultsInner) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.PublicAssignmentMessage != nil {
		return obj.PublicAssignmentMessage
	}

	if obj.PublicComment != nil {
		return obj.PublicComment
	}

	if obj.PublicConversationsMessage != nil {
		return obj.PublicConversationsMessage
	}

	if obj.PublicThreadInboxChange != nil {
		return obj.PublicThreadInboxChange
	}

	if obj.PublicThreadStatusChange != nil {
		return obj.PublicThreadStatusChange
	}

	if obj.PublicWelcomeMessage != nil {
		return obj.PublicWelcomeMessage
	}

	// all schemas are nil
	return nil
}

// Get the actual instance value
func (obj CollectionResponsePublicMessageForwardPagingResultsInner) GetActualInstanceValue() (interface{}) {
	if obj.PublicAssignmentMessage != nil {
		return *obj.PublicAssignmentMessage
	}

	if obj.PublicComment != nil {
		return *obj.PublicComment
	}

	if obj.PublicConversationsMessage != nil {
		return *obj.PublicConversationsMessage
	}

	if obj.PublicThreadInboxChange != nil {
		return *obj.PublicThreadInboxChange
	}

	if obj.PublicThreadStatusChange != nil {
		return *obj.PublicThreadStatusChange
	}

	if obj.PublicWelcomeMessage != nil {
		return *obj.PublicWelcomeMessage
	}

	// all schemas are nil
	return nil
}

type NullableCollectionResponsePublicMessageForwardPagingResultsInner struct {
	value *CollectionResponsePublicMessageForwardPagingResultsInner
	isSet bool
}

func (v NullableCollectionResponsePublicMessageForwardPagingResultsInner) Get() *CollectionResponsePublicMessageForwardPagingResultsInner {
	return v.value
}

func (v *NullableCollectionResponsePublicMessageForwardPagingResultsInner) Set(val *CollectionResponsePublicMessageForwardPagingResultsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionResponsePublicMessageForwardPagingResultsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionResponsePublicMessageForwardPagingResultsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionResponsePublicMessageForwardPagingResultsInner(val *CollectionResponsePublicMessageForwardPagingResultsInner) *NullableCollectionResponsePublicMessageForwardPagingResultsInner {
	return &NullableCollectionResponsePublicMessageForwardPagingResultsInner{value: val, isSet: true}
}

func (v NullableCollectionResponsePublicMessageForwardPagingResultsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionResponsePublicMessageForwardPagingResultsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


