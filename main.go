package main

import (
	"context"
	"fmt"

	"github.com/afarooqi8/trigger"
	"github.com/project-flogo/contrib/activity/log"
	"github.com/project-flogo/core/activity"
	"github.com/project-flogo/core/api"
	"github.com/project-flogo/core/engine"
)

func main() {

	app := myApp()

	e, err := api.NewEngine(app)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	engine.RunEngine(e)
}

func myApp() *api.App {
	app := api.NewApp()

	// trg := app.NewTrigger(&trigger.Trigger{}, &trigger.Settings{Port: "8080"})
	// h, _ := trg.NewHandler(&trigger.HandlerSettings{Event: "json"})
	trg := app.NewTrigger(&trigger.Trigger{}, &trigger.HandlerSettings{})
	h, _ := trg.NewHandler(&trigger.HandlerSettings{})
	h.NewAction(RunActivities)
	fmt.Println("hello 3 !!")

	//store in map to avoid activity instance recreation
	logAct, _ := api.NewActivity(&log.Activity{})
	activities = map[string]activity.Activity{"log": logAct}

	return app
}

var activities map[string]activity.Activity

func RunActivities(ctx context.Context, inputs map[string]interface{}) (map[string]interface{}, error) {

	fmt.Println("just got the output !!", inputs)

	m := make(map[string]interface{})
	m["message"] = "Bonjour"
	return m, nil
}
