package elastic

import (
  elastic "github.com/olivere/elastic/v7"
  "fmt"
  "context"
 _ "log"
 "time"
)
const (
index = "some_index"
mappings = `
{
"settings":{
"number_of_shards":2,
"number_of_replicas":1
},
"mappings":{
"properties":{
"field str":{
"type":"text"
},
"field int":{
"type":"integer"
},
"field bool":{
"type":"boolean"
}
}
}
}
`
)
func Run(){
	_, err := elastic.NewClient(
		elastic.SetSniff(true), // true and not work, find reason for SetSniff if wtf
		elastic.SetURL("http://103.127.206.218:9200"),
		elastic.SetHealthcheckInterval(5 * time.Second),
	)
	if err !=nil{
		fmt.Println(err)
	}
	fmt.Println(elastic.Version)
	fmt.Println(Body(mappings))
	ctx := context.Background()
	fmt.Println(ctx)
	create, err := es.CreateIndex(index).BodyString(mappings).Do(ctx)
	// if err != nil {
	// log.Fatalf("CreateIndex() ERROR: %v", err)
	// } else {
	// fmt.Println("CreateIndex():", create)
	// }

}