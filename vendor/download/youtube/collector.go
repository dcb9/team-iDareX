package youtube

import (
	"github.com/Go-Docker-Hackathon/team-iDareX/vendor/db/mongo"
	"labix.org/v2/mgo/bson"
	"fmt"
)

var WorkQueue = make(chan WorkRequest, 100)

func Collector(url , formatId string) {
	C := mongo.Connect()
	C.Update(bson.M{"fetchurl": url}, bson.M{"$set": bson.M{"status": 1}}) // start download
	fmt.Println("mongodb start download")
	
	work := WorkRequest{ Url: url, FormatId: formatId}	
	WorkQueue <- work

}