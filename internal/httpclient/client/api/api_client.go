// Passcode generated by go-swagger; DO NOT EDIT.

package api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new api API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for api API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	Decisions(params *DecisionsParams) (*DecisionsOK, error)

	GetRule(params *GetRuleParams) (*GetRuleOK, error)

	GetWellKnownJSONWebKeys(params *GetWellKnownJSONWebKeysParams) (*GetWellKnownJSONWebKeysOK, error)

	ListRules(params *ListRulesParams) (*ListRulesOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  Decisions accesses control decision API

  > This endpoint works with all HTTP Methods (GET, POST, PUT, ...) and matches every path prefixed with /decision.

This endpoint mirrors the proxy capability of ORY Oathkeeper's proxy functionality but instead of forwarding the
request to the upstream server, returns 200 (request should be allowed), 401 (unauthorized), or 403 (forbidden)
status codes. This endpoint can be used to integrate with other API Proxies like Ambassador, Kong, Envoy, and many more.
*/
func (a *Client) Decisions(params *DecisionsParams) (*DecisionsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDecisionsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "decisions",
		Method:             "GET",
		PathPattern:        "/decisions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DecisionsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DecisionsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for decisions: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetRule retrieves a rule

  Use this method to retrieve a rule from the storage. If it does not exist you will receive a 404 error.
*/
func (a *Client) GetRule(params *GetRuleParams) (*GetRuleOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetRuleParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getRule",
		Method:             "GET",
		PathPattern:        "/rules/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetRuleReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetRuleOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getRule: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetWellKnownJSONWebKeys lists cryptographic keys

  This endpoint returns cryptographic keys that are required to, for example, verify signatures of ID Tokens.
*/
func (a *Client) GetWellKnownJSONWebKeys(params *GetWellKnownJSONWebKeysParams) (*GetWellKnownJSONWebKeysOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetWellKnownJSONWebKeysParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getWellKnownJSONWebKeys",
		Method:             "GET",
		PathPattern:        "/.well-known/jwks.json",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetWellKnownJSONWebKeysReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetWellKnownJSONWebKeysOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getWellKnownJSONWebKeys: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListRules lists all rules

  This method returns an array of all rules that are stored in the backend. This is useful if you want to get a full
view of what rules you have currently in place.
*/
func (a *Client) ListRules(params *ListRulesParams) (*ListRulesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListRulesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listRules",
		Method:             "GET",
		PathPattern:        "/rules",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ListRulesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListRulesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listRules: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
