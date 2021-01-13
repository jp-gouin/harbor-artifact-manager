package notification

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"

	"backend/args"

	"github.com/bndr/gojenkins"
	"github.com/jtblin/go-ldap-client"
	"go.etcd.io/etcd/clientv3"
)

type AdminNotification struct {
	EtcdStats    []clientv3.StatusResponse `json:"etcd"`
	JenkinsNodes []gojenkins.NodeResponse  `json:"jenkins,omitempty"`
	HarborHealth HarborHealth              `json:"harbor"`
	LdapStatus   bool                      `json:"ldap,omitempty"`
	WSType       string                    `json:"mutation"`
}

type HarborHealth struct {
	Status     string `json:"status"`
	Components []struct {
		Name   string `json:"name"`
		Status string `json:"status,omitempty"`
		Error  string `json:"error,omitempty"`
	} `json:"components"`
}

/*
StartHealthCheck run global health check each minute
*/
func (nm *Manager) StartHealthCheck(jenkins *gojenkins.Jenkins, ldapClient *ldap.LDAPClient) {
	ticker := time.NewTicker(1 * time.Minute)
	go func() {
		for t := range ticker.C {
			_ = t // we don't print the ticker time, so assign this `t` variable to underscore `_` to avoid error
			go nm.notifyAdminForHealthCheck(jenkins, ldapClient)
		}
	}()
}

/*
 NotifyAdminForHealthCheck global health check of all components
*/
func (nm *Manager) notifyAdminForHealthCheck(jenkins *gojenkins.Jenkins, ldapClient *ldap.LDAPClient) {

	// Create a nil slice of integers.
	var etcdStruct []clientv3.StatusResponse
	var jenkinsStruct []gojenkins.NodeResponse
	var harborHealth HarborHealth
	dataHealth := AdminNotification{etcdStruct, jenkinsStruct, harborHealth, false, "ADMHEALTHCHECK"}

	var wg sync.WaitGroup
	wg.Add(4)

	// Check ETCD
	go func() {
		defer wg.Done()
		nm.Log("Check ETCD")
		for _, ep := range nm.Dao.Cli.Endpoints() {
			fmt.Println(ep)
			etcdStatus, err := nm.Dao.Cli.Status(nm.Dao.Ctx, ep)
			if err == nil && etcdStatus != nil {
				dataHealth.EtcdStats = append(dataHealth.EtcdStats, *etcdStatus)
			} else {
				fmt.Printf("error checking etcd %v %v", err, etcdStatus)
			}

		}
	}()
	// Check Jenkins
	go func() {
		defer wg.Done()
		if jenkins != nil {
			nm.Log("Check Jenkins")
			nodes, err := jenkins.GetAllNodes()
			nm.Log(fmt.Sprintf("%v", err))
			for _, node := range nodes {
				node.Poll()
				dataHealth.JenkinsNodes = append(dataHealth.JenkinsNodes, *node.Raw)
			}
		}
	}()
	// Check Harbor
	go func() {
		defer wg.Done()
		nm.Log("Check Harbor")
		req, err := http.NewRequest("GET", args.Args.Harbor+"/api/v2.0/health", nil)
		req.SetBasicAuth(args.Args.User, args.Args.Password)
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			fmt.Printf("there was an error performing the http request +%+v \n", err)
			return
		}
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)
		json.Unmarshal([]byte(body), &harborHealth)
		dataHealth.HarborHealth = harborHealth

	}()
	// Check LDAP
	go func() {
		defer wg.Done()
		if ldapClient != nil {
			nm.Log("Check Ldap")
			err := ldapClient.Connect()
			nm.Log(fmt.Sprintf("Error check ldap %v", err))
			if err == nil {
				dataHealth.LdapStatus = true
			}
		}
	}()
	wg.Wait()
	b, _ := json.Marshal(dataHealth)
	clients := nm.Hub.GetClients()
	for c := range clients {
		if c.Admin {
			nm.Hub.SendToClient(c, b)
		}
	}

}
