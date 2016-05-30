package main

import (
	"fmt"	
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/jsonpb"
)

func main() {
	var pb proto.Message
	a := `{"name":"test.com","records":[{"name":"@","ttl":300,"type":"SOA","value":{"mname":"ns1.test1.com.","nname":"dns-admin.test1.com.","serial":305419896,"refresh":1193046,"retry":624485,"expire":4913,"minttl":389333}},{"name":"@","ttl":300,"type":"NS","value":[{"record":["ns1.test1.com.","ns2.test2.com"]}]},{"name":"@","ttl":300,"type":"A","value":[{"record":["1.47.46.2","1.2.3.3"]}]},{"name":"view","ttl":300,"type":"A","state":1,"value":[{"record":["1.47.46.2","1.2.3.3"],"view":"any"},{"record":["1.2.3.4","1.2.4.5"],"view":"dx"}]},{"name":"weight","ttl":300,"type":"A","state":2,"value":[{"record":["1.47.46.2","1.2.3.3"],"weight":3},{"record":["1.2.3.4","1.2.4.5"],"weight":7},{"record":["7.7.7.7"],"weight":10}]},{"name":"geo","ttl":300,"type":"A","state":4,"value":[{"record":["7.7.7.7"]},{"record":["1.7.6.2"],"continent":"asia"},{"record":["1.7.6.5"],"continent":"asia","country":"cn"},{"record":["1.7.6.6"],"country":"cn"},{"record":["1.2.3.4","1.2.4.5"],"country":"kr"},{"record":["1.1.1.1","1.2.2.3","1.1.1.2"],"continent":"north-america"},{"record":["1.1.1.3"],"country":"us"}]}]}`
	jsonpb.UnmarshalString(a, pb)
	fmt.Println("Hello, playground")
}
