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

// Wallet Wallet object (V0).
//
// swagger:model Wallet
type Wallet struct {

	// wallet's address.
	// Required: true
	Address string `json:"address"`

	// key pair
	// Required: true
	KeyPair WalletKeyPair `json:"keyPair"`

	// wallet's nickname.
	// Required: true
	Nickname string `json:"nickname"`
}

// Validate validates this wallet
func (m *Wallet) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKeyPair(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNickname(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Wallet) validateAddress(formats strfmt.Registry) error {

	if err := validate.RequiredString("address", "body", m.Address); err != nil {
		return err
	}

	return nil
}

func (m *Wallet) validateKeyPair(formats strfmt.Registry) error {

	if err := m.KeyPair.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("keyPair")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("keyPair")
		}
		return err
	}

	return nil
}

func (m *Wallet) validateNickname(formats strfmt.Registry) error {

	if err := validate.RequiredString("nickname", "body", m.Nickname); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this wallet based on the context it is used
func (m *Wallet) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateKeyPair(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Wallet) contextValidateKeyPair(ctx context.Context, formats strfmt.Registry) error {

	if err := m.KeyPair.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("keyPair")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("keyPair")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Wallet) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Wallet) UnmarshalBinary(b []byte) error {
	var res Wallet
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// WalletKeyPair wallet's key pair.
//
// swagger:model WalletKeyPair
type WalletKeyPair struct {

	// Nonce used by the AES-GCM algorithm used to protect the key pair's private key.
	// Required: true
	Nonce string `json:"nonce"`

	// Key pair's private key.
	// Required: true
	PrivateKey string `json:"privateKey"`

	// Key pair's public key.
	// Required: true
	PublicKey string `json:"publicKey"`

	// Salt used by the PBKDF that generates the secret key used to protect the key pair's private key.
	// Required: true
	Salt string `json:"salt"`
}

// Validate validates this wallet key pair
func (m *WalletKeyPair) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNonce(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrivateKey(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePublicKey(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSalt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WalletKeyPair) validateNonce(formats strfmt.Registry) error {

	if err := validate.RequiredString("keyPair"+"."+"nonce", "body", m.Nonce); err != nil {
		return err
	}

	return nil
}

func (m *WalletKeyPair) validatePrivateKey(formats strfmt.Registry) error {

	if err := validate.RequiredString("keyPair"+"."+"privateKey", "body", m.PrivateKey); err != nil {
		return err
	}

	return nil
}

func (m *WalletKeyPair) validatePublicKey(formats strfmt.Registry) error {

	if err := validate.RequiredString("keyPair"+"."+"publicKey", "body", m.PublicKey); err != nil {
		return err
	}

	return nil
}

func (m *WalletKeyPair) validateSalt(formats strfmt.Registry) error {

	if err := validate.RequiredString("keyPair"+"."+"salt", "body", m.Salt); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this wallet key pair based on context it is used
func (m *WalletKeyPair) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *WalletKeyPair) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WalletKeyPair) UnmarshalBinary(b []byte) error {
	var res WalletKeyPair
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
