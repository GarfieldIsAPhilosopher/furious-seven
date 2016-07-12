package commands

import "fmt"

type AppCmd struct {
	GUID bool `long:"guid"`
}

func (cmd AppCmd) Execute(args []string) error {
	fmt.Println("CMD V2")
	return nil
	// if len(args) != 1 {
	// 	return errors.New("Incorrect Usage. Requires an argument")
	// }
	// appName := args[0]

	// config := config.NewConfig()
	// err := requirements.Login(config)
	// if err != nil {
	// 	return err
	// }

	// err = requirements.TargetedSpace(config)
	// if err != nil {
	// 	return err
	// }

	// do api version check and get appropriate client
	// actor := v3action.NewActor(api.NewClient())
	// apps, err := actor.ListApps(nil, []string{appName}, nil, []string{config.SpaceGUID})
	// if err != nil {
	// 	return err
	// }
	// if len(apps) > 1 {
	// 	return errors.New("too many apps were found with that name - this should never happen")
	// }

	// if cmd.GUID {
	// 	fmt.Println(apps[0].GUID)
	// } else {
	// 	// this needs some sort of api version specific presenter
	// 	fmt.Println(presenter.PresentApp(apps[0]))
	// }

	// // api.ListApps()
	// return nil
}
