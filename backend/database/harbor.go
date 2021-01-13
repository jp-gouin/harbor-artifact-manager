package database

import (
	"backend/args"
	dm "backend/datamodel"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

// TODO : put all harbor related stuff here (fetching project, create/delete label, robot account)
/*
GetHarborRobots test the dao
*/
func (dao *DAO) GetHarborRobots() []*dm.RobotV1 {
	req, err := http.NewRequest("GET", args.Args.Harbor+"/api/v2.0/projects/2/robots", nil)
	req.SetBasicAuth(args.Args.User, args.Args.Password)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("there was an error performing the http request +%+v \n", err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	var robots []*dm.RobotV1
	json.Unmarshal([]byte(body), &robots)
	return robots
}

/*
CreateHarborRobot test the dao
*/
func (dao *DAO) CreateHarborRobot(robot *dm.RobotV1) (string, error) {
	b, err := json.Marshal(robot)
	req, err := http.NewRequest("POST", args.Args.Harbor+"/api/v2.0/projects/2/robots", bytes.NewBuffer(b))
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(args.Args.User, args.Args.Password)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("there was an error performing the http request %+v \n", err)
		return "", err
	}
	fmt.Printf("result : %+v", resp)

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	var robotresp struct {
		Token string `json:"token"`
		Name  string `json:"name"`
	}
	json.Unmarshal([]byte(body), &robotresp)
	return robotresp.Token, nil
}

/*
DeleteHarborRobot test the dao
*/
func (dao *DAO) DeleteHarborRobot(robot *dm.RobotV1) error {
	req, err := http.NewRequest("DELETE", args.Args.Harbor+"/api/v2.0/projects/2/robots/"+strconv.Itoa(int(*robot.ID)), nil)
	req.SetBasicAuth(args.Args.User, args.Args.Password)
	client := &http.Client{}
	_, err = client.Do(req)
	if err != nil {
		fmt.Printf("there was an error performing the http request +%+v \n", err)
		return err
	}
	return nil
}
