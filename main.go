package main

import (
	
)

func main() {

    // read bot token from .env file
	err := godotenv.Load("./test_slack.env")
	if err != nil {
		log.Fatal().Msg("Error loading .env file")
	}

	// Instantiate slack socket mode client
	client, err := drivers.ConnectToSlackViaSocketmode()
	if err != nil {
		log.Error().
			Str("error", err.Error()).
			Msg("Unable to connect to slack")

		os.Exit(1)
	}

	// Inject deps in event handler
	socketmodeHandler := socketmode.NewsSocketmodeHandler(client)

    // Inject deps to Controller
	controllers.NewGreetingController(socketmodeHandler)

	socketmodeHandler.RunEventLoop()
}