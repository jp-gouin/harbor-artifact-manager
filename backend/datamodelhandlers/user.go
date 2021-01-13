package datamodelhandlers

import (
	"backend/database"
	dm "backend/datamodel"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

/*
SaveUser to the database
*/
func SaveUser(user *dm.User, dao *database.DAO) {
	fmt.Println(("\n Saving user\n "))
	b, err := json.Marshal(user)
	if err == nil {
		dao.Set("User_"+user.Username, string(b))
	} else {
		fmt.Println(err)
	}
}

/*
GetOrCreateAUser return a user from the database, if it does not exist it creates the object
*/
func GetOrCreateAUser(username string, dao *database.DAO, refreshToken bool) (*dm.User, error) {
	user, err := dao.GetUser(username)
	fmt.Printf("%v  %v", user, err)
	if err != nil {
		userToCreate := &dm.User{username, nil, make([]string, 0)}
		SaveUser(userToCreate, dao)
		user, _ = dao.GetUser(username)
		fmt.Printf("%v  %v", userToCreate, user)
	}
	if refreshToken {
		fmt.Println("refresh account")
		return RefreshUserRobotAccount(user, dao)
	}
	return user, nil
}

/*
RefreshUserRobotAccount for the provided user, if the robot account is nil or expired (or about to expire), generate a new one
*/
func RefreshUserRobotAccount(user *dm.User, dao *database.DAO) (*dm.User, error) {
	if user.Robot == nil {
		robot, err := generateRobotAccount(user.Username, dao)
		if err != nil {
			return nil, err
		}
		user.Robot = robot
		SaveUser(user, dao)
	} else if (user.Robot.ExpiresAt - time.Now().Unix()) < 3600 {
		robot, err := generateRobotAccount(user.Username, dao)
		if err != nil {
			return nil, err
		}
		user.Robot = robot
		SaveUser(user, dao)
	}
	fmt.Printf("\n refreshed user account %v \n", user)
	return user, nil
}

func generateRobotAccount(username string, dao *database.DAO) (*dm.RobotV1, error) {
	// TODO : Now the project is hardcoded to 2 , fetch all project and create a robot account for all projects
	// This could mean that a user has a list of robot account...
	// Query harbor to see if there is an existing robot account
	// if the robot account already exist , patch it (?)
	// if the case come to : a user as a running build and trigger another build after the 23th hours of the robot account generation
	// the current build will become failed if it's still pulling chart of images
	robots := dao.GetHarborRobots()
	for _, robot := range robots {
		if strings.Contains(robot.Name, username) {
			dao.DeleteHarborRobot(robot)
		}
	}
	robot := &dm.RobotV1{nil, username, "robot$" + username, "automatic robot account for " + username, "", time.Now().AddDate(0, 0, 1).Unix(), []dm.RobotAccountAccess{{"pull", "/project/2/repository"}, {"read", "/project/2/helm-chart"}}}
	token, err := dao.CreateHarborRobot(robot)
	if err != nil {
		return nil, err
	}
	robot.Token = token
	return robot, nil
}
