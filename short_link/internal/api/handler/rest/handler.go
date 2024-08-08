package rest

import (
	"errors"
	"fmt"
	"github.com/JairoRiver/short_link_app/short_link/internal/controller"
	"github.com/JairoRiver/short_link_app/short_link/internal/util"
	"github.com/gin-gonic/gin"
)

// Handler defines a HTTP short Link handler.
type Handler struct {
	ctrl   *controller.Controller
	config util.Config
}

// New creates a new service HTTP handler.
func New(ctrl *controller.Controller, config util.Config) *Handler {
	return &Handler{ctrl, config}
}

// CreateLinkType - Custom type to hold value for new short link type from 1-3
type CreateLinkType int

// Declare related constants for each short link with index 1
const (
	Short CreateLinkType = iota + 1
	Custom
	//Suggestion  uncomment when add suggestion type
)

// String - Creating common behavior - give the type a String function
func (c CreateLinkType) String() string {
	return [...]string{"Short", "Custom"}[c-1]
	//return [...]string{"Short", "Custom", "Suggestion"}[c-1]
}

// EnumIndex - Creating common behavior - give the type a EnumIndex function
func (c CreateLinkType) EnumIndex() int {
	return int(c)
}

// errorResponse is a helper function to return a error message
func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

var customLinkTokenLengthError = errors.New(fmt.Sprint("custom token length must be greater than ", util.MaxLenToken))
var linkTypeError = errors.New("link type error must be 1 for Short or 2 for Custom")
