package rest

import (
	"fmt"
	"github.com/JairoRiver/short_link_app/short_link/internal/controller"
	"github.com/JairoRiver/short_link_app/short_link/internal/repository"
	"github.com/JairoRiver/short_link_app/short_link/internal/util"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"net/url"
)

// CreateLink handler POST /create requests.
type CreateLinkRequest struct {
	Url      string         `json:"url" binding:"required"`
	LinkType CreateLinkType `json:"type" binding:"required"`
	UserId   uuid.UUID      `json:"user_id"`
	Token    string         `json:"token"`
}
type CreateLinkResponse struct {
	Url      string `json:"url"`
	Token    string `json:"token"`
	ShortUrl string `json:"short_url"`
}

func parseCreateLinkResponse(response controller.ShortLinkResponse, domainName string) (CreateLinkResponse, error) {
	shortUrl, err := url.JoinPath(domainName, response.Token)
	if err != nil {
		return CreateLinkResponse{}, err
	}

	return CreateLinkResponse{
		Url:      response.Url,
		Token:    response.Token,
		ShortUrl: shortUrl,
	}, nil
}

// @Summary Create a new short Link
// @Description generate a short link or custom link from a request
// @ID post-create-link
// @Accept  json
// @Produce  json
// @Param   request     body    rest.CreateLinkRequest     true        "Create Link Request"
// @Success 200 {object} CreateLinkResponse
// @Router /v1/create [post]
func (h *Handler) CreateLink(ctx *gin.Context) {
	var req CreateLinkRequest
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
		response, err := parseCreateLinkResponse(*newShortLink, h.config.DomainName)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(fmt.Errorf("Handler CreateLink ShortLink parseLinkResponse error: %w", err)))
			return
		}
		ctx.JSON(http.StatusOK, response)
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
		response, err := parseCreateLinkResponse(*customLink, h.config.DomainName)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(fmt.Errorf("Handler CreateLink CustomLink parseLinkResponse error: %w", err)))
			return
		}
		ctx.JSON(http.StatusOK, response)
		return
	} else {
		ctx.JSON(http.StatusBadRequest, errorResponse(fmt.Errorf("Handler CreateLink error type error: %w", linkTypeError)))
		return
	}

	//TODO add suggestion link
}
