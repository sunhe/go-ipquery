package ipquery

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"net"
	"strings"
)

const (
	IPV4 = net.IPv4len
	IPV6 = net.IPv6len
)

type IPQuery struct {
	len  int
	tree *radixTreeNode
}

func New(len int) *IPQuery {
	ipq := new(IPQuery)
	ipq.len = len

	return ipq
}

func (ipq *IPQuery) BuildFromYaml(in []byte, sep string, part int) error {
	tree := radixTreeCreate()
	db := make(map[string][]string)
	err := yaml.Unmarshal(in, &db)
	if err != nil {
		return err
	}
	for k, v := range db {
		ar := strings.Split(k, sep)
		fmt.Println(len(ar))
		if len(ar) != part {
			return fmt.Errorf("%s part is not %d", k, part)
		}
		for i := range v {
			if !strings.Contains(v[i], "/") {
				switch ipq.len {
				case IPV4:
					v[i] += "/32"
				case IPV6:
					//to do
				}
			}
			_, ipnet, err := net.ParseCIDR(v[i])
			if err != nil {
				return err
			}
			tree.insert(ipnet.IP.To4(), ipnet.Mask, ipq.len, ar)
		}
	}
	ipq.tree = tree

	return nil
}

func (ipq *IPQuery) Query(ipstr string) []string {
	ip := net.ParseIP(ipstr).To4()
	value := ipq.tree.query(ip, ipq.len)
	if value == nil {
		return nil
	} else {
		return value.([]string)
	}
}

func (ipq *IPQuery) Delete() {
	ipq.tree.destroy()
}
