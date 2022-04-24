// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ConsentRequest Contains information on an ongoing consent request.
//
// swagger:model consentRequest
type ConsentRequest struct {

	// ACR represents the Authentication AuthorizationContext Class Reference value for this authentication session. You can use it
	// to express that, for example, a user authenticated using two factor authentication.
	Acr string `json:"acr,omitempty"`

	// amr
	Amr StringSlicePipeDelimiter `json:"amr,omitempty"`

	// ID is the identifier ("authorization challenge") of the consent authorization request. It is used to
	// identify the session.
	// Required: true
	Challenge *string `json:"challenge"`

	// client
	Client *OAuth2Client `json:"client,omitempty"`

	// context
	Context JSONRawMessage `json:"context,omitempty"`

	// LoginChallenge is the login challenge this consent challenge belongs to. It can be used to associate
	// a login and consent request in the login & consent app.
	LoginChallenge string `json:"login_challenge,omitempty"`

	// LoginSessionID is the login session ID. If the user-agent reuses a login session (via cookie / remember flag)
	// this ID will remain the same. If the user-agent did not have an existing authentication session (e.g. remember is false)
	// this will be a new random value. This value is used as the "sid" parameter in the ID Token and in OIDC Front-/Back-
	// channel logout. It's value can generally be used to associate consecutive login requests by a certain user.
	LoginSessionID string `json:"login_session_id,omitempty"`

	// oidc context
	OidcContext *OpenIDConnectContext `json:"oidc_context,omitempty"`

	// RequestURL is the original OAuth 2.0 Authorization URL requested by the OAuth 2.0 client. It is the URL which
	// initiates the OAuth 2.0 Authorization Code or OAuth 2.0 Implicit flow. This URL is typically not needed, but
	// might come in handy if you want to deal with additional request parameters.
	RequestURL string `json:"request_url,omitempty"`

	// requested access token audience
	RequestedAccessTokenAudience StringSlicePipeDelimiter `json:"requested_access_token_audience,omitempty"`

	// requested scope
	RequestedScope StringSlicePipeDelimiter `json:"requested_scope,omitempty"`

	// Skip, if true, implies that the client has requested the same scopes from the same user previously.
	// If true, you must not ask the user to grant the requested scopes. You must however either allow or deny the
	// consent request using the usual API call.
	Skip bool `json:"skip,omitempty"`

	// Subject is the user ID of the end-user that authenticated. Now, that end user needs to grant or deny the scope
	// requested by the OAuth 2.0 client.
	Subject string `json:"subject,omitempty"`
}

// Validate validates this consent request
func (m *ConsentRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAmr(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateChallenge(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClient(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOidcContext(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRequestedAccessTokenAudience(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRequestedScope(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConsentRequest) validateAmr(formats strfmt.Registry) error {
	if swag.IsZero(m.Amr) { // not required
		return nil
	}

	if err := m.Amr.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("amr")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("amr")
		}
		return err
	}

	return nil
}

func (m *ConsentRequest) validateChallenge(formats strfmt.Registry) error {

	if err := validate.Required("challenge", "body", m.Challenge); err != nil {
		return err
	}

	return nil
}

func (m *ConsentRequest) validateClient(formats strfmt.Registry) error {
	if swag.IsZero(m.Client) { // not required
		return nil
	}

	if m.Client != nil {
		if err := m.Client.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("client")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("client")
			}
			return err
		}
	}

	return nil
}

func (m *ConsentRequest) validateOidcContext(formats strfmt.Registry) error {
	if swag.IsZero(m.OidcContext) { // not required
		return nil
	}

	if m.OidcContext != nil {
		if err := m.OidcContext.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("oidc_context")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("oidc_context")
			}
			return err
		}
	}

	return nil
}

func (m *ConsentRequest) validateRequestedAccessTokenAudience(formats strfmt.Registry) error {
	if swag.IsZero(m.RequestedAccessTokenAudience) { // not required
		return nil
	}

	if err := m.RequestedAccessTokenAudience.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("requested_access_token_audience")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("requested_access_token_audience")
		}
		return err
	}

	return nil
}

func (m *ConsentRequest) validateRequestedScope(formats strfmt.Registry) error {
	if swag.IsZero(m.RequestedScope) { // not required
		return nil
	}

	if err := m.RequestedScope.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("requested_scope")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("requested_scope")
		}
		return err
	}

	return nil
}

// ContextValidate validate this consent request based on the context it is used
func (m *ConsentRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAmr(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateClient(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOidcContext(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRequestedAccessTokenAudience(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRequestedScope(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConsentRequest) contextValidateAmr(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Amr.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("amr")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("amr")
		}
		return err
	}

	return nil
}

func (m *ConsentRequest) contextValidateClient(ctx context.Context, formats strfmt.Registry) error {

	if m.Client != nil {
		if err := m.Client.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("client")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("client")
			}
			return err
		}
	}

	return nil
}

func (m *ConsentRequest) contextValidateOidcContext(ctx context.Context, formats strfmt.Registry) error {

	if m.OidcContext != nil {
		if err := m.OidcContext.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("oidc_context")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("oidc_context")
			}
			return err
		}
	}

	return nil
}

func (m *ConsentRequest) contextValidateRequestedAccessTokenAudience(ctx context.Context, formats strfmt.Registry) error {

	if err := m.RequestedAccessTokenAudience.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("requested_access_token_audience")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("requested_access_token_audience")
		}
		return err
	}

	return nil
}

func (m *ConsentRequest) contextValidateRequestedScope(ctx context.Context, formats strfmt.Registry) error {

	if err := m.RequestedScope.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("requested_scope")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("requested_scope")
		}
		return err
	}

	return nil
}

// ContextValidate validate this consent request based on the context it is used
func (m *ConsentRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAmr(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateClient(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOidcContext(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRequestedAccessTokenAudience(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRequestedScope(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConsentRequest) contextValidateAmr(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Amr.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("amr")
		}
		return err
	}

	return nil
}

func (m *ConsentRequest) contextValidateClient(ctx context.Context, formats strfmt.Registry) error {

	if m.Client != nil {
		if err := m.Client.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("client")
			}
			return err
		}
	}

	return nil
}

func (m *ConsentRequest) contextValidateOidcContext(ctx context.Context, formats strfmt.Registry) error {

	if m.OidcContext != nil {
		if err := m.OidcContext.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("oidc_context")
			}
			return err
		}
	}

	return nil
}

func (m *ConsentRequest) contextValidateRequestedAccessTokenAudience(ctx context.Context, formats strfmt.Registry) error {

	if err := m.RequestedAccessTokenAudience.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("requested_access_token_audience")
		}
		return err
	}

	return nil
}

func (m *ConsentRequest) contextValidateRequestedScope(ctx context.Context, formats strfmt.Registry) error {

	if err := m.RequestedScope.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("requested_scope")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConsentRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConsentRequest) UnmarshalBinary(b []byte) error {
	var res ConsentRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
