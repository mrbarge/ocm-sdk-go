/*
Copyright (c) 2020 Red Hat, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// IMPORTANT: This file has been generated automatically, refrain from modifying it manually as all
// your changes will be lost when the file is generated again.

package v1 // github.com/openshift-online/ocm-sdk-go/authorizations/v1

// CapabilityReviewRequest represents the values of the 'capability_review_request' type.
//
// Representation of a capability review.
type CapabilityReviewRequest struct {
	accountUsername *string
	capability      *string
	clusterID       *string
	organizationID  *string
	resourceType    *string
	subscriptionID  *string
	type_           *string
}

// Empty returns true if the object is empty, i.e. no attribute has a value.
func (o *CapabilityReviewRequest) Empty() bool {
	return o == nil || (o.accountUsername == nil &&
		o.capability == nil &&
		o.clusterID == nil &&
		o.organizationID == nil &&
		o.resourceType == nil &&
		o.subscriptionID == nil &&
		o.type_ == nil &&
		true)
}

// AccountUsername returns the value of the 'account_username' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// Defines the username of the account of which capability is being reviewed.
func (o *CapabilityReviewRequest) AccountUsername() string {
	if o != nil && o.accountUsername != nil {
		return *o.accountUsername
	}
	return ""
}

// GetAccountUsername returns the value of the 'account_username' attribute and
// a flag indicating if the attribute has a value.
//
// Defines the username of the account of which capability is being reviewed.
func (o *CapabilityReviewRequest) GetAccountUsername() (value string, ok bool) {
	ok = o != nil && o.accountUsername != nil
	if ok {
		value = *o.accountUsername
	}
	return
}

// Capability returns the value of the 'capability' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// Capability to review [manage_cluster_admin].
func (o *CapabilityReviewRequest) Capability() string {
	if o != nil && o.capability != nil {
		return *o.capability
	}
	return ""
}

// GetCapability returns the value of the 'capability' attribute and
// a flag indicating if the attribute has a value.
//
// Capability to review [manage_cluster_admin].
func (o *CapabilityReviewRequest) GetCapability() (value string, ok bool) {
	ok = o != nil && o.capability != nil
	if ok {
		value = *o.capability
	}
	return
}

// ClusterID returns the value of the 'cluster_ID' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// Indicates which Cluster (internal id) the resource type belongs to.
func (o *CapabilityReviewRequest) ClusterID() string {
	if o != nil && o.clusterID != nil {
		return *o.clusterID
	}
	return ""
}

// GetClusterID returns the value of the 'cluster_ID' attribute and
// a flag indicating if the attribute has a value.
//
// Indicates which Cluster (internal id) the resource type belongs to.
func (o *CapabilityReviewRequest) GetClusterID() (value string, ok bool) {
	ok = o != nil && o.clusterID != nil
	if ok {
		value = *o.clusterID
	}
	return
}

// OrganizationID returns the value of the 'organization_ID' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// Indicates which Organization the resource type belongs to.
func (o *CapabilityReviewRequest) OrganizationID() string {
	if o != nil && o.organizationID != nil {
		return *o.organizationID
	}
	return ""
}

// GetOrganizationID returns the value of the 'organization_ID' attribute and
// a flag indicating if the attribute has a value.
//
// Indicates which Organization the resource type belongs to.
func (o *CapabilityReviewRequest) GetOrganizationID() (value string, ok bool) {
	ok = o != nil && o.organizationID != nil
	if ok {
		value = *o.organizationID
	}
	return
}

// ResourceType returns the value of the 'resource_type' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// Indicates the type of the resource.
// See uhc-account-manager/openapi/openapi.yaml for a list of possible values.
func (o *CapabilityReviewRequest) ResourceType() string {
	if o != nil && o.resourceType != nil {
		return *o.resourceType
	}
	return ""
}

// GetResourceType returns the value of the 'resource_type' attribute and
// a flag indicating if the attribute has a value.
//
// Indicates the type of the resource.
// See uhc-account-manager/openapi/openapi.yaml for a list of possible values.
func (o *CapabilityReviewRequest) GetResourceType() (value string, ok bool) {
	ok = o != nil && o.resourceType != nil
	if ok {
		value = *o.resourceType
	}
	return
}

// SubscriptionID returns the value of the 'subscription_ID' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// Indicates which Subscription the resource type belongs to.
func (o *CapabilityReviewRequest) SubscriptionID() string {
	if o != nil && o.subscriptionID != nil {
		return *o.subscriptionID
	}
	return ""
}

// GetSubscriptionID returns the value of the 'subscription_ID' attribute and
// a flag indicating if the attribute has a value.
//
// Indicates which Subscription the resource type belongs to.
func (o *CapabilityReviewRequest) GetSubscriptionID() (value string, ok bool) {
	ok = o != nil && o.subscriptionID != nil
	if ok {
		value = *o.subscriptionID
	}
	return
}

// Type returns the value of the 'type' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// Type of capability [Cluster].
func (o *CapabilityReviewRequest) Type() string {
	if o != nil && o.type_ != nil {
		return *o.type_
	}
	return ""
}

// GetType returns the value of the 'type' attribute and
// a flag indicating if the attribute has a value.
//
// Type of capability [Cluster].
func (o *CapabilityReviewRequest) GetType() (value string, ok bool) {
	ok = o != nil && o.type_ != nil
	if ok {
		value = *o.type_
	}
	return
}

// CapabilityReviewRequestListKind is the name of the type used to represent list of objects of
// type 'capability_review_request'.
const CapabilityReviewRequestListKind = "CapabilityReviewRequestList"

// CapabilityReviewRequestListLinkKind is the name of the type used to represent links to list
// of objects of type 'capability_review_request'.
const CapabilityReviewRequestListLinkKind = "CapabilityReviewRequestListLink"

// CapabilityReviewRequestNilKind is the name of the type used to nil lists of objects of
// type 'capability_review_request'.
const CapabilityReviewRequestListNilKind = "CapabilityReviewRequestListNil"

// CapabilityReviewRequestList is a list of values of the 'capability_review_request' type.
type CapabilityReviewRequestList struct {
	href  *string
	link  bool
	items []*CapabilityReviewRequest
}

// Len returns the length of the list.
func (l *CapabilityReviewRequestList) Len() int {
	if l == nil {
		return 0
	}
	return len(l.items)
}

// Empty returns true if the list is empty.
func (l *CapabilityReviewRequestList) Empty() bool {
	return l == nil || len(l.items) == 0
}

// Get returns the item of the list with the given index. If there is no item with
// that index it returns nil.
func (l *CapabilityReviewRequestList) Get(i int) *CapabilityReviewRequest {
	if l == nil || i < 0 || i >= len(l.items) {
		return nil
	}
	return l.items[i]
}

// Slice returns an slice containing the items of the list. The returned slice is a
// copy of the one used internally, so it can be modified without affecting the
// internal representation.
//
// If you don't need to modify the returned slice consider using the Each or Range
// functions, as they don't need to allocate a new slice.
func (l *CapabilityReviewRequestList) Slice() []*CapabilityReviewRequest {
	var slice []*CapabilityReviewRequest
	if l == nil {
		slice = make([]*CapabilityReviewRequest, 0)
	} else {
		slice = make([]*CapabilityReviewRequest, len(l.items))
		copy(slice, l.items)
	}
	return slice
}

// Each runs the given function for each item of the list, in order. If the function
// returns false the iteration stops, otherwise it continues till all the elements
// of the list have been processed.
func (l *CapabilityReviewRequestList) Each(f func(item *CapabilityReviewRequest) bool) {
	if l == nil {
		return
	}
	for _, item := range l.items {
		if !f(item) {
			break
		}
	}
}

// Range runs the given function for each index and item of the list, in order. If
// the function returns false the iteration stops, otherwise it continues till all
// the elements of the list have been processed.
func (l *CapabilityReviewRequestList) Range(f func(index int, item *CapabilityReviewRequest) bool) {
	if l == nil {
		return
	}
	for index, item := range l.items {
		if !f(index, item) {
			break
		}
	}
}
