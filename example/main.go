package main

import (
	"fmt"

	ia "github.com/ariefsn/ina-area"
)

func main() {
	provinces, totalProv := ia.GetProvinces(nil)
	provinces2, totalProv2 := ia.GetProvinces([]ia.AreaModel{
		{
			Code: "35",
		},
	})
	provinces3 := ia.GetProvinceByCode("35")

	fmt.Println("Provinces", len(provinces), provinces[0], totalProv, "=>", provinces2, totalProv2)
	fmt.Println("Specific Province", provinces3)

	districts, totalDist := ia.GetDistricts(nil)
	districts2, totalDist2 := ia.GetDistrictsByProvince("35", []ia.AreaModel{
		{
			Name: "Ngawi",
		},
	})
	district3 := ia.GetDistrictByCode("3521")

	fmt.Println("Districts", len(districts), districts[0], totalDist, "=>", districts2, totalDist2)
	fmt.Println("Specific District", district3)

	subDistricts, totalSubDist := ia.GetSubDistricts(nil)
	subDistricts2, totalSubDist2 := ia.GetSubDistrictsByDistrict("3521", []ia.AreaModel{
		{
			Name: "Jogorogo",
		},
	})
	subDistricts3 := ia.GetSubDistrictByCode("3521030")

	fmt.Println("SubDistricts", len(subDistricts), subDistricts[0], totalSubDist, "=>", subDistricts2, totalSubDist2)
	fmt.Println("Specific SubDistrict", subDistricts3)

	urbanVillages, totalUB := ia.GetUrbanVillages(nil)
	urbanVillages2, totalUB2 := ia.GetUrbanVillagesBySubDistrict("3521030", []ia.AreaModel{
		{
			Name: "Jogorogo",
		},
	})
	urbanVillages3 := ia.GetUrbanVillageByCode("3521030009")

	fmt.Println("UrbanVillages", len(urbanVillages), urbanVillages[0], totalUB, "=>", urbanVillages2, totalUB2)
	fmt.Println("Specific UrbanVillage", urbanVillages3)
}
