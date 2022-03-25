package views

import (
	"embed"
	"encoding/json"
	"io/ioutil"

	"github.com/slack-go/slack"
)

//go:embed greetingViewsAssets/*
var greetingAssets embed.FS

func GreetingMessage(user string) []slack.Block {

	// we need a stuct to hold template arguments
	type args struct {
		User string
	}

	tpl := renderTemplate(greetingAssets, "greetingViewsAssets/greeting.json", args{User: user})

	// we convert the View into a message struct
	view := slack.Msg{}

  // dump the renderd template into the view
	str, _ := ioutil.ReadAll(&tpl)
	json.Unmarshal(str, &view)

	// We only return the block because of the way the PostEphemeral function works
	// we are going to use slack.MsgOptionBlocks in the controller
	return view.Blocks.BlockSet
}
