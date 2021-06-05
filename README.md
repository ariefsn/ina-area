# INA-Area

Generate Area in Indonesia **(Offline)** :)

## Data

- **34** Provinces
- **499** Districts
- **6878** Sub Districts
- **79702** Urban Villages

## Install

```sh
go get -u github.com/ariefsn/ina-area
```

## Import

```go
import ia "github.com/ariefsn/ina-area"
```

## Model

```go
type AreaModel struct {
  Code   string `json:"code"`
  Name   string `json:"name"`
  Parent string `json:"parent"`
}
```

## API

- **Get Provinces**

  ```go
  // Will return slice provinces and total data
  data, _ := ia.GetProvinces(nil)

  fmt.Println(data[0]) // {11 Aceh }

  // With filter
  data, _ := ia.GetProvinces([]ia.AreaModel{
    {
      Code: "35",
    },
  })

  fmt.Println(data) // [{35 Jawa Timur }]

  // Find by code ==> if not found it will return nil
  data2 := ia.GetProvinceByCode("35")

  fmt.Println(data2) // &{35 Jawa Timur }
  ```

- **Get Districts**

  ```go
  // Will return slice districs and total data
  data, _ := ia.GetDistrics(nil)

  fmt.Println(data[0]) // {1101 Kab. Simeulue 11}

  // With filter
  data, _ := ia.GetDistricts([]ia.AreaModel{
    {
      Code: "3521",
    },
  })

  fmt.Println(data) // [{3521 Kab. Ngawi 35}]

  // Or find by specific province code
  data, _totalDist2_ := ia.GetDistrictsByProvince("35", []ia.AreaModel{
    {
      Name: "Ngawi",
    },
  })

  fmt.Println(data) // [{3521 Kab. Ngawi 35}]

  // Find by code ==> if not found it will return nil
  data2 := ia.GetDistrictByCode("3521")

  fmt.Println(data2) // &{3521 Kab. Ngawi 35}
  ```

- **Get Sub Districts**

  ```go
  // Will return slice sub districs and total data
  data, _ := ia.GetSubDistrics(nil)

  fmt.Println(data[0]) // {1101010 Teupah Selatan 1101}

  // With filter
  data, _ := ia.GetSubDistricts([]ia.AreaModel{
    {
      Code: "3521030",
    },
  })

  fmt.Println(data) // [{3521030 Jogorogo 3521}]

  // Or find by specific district code
  data, _totalDist2_ := ia.GetSubDistrictsByDistrict("3521", []ia.AreaModel{
    {
      Name: "Jogorogo",
    },
  })

  fmt.Println(data) // [{3521030 Jogorogo 3521}]

  // Find by code ==> if not found it will return nil
  data2 := ia.GetSubDistrictByCode("3521030")

  fmt.Println(data2) // &{3521030 Jogorogo 3521}
  ```

- **Get Urban Villages**

  ```go
  // Will return slice urban villages and total data
  data, _ := ia.GetUrbanVillages(nil)

  fmt.Println(data[0]) // {9104073001 Weriagar 9104073}

  // With filter
  data, _ := ia.GetUrbanVillages([]ia.AreaModel{
    {
      Code: "3521030009",
    },
  })

  fmt.Println(data) // [{3521030009 Jogorogo 3521030}]

  // Or find by specific sub district code
  data, _totalDist2_ := ia.GetUrbanVillagesBySubDistrict("3521", []ia.AreaModel{
    {
      Name: "Jogorogo",
    },
  })

  fmt.Println(data) // [{3521030009 Jogorogo 3521030}]


  // Find by code ==> if not found it will return nil
  data2 := ia.GeturbanVillageByCode("3521030009")

  fmt.Println(data2) // &{3521030009 Jogorogo 3521030}
  ```
