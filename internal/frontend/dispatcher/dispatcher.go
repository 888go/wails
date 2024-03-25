package dispatcher

import (
	"context"

	"github.com/pkg/errors"
	"github.com/888go/wails/internal/binding"
	"github.com/888go/wails/internal/frontend"
	"github.com/888go/wails/internal/logger"
	"github.com/888go/wails/pkg/options"
)

type Dispatcher struct {
	log        *logger.Logger
	bindings   *binding.Bindings
	events     frontend.Events
	bindingsDB *binding.DB
	ctx        context.Context
	errfmt     options.ErrorFormatter
}


// ff:
// errfmt:
// events:
// bindings:
// log:
// ctx:
func NewDispatcher(ctx context.Context, log *logger.Logger, bindings *binding.Bindings, events frontend.Events, errfmt options.ErrorFormatter) *Dispatcher {
	return &Dispatcher{
		log:        log,
		bindings:   bindings,
		events:     events,
		bindingsDB: bindings.DB(),
		ctx:        ctx,
		errfmt:     errfmt,
	}
}


// ff:
// sender:
// message:
func (d *Dispatcher) ProcessMessage(message string, sender frontend.Frontend) (string, error) {
	if message == "" {
		return "", errors.New("No message to process")
	}
	switch message[0] {
	case 'L':
		return d.processLogMessage(message)
	case 'E':
		return d.processEventMessage(message, sender)
	case 'C':
		return d.processCallMessage(message, sender)
	case 'c':
		return d.processSecureCallMessage(message, sender)
	case 'W':
		return d.processWindowMessage(message, sender)
	case 'B':
		return d.processBrowserMessage(message, sender)
	case 'Q':
		sender.X退出()
		return "", nil
	case 'S':
		sender.X显示()
		return "", nil
	case 'H':
		sender.X隐藏()
		return "", nil
	default:
		return "", errors.New("Unknown message from front end: " + message)
	}
}
