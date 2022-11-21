package api

import (
	"net/http"

	"github.com/chrmang-jambit/demo-container-go/pkg/counter"
	"github.com/labstack/echo/v4"
)

type Api struct {
	counter *counter.Counter
}

var _ ServerInterface = (*Api)(nil)

func New() *Api {
	return &Api{
		counter: counter.New(),
	}
}

func (x *Api) GetCounter(ctx echo.Context) error {
	log := ctx.Logger()
	res := CounterResponse{
		Counter: int(x.counter.Inc()),
	}
	log.Infof("counter: %d", x.counter.Get())
	return ctx.JSON(http.StatusOK, res)
}

func (x *Api) SetCounter(ctx echo.Context, params SetCounterParams) error {
	log := ctx.Logger()
	x.counter.Set(int64(*params.Value))
	res := CounterResponse{
		Counter: (int)(x.counter.Get()),
	}
	log.Warnf("set counter to %d", *params.Value)
	return ctx.JSON(http.StatusOK, res)
}
