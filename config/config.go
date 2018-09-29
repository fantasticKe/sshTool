package config

import (
	"errors"
	"strings"
)

type SSHHost struct {
	Host      string
	Port      int
	Username  string
	Password  string
	Cmds      string
	Key       string
}

//获取ip集合
func GetIpList(ips string) ([]string, error) {
	var ipList []string
	if strings.Contains(ips, ",") {
		for _, ip := range strings.Split(ips,","){
			ipList = append(ipList, ip)
		}
	}else {
		return nil,errors.New("ips must join \",\"")
	}
	return ipList, nil
}

