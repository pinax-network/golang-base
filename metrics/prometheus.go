package metrics

import (
	"github.com/Depado/ginprom"
	"github.com/gin-gonic/gin"
)

func NewPrometheusExporter(engine *gin.Engine, path string) *ginprom.Prometheus {
	prometheusExporter := ginprom.New(
		ginprom.Engine(engine),
		ginprom.Subsystem("gin"),
		ginprom.Path(path),
	)

	return prometheusExporter
}
