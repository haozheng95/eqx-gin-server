package startapi

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	metal "github.com/haozheng95/eqx-gin-server/metal/v1"
	"golang.org/x/crypto/ssh"
	"log"
	"net"
	"net/http"
	"strings"
	"sync"
	"time"
)

// timeout middleware wraps the request context with a timeout
func timeoutMiddleware(timeout time.Duration) func(c *gin.Context) {
	return func(c *gin.Context) {

		// wrap the request context with a timeout
		ctx, cancel := context.WithTimeout(c.Request.Context(), timeout)

		defer func() {
			// check if context timeout was reached
			if ctx.Err() == context.DeadlineExceeded {

				// write response and abort the request
				c.Writer.WriteHeader(http.StatusGatewayTimeout)
				c.Abort()
			}

			//cancel to clear resources after finished
			cancel()
		}()

		// replace request with context wrapped request
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

func timedHandler(duration time.Duration) func(c *gin.Context) {
	return func(c *gin.Context) {

		// get the underlying request context
		ctx := c.Request.Context()

		// create the response data type to use as a channel type
		type responseData struct {
			status int
			body   map[string]interface{}
		}

		// create a done channel to tell the request it's done
		doneChan := make(chan responseData)

		// here you put the actual work needed for the request
		// and then send the doneChan with the status and body
		// to finish the request by writing the response
		go func() {
			time.Sleep(duration)
			doneChan <- responseData{
				status: 200,
				body:   gin.H{"hello": "world"},
			}
		}()

		// non-blocking select on two channels see if the request
		// times out or finishes
		select {

		// if the context is done it timed out or was cancelled
		// so don't return anything
		case <-ctx.Done():
			return

			// if the request finished then finish the request by
			// writing the response
		case res := <-doneChan:
			c.JSON(res.status, res.body)
		}
	}

}

var sshPool sync.Map

func execCmd(sshHost string, sshPort int, cmd string) ([]byte, error) {
	sshConfig := &ssh.ClientConfig{
		Timeout:         time.Second,
		User:            "root",
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	//if false {
	//	sshConfig.Auth = []ssh.AuthMethod{ssh.Password(data.SshPassword)}
	//} else {
	key := "-----BEGIN OPENSSH PRIVATE KEY-----\n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\n-----END OPENSSH PRIVATE KEY-----"
	sshConfig.Auth = []ssh.AuthMethod{publicKeyAuthFunc([]byte(key))}
	//}

	session := getSession(sshHost, sshPort, sshConfig)
	log.Println("cmd ", cmd)
	return session.CombinedOutput(cmd)
}

func publicKeyAuthFunc(key []byte) ssh.AuthMethod {

	//key, err := ioutil.ReadFile(kPath)
	//if err != nil {
	//	log.Println("ssh key file read failed", err)
	//}
	// Create the Signer for this private key.
	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		log.Println("ssh key signer failed", err)
	}
	return ssh.PublicKeys(signer)
}

func getSession(sshHost string, sshPort int, config *ssh.ClientConfig) (session *ssh.Session) {

	addr := fmt.Sprintf("%s:%d", sshHost, sshPort)
	log.Println("ssh addr ", addr)
	sshClientStock, has := sshPool.Load(addr)
	var sshClient *ssh.Client

	if !has {
		sshClient, err := ssh.Dial("tcp", addr, config)
		if err != nil {
			log.Println("created ssh client failed", err)
			if err := sshClient.Close(); err != nil {
				log.Println("close ssh client failed", err)
			}
		} else {
			log.Println("existing sshClient")
			sshPool.Store(addr, sshClient)
		}
	} else {
		log.Println("map existing sshClient")
		sshClient = sshClientStock.(*ssh.Client)
	}

	defer func(addr string) {
		err := recover()
		if err != nil {
			log.Println("recover ", err)
			sshClient, err := ssh.Dial("tcp", addr, config)
			log.Println("defer", err)
			log.Println("defer existing sshClient")
			sshPool.Delete(addr)
			sshPool.Store(addr, sshClient)
			session, err = sshClient.NewSession()
			log.Println("defer", err)
		}
	}(addr)

	//create ssh-session
	session, err := sshClient.NewSession()
	if err != nil {
		log.Println("created session failed", err)
		if err := session.Close(); err != nil {
			log.Println("close session failed", err)
		}
	}

	return session
}

func IsIPv4(ipAddr string) bool {
	ip := net.ParseIP(ipAddr)
	return ip != nil && strings.Contains(ipAddr, ".")
}

func addAuthorizedKeys(ip string, body *metal.CreateDeviceRequestOneOf) {
	cmd := "cat ~/.ssh/authorized_keys"
	combo, err := execCmd(ip, 22, cmd)
	if err != nil {
		log.Println("remote run cmd failed", cmd, err)
	}
	log.Println("cmd output:", string(combo))
	for i, v := range body.SshKeys {
		log.Println(i, " check key:", v.GetKey())
		if !strings.Contains(string(combo), v.GetKey()) {
			log.Println(i, " add key:", v.GetKey())
			cmd = "echo \"" + v.GetKey() + "\" >> ~/.ssh/authorized_keys"
			combo, err = execCmd(ip, 22, cmd)
			if err != nil {
				log.Println("remote run cmd failed", cmd, err)
			}
		}
	}
}
