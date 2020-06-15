package cmd

import (
	"fmt"

	"github.com/skckadiyala/apimanager/apimgr"
	"github.com/spf13/viper"
)

func getSecurityProfilePassThrough() []apimgr.SecurityProfile {
	securityProfile := make([]apimgr.SecurityProfile, 1)
	device := make([]apimgr.SecurityDevice, 1)

	props := map[string]string{
		"subjectIdFieldName":         "Pass Through",
		"removeCredentialsOnSuccess": "true",
	}

	device[0].Name = "Pass Through"
	device[0].Type = "passThrough"
	device[0].Order = 1
	device[0].Properties = props

	securityProfile[0].Devices = device
	securityProfile[0].IsDefault = true
	securityProfile[0].Name = "_default"

	return securityProfile
}

func getSecurityProfileAPIKey() []apimgr.SecurityProfile {
	securityProfile := make([]apimgr.SecurityProfile, 1)
	device := make([]apimgr.SecurityDevice, 1)

	props := map[string]string{
		"apiKeyFieldName":            "KeyId",
		"removeCredentialsOnSuccess": "true",
		"takeFrom":                   "HEADER",
		"typeDisplayName":            "API Key only",
	}

	device[0].Name = "API Key"
	device[0].Type = "apiKey"
	device[0].Order = 1
	device[0].Properties = props

	securityProfile[0].Devices = device
	securityProfile[0].IsDefault = true
	securityProfile[0].Name = "_default"

	return securityProfile

}

func getSecurityProfileHTTPBasic() []apimgr.SecurityProfile {
	securityProfile := make([]apimgr.SecurityProfile, 1)
	device := make([]apimgr.SecurityDevice, 1)

	props := map[string]string{
		"realm":                      "localhost",
		"removeCredentialsOnSuccess": "true",
	}

	device[0].Name = "HTTP Basic"
	device[0].Type = "basic"
	device[0].Order = 1
	device[0].Properties = props

	securityProfile[0].Devices = device
	securityProfile[0].IsDefault = true
	securityProfile[0].Name = "_default"

	return securityProfile
}

func getSecurityProfileOAuth() []apimgr.SecurityProfile {

	apiHost := viper.GetString("apimanagerhost")
	securityProfile := make([]apimgr.SecurityProfile, 1)
	device := make([]apimgr.SecurityDevice, 1)

	endPointURL := fmt.Sprintf("http://%s:%s/api/oauth/authorize", apiHost, "8089")
	tokenEndpointURL := fmt.Sprintf("http://%s:%s/api/oauth/token", apiHost, "8089")

	props := map[string]string{
		"tokenStore":                              "<key type='OAuth2StoresGroup'><id field='name' value='OAuth2 Stores'/><key type='AccessTokenStoreGroup'><id field='name' value='Access Token Stores'/><key type='AccessTokenPersist'><id field='name' value='OAuth Access Token Store'/></key></key></key>",
		"accessTokenLocation":                     "HEADER",
		"authorizationHeaderPrefix":               "Bearer",
		"accessTokenLocationQueryString":          "",
		"scopesMustMatch":                         "Any",
		"scopes":                                  "create-orders",
		"removeCredentialsOnSuccess":              "true",
		"implicitGrantEnabled":                    "true",
		"implicitGrantLoginEndpointUrl":           endPointURL,
		"implicitGrantLoginTokenName":             "access_token",
		"authCodeGrantTypeEnabled":                "true",
		"authCodeGrantTypeRequestEndpointUrl":     endPointURL,
		"authCodeGrantTypeRequestClientIdName":    "client_id",
		"authCodeGrantTypeRequestSecretName":      "client_secret",
		"authCodeGrantTypeTokenEndpointUrl":       tokenEndpointURL,
		"authCodeGrantTypeTokenEndpointTokenName": "access_code",
	}

	device[0].Name = "OAuth"
	device[0].Type = "oauth"
	device[0].Order = 1
	device[0].Properties = props

	securityProfile[0].Devices = device
	securityProfile[0].IsDefault = true
	securityProfile[0].Name = "_default"

	return securityProfile
}
