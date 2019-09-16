package indexes

const (
	indexes = "/crx/server/crx.default/jcr:root/oak:index.1.json"
	reindexUrl = "/oak:index/%s"
)

type IndexList struct {
	JcrPrimaryType string   `json:":jcr:primaryType"`
	JcrMixinTypes  []string `json:"jcr:mixinTypes"`
}

type Index struct {
	Name          string   `json:"name"`
	Info          string   `json:"info"`
	Type          string   `json:"type"`
	QueryPaths    []string `json:"queryPaths"`
	Async         []string `json:"async"`
	ReindexCount  int64    `json:"reindexCount"`
	IncludedPaths []string `json:"includedPaths"`
	ExcludedPaths []string `json:"excludedPaths"`
}
