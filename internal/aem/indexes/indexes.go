package indexes

import (
	"fmt"
	"github.com/jlentink/aem/internal/aem"
	"github.com/jlentink/aem/internal/aem/objects"
	"github.com/jlentink/aem/internal/http"
	"github.com/tidwall/gjson"
)

func getStringArrayValue(result gjson.Result, path string) []string {
	if !result.Get(path).Exists() {
		return nil
	}

	if result.Get(path).IsArray() {
		values := make([]string, 0)
		for _, arrValue := range result.Get(path).Array() {
			values = append(values, arrValue.String())
		}
		return values
	}

	return []string{result.Get(path).String()}
}

// GetIndexes retrieves the indexes from a aem instance
func GetIndexes(instance *objects.Instance) ([]*Index, error) {

	pw, err := instance.GetPassword()
	if err != nil {
		return nil, err
	}

	if !aem.Cnf.ValidateSSL {
		http.DisableSSLValidation()
	}

	resp, err := http.GetPlain(instance.URLString()+indexes, instance.Username, pw)
	if err != nil {
		return nil, err
	}

	result := gjson.ParseBytes(resp)
	indexes := make([]*Index, 0)
	rMap := result.Map()
	for k := range rMap {
		cResult := result.Get(k)
		if cResult.Type == gjson.JSON {
			isIndex := cResult.Get("jcr:primaryType")
			if isIndex.Type == gjson.String && isIndex.Str == "oak:QueryIndexDefinition" {
				index := Index{}
				index.Info = cResult.Get("info").String()
				index.Name = k
				index.Type = cResult.Get("type").String()
				index.ReindexCount = cResult.Get("reindexCount").Int()
				index.Async = getStringArrayValue(cResult, "async")
				index.ExcludedPaths = getStringArrayValue(cResult, "excludedPaths")
				index.IncludedPaths = getStringArrayValue(cResult, "includedPaths")
				indexes = append(indexes, &index)
			}
		}
	}
	return indexes, nil
}

// Reindex start reindex of indexed on aem instance
func Reindex(instance *objects.Instance, index string) error {
	pw, err := instance.GetPassword()
	if err != nil {
		return err
	}

	if !aem.Cnf.ValidateSSL {
		http.DisableSSLValidation()
	}

	_, _, err = http.PostMultiPart(instance.URLString()+fmt.Sprintf(reindexURL, index), instance.Username, pw, map[string]string{"reindex" : "true"})
	
	if err != nil {
		return err
	}
	return nil
}
