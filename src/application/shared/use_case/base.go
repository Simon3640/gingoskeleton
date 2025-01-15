package usecase

import (
	"context"

	"gingoskeleton/src/application/shared/locales"
)

type BaseUseCase[Input any, Output any] interface {
	SetLocale(locale locales.LocaleTypeEnum)
	Execute(ctx context.Context, input Input) *UseCaseResult[Output]
}
