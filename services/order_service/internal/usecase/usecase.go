package usecase

import (
	"context"
)

type UseCase struct {
	Order
}

func New() *UseCase {
	return uc
}

func (u *UseCase) Start(ctx context.Context) error {
	return nil
}

func (u *UseCase) Stop(ctx context.Context) error {
	return nil
}
