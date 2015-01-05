package marathon

import (
	"testing"

	log "github.com/Sirupsen/logrus"
)

var (
	appID = "nginx"
)

func TestCreateApp(t *testing.T) {
	if respApp, err := client.CreateApp(&Application{
		ID:        appID,
		Instances: 1,
		CPUs:      0.1,
		Mem:       128.0,
		Container: &Container{
			Type: "DOCKER",
			Docker: &Docker{
				Image: "nginx:latest",
			},
		},
	}); err != nil {
		t.Error(err)
	} else {
		log.Debugf("Response App: %#v\n", respApp)
	}
}

func TestGetApps(t *testing.T) {
	if apps, err := client.GetApps(); err != nil {
		t.Error(err)
	} else {
		for _, app := range apps {
			log.Debugf("Response App: %#v\n", app)
		}
	}
}

func TestGetAppsWithParams(t *testing.T) {
	if apps, err := client.GetAppsWithParams(&GetAppsParams{Embed: AppsTasks}); err != nil {
		t.Error(err)
	} else {
		for _, app := range apps {
			log.Debugf("Response App: %#v\n", app)
		}
	}
}

func TestGetApp(t *testing.T) {
	if app, err := client.GetApp(appID); err != nil {
		t.Error(err)
	} else {
		log.Debugf("Response App: %#v\n", app)
	}

	if _, err := client.GetApp("notfound"); err == nil {
		t.Error("found app that not exists")
	}
}

func TestGetAppVersion(t *testing.T) {
	if versions, err := client.GetAppVersions(appID); err != nil {
		t.Error(err)
	} else {
		log.Debugf("Versions: %#v\n", versions)
	}

	if _, err := client.GetAppVersions("notfound"); err == nil {
		t.Error("found app's versions that not exists")
	}
}

func TestGetAppByVersion(t *testing.T) {
	versions, err := client.GetAppVersions(appID)
	if err != nil {
		t.Error(err)
	}

	if app, err := client.GetAppByVersion(appID, versions[0]); err != nil {
		t.Error(err)
	} else {
		log.Debugf("Response App: %#v\n", app)
	}

	if _, err := client.GetAppByVersion(appID, "notfoundversion"); err == nil {
		t.Error("found app that not exists")
	}
}

func TestUpdateApp(t *testing.T) {
	if deploymentID, version, err := client.UpdateApp(appID, &Application{
		Instances: 2,
		CPUs:      0.2,
		Mem:       256.0,
	}); err != nil {
		t.Error(err)
	} else {
		log.Debugf("deploymentID: %s\n", deploymentID)
		log.Debugf("version: %s\n", version)
	}
}

func TestUpdateAppWithParams(t *testing.T) {
	if deploymentID, version, err := client.UpdateAppWithParams(appID, &Application{
		Instances: 1,
		CPUs:      0.1,
		Mem:       128.0,
	}, true); err != nil {
		t.Error(err)
	} else {
		log.Debugf("deploymentID: %s\n", deploymentID)
		log.Debugf("version: %s\n", version)
	}
}

func TestGetAppTasks(t *testing.T) {
	if tasks, err := client.GetAppTasks(appID); err != nil {
		t.Error(err)
	} else {
		log.Debugf("Tasks: %#v\n", tasks)
	}
}

func TestKillTasks(t *testing.T) {
	if _, err := client.KillTasks("notfound"); err == nil {
		t.Error("found app that not exists")
	}

	if tasks, err := client.KillTasks(appID); err != nil {
		t.Error(err)
	} else {
		log.Debugf("Tasks: %#v\n", tasks)
	}
}

func TestKillTasksWithParams(t *testing.T) {
	if _, err := client.KillTasks("notfound"); err == nil {
		t.Error("found app that not exists")
	}

	if tasks, err := client.KillTasksWithParams(appID, &KillTasksParams{
		Scale: true,
	}); err != nil {
		t.Error(err)
	} else {
		log.Debugf("Tasks: %#v\n", tasks)
	}
}

func TestKillTask(t *testing.T) {
	if _, err := client.KillTask(appID, "notfound"); err == nil {
		t.Error("found task that not exists")
	}

	tasks, err := client.GetTasks()
	if err != nil {
		t.Error(err)
	}

	if task, err := client.KillTask(appID, tasks[0].ID); err != nil {
		t.Error(err)
	} else {
		log.Debugf("Task: %#v\n", task)
	}
}

func TestKillTaskWithParams(t *testing.T) {
	if _, err := client.KillTaskWithParams(appID, "notfound", false); err == nil {
		t.Error("found task that not exists")
	}

	tasks, err := client.GetTasks()
	if err != nil {
		t.Error(err)
	}

	if task, err := client.KillTaskWithParams(appID, tasks[0].ID, true); err != nil {
		t.Error(err)
	} else {
		log.Debugf("Task: %#v\n", task)
	}
}

func TestDestroyApp(t *testing.T) {
	if _, _, err := client.DestroyApp("notfound"); err == nil {
		t.Error("found app that not exists")
	}

	if deploymentID, version, err := client.DestroyApp(appID); err != nil {
		t.Error(err)
	} else {
		log.Debugf("DeploymentID: %s\n", deploymentID)
		log.Debugf("Version: %s\n", version)
	}
}
