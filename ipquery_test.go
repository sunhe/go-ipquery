package ipquery

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestUsage(t *testing.T) {
	ipq := New(IPV4)
	out, err := ioutil.ReadFile("./ip.yaml")
	if err != nil {
		t.Error(err)
		return
	}
	err = ipq.BuildFromYaml(out, ".", 5)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(ipq.Query("192.168.1.1"))
	fmt.Println(ipq.Query("192.168.2.1"))
	fmt.Println(ipq.Query("192.168.1.2"))
}
