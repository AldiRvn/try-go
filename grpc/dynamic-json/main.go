// Using arbitary JSON / map[string]interface in grpc
// ref: https://stackoverflow.com/a/65026774/11983973
package main

import (
	"google.golang.org/protobuf/types/known/structpb"
)

func main() {
	data, err := structpb.NewValue(map[string]interface{}{
		"name":    "Ricardo Hogan",
		"country": "Guam",
		"date":    "29/05/2021",
	})
	if err != nil {
		log.Fatalln(err)
	}

	dataAsModel := stub.Response{
		Message: "OK",
		Data:    data.GetStructValue(),
	}
	asBytes, _ := json.MarshalIndent(
		&dataAsModel, "", "  ")
	fmt.Println(string(asBytes))

	for key, field := range dataAsModel.Data.Fields {
		fmt.Println(key, field)
	}

  //* output
  // {
  //   "message": "OK",
  //   "data": {
  //     "country": "Guam",
  //     "date": "29/05/2021",
  //     "name": "Ricardo Hogan"
  //   }
  // }
  // name string_value:"Ricardo Hogan"
  // country string_value:"Guam"
  // date string_value:"29/05/2021"
}
