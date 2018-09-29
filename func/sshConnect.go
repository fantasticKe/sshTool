package _func

import (
	"fmt"
	"golang.org/x/crypto/ssh"
	"io/ioutil"
	"log"
	"net"
	"os"
	"sshTool/config"
	"time"
)

//执行命令
func DoCmd(hostCfg config.SSHHost) {
	session, err := SSHConnect(hostCfg,[]string{})
	if err != nil {
		log.Printf("connect host err,%s\n",err.Error())
	}
	defer session.Close()
	session.Stderr = os.Stderr
	session.Stdout = os.Stdout
	session.Run(hostCfg.Cmds)
}

func SSHConnect(hostCfg config.SSHHost, cipherList []string) (*ssh.Session, error) {
	var (
		auth         []ssh.AuthMethod
		addr         string
		clientConfig *ssh.ClientConfig
		client       *ssh.Client
		config       ssh.Config
		session      *ssh.Session
		err          error
	)

	//get auth method
	auth = make([]ssh.AuthMethod, 0)
	if hostCfg.Key == "" {
		auth = append(auth, ssh.Password(hostCfg.Password))
	} else {
		pemBytes, err := ioutil.ReadFile(hostCfg.Key)
		if err != nil {
			return nil, err
		}

		var signer ssh.Signer
		if hostCfg.Password == "" {
			signer, err = ssh.ParsePrivateKey(pemBytes)
		} else {
			signer, err = ssh.ParsePrivateKeyWithPassphrase(pemBytes, []byte(hostCfg.Password))
		}
		if err != nil {
			return nil, err
		}
		auth = append(auth, ssh.PublicKeys(signer))
	}
	if len(cipherList) == 0 {
		config = ssh.Config{
			Ciphers: []string{"aes128-ctr", "aes192-ctr", "aes256-ctr", "aes128-gcm@openssh.com", "arcfour256",
			"arcfour128", "aes128-cbc", "3des-cbc", "aes192-cbc", "aes256-cbc"},
		}
	} else {
		config = ssh.Config{
			Ciphers: cipherList,
		}
	}
	clientConfig = &ssh.ClientConfig{
		User:    hostCfg.Username,
		Auth:    auth,
		Timeout: 30 * time.Second,
		Config:  config,
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
	}
	//connect ssh
	addr = fmt.Sprintf("%s:%d", hostCfg.Host, hostCfg.Port)
	if client, err = ssh.Dial("tcp", addr, clientConfig); err != nil {
		return nil, err
	}
	//create session
	if session, err = client.NewSession(); err != nil {
		return nil, err
	}

	modes := ssh.TerminalModes{
		ssh.ECHO:          0,     // disable echoing
		ssh.TTY_OP_ISPEED: 14400, // input speed = 14.4kbaud
		ssh.TTY_OP_OSPEED: 14400, // output speed = 14.4kbaud
	}

	if err := session.RequestPty("xterm", 80, 40, modes); err != nil {
		return nil, err
	}

	return session, nil
}
