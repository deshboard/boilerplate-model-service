package main

// ServiceName is an identifier-like name used anywhere this app needs to be identified.
//
// It identifies the service itself, the actual instance needs to be identified via environment
// and other details.
const ServiceName = "boilerplate.model.service"

// FriendlyServiceName is the visible name of the service.
const FriendlyServiceName = "Boilerplate Model service"

// LogTag is usually a static value across all instances of the same application
// as such it is set here as a constant value.
//
// It represents an identifier which can be used to separate logs from different sources.
//
// Falls back to the ServiceName.
const LogTag string = ServiceName
