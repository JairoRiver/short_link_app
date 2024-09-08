package rest

import (
	"errors"
	"fmt"
	"github.com/JairoRiver/short_link_app/short_link/internal/repository"
	"github.com/JairoRiver/short_link_app/short_link/internal/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

// AvailabilityLink handler Post / requests check the custom token availability.
type availabilityLinkRequest struct {
	Token string `uri:"token" binding:"required"`
}

// @Summary Check if a custom link are available
// @Description AvailabilityLink check if one custom token no are in use yet
// @ID post-check-link
// @Accept  json
// @Produce  json
// @Param   token     path    string     true        "token"
// @Success 200 string token availability
// @Router /check/{token} [post]
func (h *Handler) AvailabilityLink(ctx *gin.Context) {
	var req availabilityLinkRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(fmt.Errorf("Handler AvailabilityLink get uri error: %w", err)))
		return
	}

	//Check the token length
	if len(req.Token) > util.MaxLenToken {
		_, err := h.ctrl.GetByCustomToken(ctx, req.Token)
		if err != nil {
			if errors.Is(err, repository.ErrNotFound) {
				ctx.JSON(http.StatusOK, gin.H{"message": "token availability", "status": "available"})
				return
			}

			ctx.JSON(http.StatusInternalServerError, errorResponse(fmt.Errorf("Handler AvailabilityLink GetByCustomToken error: %w", err)))
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"message": "token not availability", "status": "busy"})
		return
	}
	ctx.JSON(http.StatusBadRequest, errorResponse(fmt.Errorf("Handler AvailabilityLink token len error: %w", customLinkTokenLengthError)))
	return
}
