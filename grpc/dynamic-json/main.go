// Using arbitary JSON / map[string]interface in grpc
// ref: https://stackoverflow.com/a/65026774/11983973
package main

import (
	"google.golang.org/protobuf/types/known/structpb"
)

func main() {
	response := &stub.Response{}
	detail := `[{"id":1,"name":"Leanne Graham","username":"Bret","email":"Sincere@april.biz","address":{"street":"Kulas Light","suite":"Apt. 556","city":"Gwenborough","zipcode":"92998-3874","geo":{"lat":"-37.3159","lng":"81.1496"}},"phone":"1-770-736-8031 x56442","website":"hildegard.org","company":{"name":"Romaguera-Crona","catchPhrase":"Multi-layered client-server neural-net","bs":"harness real-time e-markets"}},{"id":2,"name":"Ervin Howell","username":"Antonette","email":"Shanna@melissa.tv","address":{"street":"Victor Plains","suite":"Suite 879","city":"Wisokyburgh","zipcode":"90566-7771","geo":{"lat":"-43.9509","lng":"-34.4618"}},"phone":"010-692-6593 x09125","website":"anastasia.net","company":{"name":"Deckow-Crist","catchPhrase":"Proactive didactic contingency","bs":"synergize scalable supply-chains"}}]`

	//? New object using NewValue
	data, err := structpb.NewValue(map[string]interface{}{
		"country": "Guam",
		"date":    "29/05/2021",
	})
	if err != nil {
		log.Println(err)
	}
	//? New list of object using unmarshal
	_ = json.Unmarshal([]byte(detail), &response.Detail)

	response.Message = "OK"
	response.Data = data.GetStructValue()
	responseAsBytes, _ := json.MarshalIndent(response, "", "  ")
	fmt.Println(string(responseAsBytes))

//* output
// {
//   "message": "OK",
//   "data": {
//     "country": "Guam",
//     "date": "29/05/2021"
//   },
//   "detail": [
//     {
//       "address": {
//         "city": "Gwenborough",
//         "geo": {
//           "lat": "-37.3159",
//           "lng": "81.1496"
//         },
//         "street": "Kulas Light",
//         "suite": "Apt. 556",
//         "zipcode": "92998-3874"
//       },
//       "company": {
//         "bs": "harness real-time e-markets",
//         "catchPhrase": "Multi-layered client-server neural-net",
//         "name": "Romaguera-Crona"
//       },
//       "email": "Sincere@april.biz",
//       "id": 1,
//       "name": "Leanne Graham",
//       "phone": "1-770-736-8031 x56442",
//       "username": "Bret",
//       "website": "hildegard.org"
//     },
//     {
//       "address": {
//         "city": "Wisokyburgh",
//         "geo": {
//           "lat": "-43.9509",
//           "lng": "-34.4618"
//         },
//         "street": "Victor Plains",
//         "suite": "Suite 879",
//         "zipcode": "90566-7771"
//       },
//       "company": {
//         "bs": "synergize scalable supply-chains",
//         "catchPhrase": "Proactive didactic contingency",
//         "name": "Deckow-Crist"
//       },
//       "email": "Shanna@melissa.tv",
//       "id": 2,
//       "name": "Ervin Howell",
//       "phone": "010-692-6593 x09125",
//       "username": "Antonette",
//       "website": "anastasia.net"
//     }
//   ]
// }
}
