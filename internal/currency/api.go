package currency

import (
	"net/http"

	p "github.com/IlmarLopez/currency/pkg/pagination"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// RegisterHandlers registers the handlers for the currencies.
func RegisterHandlers(r *gin.RouterGroup, service Service, logger *zap.SugaredLogger) {
	res := resource{service, logger}

	r.GET("/currencies/:currency", res.query())
}

// resource is the implementation of the resource interface.
type resource struct {
	service Service
	logger  *zap.SugaredLogger
}

// query handles the GET request for the currencies.
func (r resource) query() func(c *gin.Context) {
	return func(c *gin.Context) {
		queryParameters := make(map[string]string)

		currency := c.Param("currency")
		queryParameters["finit"] = c.Request.URL.Query().Get("finit")
		queryParameters["fend"] = c.Request.URL.Query().Get("fend")

		ctx := c.Request.Context()
		count, err := r.service.Count(ctx, currency, queryParameters)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to count currencies"})
			return
		}

		pages := p.NewFromRequest(c.Request, count)

		currencies, err := r.service.Query(ctx, currency, queryParameters, pages.Offset(), pages.Limit())
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get currencies"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"currencies": currencies})
	}
}
