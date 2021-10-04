package middleware

import (
	"github.com/eosnationftw/eosn-base-api/log"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"net/http/httputil"
	"net/url"
)

type ReverseProxyMiddleware struct {
	targetUrl *url.URL
}

func NewReverseProxyMiddleware(target string) (*ReverseProxyMiddleware, error) {

	targetUrl, err := url.Parse(target)
	if err != nil {
		return nil, err
	}

	return &ReverseProxyMiddleware{targetUrl: targetUrl}, nil
}

func (r *ReverseProxyMiddleware) ProxyRequest(responseHandler func(*gin.Context) func(*http.Response) error) gin.HandlerFunc {

	return func(c *gin.Context) {
		proxy := httputil.NewSingleHostReverseProxy(r.targetUrl)

		if responseHandler != nil {
			proxy.ModifyResponse = responseHandler(c)
		}

		proxy.ErrorHandler = func(writer http.ResponseWriter, request *http.Request, err error) {
			log.Panic("failed to reverse proxy request", zap.Error(err), zap.String("request", request.RequestURI))
		}

		proxy.ServeHTTP(c.Writer, c.Request)
	}
}

