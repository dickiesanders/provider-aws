/*
Copyright 2021 The Crossplane Authors.

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

// Code generated by ack-generate. DO NOT EDIT.

package v1alpha1

type AccessDeniedForDependencyExceptionReason string

const (
	AccessDeniedForDependencyExceptionReason_ACCESS_DENIED_DURING_CREATE_SERVICE_LINKED_ROLE AccessDeniedForDependencyExceptionReason = "ACCESS_DENIED_DURING_CREATE_SERVICE_LINKED_ROLE"
)

type AccountJoinedMethod string

const (
	AccountJoinedMethod_INVITED AccountJoinedMethod = "INVITED"
	AccountJoinedMethod_CREATED AccountJoinedMethod = "CREATED"
)

type AccountStatus_SDK string

const (
	AccountStatus_SDK_ACTIVE    AccountStatus_SDK = "ACTIVE"
	AccountStatus_SDK_SUSPENDED AccountStatus_SDK = "SUSPENDED"
)

type ActionType string

const (
	ActionType_INVITE                                ActionType = "INVITE"
	ActionType_ENABLE_ALL_FEATURES                   ActionType = "ENABLE_ALL_FEATURES"
	ActionType_APPROVE_ALL_FEATURES                  ActionType = "APPROVE_ALL_FEATURES"
	ActionType_ADD_ORGANIZATIONS_SERVICE_LINKED_ROLE ActionType = "ADD_ORGANIZATIONS_SERVICE_LINKED_ROLE"
)

type ChildType string

const (
	ChildType_ACCOUNT             ChildType = "ACCOUNT"
	ChildType_ORGANIZATIONAL_UNIT ChildType = "ORGANIZATIONAL_UNIT"
)

type ConstraintViolationExceptionReason string

const (
	ConstraintViolationExceptionReason_ACCOUNT_NUMBER_LIMIT_EXCEEDED                           ConstraintViolationExceptionReason = "ACCOUNT_NUMBER_LIMIT_EXCEEDED"
	ConstraintViolationExceptionReason_HANDSHAKE_RATE_LIMIT_EXCEEDED                           ConstraintViolationExceptionReason = "HANDSHAKE_RATE_LIMIT_EXCEEDED"
	ConstraintViolationExceptionReason_OU_NUMBER_LIMIT_EXCEEDED                                ConstraintViolationExceptionReason = "OU_NUMBER_LIMIT_EXCEEDED"
	ConstraintViolationExceptionReason_OU_DEPTH_LIMIT_EXCEEDED                                 ConstraintViolationExceptionReason = "OU_DEPTH_LIMIT_EXCEEDED"
	ConstraintViolationExceptionReason_POLICY_NUMBER_LIMIT_EXCEEDED                            ConstraintViolationExceptionReason = "POLICY_NUMBER_LIMIT_EXCEEDED"
	ConstraintViolationExceptionReason_POLICY_CONTENT_LIMIT_EXCEEDED                           ConstraintViolationExceptionReason = "POLICY_CONTENT_LIMIT_EXCEEDED"
	ConstraintViolationExceptionReason_MAX_POLICY_TYPE_ATTACHMENT_LIMIT_EXCEEDED               ConstraintViolationExceptionReason = "MAX_POLICY_TYPE_ATTACHMENT_LIMIT_EXCEEDED"
	ConstraintViolationExceptionReason_MIN_POLICY_TYPE_ATTACHMENT_LIMIT_EXCEEDED               ConstraintViolationExceptionReason = "MIN_POLICY_TYPE_ATTACHMENT_LIMIT_EXCEEDED"
	ConstraintViolationExceptionReason_ACCOUNT_CANNOT_LEAVE_ORGANIZATION                       ConstraintViolationExceptionReason = "ACCOUNT_CANNOT_LEAVE_ORGANIZATION"
	ConstraintViolationExceptionReason_ACCOUNT_CANNOT_LEAVE_WITHOUT_EULA                       ConstraintViolationExceptionReason = "ACCOUNT_CANNOT_LEAVE_WITHOUT_EULA"
	ConstraintViolationExceptionReason_ACCOUNT_CANNOT_LEAVE_WITHOUT_PHONE_VERIFICATION         ConstraintViolationExceptionReason = "ACCOUNT_CANNOT_LEAVE_WITHOUT_PHONE_VERIFICATION"
	ConstraintViolationExceptionReason_MASTER_ACCOUNT_PAYMENT_INSTRUMENT_REQUIRED              ConstraintViolationExceptionReason = "MASTER_ACCOUNT_PAYMENT_INSTRUMENT_REQUIRED"
	ConstraintViolationExceptionReason_MEMBER_ACCOUNT_PAYMENT_INSTRUMENT_REQUIRED              ConstraintViolationExceptionReason = "MEMBER_ACCOUNT_PAYMENT_INSTRUMENT_REQUIRED"
	ConstraintViolationExceptionReason_ACCOUNT_CREATION_RATE_LIMIT_EXCEEDED                    ConstraintViolationExceptionReason = "ACCOUNT_CREATION_RATE_LIMIT_EXCEEDED"
	ConstraintViolationExceptionReason_MASTER_ACCOUNT_ADDRESS_DOES_NOT_MATCH_MARKETPLACE       ConstraintViolationExceptionReason = "MASTER_ACCOUNT_ADDRESS_DOES_NOT_MATCH_MARKETPLACE"
	ConstraintViolationExceptionReason_MASTER_ACCOUNT_MISSING_CONTACT_INFO                     ConstraintViolationExceptionReason = "MASTER_ACCOUNT_MISSING_CONTACT_INFO"
	ConstraintViolationExceptionReason_MASTER_ACCOUNT_NOT_GOVCLOUD_ENABLED                     ConstraintViolationExceptionReason = "MASTER_ACCOUNT_NOT_GOVCLOUD_ENABLED"
	ConstraintViolationExceptionReason_ORGANIZATION_NOT_IN_ALL_FEATURES_MODE                   ConstraintViolationExceptionReason = "ORGANIZATION_NOT_IN_ALL_FEATURES_MODE"
	ConstraintViolationExceptionReason_CREATE_ORGANIZATION_IN_BILLING_MODE_UNSUPPORTED_REGION  ConstraintViolationExceptionReason = "CREATE_ORGANIZATION_IN_BILLING_MODE_UNSUPPORTED_REGION"
	ConstraintViolationExceptionReason_EMAIL_VERIFICATION_CODE_EXPIRED                         ConstraintViolationExceptionReason = "EMAIL_VERIFICATION_CODE_EXPIRED"
	ConstraintViolationExceptionReason_WAIT_PERIOD_ACTIVE                                      ConstraintViolationExceptionReason = "WAIT_PERIOD_ACTIVE"
	ConstraintViolationExceptionReason_MAX_TAG_LIMIT_EXCEEDED                                  ConstraintViolationExceptionReason = "MAX_TAG_LIMIT_EXCEEDED"
	ConstraintViolationExceptionReason_TAG_POLICY_VIOLATION                                    ConstraintViolationExceptionReason = "TAG_POLICY_VIOLATION"
	ConstraintViolationExceptionReason_MAX_DELEGATED_ADMINISTRATORS_FOR_SERVICE_LIMIT_EXCEEDED ConstraintViolationExceptionReason = "MAX_DELEGATED_ADMINISTRATORS_FOR_SERVICE_LIMIT_EXCEEDED"
	ConstraintViolationExceptionReason_CANNOT_REGISTER_MASTER_AS_DELEGATED_ADMINISTRATOR       ConstraintViolationExceptionReason = "CANNOT_REGISTER_MASTER_AS_DELEGATED_ADMINISTRATOR"
	ConstraintViolationExceptionReason_CANNOT_REMOVE_DELEGATED_ADMINISTRATOR_FROM_ORG          ConstraintViolationExceptionReason = "CANNOT_REMOVE_DELEGATED_ADMINISTRATOR_FROM_ORG"
	ConstraintViolationExceptionReason_DELEGATED_ADMINISTRATOR_EXISTS_FOR_THIS_SERVICE         ConstraintViolationExceptionReason = "DELEGATED_ADMINISTRATOR_EXISTS_FOR_THIS_SERVICE"
	ConstraintViolationExceptionReason_MASTER_ACCOUNT_MISSING_BUSINESS_LICENSE                 ConstraintViolationExceptionReason = "MASTER_ACCOUNT_MISSING_BUSINESS_LICENSE"
)

type CreateAccountFailureReason string

const (
	CreateAccountFailureReason_ACCOUNT_LIMIT_EXCEEDED          CreateAccountFailureReason = "ACCOUNT_LIMIT_EXCEEDED"
	CreateAccountFailureReason_EMAIL_ALREADY_EXISTS            CreateAccountFailureReason = "EMAIL_ALREADY_EXISTS"
	CreateAccountFailureReason_INVALID_ADDRESS                 CreateAccountFailureReason = "INVALID_ADDRESS"
	CreateAccountFailureReason_INVALID_EMAIL                   CreateAccountFailureReason = "INVALID_EMAIL"
	CreateAccountFailureReason_CONCURRENT_ACCOUNT_MODIFICATION CreateAccountFailureReason = "CONCURRENT_ACCOUNT_MODIFICATION"
	CreateAccountFailureReason_INTERNAL_FAILURE                CreateAccountFailureReason = "INTERNAL_FAILURE"
	CreateAccountFailureReason_GOVCLOUD_ACCOUNT_ALREADY_EXISTS CreateAccountFailureReason = "GOVCLOUD_ACCOUNT_ALREADY_EXISTS"
	CreateAccountFailureReason_MISSING_BUSINESS_VALIDATION     CreateAccountFailureReason = "MISSING_BUSINESS_VALIDATION"
	CreateAccountFailureReason_MISSING_PAYMENT_INSTRUMENT      CreateAccountFailureReason = "MISSING_PAYMENT_INSTRUMENT"
)

type CreateAccountState string

const (
	CreateAccountState_IN_PROGRESS CreateAccountState = "IN_PROGRESS"
	CreateAccountState_SUCCEEDED   CreateAccountState = "SUCCEEDED"
	CreateAccountState_FAILED      CreateAccountState = "FAILED"
)

type EffectivePolicyType string

const (
	EffectivePolicyType_TAG_POLICY                EffectivePolicyType = "TAG_POLICY"
	EffectivePolicyType_BACKUP_POLICY             EffectivePolicyType = "BACKUP_POLICY"
	EffectivePolicyType_AISERVICES_OPT_OUT_POLICY EffectivePolicyType = "AISERVICES_OPT_OUT_POLICY"
)

type HandshakeConstraintViolationExceptionReason string

const (
	HandshakeConstraintViolationExceptionReason_ACCOUNT_NUMBER_LIMIT_EXCEEDED                      HandshakeConstraintViolationExceptionReason = "ACCOUNT_NUMBER_LIMIT_EXCEEDED"
	HandshakeConstraintViolationExceptionReason_HANDSHAKE_RATE_LIMIT_EXCEEDED                      HandshakeConstraintViolationExceptionReason = "HANDSHAKE_RATE_LIMIT_EXCEEDED"
	HandshakeConstraintViolationExceptionReason_ALREADY_IN_AN_ORGANIZATION                         HandshakeConstraintViolationExceptionReason = "ALREADY_IN_AN_ORGANIZATION"
	HandshakeConstraintViolationExceptionReason_ORGANIZATION_ALREADY_HAS_ALL_FEATURES              HandshakeConstraintViolationExceptionReason = "ORGANIZATION_ALREADY_HAS_ALL_FEATURES"
	HandshakeConstraintViolationExceptionReason_INVITE_DISABLED_DURING_ENABLE_ALL_FEATURES         HandshakeConstraintViolationExceptionReason = "INVITE_DISABLED_DURING_ENABLE_ALL_FEATURES"
	HandshakeConstraintViolationExceptionReason_PAYMENT_INSTRUMENT_REQUIRED                        HandshakeConstraintViolationExceptionReason = "PAYMENT_INSTRUMENT_REQUIRED"
	HandshakeConstraintViolationExceptionReason_ORGANIZATION_FROM_DIFFERENT_SELLER_OF_RECORD       HandshakeConstraintViolationExceptionReason = "ORGANIZATION_FROM_DIFFERENT_SELLER_OF_RECORD"
	HandshakeConstraintViolationExceptionReason_ORGANIZATION_MEMBERSHIP_CHANGE_RATE_LIMIT_EXCEEDED HandshakeConstraintViolationExceptionReason = "ORGANIZATION_MEMBERSHIP_CHANGE_RATE_LIMIT_EXCEEDED"
)

type HandshakePartyType string

const (
	HandshakePartyType_ACCOUNT      HandshakePartyType = "ACCOUNT"
	HandshakePartyType_ORGANIZATION HandshakePartyType = "ORGANIZATION"
	HandshakePartyType_EMAIL        HandshakePartyType = "EMAIL"
)

type HandshakeResourceType string

const (
	HandshakeResourceType_ACCOUNT                  HandshakeResourceType = "ACCOUNT"
	HandshakeResourceType_ORGANIZATION             HandshakeResourceType = "ORGANIZATION"
	HandshakeResourceType_ORGANIZATION_FEATURE_SET HandshakeResourceType = "ORGANIZATION_FEATURE_SET"
	HandshakeResourceType_EMAIL                    HandshakeResourceType = "EMAIL"
	HandshakeResourceType_MASTER_EMAIL             HandshakeResourceType = "MASTER_EMAIL"
	HandshakeResourceType_MASTER_NAME              HandshakeResourceType = "MASTER_NAME"
	HandshakeResourceType_NOTES                    HandshakeResourceType = "NOTES"
	HandshakeResourceType_PARENT_HANDSHAKE         HandshakeResourceType = "PARENT_HANDSHAKE"
)

type HandshakeState string

const (
	HandshakeState_REQUESTED HandshakeState = "REQUESTED"
	HandshakeState_OPEN      HandshakeState = "OPEN"
	HandshakeState_CANCELED  HandshakeState = "CANCELED"
	HandshakeState_ACCEPTED  HandshakeState = "ACCEPTED"
	HandshakeState_DECLINED  HandshakeState = "DECLINED"
	HandshakeState_EXPIRED   HandshakeState = "EXPIRED"
)

type IAMUserAccessToBilling string

const (
	IAMUserAccessToBilling_ALLOW IAMUserAccessToBilling = "ALLOW"
	IAMUserAccessToBilling_DENY  IAMUserAccessToBilling = "DENY"
)

type InvalidInputExceptionReason string

const (
	InvalidInputExceptionReason_INVALID_PARTY_TYPE_TARGET              InvalidInputExceptionReason = "INVALID_PARTY_TYPE_TARGET"
	InvalidInputExceptionReason_INVALID_SYNTAX_ORGANIZATION_ARN        InvalidInputExceptionReason = "INVALID_SYNTAX_ORGANIZATION_ARN"
	InvalidInputExceptionReason_INVALID_SYNTAX_POLICY_ID               InvalidInputExceptionReason = "INVALID_SYNTAX_POLICY_ID"
	InvalidInputExceptionReason_INVALID_ENUM                           InvalidInputExceptionReason = "INVALID_ENUM"
	InvalidInputExceptionReason_INVALID_ENUM_POLICY_TYPE               InvalidInputExceptionReason = "INVALID_ENUM_POLICY_TYPE"
	InvalidInputExceptionReason_INVALID_LIST_MEMBER                    InvalidInputExceptionReason = "INVALID_LIST_MEMBER"
	InvalidInputExceptionReason_MAX_LENGTH_EXCEEDED                    InvalidInputExceptionReason = "MAX_LENGTH_EXCEEDED"
	InvalidInputExceptionReason_MAX_VALUE_EXCEEDED                     InvalidInputExceptionReason = "MAX_VALUE_EXCEEDED"
	InvalidInputExceptionReason_MIN_LENGTH_EXCEEDED                    InvalidInputExceptionReason = "MIN_LENGTH_EXCEEDED"
	InvalidInputExceptionReason_MIN_VALUE_EXCEEDED                     InvalidInputExceptionReason = "MIN_VALUE_EXCEEDED"
	InvalidInputExceptionReason_IMMUTABLE_POLICY                       InvalidInputExceptionReason = "IMMUTABLE_POLICY"
	InvalidInputExceptionReason_INVALID_PATTERN                        InvalidInputExceptionReason = "INVALID_PATTERN"
	InvalidInputExceptionReason_INVALID_PATTERN_TARGET_ID              InvalidInputExceptionReason = "INVALID_PATTERN_TARGET_ID"
	InvalidInputExceptionReason_INPUT_REQUIRED                         InvalidInputExceptionReason = "INPUT_REQUIRED"
	InvalidInputExceptionReason_INVALID_NEXT_TOKEN                     InvalidInputExceptionReason = "INVALID_NEXT_TOKEN"
	InvalidInputExceptionReason_MAX_LIMIT_EXCEEDED_FILTER              InvalidInputExceptionReason = "MAX_LIMIT_EXCEEDED_FILTER"
	InvalidInputExceptionReason_MOVING_ACCOUNT_BETWEEN_DIFFERENT_ROOTS InvalidInputExceptionReason = "MOVING_ACCOUNT_BETWEEN_DIFFERENT_ROOTS"
	InvalidInputExceptionReason_INVALID_FULL_NAME_TARGET               InvalidInputExceptionReason = "INVALID_FULL_NAME_TARGET"
	InvalidInputExceptionReason_UNRECOGNIZED_SERVICE_PRINCIPAL         InvalidInputExceptionReason = "UNRECOGNIZED_SERVICE_PRINCIPAL"
	InvalidInputExceptionReason_INVALID_ROLE_NAME                      InvalidInputExceptionReason = "INVALID_ROLE_NAME"
	InvalidInputExceptionReason_INVALID_SYSTEM_TAGS_PARAMETER          InvalidInputExceptionReason = "INVALID_SYSTEM_TAGS_PARAMETER"
	InvalidInputExceptionReason_DUPLICATE_TAG_KEY                      InvalidInputExceptionReason = "DUPLICATE_TAG_KEY"
	InvalidInputExceptionReason_TARGET_NOT_SUPPORTED                   InvalidInputExceptionReason = "TARGET_NOT_SUPPORTED"
)

type OrganizationFeatureSet string

const (
	OrganizationFeatureSet_ALL                  OrganizationFeatureSet = "ALL"
	OrganizationFeatureSet_CONSOLIDATED_BILLING OrganizationFeatureSet = "CONSOLIDATED_BILLING"
)

type ParentType string

const (
	ParentType_ROOT                ParentType = "ROOT"
	ParentType_ORGANIZATIONAL_UNIT ParentType = "ORGANIZATIONAL_UNIT"
)

type PolicyType string

const (
	PolicyType_SERVICE_CONTROL_POLICY    PolicyType = "SERVICE_CONTROL_POLICY"
	PolicyType_TAG_POLICY                PolicyType = "TAG_POLICY"
	PolicyType_BACKUP_POLICY             PolicyType = "BACKUP_POLICY"
	PolicyType_AISERVICES_OPT_OUT_POLICY PolicyType = "AISERVICES_OPT_OUT_POLICY"
)

type PolicyTypeStatus string

const (
	PolicyTypeStatus_ENABLED         PolicyTypeStatus = "ENABLED"
	PolicyTypeStatus_PENDING_ENABLE  PolicyTypeStatus = "PENDING_ENABLE"
	PolicyTypeStatus_PENDING_DISABLE PolicyTypeStatus = "PENDING_DISABLE"
)

type TargetType string

const (
	TargetType_ACCOUNT             TargetType = "ACCOUNT"
	TargetType_ORGANIZATIONAL_UNIT TargetType = "ORGANIZATIONAL_UNIT"
	TargetType_ROOT                TargetType = "ROOT"
)
