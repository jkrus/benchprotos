// Copyright © 2020-2021 The EVEN Solutions Developers Team

package app

import (
	"github.com/evenlab/go-kit/errors"
)

const (
	errStartApplicationMsg  = "start application"
	errProvideCliContextMsg = "provide cli context"
	errStartServiceGRPCMsg  = "start service grpc"
)

func ErrStartApplication(w error) error {
	return errors.WrapErr(errStartApplicationMsg, w)
}

func ErrProvideCliContext(w error) error {
	return errors.WrapErr(errProvideCliContextMsg, w)
}

func ErrStartServiceGRPC(w error) error {
	return errors.WrapErr(errStartServiceGRPCMsg, w)
}
