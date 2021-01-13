package database

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	dm "backend/datamodel"
	"backend/helpers"

	"backend/args"

	"go.etcd.io/etcd/clientv3"
)

var (
	dialTimeout = 3 * time.Second
	kv          clientv3.KV
	cli         *clientv3.Client
)

/*
DAO for handling database operation
*/
type DAO struct {
	Kv  clientv3.KV
	Cli *clientv3.Client
	Ctx context.Context
}

/*
Init the package
*/
func Init(ctx context.Context) *DAO {
	fmt.Println("init database cli")
	cli, _ := clientv3.New(clientv3.Config{
		Endpoints:   []string{args.Args.Etcd},
		DialTimeout: dialTimeout,
	})
	//defer cli.Close()
	kv = clientv3.NewKV(cli)
	return &DAO{
		Ctx: ctx,
		Kv:  kv,
		Cli: cli,
	}
}

/*
TestDAO test the dao
*/
func (dao *DAO) TestDAO() []byte {
	opts := []clientv3.OpOption{clientv3.WithProgressNotify()}
	dao.Kv.Put(context.TODO(), "dataTest", "{test}", opts...)
	gr, _ := dao.Kv.Get(dao.Ctx, "dataTest")
	return gr.Kvs[0].Value
}

/*
Get data from etcd database
*/
func (dao *DAO) Get(id string) (*clientv3.GetResponse, error) {
	gr, err := dao.Kv.Get(dao.Ctx, id)
	return gr, err
}

/*
Set data to etcd database
*/
func (dao *DAO) Set(id string, data string) {
	r, err := dao.Kv.Put(dao.Ctx, id, data)
	fmt.Printf("\n %v %v \n", r, err)
}

/*
Delete the id entry
*/
func (dao *DAO) Delete(id string) {
	dao.Kv.Delete(dao.Ctx, id)
}

/*
 GetArtifactDB from database
*/
func (dao *DAO) GetArtifactDB(id string) (*dm.Artifact, error) {
	var artifactDB *dm.Artifact
	gr, _ := dao.Get(id)
	if len(gr.Kvs) == 0 {
		return nil, errors.New("no artifact found")
	}
	json.Unmarshal(gr.Kvs[0].Value, &artifactDB)
	return artifactDB, nil
}

/*
 GetChartDB from database
*/
func (dao *DAO) GetChartDB(id string) (*dm.ChartDB, error) {
	var chartDB *dm.ChartDB
	gr, _ := dao.Get(id)
	if len(gr.Kvs) == 0 {
		return nil, errors.New("no artifact found")
	}
	json.Unmarshal(gr.Kvs[0].Value, &chartDB)
	return chartDB, nil
}

/*
 GetDockerImage from database
*/
func (dao *DAO) GetDockerImage(id string) (*dm.DockerImage, error) {
	var dockerImage *dm.DockerImage
	gr, _ := dao.Get(id)
	if len(gr.Kvs) == 0 {
		return nil, errors.New("no artifact found")
	}
	json.Unmarshal(gr.Kvs[0].Value, &dockerImage)
	return dockerImage, nil
}

/*
 GetProjectDB from database
*/
func (dao *DAO) GetProjectDB(id string) (*dm.ProjectDB, error) {
	var projectDB *dm.ProjectDB
	gr, _ := dao.Get(id)
	if len(gr.Kvs) == 0 {
		return nil, errors.New("no artifact found")
	}
	json.Unmarshal(gr.Kvs[0].Value, &projectDB)
	return projectDB, nil
}

/*
 GetVersionDB from database
*/
func (dao *DAO) GetVersionDB(id string) (*dm.VersionDB, error) {
	var versionDB *dm.VersionDB
	gr, _ := dao.Get(id)
	if len(gr.Kvs) == 0 {
		return nil, errors.New("no artifact found")
	}
	json.Unmarshal(gr.Kvs[0].Value, &versionDB)
	return versionDB, nil
}

/*
 GetProjectsDB from database
*/
func (dao *DAO) GetProjectsDB(key string) ([]*dm.ProjectDB, error) {
	gr, _ := dao.Kv.Get(dao.Ctx, key, clientv3.WithPrefix())
	if len(gr.Kvs) == 0 {
		return nil, errors.New("no artifact found")
	}
	result := make([]*dm.ProjectDB, 0)
	for _, v := range gr.Kvs {
		var projectDB *dm.ProjectDB
		json.Unmarshal(v.Value, &projectDB)
		result = append(result, projectDB)
	}
	return result, nil
}

/*
GetNotifications from the database
*/
func (dao *DAO) GetNotifications(notificationType string, limit int64, owner string) ([]*dm.Notification, error) {
	result := make([]*dm.Notification, 0)
	gr, _ := dao.Kv.Get(dao.Ctx, notificationType, clientv3.WithPrefix(), clientv3.WithLimit(limit))
	if len(gr.Kvs) == 0 {
		return nil, errors.New("no notification found")
	}
	for _, v := range gr.Kvs {
		var notif *dm.Notification
		json.Unmarshal(v.Value, &notif)
		if helpers.IsAdmin(owner) || owner == notif.Owner {
			result = append(result, notif)
		}
	}
	return result, nil
}

/*
 GetUsers from database
*/
func (dao *DAO) GetUsers() ([]*dm.User, error) {
	gr, _ := dao.Kv.Get(dao.Ctx, "User_", clientv3.WithPrefix())
	if len(gr.Kvs) == 0 {
		return nil, errors.New("no artifact found")
	}
	result := make([]*dm.User, 0)
	for _, v := range gr.Kvs {
		var user *dm.User
		json.Unmarshal(v.Value, &user)
		result = append(result, user)
	}
	return result, nil
}

/*
 GetUser from database
*/
func (dao *DAO) GetUser(id string) (*dm.User, error) {
	var user *dm.User
	gr, _ := dao.Get("User_" + id)
	if len(gr.Kvs) == 0 {
		return nil, errors.New("no artifact found")
	}
	json.Unmarshal(gr.Kvs[0].Value, &user)
	return user, nil
}

/*
 GetStarredArtifacts from database
*/
func (dao *DAO) GetStarredArtifacts() ([]*dm.StarredArtifacts, error) {
	var starredArtifacts []*dm.StarredArtifacts
	gr, _ := dao.Get("StarredArtifacts")
	if len(gr.Kvs) == 0 {
		return nil, errors.New("no artifact found")
	}
	json.Unmarshal(gr.Kvs[0].Value, &starredArtifacts)
	return starredArtifacts, nil
}
