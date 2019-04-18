package sfomuseum

import (
	"errors"
	"github.com/whosonfirst/go-whosonfirst-flags"
	"github.com/whosonfirst/go-whosonfirst-flags/existential"
	"github.com/whosonfirst/go-whosonfirst-geojson-v2"
	"github.com/whosonfirst/go-whosonfirst-geojson-v2/properties/whosonfirst"
	"github.com/whosonfirst/go-whosonfirst-geojson-v2/utils"
	"github.com/whosonfirst/go-whosonfirst-placetypes"
	_ "log"
	"strings"
)

func Placetype(f geojson.Feature) string {

	possible := []string{
		"properties.sfomuseum:placetype",
	}

	return utils.StringProperty(f.Bytes(), possible, f.Placetype())
}

func IsSFO(f geojson.Feature) (flags.ExistentialFlag, error) {

	possible := []string{
		"properties.sfomuseum:is_sfo",
	}

	v := utils.Int64Property(f.Bytes(), possible, -1)

	if v == 1 || v == 0 {
		return existential.NewKnownUnknownFlag(v)
	}

	return existential.NewKnownUnknownFlag(-1)
}

func Depicts(f geojson.Feature) ([]int64, error) {

	switch Placetype(f) {

	case "exhibition":
		return DepictsCampus(f)
	case "publicart":
		return DepictsCampus(f)
	default:
		return nil, errors.New("Unsupported placetype")
	}
}

func DepictsCampus(f geojson.Feature) ([]int64, error) {

	pt, err := placetypes.GetPlacetypeByName("campus")

	if err != nil {
		return nil, err
	}

	valid := make(map[string]bool)
	depicts := make(map[int64]bool)

	children := placetypes.Children(pt)

	for _, p := range children {
		valid[p.Name] = bool
	}

	hierarchies := whosonfirst.Hierarchies(f)

	for _, h := range hierarchies {

		for k, id := range h {

			_, ok := depicts[id]

			if ok {
				continue
			}

			k = strings.Replace(k, "_id", "", 1)

			_, ok = valid[k]

			if ok {
				depicts[id] = true
			}
		}
	}

	ids := make([]int64, 0)

	for id, _ := range depicts {
		ids = append(ids, id)
	}

	return ids
}
