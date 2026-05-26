package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strings"
)

const (
	exitOK                = 0
	exitGenericFailure    = 1
	exitUsageError        = 2
	exitReadinessBlocked  = 3
	exitStaleLock         = 4
	exitInvalidContract   = 5
	exitUnsupportedTarget = 6
	exitSemanticConflict  = 7
)

type commandError struct {
	Code     string            `json:"code"`
	ExitCode int               `json:"exit_code"`
	Message  string            `json:"message"`
	Details  map[string]string `json:"details,omitempty"`
}

type errorEnvelope struct {
	Error commandError `json:"error"`
}

func (e commandError) Error() string {
	return e.Message
}

func usageErrorf(format string, args ...any) commandError {
	return commandError{
		Code:     "usage_error",
		ExitCode: exitUsageError,
		Message:  fmt.Sprintf(format, args...),
	}
}

func jsonRequested(args []string) bool {
	for _, arg := range args {
		if arg == "--json" {
			return true
		}
	}
	return false
}

func classifyError(err error) commandError {
	var typed commandError
	if errors.As(err, &typed) {
		return typed
	}

	message := err.Error()
	switch {
	case strings.Contains(message, "readiness is BLOCKED"):
		return commandError{Code: "readiness_blocked", ExitCode: exitReadinessBlocked, Message: message}
	case strings.Contains(message, "lock hash mismatch"):
		return commandError{Code: "stale_lock", ExitCode: exitStaleLock, Message: message}
	case strings.Contains(message, "unsupported target"), strings.Contains(message, "unsupported export target"):
		return commandError{Code: "unsupported_target", ExitCode: exitUnsupportedTarget, Message: message}
	case isInvalidContractMessage(message):
		return commandError{Code: "invalid_contract", ExitCode: exitInvalidContract, Message: message}
	default:
		return commandError{Code: "generic_failure", ExitCode: exitGenericFailure, Message: message}
	}
}

func isInvalidContractMessage(message string) bool {
	markers := []string{
		"malformed contract JSON",
		"contract missing required field",
		"unsupported contract schema",
		"unsupported product_type",
		"delivery_surfaces must contain",
		"unsupported delivery surface",
		"interaction_mode must be",
		"readiness profile",
		"must use ",
		"missing NG id",
		"missing CAP id",
		"missing REQ id",
		"missing DEC id",
		"missing RISK id",
		"missing EVAL id",
		"missing ART id",
		"missing OQ id",
	}
	for _, marker := range markers {
		if strings.Contains(message, marker) {
			return true
		}
	}
	return false
}

func failCommand(stdout io.Writer, stderr io.Writer, command string, err error, jsonOutput bool) int {
	typed := classifyError(err)
	if command != "" {
		typed.Message = fmt.Sprintf("%s failed: %s", command, typed.Message)
		if typed.Details == nil {
			typed.Details = map[string]string{}
		}
		typed.Details["command"] = command
	}
	return failStructured(stdout, stderr, typed, jsonOutput)
}

func failStructured(stdout io.Writer, stderr io.Writer, err commandError, jsonOutput bool) int {
	if jsonOutput {
		encoder := json.NewEncoder(stdout)
		encoder.SetIndent("", "  ")
		if encodeErr := encoder.Encode(errorEnvelope{Error: err}); encodeErr != nil {
			fmt.Fprintf(stderr, "failed to encode JSON error: %v\n", encodeErr)
			return exitGenericFailure
		}
		return err.ExitCode
	}
	fmt.Fprintln(stderr, err.Message)
	return err.ExitCode
}
