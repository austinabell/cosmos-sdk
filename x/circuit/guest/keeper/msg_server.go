package keeper

import (
	"bytes"
	"context"
	"errors"
	"fmt"

	"cosmossdk.io/collections"
	errorsmod "cosmossdk.io/errors/v2"
	"cosmossdk.io/x/circuit/types"
)

var (
	ErrUnauthorized   = errorsmod.New("circuit", 0, "unauthorized")
	ErrInvalidRequest = errorsmod.New("circuit", 1, "invalid request")
)

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the circuit MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) *msgServer {
	return &msgServer{Keeper: keeper}
}

func (srv msgServer) AuthorizeCircuitBreaker(ctx context.Context, msg *types.MsgAuthorizeCircuitBreaker) (*types.MsgAuthorizeCircuitBreakerResponse, error) {
	address, err := srv.addressCodec.StringToBytes(msg.Granter)
	if err != nil {
		return nil, err
	}

	// if the granter is the module authority no need to check perms
	if !bytes.Equal(address, srv.GetAuthority()) {
		// Check that the authorizer has the permission level of "super admin"
		perms, err := srv.Permissions.Get(ctx, address)
		if err != nil {
			if errors.Is(err, collections.ErrNotFound) {
				return nil, errorsmod.Wrap(ErrUnauthorized, "only super admins can authorize users")
			}

			return nil, err
		}

		if perms.Level != types.Permissions_LEVEL_SUPER_ADMIN {
			return nil, errorsmod.Wrap(ErrUnauthorized, "only super admins can authorize users")
		}
	}

	grantee, err := srv.addressCodec.StringToBytes(msg.Grantee)
	if err != nil {
		return nil, err
	}

	if msg.Permissions == nil {
		return nil, errorsmod.Wrap(ErrInvalidRequest, "permissions cannot be nil")
	}

	// Append the account in the msg to the store's set of authorized super admins
	if err = srv.Permissions.Set(ctx, grantee, *msg.Permissions); err != nil {
		return nil, err
	}

	return &types.MsgAuthorizeCircuitBreakerResponse{
		Success: true,
	}, nil
}

func (srv msgServer) TripCircuitBreaker(ctx context.Context, msg *types.MsgTripCircuitBreaker) (*types.MsgTripCircuitBreakerResponse, error) {
	address, err := srv.addressCodec.StringToBytes(msg.Authority)
	if err != nil {
		return nil, err
	}

	// Check that the account has the permissions
	perms, err := srv.Permissions.Get(ctx, address)
	if err != nil && !errors.Is(err, collections.ErrNotFound) {
		return nil, err
	}

	for _, msgTypeURL := range msg.MsgTypeUrls {
		// check if the message is in the list of allowed messages
		isAllowed, err := srv.IsAllowed(ctx, msgTypeURL)
		if err != nil {
			return nil, err
		}

		if !isAllowed {
			return nil, fmt.Errorf("message %s is already disabled", msgTypeURL)
		}

		switch {
		case perms.Level == types.Permissions_LEVEL_SUPER_ADMIN || perms.Level == types.Permissions_LEVEL_ALL_MSGS || bytes.Equal(address, srv.GetAuthority()):
			// if the sender is a super admin or the module authority, no need to check perms
		case perms.Level == types.Permissions_LEVEL_SOME_MSGS:
			// if the sender has permission for some messages, check if the sender has permission for this specific message
			if !hasPermissionForMsg(perms, msgTypeURL) {
				return nil, errorsmod.Wrapf(ErrUnauthorized, "account does not have permission to trip circuit breaker for message %s", msgTypeURL)
			}
		default:
			return nil, errorsmod.Wrap(ErrUnauthorized, "account does not have permission to trip circuit breaker")
		}

		if err = srv.DisableList.Set(ctx, msgTypeURL); err != nil {
			return nil, err
		}

	}

	return &types.MsgTripCircuitBreakerResponse{
		Success: true,
	}, nil
}

// ResetCircuitBreaker resumes processing of Msg's in the state machine that
// have been been paused using TripCircuitBreaker.
func (srv msgServer) ResetCircuitBreaker(ctx context.Context, msg *types.MsgResetCircuitBreaker) (*types.MsgResetCircuitBreakerResponse, error) {
	keeper := srv.Keeper
	address, err := srv.addressCodec.StringToBytes(msg.Authority)
	if err != nil {
		return nil, err
	}

	// Get the permissions for the account specified in the msg.Authority field
	perms, err := keeper.Permissions.Get(ctx, address)
	if err != nil && !errors.Is(err, collections.ErrNotFound) {
		return nil, err
	}

	for _, msgTypeURL := range msg.MsgTypeUrls {
		// check if the message is in the list of allowed messages
		isAllowed, err := srv.IsAllowed(ctx, msgTypeURL)
		if err != nil {
			return nil, err
		}

		if isAllowed {
			return nil, fmt.Errorf("message %s is not disabled", msgTypeURL)
		}

		switch {
		case perms.Level == types.Permissions_LEVEL_SUPER_ADMIN || perms.Level == types.Permissions_LEVEL_ALL_MSGS || bytes.Equal(address, srv.GetAuthority()):
			// if the sender is a super admin or the module authority, no need to check perms
		case perms.Level == types.Permissions_LEVEL_SOME_MSGS:
			// if the sender has permission for some messages, check if the sender has permission for this specific message
			if !hasPermissionForMsg(perms, msgTypeURL) {
				return nil, errorsmod.Wrapf(ErrUnauthorized, "account does not have permission to reset circuit breaker for message %s", msgTypeURL)
			}
		default:
			return nil, errorsmod.Wrap(ErrUnauthorized, "account does not have permission to reset circuit breaker")
		}

		if err = srv.DisableList.Remove(ctx, msgTypeURL); err != nil {
			return nil, err
		}
	}

	return &types.MsgResetCircuitBreakerResponse{Success: true}, nil
}

// hasPermissionForMsg returns true if the account can trip or reset the message.
func hasPermissionForMsg(perms types.Permissions, msg string) bool {
	for _, msgurl := range perms.LimitTypeUrls {
		if msg == msgurl {
			return true
		}
	}
	return false
}