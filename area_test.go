package inaarea

import "testing"

type testCase struct {
	name    string
	area    Area
	filters []AreaModel
	expect  string
	message string
}

func TestArea(t *testing.T) {
	testCases := []testCase{
		{
			name: "Province Test",
			area: PROVINCE,
			filters: []AreaModel{
				{
					Code: "35",
				},
			},
			expect:  "Jawa Timur",
			message: `should check province "Jawa Timur"`,
		},
		{
			name: "District Test",
			area: DISTRICT,
			filters: []AreaModel{
				{
					Code: "3521",
				},
			},
			expect:  "Kab. Ngawi",
			message: `should check district "Kab. Ngawi"`,
		},
		{
			name: "Sub District Test",
			area: SUBDISTRICT,
			filters: []AreaModel{
				{
					Code: "3521030",
				},
			},
			expect:  "Jogorogo",
			message: `should check district "Jogorogo"`,
		},
		{
			name: "Urban Village Test",
			area: URBANVILLAGE,
			filters: []AreaModel{
				{
					Code: "3521030009",
				},
			},
			expect:  "Jogorogo",
			message: `should check urban villages "Jogorogo"`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			datas := []AreaModel{}

			if tc.area == PROVINCE {
				datas, _ = GetProvinces(tc.filters)
			} else if tc.area == DISTRICT {
				datas, _ = GetDistricts(tc.filters)
			} else if tc.area == SUBDISTRICT {
				datas, _ = GetSubDistricts(tc.filters)
			} else {
				datas, _ = GetUrbanVillages(tc.filters)
			}

			if datas[0].Name != tc.expect {
				t.Errorf(tc.message)
			}
		})
	}
}
