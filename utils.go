package inaarea

import (
	"encoding/json"
	"fmt"
	"path"
	"runtime"
	"strings"

	"github.com/tidwall/gjson"
)

type Area string

const (
	PROVINCE     Area = "PROVINCE"
	DISTRICT          = "DISTRICT"
	SUBDISTRICT       = "SUBDISTRICT"
	URBANVILLAGE      = "URBANVILLAGE"
)

func getPath() (dirPath, fileName string) {
	_, fileName, _, ok := runtime.Caller(0)

	if !ok {
		panic("No caller information")
	}

	dirPath = path.Dir(fileName)

	return
}

func getAsset(area Area) string {
	dir, _ := getPath()
	file := "provinces.json"

	switch area {
	case DISTRICT:
		file = "districts.json"
	case SUBDISTRICT:
		file = "subDistricts.json"
	case URBANVILLAGE:
		file = "urbanVillages.json"
	}

	return path.Join(dir, "assets", file)
}

func parse(obj gjson.Result) AreaModel {
	model := AreaModel{}

	_ = json.Unmarshal([]byte(obj.Raw), &model)

	return model
}

func buildQuery(filter AreaModel) string {
	queries := []string{}

	if filter.Code != "" {
		queries = append(queries, fmt.Sprintf(`#(code%s"*%s*")#`, "%", filter.Code))
	}

	if filter.Name != "" {
		queries = append(queries, fmt.Sprintf(`#(name%s"*%s*")#`, "%", filter.Name))
	}

	if filter.Parent != "" {
		queries = append(queries, fmt.Sprintf(`#(parent%s"*%s*")#`, "%", filter.Parent))
	}

	query := "*"

	if len(queries) > 0 {
		query = "data." + strings.Join(queries, "|")
	}

	return query
}

func buildFilterParent(parentCode string, filters []AreaModel) []AreaModel {
	if len(filters) == 0 {
		filters = append(filters, AreaModel{
			Parent: parentCode,
		})
	} else {
		for i, _ := range filters {
			filters[i].Parent = parentCode
		}
	}

	return filters
}

func unique(areas []AreaModel) []AreaModel {
	keys := make(map[string]bool)
	list := []AreaModel{}

	for _, a := range areas {
		if _, value := keys[a.Code]; !value {
			keys[a.Code] = true
			list = append(list, a)
		}
	}

	return list
}
