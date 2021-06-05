package inaarea

import (
	"io/ioutil"

	"github.com/tidwall/gjson"
)

func getData(area Area, filters []AreaModel) ([]AreaModel, int) {
	file, _ := ioutil.ReadFile(getAsset(area))

	if len(filters) == 0 {
		filters = []AreaModel{
			{},
		}
	}

	all := gjson.GetBytes(file, "data.#")

	res := []gjson.Result{}

	for _, v := range filters {
		res = append(res, gjson.Get(string(file), buildQuery(v)).Array()...)
	}

	datas := []AreaModel{}

	for _, v := range res {
		datas = append(datas, parse(v))
	}

	return unique(datas), int(all.Int())
}

// GetProvinces generate province datas
func GetProvinces(filters []AreaModel) ([]AreaModel, int) {
	return getData(PROVINCE, filters)
}

// GetDistricts generate district datas
func GetDistricts(filters []AreaModel) ([]AreaModel, int) {
	return getData(DISTRICT, filters)
}

// GetDistrictsByProvince generate district by province code
func GetDistrictsByProvince(provinceCode string, filters []AreaModel) ([]AreaModel, int) {
	filters = buildFilterParent(provinceCode, filters)
	return getData(DISTRICT, filters)
}

// GetSUbDistricts generate sub districts datas
func GetSubDistricts(filters []AreaModel) ([]AreaModel, int) {
	return getData(SUBDISTRICT, filters)
}

// GetSubDistrictsByDistrict generate sub district by district code
func GetSubDistrictsByDistrict(districtCode string, filters []AreaModel) ([]AreaModel, int) {
	filters = buildFilterParent(districtCode, filters)
	return getData(SUBDISTRICT, filters)
}

// GetUrbanVillages generate urban villages datas
func GetUrbanVillages(filters []AreaModel) ([]AreaModel, int) {
	return getData(URBANVILLAGE, filters)
}

// GetUrbanVillageBySubDistrict generate urban village by sub district code
func GetUrbanVillagesBySubDistrict(subDistrictCode string, filters []AreaModel) ([]AreaModel, int) {
	filters = buildFilterParent(subDistrictCode, filters)
	return getData(URBANVILLAGE, filters)
}
