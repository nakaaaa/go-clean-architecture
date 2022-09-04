package usecase

import "context"

type Usecase[TInputPort any, TOutputPort any] interface {
	Execute(ctx context.Context, input TInputPort) (TOutputPort, error)
}
