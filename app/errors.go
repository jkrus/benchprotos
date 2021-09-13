// Copyright © 2020-2021 The EVEN Solutions Developers Team

package app

import (
	"github.com/evenlab/go-kit/errors"
)

const (
	errProvideCliContextMsg = "provide cli context"
	errStartServiceGRPCMsg  = "start service grpc"
	errStartServiceCRPCMsg  = "start service crpc"
)

func ErrProvideCliContext(w error) error {
	return errors.WrapErr(errProvideCliContextMsg, w)
}

func ErrStartServiceGRPC(w error) error {
	return errors.WrapErr(errStartServiceGRPCMsg, w)
}

func ErrStartServiceCRPC(w error) error {
	return errors.WrapErr(errStartServiceCRPCMsg, w)
}
