package session

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type URL struct {
	URL *url.URL
	BaseURLForParameterizedURLS *url.URL
	PostParameterizedEndpoint string
	BaseEndpointForParameterizedEndpoints string
	Method string
}

type URLS struct {
	OAuth2Tokens struct {
		V1 struct {
			GenerateTokens *URL
			RevokeToken *URL
			GetRateLimit *URL
		}
		V2 struct {
			GenerateTokens *URL
			RevokeToken *URL
			GetRateLimit *URL
		}
	}
	ApiAuthorization struct {
		V2 struct {
			ListAuthorizationServers *URL
			GetAuthorizationServer *URL
			CreateAuthorizationServer *URL
			UpdateAuthorizationServer *URL
			DeleteAuthorizationServer *URL
			ListAccessTokenClaims *URL
			AddAccessTokenClaim *URL
			UpdateAccessTokenClaim *URL
			DeleteAccessTokenClaim *URL
			ListScopes *URL
			AddScope *URL
			UpdateScope *URL
			DeleteScope *URL
			GetClientApps *URL
			AddClientApp *URL
			UpdateClientApp *URL
			RemoveClientApp *URL
		}
	}
	Apps struct {
		V1 struct {
			GetApps *URL
		}
		V2 struct {
			ListApps *URL
			GetApp *URL
			CreateApp *URL
			UpdateApp *URL
			DeleteAppParameter *URL
			DeleteApp *URL
			ListAppUsers *URL
		}
	}
	Connectors struct {
		V2 struct {
			ListConnectors *URL
		}
	}
	SAMLAssertions struct {
		V1 struct {
			GenerateSAMLAssertion *URL
			VerifyFactor *URL
		}
		V2 struct {
			GenerateSAMLAssertion *URL
			VerifyFactor *URL
		}
	}
	UserMappings struct {
		V2 struct {
			ListMappings *URL
			GetMapping *URL
			CreateMapping *URL
			UpdateMapping *URL
			DeleteMapping *URL
			ListConditions *URL
			ListConditionOperators *URL
			ListConditionValues *URL
			ListActions *URL
			ListActionValues *URL
			BulkSort *URL
		}
	}
	VigilanceAI struct {
		V2 struct {
			TrackEvent *URL
			GetRiskScore *URL
			CreateRule *URL
			ListRules *URL
			GetRule *URL
			UpdateRule *URL
			DeleteRule *URL
			GetScoreInsights *URL
		}
	}
	Events struct {
		V1 struct {
			GetEventTypes *URL
			GetEvents *URL
			GetEvent *URL
		}
	}
	Groups struct {
		V1 struct {
			GetGroups *URL
			GetGroup *URL
		}
	}
	LoginPages struct {
		V1 struct {
			CreateSessionLoginToken *URL
			VerifyFactor *URL
			CreateSession *URL
		}
	}
	MultiFactorAuthentication struct {
		V1 struct {
			GetAvailableFactors *URL
			EnrollFactor *URL
			GetEnrolledFactors *URL
			ActivateFactor *URL
			VerifyFactor *URL
			RemoveFactor *URL
			GenerateMFAToken *URL
		}
	}
	Privileges struct {
		V1 struct {
			ListPrivileges *URL
			CreatePrivilege *URL
			UpdatePrivilege *URL
			GetPrivilege *URL
			DeletePrivilege *URL
			GetAssignedRoles *URL
			AssignRoles *URL
			RemoveRole *URL
			GetAssignedUsers *URL
			AssignUsers *URL
			RemoveUser *URL
		}
	}
	Roles struct {
		V1 struct {
			GetRoles *URL
			GetRole *URL
		}
	}
	Users struct {
		V1 struct {
			GetUsers *URL
			GetUser *URL
			GetApps *URL
			GetRoles *URL
			GetCustomAttributes *URL
			CreateUser *URL
			UpdateUser *URL
			AssignRole *URL
			RemoveRole *URL
			SetCleartextPassword *URL
			SetSHA256Password *URL
			SetCustomAttribute *URL
			SetUserState *URL
			LogUserOut *URL
			LockUserAccount *URL
			DeleteUser *URL
		}
	}
	InviteLinks struct {
		V1 struct {
			GenerateInviteLink *URL
			SendInviteLink *URL
		}
	}
	EmbedApps struct {
		V1 struct {
			GetAppsToEmbedForAUser *URL
		}
	}
}

type ClientID string

func (C ClientID) Validate() (ClientID, error) {
	switch strings.ToLower(string(C)) {
	case "":
		return "", fmt.Errorf("invalid value, ClientID must not be blank")
	default:
		return C, nil
	}
}

type ClientSecret string

func (C ClientSecret) Validate() (ClientSecret, error) {
	switch strings.ToLower(string(C)) {
	case "":
		return "", fmt.Errorf("invalid value, ClientSecret must not be blank")
	default:
		return C, nil
	}
}

type Region string

func (R Region) Validate() (Region, error) {
	switch strings.ToLower(string(R)) {
	case "us":
		return R, nil
	case "eu":
		return R, nil
	default:
		return "", fmt.Errorf("invalid value: Region must be one of US or EU")
	}
}

type Session struct {
	ClientID ClientID
	ClientSecret ClientSecret
	Region Region
	SubDomain string
	URLS *URLS
}

func New(id ClientID, secret ClientSecret, region Region, subDomain string) (*Session, error) {
	S := &Session{}

	validatedID, err := id.Validate()
	if err != nil {
		return nil, err
	}

	validatedSecret, err := secret.Validate()
	if err != nil {
		return nil, err
	}

	validatedRegion, err := region.Validate()
	if err != nil {
		return nil, err
	}

	if subDomain != "" {
		S.SubDomain = subDomain
	} else {
		return nil, fmt.Errorf("subDomain must not be blank")
	}

	S.ClientID = validatedID
	S.ClientSecret = validatedSecret
	S.Region = validatedRegion
	S.URLS = S.generateURLS()

	return S, nil
}

func (S *Session) generateURLS() *URLS {
	U := &URLS{}

	U.OAuth2Tokens.V1.GenerateTokens = &URL{
		URL: S.mustGenerateAPIDomainURL("https://api.%s.onelogin.com/auth/oauth2/v2/token"),
		Method: http.MethodPost,
	}

	U.OAuth2Tokens.V1.RevokeToken = &URL{
		URL: S.mustGenerateAPIDomainURL("https://api.%s.onelogin.com/auth/oauth2/revoke"),
		Method: http.MethodPost,
	}

	U.OAuth2Tokens.V1.RevokeToken = &URL{
		URL: S.mustGenerateAPIDomainURL("https://api.%s.onelogin.com/auth/rate_limit"),
		Method: http.MethodGet,
	}

	U.OAuth2Tokens.V2.GenerateTokens = &URL{
		URL: S.mustGenerateAPIDomainURL("https://api.%s.onelogin.com/auth/oauth2/v2/token"),
	}

	U.OAuth2Tokens.V2.RevokeToken = &URL{
		URL: S.mustGenerateAPIDomainURL("https://api.%s.onelogin.com/auth/oauth2/revoke"),
		Method: http.MethodPost,
	}

	U.OAuth2Tokens.V2.RevokeToken = &URL{
		URL: S.mustGenerateAPIDomainURL("https://api.%s.onelogin.com/auth/rate_limit"),
		Method: http.MethodGet,
	}

	U.ApiAuthorization.V2.ListAuthorizationServers = &URL{
		URL: S.mustGenerateAPIDomainURL("https://api.%s.onelogin.com/api/2/api_authorizations"),
		Method: http.MethodGet,
	}

	U.ApiAuthorization.V2.GetAuthorizationServer = &URL{
		BaseURLForParameterizedURLS: S.mustGenerateAPIDomainURL("https://api.%s.onelogin.com/api/2/api_authorizations"),
		Method: http.MethodGet,
	}

	U.ApiAuthorization.V2.CreateAuthorizationServer = &URL{
		URL: S.mustGenerateAPIDomainURL("https://api.%s.onelogin.com/api/2/api_authorizations"),
		Method: http.MethodPost,
	}

	U.ApiAuthorization.V2.UpdateAuthorizationServer = &URL{
		BaseURLForParameterizedURLS: S.mustGenerateAPIDomainURL("https://api.%s.onelogin.com/api/2/api_authorizations"),
		Method: http.MethodPut,
	}

	U.ApiAuthorization.V2.DeleteAuthorizationServer = &URL{
		BaseURLForParameterizedURLS: S.mustGenerateAPIDomainURL("https://api.%s.onelogin.com/api/2/api_authorizations"),
		Method: http.MethodDelete,
	}

	U.ApiAuthorization.V2.ListAccessTokenClaims = &URL{
		BaseURLForParameterizedURLS: S.mustGenerateAPIDomainURL("https://api.%s.onelogin.com/api/2/api_authorizations"),
		PostParameterizedEndpoint: "claims",
		Method: http.MethodGet,
	}

	U.ApiAuthorization.V2.AddAccessTokenClaim = &URL{
		BaseURLForParameterizedURLS: S.mustGenerateAPIDomainURL("https://api.%s.onelogin.com/api/2/api_authorizations"),
		PostParameterizedEndpoint: "claims",
		Method: http.MethodPost,
	}

	U.ApiAuthorization.V2.UpdateAccessTokenClaim = &URL{
		BaseURLForParameterizedURLS: S.mustGenerateAPIDomainURL("https://api.%s.onelogin.com/api/2/api_authorizations"),
		BaseEndpointForParameterizedEndpoints: "claims",
		Method: http.MethodPut,
	}

	U.ApiAuthorization.V2.DeleteAccessTokenClaim = &URL{
		BaseURLForParameterizedURLS: S.mustGenerateAPIDomainURL("https://api.%s.onelogin.com/api/2/api_authorizations"),
		BaseEndpointForParameterizedEndpoints: "claims",
		Method: http.MethodDelete,
	}

	U.ApiAuthorization.V2.ListScopes = &URL{
		BaseURLForParameterizedURLS: S.mustGenerateAPIDomainURL("https://api.%s.onelogin.com/api/2/api_authorizations"),
		PostParameterizedEndpoint: "scopes",
		Method: http.MethodGet,
	}

	U.ApiAuthorization.V2.AddScope = &URL{
		BaseURLForParameterizedURLS: S.mustGenerateAPIDomainURL("https://api.%s.onelogin.com/api/2/api_authorizations"),
		PostParameterizedEndpoint: "scopes",
		Method: http.MethodPost,
	}

	U.ApiAuthorization.V2.UpdateScope = &URL{
		BaseURLForParameterizedURLS: S.mustGenerateAPIDomainURL("https://api.%s.onelogin.com/api/2/api_authorizations"),
		BaseEndpointForParameterizedEndpoints: "scopes",
		Method: http.MethodPut,
	}

	U.ApiAuthorization.V2.DeleteScope = &URL{
		BaseURLForParameterizedURLS: S.mustGenerateAPIDomainURL("https://api.%s.onelogin.com/api/2/api_authorizations"),
		BaseEndpointForParameterizedEndpoints: "scopes",
		Method: http.MethodDelete,
	}

	U.ApiAuthorization.V2.GetClientApps = &URL{
		BaseURLForParameterizedURLS: S.mustGenerateAPIDomainURL("https://api.%s.onelogin.com/api/2/api_authorizations"),
		PostParameterizedEndpoint: "clients",
		Method: http.MethodGet,
	}

	U.ApiAuthorization.V2.AddClientApp = &URL{
		BaseURLForParameterizedURLS: S.mustGenerateAPIDomainURL("https://api.%s.onelogin.com/api/2/api_authorizations"),
		PostParameterizedEndpoint: "clients",
		Method: http.MethodPost,
	}

	U.ApiAuthorization.V2.UpdateClientApp = &URL{
		BaseURLForParameterizedURLS: S.mustGenerateAPIDomainURL("https://api.%s.onelogin.com/api/2/api_authorizations"),
		BaseEndpointForParameterizedEndpoints: "clients",
		Method: http.MethodPut,
	}

	U.ApiAuthorization.V2.RemoveClientApp = &URL{
		BaseURLForParameterizedURLS: S.mustGenerateAPIDomainURL("https://api.%s.onelogin.com/api/2/api_authorizations"),
		BaseEndpointForParameterizedEndpoints: "clients",
		Method: http.MethodDelete,
	}

	U.Apps.V1.GetApps = &URL{
		URL: S.mustGenerateAPIDomainURL("https://api.%s.onelogin.com/api/2/apps"),
		Method: http.MethodGet,
	}

	U.Apps.V2.ListApps = &URL{
		URL: S.mustGenerateAPIDomainURL("https://api.%s.onelogin.com/api/2/apps"),
		Method: http.MethodGet,
	}

	U.Apps.V2.GetApp = &URL{
		BaseURLForParameterizedURLS: S.mustGenerateAPIDomainURL("https://api.%s.onelogin.com/api/2/apps"),
		Method: http.MethodGet,
	}

	U.Apps.V2.CreateApp = &URL{
		URL: S.mustGenerateAPIDomainURL("https://api.%s.onelogin.com/api/2/apps"),
		Method: http.MethodPost,
	}

	U.Apps.V2.UpdateApp = &URL{
		BaseURLForParameterizedURLS: S.mustGenerateAPIDomainURL("https://api.%s.onelogin.com/api/2/apps"),
		Method: http.MethodPut,
	}

	U.Apps.V2.DeleteAppParameter = &URL{
		BaseURLForParameterizedURLS: S.mustGenerateAPIDomainURL("https://api.%s.onelogin.com/api/2/apps"),
		BaseEndpointForParameterizedEndpoints: "parameters",
		Method: http.MethodDelete,
	}

	U.Apps.V2.DeleteApp = &URL{
		BaseURLForParameterizedURLS: S.mustGenerateAPIDomainURL("https://api.%s.onelogin.com/api/2/apps"),
		Method: http.MethodDelete,
	}

	U.Apps.V2.ListAppUsers = &URL{
		BaseURLForParameterizedURLS: S.mustGenerateAPIDomainURL("https://api.%s.onelogin.com/api/2/apps"),
		PostParameterizedEndpoint: "users",
		Method: http.MethodGet,
	}

	U.Connectors.V2.ListConnectors = &URL{
		URL: S.mustGenerateAPIDomainURL("https://api.%s.onelogin.com/api/2/connectors"),
	}

	U.SAMLAssertions.V1.GenerateSAMLAssertion = &URL{
		URL: S.mustGenerateAPIDomainURL("https://api.%s.onelogin.com/api/1/saml_assertion"),
		Method: http.MethodPost,
	}

	U.SAMLAssertions.V1.VerifyFactor = &URL{
		URL: S.mustGenerateAPIDomainURL("https://api.%s.onelogin.com/api/1/saml_assertion/verify_factor"),
		Method: http.MethodPost,
	}

	U.SAMLAssertions.V2.GenerateSAMLAssertion = &URL{
		URL: S.mustGenerateAPIDomainURL("https://api.%s.onelogin.com/api/2/saml_assertion"),
		Method: http.MethodPost,
	}

	U.SAMLAssertions.V2.VerifyFactor = &URL{
		URL: S.mustGenerateAPIDomainURL("https://api.%s.onelogin.com/api/2/saml_assertion/verify_factor"),
		Method: http.MethodPost,
	}

	U.UserMappings.V2.ListMappings = &URL{
		URL: S.mustGenerateAPIDomainURL("https://api.%s.onelogin.com/api/2/mappings"),
		Method: http.MethodGet,
	}

	U.UserMappings.V2.GetMapping = &URL{
		BaseURLForParameterizedURLS: S.mustGenerateAPIDomainURL("https://api.%s.onelogin.com/api/2/mappings"),
		Method: http.MethodGet,
	}

	U.UserMappings.V2.CreateMapping = &URL{
		URL: S.mustGenerateAPIDomainURL("https://api.%s.onelogin.com/api/2/mappings"),
		Method: http.MethodPost,
	}

	U.UserMappings.V2.UpdateMapping = &URL{
		BaseURLForParameterizedURLS: S.mustGenerateAPIDomainURL("https://api.%s.onelogin.com/api/2/mappings"),
		Method: http.MethodPut,
	}

	U.UserMappings.V2.DeleteMapping = &URL{
		BaseURLForParameterizedURLS: S.mustGenerateAPIDomainURL("https://api.%s.onelogin.com/api/2/mappings"),
		Method: http.MethodDelete,
	}

	U.UserMappings.V2.ListConditions = &URL{
		URL: S.mustGenerateAPIDomainURL("https://api.%s.onelogin.com/api/2/mappings/conditions"),
		Method: http.MethodGet,
	}

	U.UserMappings.V2.ListConditionOperators = &URL{
		BaseURLForParameterizedURLS: S.mustGenerateAPIDomainURL("https://api.%s.onelogin.com/api/2/mappings"),
		PostParameterizedEndpoint: "operators",
		Method: http.MethodGet,
	}

	U.UserMappings.V2.ListConditionValues = &URL{
		BaseURLForParameterizedURLS: S.mustGenerateAPIDomainURL("https://api.%s.onelogin.com/api/2/mappings"),
		PostParameterizedEndpoint: "values",
		Method: http.MethodGet,
	}

	U.UserMappings.V2.ListActions = &URL{
		URL: S.mustGenerateAPIDomainURL("https://api.%s.onelogin.com/api/2/mappings/actions"),
		Method: http.MethodGet,
	}

	U.UserMappings.V2.ListActionValues = &URL{
		BaseURLForParameterizedURLS: S.mustGenerateAPIDomainURL("https://api.%s.onelogin.com/api/2/mappings"),
		PostParameterizedEndpoint: "values",
		Method: http.MethodGet,
	}

	U.UserMappings.V2.BulkSort = &URL{
		URL: S.mustGenerateAPIDomainURL("https://api.%s.onelogin.com/api/2/mappings/sort"),
		Method: http.MethodPut,
	}

	U.VigilanceAI.V2.TrackEvent = &URL{
		URL: S.mustGenerateAPIDomainURL("https://api.%s.onelogin.com/api/2/risk/events"),
		Method: http.MethodPost,
	}

	U.VigilanceAI.V2.GetRiskScore = &URL{
		URL: S.mustGenerateAPIDomainURL("https://api.%s.onelogin.com/api/2/risk/verify"),
		Method: http.MethodPost,
	}

	U.VigilanceAI.V2.CreateRule = &URL{
		URL: S.mustGenerateAPIDomainURL("https://api.%s.onelogin.com/api/2/risk/rules"),
		Method: http.MethodPost,
	}

	U.VigilanceAI.V2.ListRules = &URL{
		URL: S.mustGenerateAPIDomainURL("https://api.%s.onelogin.com/api/2/risk/rules"),
		Method: http.MethodGet,
	}

	U.VigilanceAI.V2.GetRule = &URL{
		BaseURLForParameterizedURLS: S.mustGenerateAPIDomainURL("https://api.%s.onelogin.com/api/2/risk/rules"),
		Method: http.MethodGet,
	}

	U.VigilanceAI.V2.UpdateRule = &URL{
		BaseURLForParameterizedURLS: S.mustGenerateAPIDomainURL("https://api.%s.onelogin.com/api/2/risk/rules"),
		Method: http.MethodPut,
	}

	U.VigilanceAI.V2.DeleteRule = &URL{
		BaseURLForParameterizedURLS: S.mustGenerateAPIDomainURL("https://api.%s.onelogin.com/api/2/risk/rules"),
		Method: http.MethodDelete,
	}

	U.VigilanceAI.V2.GetScoreInsights = &URL{
		URL: S.mustGenerateAPIDomainURL("https://api.%s.onelogin.com/api/2/risk/scores"),
		Method: http.MethodGet,
	}

	U.Events.V1.GetEventTypes = &URL{
		URL: S.mustGenerateAPIDomainURL("https://api.%s.onelogin.com/api/1/events/types"),
		Method: http.MethodGet,
	}

	U.Events.V1.GetEvents = &URL{
		URL: S.mustGenerateAPIDomainURL("https://api.%s.onelogin.com/api/1/events"),
		Method: http.MethodGet,
	}

	U.Events.V1.GetEvent = &URL{
		BaseURLForParameterizedURLS: S.mustGenerateAPIDomainURL("https://api.%s.onelogin.com/api/1/events/types"),
		Method: http.MethodGet,
	}

	U.Groups.V1.GetGroups = &URL{
		URL: S.mustGenerateAPIDomainURL("https://api.%s.onelogin.com/api/1/groups"),
		Method: http.MethodGet,
	}

	U.Groups.V1.GetGroup = &URL{
		BaseURLForParameterizedURLS: S.mustGenerateAPIDomainURL("https://api.%s.onelogin.com/api/1/groups"),
		Method: http.MethodGet,
	}

	U.Privileges.V1.ListPrivileges = &URL{
		URL: S.mustGenerateAPIDomainURL("https://api.%s.onelogin.com/api/1/privileges"),
		Method: http.MethodGet,
	}

	U.Privileges.V1.CreatePrivilege = &URL{
		URL: S.mustGenerateAPIDomainURL("https://api.%s.onelogin.com/api/1/privileges"),
		Method: http.MethodPost,
	}

	U.Privileges.V1.UpdatePrivilege = &URL{
		BaseURLForParameterizedURLS: S.mustGenerateAPIDomainURL("https://api.%s.onelogin.com/api/1/privileges"),
		Method: http.MethodPut,
	}

	U.Privileges.V1.GetPrivilege = &URL{
		BaseURLForParameterizedURLS: S.mustGenerateAPIDomainURL("https://api.%s.onelogin.com/api/1/privileges"),
		Method: http.MethodGet,
	}

	U.Privileges.V1.DeletePrivilege = &URL{
		BaseURLForParameterizedURLS: S.mustGenerateAPIDomainURL("https://api.%s.onelogin.com/api/1/privileges"),
		Method: http.MethodDelete,
	}

	U.Privileges.V1.GetAssignedRoles = &URL{
		BaseURLForParameterizedURLS: S.mustGenerateAPIDomainURL("https://api.%s.onelogin.com/api/1/privileges"),
		PostParameterizedEndpoint: "roles",
		Method: http.MethodGet,
	}

	U.Privileges.V1.AssignRoles = &URL{
		BaseURLForParameterizedURLS: S.mustGenerateAPIDomainURL("https://api.%s.onelogin.com/api/1/privileges"),
		PostParameterizedEndpoint: "roles",
		Method: http.MethodPost,
	}

	U.Privileges.V1.RemoveRole = &URL{
		BaseURLForParameterizedURLS: S.mustGenerateAPIDomainURL("https://api.%s.onelogin.com/api/1/privileges"),
		BaseEndpointForParameterizedEndpoints: "roles",
		Method: http.MethodDelete,
	}

	U.Privileges.V1.GetAssignedUsers = &URL{
		BaseURLForParameterizedURLS: S.mustGenerateAPIDomainURL("https://api.%s.onelogin.com/api/1/privileges"),
		PostParameterizedEndpoint: "users",
		Method: http.MethodGet,
	}

	U.Privileges.V1.AssignUsers = &URL{
		BaseURLForParameterizedURLS: S.mustGenerateAPIDomainURL("https://api.%s.onelogin.com/api/1/privileges"),
		PostParameterizedEndpoint: "users",
		Method: http.MethodPost,
	}

	U.Privileges.V1.RemoveUser = &URL{
		BaseURLForParameterizedURLS: S.mustGenerateAPIDomainURL("https://api.%s.onelogin.com/api/1/privileges"),
		BaseEndpointForParameterizedEndpoints: "users",
		Method: http.MethodDelete,
	}

	U.Roles.V1.GetRoles = &URL{
		URL: S.mustGenerateAPIDomainURL("https://api.%s.onelogin.com/api/1/roles"),
		Method: http.MethodGet,
	}

	U.Roles.V1.GetRole = &URL{
		BaseURLForParameterizedURLS: S.mustGenerateAPIDomainURL("https://api.%s.onelogin.com/api/1/roles"),
		Method: http.MethodGet,
	}

	U.Users.V1.GetUsers = &URL{
		URL: S.mustGenerateAPIDomainURL("https://api.%s.onelogin.com/api/1/users"),
		Method: http.MethodGet,
	}

	U.Users.V1.GetUser = &URL{
		BaseURLForParameterizedURLS: S.mustGenerateAPIDomainURL("https://api.%s.onelogin.com/api/1/users"),
		Method: http.MethodGet,
	}

	U.Users.V1.GetApps = &URL{
		BaseURLForParameterizedURLS: S.mustGenerateAPIDomainURL("https://api.%s.onelogin.com/api/1/users"),
		PostParameterizedEndpoint: "apps",
		Method: http.MethodGet,
	}

	U.Users.V1.GetRoles = &URL{
		BaseURLForParameterizedURLS: S.mustGenerateAPIDomainURL("https://api.%s.onelogin.com/api/1/users"),
		PostParameterizedEndpoint: "roles",
		Method: http.MethodGet,
	}

	U.Users.V1.GetCustomAttributes = &URL{
		URL: S.mustGenerateAPIDomainURL("https://api.%s.onelogin.com/api/1/users/custom_attributes"),
		Method: http.MethodGet,
	}

	U.Users.V1.CreateUser = &URL{
		URL: S.mustGenerateAPIDomainURL("https://api.%s.onelogin.com/api/1/users"),
		Method: http.MethodPost,
	}

	U.Users.V1.UpdateUser = &URL{
		BaseURLForParameterizedURLS: S.mustGenerateAPIDomainURL("https://api.%s.onelogin.com/api/1/users"),
		Method: http.MethodPut,
	}

	U.Users.V1.AssignRole = &URL{
		BaseURLForParameterizedURLS: S.mustGenerateAPIDomainURL("https://api.%s.onelogin.com/api/1/users"),
		PostParameterizedEndpoint: "add_roles",
		Method: http.MethodPut,
	}

	U.Users.V1.RemoveRole = &URL{
		BaseURLForParameterizedURLS: S.mustGenerateAPIDomainURL("https://api.%s.onelogin.com/api/1/users"),
		PostParameterizedEndpoint: "remove_roles",
		Method: http.MethodPut,
	}

	U.Users.V1.SetCleartextPassword = &URL{
		BaseURLForParameterizedURLS: S.mustGenerateAPIDomainURL("https://api.%s.onelogin.com/api/1/users/set_password_clear_text"),
		Method: http.MethodPut,
	}

	U.Users.V1.SetSHA256Password = &URL{
		BaseURLForParameterizedURLS: S.mustGenerateAPIDomainURL("https://api.%s.onelogin.com/api/1/users/set_password_using_salt"),
		Method: http.MethodPut,
	}

	U.Users.V1.SetCustomAttribute = &URL{
		BaseURLForParameterizedURLS: S.mustGenerateAPIDomainURL("https://api.%s.onelogin.com/api/1/users"),
		PostParameterizedEndpoint: "set_custom_attributes",
		Method: http.MethodPut,
	}

	U.Users.V1.SetUserState = &URL{
		BaseURLForParameterizedURLS: S.mustGenerateAPIDomainURL("https://api.%s.onelogin.com/api/1/users"),
		PostParameterizedEndpoint: "set_state",
		Method: http.MethodPut,
	}

	U.Users.V1.LogUserOut = &URL{
		BaseURLForParameterizedURLS: S.mustGenerateAPIDomainURL("https://api.%s.onelogin.com/api/1/users"),
		PostParameterizedEndpoint: "logout",
		Method: http.MethodPut,
	}

	U.Users.V1.LockUserAccount = &URL{
		BaseURLForParameterizedURLS: S.mustGenerateAPIDomainURL("https://api.%s.onelogin.com/api/1/users"),
		PostParameterizedEndpoint: "lock_user",
		Method: http.MethodPut,
	}

	U.Users.V1.DeleteUser = &URL{
		BaseURLForParameterizedURLS: S.mustGenerateAPIDomainURL("https://api.%s.onelogin.com/api/1/users"),
		Method: http.MethodDelete,
	}

	U.InviteLinks.V1.GenerateInviteLink = &URL{
		URL: S.mustGenerateAPIDomainURL("https://api.%s.onelogin.com/api/1/invites/get_invite_link"),
		Method: http.MethodPost,
	}

	U.InviteLinks.V1.SendInviteLink = &URL{
		URL: S.mustGenerateAPIDomainURL("https://api.%s.onelogin.com/api/1/invites/send_invite_link"),
		Method: http.MethodPost,
	}

	U.EmbedApps.V1.GetAppsToEmbedForAUser = &URL{
		URL: S.mustGenerateSubDomainURL("https://%s.onelogin.com/client/apps/embed2"),
		Method: http.MethodGet,
	}

	return U
}

func (S *Session) mustGenerateAPIDomainURL(format string) *url.URL {
	u, err := url.Parse(fmt.Sprintf(format, S.Region))
	if err != nil {
		panic(err)
	}
	return u
}

func (S *Session) mustGenerateSubDomainURL(format string) *url.URL {
	u, err := url.Parse(fmt.Sprintf(format, S.SubDomain))
	if err != nil {
		panic(err)
	}
	return u
}