package main

import (
	"flag"
	"log"
	"sshTool/config"
	"sshTool/func"
)

func main() {
	ips := flag.String("ips", "", "IP address list")
	port := flag.Int("p", 22, "Port")
	username := flag.String("u", "", "Username")
	password := flag.String("pwd", "", "Password")
	key := flag.String("k", "", "SSH private key")
	cmds := flag.String("cmds", "", "Command that needs to be executed")
	flag.Parse()

	var ipList []string
	var sshHosts []config.SSHHost
	var hostCfg config.SSHHost
	var err error

	if *ips != "" {
		if ipList, err = config.GetIpList(*ips); err != nil{
			log.Printf("get ip err,%s\n",err.Error())
		}
	}
	for _, i := range ipList{
		hostCfg.Host = i
		hostCfg.Username = *username
		hostCfg.Password = *password
		hostCfg.Port = *port
		hostCfg.Key = *key
		hostCfg.Cmds = *cmds
		sshHosts = append(sshHosts, hostCfg)
	}
	for _, h := range sshHosts{
		go _func.DoCmd(h)
	}
}
