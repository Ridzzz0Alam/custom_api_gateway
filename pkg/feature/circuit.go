package feature

import (
	"log/slog"

	"github.com/ArmaanKaytal/pkg/config"
	"github.com/Ridzzz0Alam/api_gateway/pkg/config"
	"github.com/sony/gobreaker/v2"
)

type CirciutBreaker struct {
	Settings config.CircuitSettings `json:"settings"`
	breaker  *gobreaker.CircuitBreaker[[]byte]
}

func NewCirciutBreaker(svcName string, settings config.CircuitSettings) *CirciutBreaker {
	return &CirciutBreaker{
		Settings: settings,
		breaker:  gobreaker.NewCircuitBreaker[[]byte](settings.Into(svcName)),
	}
}

func (cb *CirciutBreaker) Execute(service string, f func() ([]byte, error)) ([]byte, error) {
	slog.Info("Forwarding request using circuit breaker", "service", service, "breaker", cb.breaker.Name)
	return cb.breaker.Execute(f)
}

func (cb *CirciutBreaker) IsOpen() bool {
	return cb.breaker.State() == gobreaker.StateOpen
}

func (cb *CirciutBreaker) isEnabled() bool {
	return cb.Settings.isEnabled
}
