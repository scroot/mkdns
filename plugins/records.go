package plugins

import (
	"log"
	"net"
	//"fmt"
	//"strings"
)

/*
 *
 type & 1  view
 type & 2  weight
 type & 4  geo
*/

const (
	VIEW   = 1
	WEIGHT = 2
	GEO    = 4
)

var upChooseRecord = -1
var currentWeight int = 0

type BaseRecords struct {
	Addr    net.IP
	Records []interface{}
	RType   uint64
}

func newBaseRecords(addr net.IP, rtype uint64, records []interface{}) *BaseRecords {
	return &BaseRecords{
		Addr:    addr,
		Records: records,
		RType:   rtype,
	}
}

func (this *BaseRecords) GetRecords() (answer []interface{}) {
	var records []interface{}
	records = this.Records
	if this.RType&GEO == GEO {
		records = this.GeoRecord(records)
	}
	/*
		if this.RType & VIEW == VIEW {
			return this.ViewRecord()
		}
	*/
	if this.RType&WEIGHT == WEIGHT {
		records = this.WeightRecord(records)
	}
	return records
}

func (this *BaseRecords) getMaxWeight(records []interface{}) int {
	var ok bool
	var w uint64
	maxweight := 0
	for _, v := range records {
		if _, ok = v.(map[string]interface{})["weight"]; !ok {
			w = 0
		} else {
			w = v.(map[string]interface{})["weight"].(uint64)
		}
		weight := int(w)
		if weight > maxweight {
			maxweight = weight
		}
	}
	return maxweight
}

func (this *BaseRecords) GeoRecord(records []interface{}) (answer []interface{}) {
	var _country, _continent string
	var def_answer []interface{}
	hitGeo := false
	country, continent, netmask := geoIP.GetCountry(this.Addr)
	log.Printf("geoip= %s, country= %s, continent=%s, netmask=%d", this.Addr, country, continent, netmask)

	for _, v := range records {
		vv := v.(map[string]interface{})
		if _, ok := vv["country"]; ok {
			_country = vv["country"].(string)
		} else {
			_country = ""
		}
		if _, ok := vv["continent"]; ok {
			_continent = vv["continent"].(string)
		} else {
			_continent = ""
		}
		if _country == "" && _continent == "" {
			def_answer = append(def_answer, v)
		}
		if (_country != "" && _country == country) || (_continent != "" && _continent == continent) {
			hitGeo = true
			answer = append(answer, v)
		}
		//log.Printf("%+v, %+v", _country, _continent)
	}
	if !hitGeo {
		answer = def_answer
	}
	return
}

//http://www.wangshangyou.com/go/126.html
func (this *BaseRecords) WeightRecord(records []interface{}) (answer []interface{}) {
	var ok bool
	var w uint64
	rlen := len(records)
	if rlen == 0 {
		return
	}

	maxweight := this.getMaxWeight(records)
	log.Printf("maxweight : %d ", maxweight)
	for {
		upChooseRecord = (upChooseRecord + 1) % rlen
		if upChooseRecord == 0 {
			currentWeight = currentWeight - 2
			if currentWeight <= 0 {
				currentWeight = maxweight
			}
		}
		if _, ok = records[upChooseRecord].(map[string]interface{})["weight"]; !ok {
			w = 0
		} else {
			w = records[upChooseRecord].(map[string]interface{})["weight"].(uint64)
		}
		if weight := int(w); weight >= currentWeight {
			//log.Printf("%+v", records[upChooseRecord])
			answer = append(answer, records[upChooseRecord])
			break
		}
	}

	return
}