package main

import (
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/sagikazarmark/healthz"
)

// Creates the health service handler and the status checker
func newHealthServiceHandler(db *sqlx.DB) (http.Handler, *healthz.StatusChecker) {
	status := healthz.NewStatusChecker(healthz.Healthy)
	healthMux := healthz.NewHealthServiceHandler(healthz.NewCheckers(), healthz.NewCheckers(status, healthz.NewPingChecker(db)))

	return healthMux, status
}
