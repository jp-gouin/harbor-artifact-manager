package notification

import (
	"encoding/json"
	"fmt"
	"time"

	"backend/database"
	"backend/datamodel"
	dm "backend/datamodel"
	"backend/hub"

	guuid "github.com/google/uuid"
)

//Manager struct
type Manager struct {
	Dao *database.DAO
	Hub *hub.Hub
}

/*
Init the package
*/
func Init(dao *database.DAO, hub *hub.Hub) *Manager {
	return &Manager{
		Dao: dao,
		Hub: hub,
	}
}

// NotifyProject send a project notification, can be
// Creation of a project/version
// Modification of project
// deletion of project
// All member of a project shall be notified along administrator
// This type of notification is also backup in ETCD
func (nm *Manager) NotifyProject(action string, owner string, severity string, title string, project *datamodel.ProjectDB) (*dm.Notification, error) {
	nm.Log("Notifying user of project change")
	currentTime := time.Now()
	notification := dm.Notification{guuid.New().String(), currentTime.Format("2006.01.02 15:04:05"), title, owner, "project", action, 0, severity, "NOTIFICATION"}
	b, err := json.Marshal(notification)
	if err != nil {
		nm.Log(err.Error())
		return nil, err
	}
	nm.sendNotification(b, project)
	nm.Dao.Set("Notification_"+notification.ID, string(b))
	return &notification, nil
}

// NotifyBuild send a notification that project are being build
// All member of a project shall be notified along administrator
// This type of notification is also backup in ETCD
// Return the notification for progress tracking
func (nm *Manager) NotifyBuild(action string, owner string, title string, project *datamodel.ProjectDB) (*dm.Notification, error) {
	currentTime := time.Now()
	notification := dm.Notification{guuid.New().String(), currentTime.Format("2006.01.02 15:04:05"), title, owner, "progress", action, 0, "", "NOTIFICATION"}
	b, err := json.Marshal(notification)
	nm.sendNotification(b, project)
	if err != nil {
		nm.Log(err.Error())
		return nil, err
	}
	return &notification, nil
}

// UpdateNotification send an update on a notification
func (nm *Manager) UpdateNotification(notification *dm.Notification, project *datamodel.ProjectDB) (*dm.Notification, error) {
	// Do stuff
	b, err := json.Marshal(notification)
	if err != nil {
		nm.Log(err.Error())
		return nil, err
	}
	nm.sendNotification(b, project)
	if notification.Type == "project" {
		nm.Dao.Set("Notification_"+notification.ID, string(b))
	}
	return notification, nil
}

func (nm *Manager) sendNotification(notification []byte, project *datamodel.ProjectDB) {
	nm.Log("Sending Notification")
	clients := nm.Hub.GetClients()
	nm.Log(fmt.Sprintf("clients : %+v", clients))
	for c := range clients {
		if c.Admin {
			nm.Hub.SendToClient(c, notification)
		} else {
			for _, user := range project.Owners {
				if c.User == user {
					nm.Hub.SendToClient(c, notification)
				}
			}
			for _, user := range project.Members {
				if c.User == user {
					nm.Hub.SendToClient(c, notification)
				}
			}
		}

	}
}

/*
CreateAnnoncement message in the database
*/
func (nm *Manager) CreateAnnoncement(message string) error {
	currentTime := time.Now()
	notification := dm.Notification{guuid.New().String(), currentTime.Format("2006.01.02 15:04:05"), "Global message", "", "project", message, 0, "Success", "ANNOUCEMENT"}
	b, err := json.Marshal(notification)
	if err != nil {
		nm.Log(err.Error())
		return err
	}
	nm.Dao.Set("Annoucement", string(b))
	// TODO Maybe broadcast the announcement ?
	return nil
}
func (nm *Manager) BroadcastMessage(message string, user string) {
	currentTime := time.Now()
	clients := nm.Hub.GetClients()
	notification := dm.Notification{guuid.New().String(), currentTime.Format("2006.01.02 15:04:05"), "Global message", user, "project", message + " has been validated", 0, "Success", "NOTIFICATION"}
	b, _ := json.Marshal(notification)
	for c := range clients {
		nm.Hub.SendToClient(c, b)
	}
}

func (nm *Manager) Log(message string) {
	go func() {
		currentTime := time.Now()
		fmt.Printf("%s \n", message)
		clients := nm.Hub.GetClients()
		notification := dm.Notification{guuid.New().String(), currentTime.Format("2006.01.02 15:04:05"), "", "", "", message, 0, "", "LOG"}
		b, _ := json.Marshal(notification)
		for c := range clients {
			if c.Admin {
				nm.Hub.SendToClient(c, b)
			}
		}
	}()
}
