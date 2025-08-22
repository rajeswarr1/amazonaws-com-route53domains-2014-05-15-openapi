package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// DisassociateDelegationSignerFromDomainResponse represents the DisassociateDelegationSignerFromDomainResponse schema from the OpenAPI specification
type DisassociateDelegationSignerFromDomainResponse struct {
	Operationid interface{} `json:"OperationId,omitempty"`
}

// ExtraParam represents the ExtraParam schema from the OpenAPI specification
type ExtraParam struct {
	Name interface{} `json:"Name"`
	Value interface{} `json:"Value"`
}

// DeleteDomainResponse represents the DeleteDomainResponse schema from the OpenAPI specification
type DeleteDomainResponse struct {
	Operationid interface{} `json:"OperationId,omitempty"`
}

// RejectDomainTransferFromAnotherAwsAccountResponse represents the RejectDomainTransferFromAnotherAwsAccountResponse schema from the OpenAPI specification
type RejectDomainTransferFromAnotherAwsAccountResponse struct {
	Operationid interface{} `json:"OperationId,omitempty"`
}

// AcceptDomainTransferFromAnotherAwsAccountRequest represents the AcceptDomainTransferFromAnotherAwsAccountRequest schema from the OpenAPI specification
type AcceptDomainTransferFromAnotherAwsAccountRequest struct {
	Domainname interface{} `json:"DomainName"`
	Password interface{} `json:"Password"`
}

// DomainSuggestion represents the DomainSuggestion schema from the OpenAPI specification
type DomainSuggestion struct {
	Availability interface{} `json:"Availability,omitempty"`
	Domainname interface{} `json:"DomainName,omitempty"`
}

// ListDomainsRequest represents the ListDomainsRequest schema from the OpenAPI specification
type ListDomainsRequest struct {
	Sortcondition interface{} `json:"SortCondition,omitempty"`
	Filterconditions interface{} `json:"FilterConditions,omitempty"`
	Marker interface{} `json:"Marker,omitempty"`
	Maxitems interface{} `json:"MaxItems,omitempty"`
}

// GetDomainDetailRequest represents the GetDomainDetailRequest schema from the OpenAPI specification
type GetDomainDetailRequest struct {
	Domainname interface{} `json:"DomainName"`
}

// DomainSummary represents the DomainSummary schema from the OpenAPI specification
type DomainSummary struct {
	Transferlock interface{} `json:"TransferLock,omitempty"`
	Autorenew interface{} `json:"AutoRenew,omitempty"`
	Domainname interface{} `json:"DomainName,omitempty"`
	Expiry interface{} `json:"Expiry,omitempty"`
}

// OperationSummary represents the OperationSummary schema from the OpenAPI specification
type OperationSummary struct {
	Lastupdateddate interface{} `json:"LastUpdatedDate,omitempty"`
	Message interface{} `json:"Message,omitempty"`
	Operationid interface{} `json:"OperationId,omitempty"`
	Status interface{} `json:"Status,omitempty"`
	Statusflag interface{} `json:"StatusFlag,omitempty"`
	Submitteddate interface{} `json:"SubmittedDate,omitempty"`
	TypeField interface{} `json:"Type,omitempty"`
	Domainname interface{} `json:"DomainName,omitempty"`
}

// PushDomainRequest represents the PushDomainRequest schema from the OpenAPI specification
type PushDomainRequest struct {
	Domainname interface{} `json:"DomainName"`
	Target interface{} `json:"Target"`
}

// DeleteTagsForDomainResponse represents the DeleteTagsForDomainResponse schema from the OpenAPI specification
type DeleteTagsForDomainResponse struct {
}

// RetrieveDomainAuthCodeResponse represents the RetrieveDomainAuthCodeResponse schema from the OpenAPI specification
type RetrieveDomainAuthCodeResponse struct {
	Authcode interface{} `json:"AuthCode,omitempty"`
}

// UpdateDomainNameserversRequest represents the UpdateDomainNameserversRequest schema from the OpenAPI specification
type UpdateDomainNameserversRequest struct {
	Domainname interface{} `json:"DomainName"`
	Fiauthkey interface{} `json:"FIAuthKey,omitempty"`
	Nameservers interface{} `json:"Nameservers"`
}

// EnableDomainAutoRenewResponse represents the EnableDomainAutoRenewResponse schema from the OpenAPI specification
type EnableDomainAutoRenewResponse struct {
}

// AssociateDelegationSignerToDomainRequest represents the AssociateDelegationSignerToDomainRequest schema from the OpenAPI specification
type AssociateDelegationSignerToDomainRequest struct {
	Domainname interface{} `json:"DomainName"`
	Signingattributes interface{} `json:"SigningAttributes"`
}

// CancelDomainTransferToAnotherAwsAccountRequest represents the CancelDomainTransferToAnotherAwsAccountRequest schema from the OpenAPI specification
type CancelDomainTransferToAnotherAwsAccountRequest struct {
	Domainname interface{} `json:"DomainName"`
}

// UpdateTagsForDomainResponse represents the UpdateTagsForDomainResponse schema from the OpenAPI specification
type UpdateTagsForDomainResponse struct {
}

// DisableDomainAutoRenewRequest represents the DisableDomainAutoRenewRequest schema from the OpenAPI specification
type DisableDomainAutoRenewRequest struct {
	Domainname interface{} `json:"DomainName"`
}

// UpdateDomainContactResponse represents the UpdateDomainContactResponse schema from the OpenAPI specification
type UpdateDomainContactResponse struct {
	Operationid interface{} `json:"OperationId,omitempty"`
}

// RenewDomainResponse represents the RenewDomainResponse schema from the OpenAPI specification
type RenewDomainResponse struct {
	Operationid interface{} `json:"OperationId,omitempty"`
}

// FilterCondition represents the FilterCondition schema from the OpenAPI specification
type FilterCondition struct {
	Operator interface{} `json:"Operator"`
	Values interface{} `json:"Values"`
	Name interface{} `json:"Name"`
}

// BillingRecord represents the BillingRecord schema from the OpenAPI specification
type BillingRecord struct {
	Domainname interface{} `json:"DomainName,omitempty"`
	Invoiceid interface{} `json:"InvoiceId,omitempty"`
	Operation interface{} `json:"Operation,omitempty"`
	Price interface{} `json:"Price,omitempty"`
	Billdate interface{} `json:"BillDate,omitempty"`
}

// CheckDomainTransferabilityRequest represents the CheckDomainTransferabilityRequest schema from the OpenAPI specification
type CheckDomainTransferabilityRequest struct {
	Authcode interface{} `json:"AuthCode,omitempty"`
	Domainname interface{} `json:"DomainName"`
}

// TransferDomainRequest represents the TransferDomainRequest schema from the OpenAPI specification
type TransferDomainRequest struct {
	Nameservers interface{} `json:"Nameservers,omitempty"`
	Privacyprotecttechcontact interface{} `json:"PrivacyProtectTechContact,omitempty"`
	Durationinyears interface{} `json:"DurationInYears"`
	Idnlangcode interface{} `json:"IdnLangCode,omitempty"`
	Registrantcontact interface{} `json:"RegistrantContact"`
	Domainname interface{} `json:"DomainName"`
	Privacyprotectadmincontact interface{} `json:"PrivacyProtectAdminContact,omitempty"`
	Authcode interface{} `json:"AuthCode,omitempty"`
	Autorenew interface{} `json:"AutoRenew,omitempty"`
	Privacyprotectregistrantcontact interface{} `json:"PrivacyProtectRegistrantContact,omitempty"`
	Techcontact interface{} `json:"TechContact"`
	Admincontact interface{} `json:"AdminContact"`
}

// EnableDomainAutoRenewRequest represents the EnableDomainAutoRenewRequest schema from the OpenAPI specification
type EnableDomainAutoRenewRequest struct {
	Domainname interface{} `json:"DomainName"`
}

// TransferDomainResponse represents the TransferDomainResponse schema from the OpenAPI specification
type TransferDomainResponse struct {
	Operationid interface{} `json:"OperationId,omitempty"`
}

// DomainPrice represents the DomainPrice schema from the OpenAPI specification
type DomainPrice struct {
	Transferprice interface{} `json:"TransferPrice,omitempty"`
	Changeownershipprice interface{} `json:"ChangeOwnershipPrice,omitempty"`
	Name interface{} `json:"Name,omitempty"`
	Registrationprice interface{} `json:"RegistrationPrice,omitempty"`
	Renewalprice interface{} `json:"RenewalPrice,omitempty"`
	Restorationprice interface{} `json:"RestorationPrice,omitempty"`
}

// EnableDomainTransferLockRequest represents the EnableDomainTransferLockRequest schema from the OpenAPI specification
type EnableDomainTransferLockRequest struct {
	Domainname interface{} `json:"DomainName"`
}

// CheckDomainAvailabilityRequest represents the CheckDomainAvailabilityRequest schema from the OpenAPI specification
type CheckDomainAvailabilityRequest struct {
	Domainname interface{} `json:"DomainName"`
	Idnlangcode interface{} `json:"IdnLangCode,omitempty"`
}

// EnableDomainTransferLockResponse represents the EnableDomainTransferLockResponse schema from the OpenAPI specification
type EnableDomainTransferLockResponse struct {
	Operationid interface{} `json:"OperationId,omitempty"`
}

// CheckDomainAvailabilityResponse represents the CheckDomainAvailabilityResponse schema from the OpenAPI specification
type CheckDomainAvailabilityResponse struct {
	Availability interface{} `json:"Availability,omitempty"`
}

// DisableDomainTransferLockResponse represents the DisableDomainTransferLockResponse schema from the OpenAPI specification
type DisableDomainTransferLockResponse struct {
	Operationid interface{} `json:"OperationId,omitempty"`
}

// DisableDomainAutoRenewResponse represents the DisableDomainAutoRenewResponse schema from the OpenAPI specification
type DisableDomainAutoRenewResponse struct {
}

// ListTagsForDomainResponse represents the ListTagsForDomainResponse schema from the OpenAPI specification
type ListTagsForDomainResponse struct {
	Taglist interface{} `json:"TagList,omitempty"`
}

// RenewDomainRequest represents the RenewDomainRequest schema from the OpenAPI specification
type RenewDomainRequest struct {
	Durationinyears interface{} `json:"DurationInYears,omitempty"`
	Currentexpiryyear interface{} `json:"CurrentExpiryYear"`
	Domainname interface{} `json:"DomainName"`
}

// GetOperationDetailRequest represents the GetOperationDetailRequest schema from the OpenAPI specification
type GetOperationDetailRequest struct {
	Operationid interface{} `json:"OperationId"`
}

// ListDomainsResponse represents the ListDomainsResponse schema from the OpenAPI specification
type ListDomainsResponse struct {
	Domains interface{} `json:"Domains,omitempty"`
	Nextpagemarker interface{} `json:"NextPageMarker,omitempty"`
}

// UpdateDomainContactPrivacyRequest represents the UpdateDomainContactPrivacyRequest schema from the OpenAPI specification
type UpdateDomainContactPrivacyRequest struct {
	Techprivacy interface{} `json:"TechPrivacy,omitempty"`
	Adminprivacy interface{} `json:"AdminPrivacy,omitempty"`
	Domainname interface{} `json:"DomainName"`
	Registrantprivacy interface{} `json:"RegistrantPrivacy,omitempty"`
}

// ListOperationsResponse represents the ListOperationsResponse schema from the OpenAPI specification
type ListOperationsResponse struct {
	Nextpagemarker interface{} `json:"NextPageMarker,omitempty"`
	Operations interface{} `json:"Operations,omitempty"`
}

// ResendContactReachabilityEmailResponse represents the ResendContactReachabilityEmailResponse schema from the OpenAPI specification
type ResendContactReachabilityEmailResponse struct {
	Isalreadyverified interface{} `json:"isAlreadyVerified,omitempty"`
	Domainname interface{} `json:"domainName,omitempty"`
	Emailaddress interface{} `json:"emailAddress,omitempty"`
}

// RegisterDomainRequest represents the RegisterDomainRequest schema from the OpenAPI specification
type RegisterDomainRequest struct {
	Domainname interface{} `json:"DomainName"`
	Privacyprotectregistrantcontact interface{} `json:"PrivacyProtectRegistrantContact,omitempty"`
	Registrantcontact interface{} `json:"RegistrantContact"`
	Autorenew interface{} `json:"AutoRenew,omitempty"`
	Techcontact interface{} `json:"TechContact"`
	Admincontact interface{} `json:"AdminContact"`
	Durationinyears interface{} `json:"DurationInYears"`
	Idnlangcode interface{} `json:"IdnLangCode,omitempty"`
	Privacyprotectadmincontact interface{} `json:"PrivacyProtectAdminContact,omitempty"`
	Privacyprotecttechcontact interface{} `json:"PrivacyProtectTechContact,omitempty"`
}

// TransferDomainToAnotherAwsAccountRequest represents the TransferDomainToAnotherAwsAccountRequest schema from the OpenAPI specification
type TransferDomainToAnotherAwsAccountRequest struct {
	Accountid interface{} `json:"AccountId"`
	Domainname interface{} `json:"DomainName"`
}

// TransferDomainToAnotherAwsAccountResponse represents the TransferDomainToAnotherAwsAccountResponse schema from the OpenAPI specification
type TransferDomainToAnotherAwsAccountResponse struct {
	Password interface{} `json:"Password,omitempty"`
	Operationid interface{} `json:"OperationId,omitempty"`
}

// GetOperationDetailResponse represents the GetOperationDetailResponse schema from the OpenAPI specification
type GetOperationDetailResponse struct {
	TypeField interface{} `json:"Type,omitempty"`
	Domainname interface{} `json:"DomainName,omitempty"`
	Lastupdateddate interface{} `json:"LastUpdatedDate,omitempty"`
	Message interface{} `json:"Message,omitempty"`
	Operationid interface{} `json:"OperationId,omitempty"`
	Status interface{} `json:"Status,omitempty"`
	Statusflag interface{} `json:"StatusFlag,omitempty"`
	Submitteddate interface{} `json:"SubmittedDate,omitempty"`
}

// ViewBillingResponse represents the ViewBillingResponse schema from the OpenAPI specification
type ViewBillingResponse struct {
	Billingrecords interface{} `json:"BillingRecords,omitempty"`
	Nextpagemarker interface{} `json:"NextPageMarker,omitempty"`
}

// Nameserver represents the Nameserver schema from the OpenAPI specification
type Nameserver struct {
	Glueips interface{} `json:"GlueIps,omitempty"`
	Name interface{} `json:"Name"`
}

// ListPricesRequest represents the ListPricesRequest schema from the OpenAPI specification
type ListPricesRequest struct {
	Marker interface{} `json:"Marker,omitempty"`
	Maxitems interface{} `json:"MaxItems,omitempty"`
	Tld interface{} `json:"Tld,omitempty"`
}

// GetDomainSuggestionsResponse represents the GetDomainSuggestionsResponse schema from the OpenAPI specification
type GetDomainSuggestionsResponse struct {
	Suggestionslist interface{} `json:"SuggestionsList,omitempty"`
}

// Tag represents the Tag schema from the OpenAPI specification
type Tag struct {
	Key interface{} `json:"Key,omitempty"`
	Value interface{} `json:"Value,omitempty"`
}

// UpdateDomainNameserversResponse represents the UpdateDomainNameserversResponse schema from the OpenAPI specification
type UpdateDomainNameserversResponse struct {
	Operationid interface{} `json:"OperationId,omitempty"`
}

// DnssecSigningAttributes represents the DnssecSigningAttributes schema from the OpenAPI specification
type DnssecSigningAttributes struct {
	Algorithm interface{} `json:"Algorithm,omitempty"`
	Flags interface{} `json:"Flags,omitempty"`
	Publickey interface{} `json:"PublicKey,omitempty"`
}

// GetContactReachabilityStatusRequest represents the GetContactReachabilityStatusRequest schema from the OpenAPI specification
type GetContactReachabilityStatusRequest struct {
	Domainname interface{} `json:"domainName,omitempty"`
}

// ContactDetail represents the ContactDetail schema from the OpenAPI specification
type ContactDetail struct {
	Fax interface{} `json:"Fax,omitempty"`
	Firstname interface{} `json:"FirstName,omitempty"`
	Addressline1 interface{} `json:"AddressLine1,omitempty"`
	Contacttype interface{} `json:"ContactType,omitempty"`
	Lastname interface{} `json:"LastName,omitempty"`
	Organizationname interface{} `json:"OrganizationName,omitempty"`
	City interface{} `json:"City,omitempty"`
	Email interface{} `json:"Email,omitempty"`
	Phonenumber interface{} `json:"PhoneNumber,omitempty"`
	State interface{} `json:"State,omitempty"`
	Zipcode interface{} `json:"ZipCode,omitempty"`
	Addressline2 interface{} `json:"AddressLine2,omitempty"`
	Countrycode interface{} `json:"CountryCode,omitempty"`
	Extraparams interface{} `json:"ExtraParams,omitempty"`
}

// RetrieveDomainAuthCodeRequest represents the RetrieveDomainAuthCodeRequest schema from the OpenAPI specification
type RetrieveDomainAuthCodeRequest struct {
	Domainname interface{} `json:"DomainName"`
}

// ViewBillingRequest represents the ViewBillingRequest schema from the OpenAPI specification
type ViewBillingRequest struct {
	End interface{} `json:"End,omitempty"`
	Marker interface{} `json:"Marker,omitempty"`
	Maxitems interface{} `json:"MaxItems,omitempty"`
	Start interface{} `json:"Start,omitempty"`
}

// RegisterDomainResponse represents the RegisterDomainResponse schema from the OpenAPI specification
type RegisterDomainResponse struct {
	Operationid interface{} `json:"OperationId,omitempty"`
}

// ListOperationsRequest represents the ListOperationsRequest schema from the OpenAPI specification
type ListOperationsRequest struct {
	Marker interface{} `json:"Marker,omitempty"`
	Maxitems interface{} `json:"MaxItems,omitempty"`
	Sortby interface{} `json:"SortBy,omitempty"`
	Sortorder interface{} `json:"SortOrder,omitempty"`
	Status interface{} `json:"Status,omitempty"`
	Submittedsince interface{} `json:"SubmittedSince,omitempty"`
	TypeField interface{} `json:"Type,omitempty"`
}

// AssociateDelegationSignerToDomainResponse represents the AssociateDelegationSignerToDomainResponse schema from the OpenAPI specification
type AssociateDelegationSignerToDomainResponse struct {
	Operationid interface{} `json:"OperationId,omitempty"`
}

// DisableDomainTransferLockRequest represents the DisableDomainTransferLockRequest schema from the OpenAPI specification
type DisableDomainTransferLockRequest struct {
	Domainname interface{} `json:"DomainName"`
}

// Consent represents the Consent schema from the OpenAPI specification
type Consent struct {
	Currency interface{} `json:"Currency"`
	Maxprice interface{} `json:"MaxPrice"`
}

// GetDomainDetailResponse represents the GetDomainDetailResponse schema from the OpenAPI specification
type GetDomainDetailResponse struct {
	Techprivacy interface{} `json:"TechPrivacy,omitempty"`
	Dnsseckeys interface{} `json:"DnssecKeys,omitempty"`
	Registrantcontact interface{} `json:"RegistrantContact,omitempty"`
	Adminprivacy interface{} `json:"AdminPrivacy,omitempty"`
	Registrydomainid interface{} `json:"RegistryDomainId,omitempty"`
	Reseller interface{} `json:"Reseller,omitempty"`
	Registrantprivacy interface{} `json:"RegistrantPrivacy,omitempty"`
	Abusecontactemail interface{} `json:"AbuseContactEmail,omitempty"`
	Autorenew interface{} `json:"AutoRenew,omitempty"`
	Creationdate interface{} `json:"CreationDate,omitempty"`
	Statuslist interface{} `json:"StatusList,omitempty"`
	Abusecontactphone interface{} `json:"AbuseContactPhone,omitempty"`
	Updateddate interface{} `json:"UpdatedDate,omitempty"`
	Admincontact interface{} `json:"AdminContact,omitempty"`
	Dnssec interface{} `json:"DnsSec,omitempty"`
	Registrarurl interface{} `json:"RegistrarUrl,omitempty"`
	Expirationdate interface{} `json:"ExpirationDate,omitempty"`
	Registrarname interface{} `json:"RegistrarName,omitempty"`
	Whoisserver interface{} `json:"WhoIsServer,omitempty"`
	Nameservers interface{} `json:"Nameservers,omitempty"`
	Domainname interface{} `json:"DomainName,omitempty"`
	Techcontact interface{} `json:"TechContact,omitempty"`
}

// GetDomainSuggestionsRequest represents the GetDomainSuggestionsRequest schema from the OpenAPI specification
type GetDomainSuggestionsRequest struct {
	Domainname interface{} `json:"DomainName"`
	Onlyavailable interface{} `json:"OnlyAvailable"`
	Suggestioncount interface{} `json:"SuggestionCount"`
}

// GetContactReachabilityStatusResponse represents the GetContactReachabilityStatusResponse schema from the OpenAPI specification
type GetContactReachabilityStatusResponse struct {
	Domainname interface{} `json:"domainName,omitempty"`
	Status interface{} `json:"status,omitempty"`
}

// UpdateDomainContactRequest represents the UpdateDomainContactRequest schema from the OpenAPI specification
type UpdateDomainContactRequest struct {
	Techcontact interface{} `json:"TechContact,omitempty"`
	Admincontact interface{} `json:"AdminContact,omitempty"`
	Consent interface{} `json:"Consent,omitempty"`
	Domainname interface{} `json:"DomainName"`
	Registrantcontact interface{} `json:"RegistrantContact,omitempty"`
}

// AcceptDomainTransferFromAnotherAwsAccountResponse represents the AcceptDomainTransferFromAnotherAwsAccountResponse schema from the OpenAPI specification
type AcceptDomainTransferFromAnotherAwsAccountResponse struct {
	Operationid interface{} `json:"OperationId,omitempty"`
}

// UpdateDomainContactPrivacyResponse represents the UpdateDomainContactPrivacyResponse schema from the OpenAPI specification
type UpdateDomainContactPrivacyResponse struct {
	Operationid interface{} `json:"OperationId,omitempty"`
}

// DisassociateDelegationSignerFromDomainRequest represents the DisassociateDelegationSignerFromDomainRequest schema from the OpenAPI specification
type DisassociateDelegationSignerFromDomainRequest struct {
	Domainname interface{} `json:"DomainName"`
	Id interface{} `json:"Id"`
}

// DeleteTagsForDomainRequest represents the DeleteTagsForDomainRequest schema from the OpenAPI specification
type DeleteTagsForDomainRequest struct {
	Tagstodelete interface{} `json:"TagsToDelete"`
	Domainname interface{} `json:"DomainName"`
}

// PriceWithCurrency represents the PriceWithCurrency schema from the OpenAPI specification
type PriceWithCurrency struct {
	Currency interface{} `json:"Currency"`
	Price interface{} `json:"Price"`
}

// ListPricesResponse represents the ListPricesResponse schema from the OpenAPI specification
type ListPricesResponse struct {
	Nextpagemarker interface{} `json:"NextPageMarker,omitempty"`
	Prices interface{} `json:"Prices,omitempty"`
}

// RejectDomainTransferFromAnotherAwsAccountRequest represents the RejectDomainTransferFromAnotherAwsAccountRequest schema from the OpenAPI specification
type RejectDomainTransferFromAnotherAwsAccountRequest struct {
	Domainname interface{} `json:"DomainName"`
}

// ListTagsForDomainRequest represents the ListTagsForDomainRequest schema from the OpenAPI specification
type ListTagsForDomainRequest struct {
	Domainname interface{} `json:"DomainName"`
}

// ResendContactReachabilityEmailRequest represents the ResendContactReachabilityEmailRequest schema from the OpenAPI specification
type ResendContactReachabilityEmailRequest struct {
	Domainname interface{} `json:"domainName,omitempty"`
}

// CancelDomainTransferToAnotherAwsAccountResponse represents the CancelDomainTransferToAnotherAwsAccountResponse schema from the OpenAPI specification
type CancelDomainTransferToAnotherAwsAccountResponse struct {
	Operationid interface{} `json:"OperationId,omitempty"`
}

// DnssecKey represents the DnssecKey schema from the OpenAPI specification
type DnssecKey struct {
	Algorithm interface{} `json:"Algorithm,omitempty"`
	Digest interface{} `json:"Digest,omitempty"`
	Digesttype interface{} `json:"DigestType,omitempty"`
	Flags interface{} `json:"Flags,omitempty"`
	Id interface{} `json:"Id,omitempty"`
	Keytag interface{} `json:"KeyTag,omitempty"`
	Publickey interface{} `json:"PublicKey,omitempty"`
}

// ResendOperationAuthorizationRequest represents the ResendOperationAuthorizationRequest schema from the OpenAPI specification
type ResendOperationAuthorizationRequest struct {
	Operationid interface{} `json:"OperationId"`
}

// SortCondition represents the SortCondition schema from the OpenAPI specification
type SortCondition struct {
	Sortorder interface{} `json:"SortOrder"`
	Name interface{} `json:"Name"`
}

// UpdateTagsForDomainRequest represents the UpdateTagsForDomainRequest schema from the OpenAPI specification
type UpdateTagsForDomainRequest struct {
	Tagstoupdate interface{} `json:"TagsToUpdate,omitempty"`
	Domainname interface{} `json:"DomainName"`
}

// DeleteDomainRequest represents the DeleteDomainRequest schema from the OpenAPI specification
type DeleteDomainRequest struct {
	Domainname interface{} `json:"DomainName"`
}

// DomainTransferability represents the DomainTransferability schema from the OpenAPI specification
type DomainTransferability struct {
	Transferable string `json:"Transferable,omitempty"` // <p>Whether the domain name can be transferred to Route 53.</p> <note> <p>You can transfer only domains that have a value of <code>TRANSFERABLE</code> or <code>Transferable</code>.</p> </note> <p>Valid values:</p> <dl> <dt>TRANSFERABLE</dt> <dd> <p>The domain name can be transferred to Route 53.</p> </dd> <dt>UNTRANSFERRABLE</dt> <dd> <p>The domain name can't be transferred to Route 53.</p> </dd> <dt>DONT_KNOW</dt> <dd> <p>Reserved for future use.</p> </dd> <dt>DOMAIN_IN_OWN_ACCOUNT</dt> <dd> <p>The domain already exists in the current Amazon Web Services account.</p> </dd> <dt>DOMAIN_IN_ANOTHER_ACCOUNT</dt> <dd> <p> the domain exists in another Amazon Web Services account.</p> </dd> <dt>PREMIUM_DOMAIN</dt> <dd> <p>Premium domain transfer is not supported.</p> </dd> </dl>
}

// CheckDomainTransferabilityResponse represents the CheckDomainTransferabilityResponse schema from the OpenAPI specification
type CheckDomainTransferabilityResponse struct {
	Transferability interface{} `json:"Transferability,omitempty"`
}
