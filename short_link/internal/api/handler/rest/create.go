package rest

import (
	"fmt"
	"github.com/JairoRiver/short_link_app/short_link/internal/repository"
	"github.com/JairoRiver/short_link_app/short_link/internal/util"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

// CreateLink handler POST /create requests.
type createLinkRequest struct {
	Url      string         `json:"url" binding:"required"`
	LinkType CreateLinkType `json:"type" binding:"required"`
	UserId   uuid.UUID      `json:"user_id"`
	Token    string         `json:"token"`
}

func (h *Handler) CreateLink(ctx *gin.Context) {
	var req createLinkRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(fmt.Errorf("Handler CreateLink get body params error: %w", err)))
		return
	}

	//Validate if we have userId
	userId := repository.HasUserID{}
	if req.UserId == uuid.Nil {
		userId.Valid = false
	} else {
		userId.ID = req.UserId
		userId.Valid = true
	}

	//Handler to shortLink type
	if req.LinkType == Short {
		newShortLink, err := h.ctrl.CreateShortLink(ctx, req.Url, userId)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(fmt.Errorf("Handler CreateLink ShortLink type error: %w", err)))
			return
		}
		ctx.JSON(http.StatusOK, newShortLink)
		return
	}
	//Handler to customLink type
	if req.LinkType == Custom {
		//Validate the token len must be gt 6
		if len(req.Token) <= util.MaxLenToken {
			ctx.JSON(http.StatusBadRequest, errorResponse(fmt.Errorf("Handler CreateLink CustomLink Token len error: %w", customLinkTokenLengthError)))
			return
		}
		customLink, err := h.ctrl.CreateCustomLink(ctx, req.Url, userId, req.Token)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(fmt.Errorf("Handler CreateLink CustomLink type error: %w", err)))
			return
		}
		ctx.JSON(http.StatusOK, customLink)
		return
	} else {
		ctx.JSON(http.StatusBadRequest, errorResponse(fmt.Errorf("Handler CreateLink error type error: %w", linkTypeError)))
		return
	}

	//TODO add suggestion link
}
