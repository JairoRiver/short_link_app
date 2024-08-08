package rest

import (
	"errors"
	"fmt"
	"github.com/JairoRiver/short_link_app/short_link/internal/repository"
	"github.com/JairoRiver/short_link_app/short_link/internal/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetLink handler Get / requests.
type getLinkRequest struct {
	Token string `uri:"token" binding:"required"`
}

// @Summary GetLink get a token an retuen a redirect
// @Description get a short link or custom link and return a redirect
// @ID get-link
// @Accept  json
// @Produce  json
// @Param   token     path    string     true        "token"
// @Success 307
// @Router /{token} [get]
func (h *Handler) GetLink(ctx *gin.Context) {
	var req getLinkRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(fmt.Errorf("Handler GetLink get uri error: %w", err)))
		return
	}

	//Check is Short Link
	if len(req.Token) <= util.MaxLenToken {
		link, err := h.ctrl.GetByToken(ctx, req.Token)
		if err != nil {

			if errors.Is(err, repository.ErrNotFound) {
				ctx.HTML(http.StatusNotFound, "404.html", gin.H{
					"title": "404 Not Found",
				})
				return
			}

			ctx.JSON(http.StatusInternalServerError, errorResponse(fmt.Errorf("Handler GetLink GetByToken error: %w", err)))
			return
		}
		ctx.Redirect(http.StatusTemporaryRedirect, link.Url)
		return
	}

	//Case When is a Custom Token
	link, err := h.ctrl.GetByCustomToken(ctx, req.Token)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			ctx.HTML(http.StatusNotFound, "404.html", gin.H{
				"title": "404 Not Found",
			})
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(fmt.Errorf("Handler GetLink GetByCustomToken error: %w", err)))
		return
	}
	ctx.Redirect(http.StatusTemporaryRedirect, link.Url)
	return
}
