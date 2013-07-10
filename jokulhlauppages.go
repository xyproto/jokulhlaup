package main

import (
	"strconv"
	"time"

	. "github.com/xyproto/genericsite"
	. "github.com/xyproto/siteengines"
)

// TODO: Font for headline: IM Fell Double Pica SC

// The default settings for Jokulhlaup content pages
func JokulhlaupBaseCP(state *UserState) *ContentPage {
	cp := DefaultCP(state)
	cp.Title = "Jøkulhlaup ."
	cp.Subtitle = "."

	y := time.Now().Year()

	// TODO: Use templates for the footer, for more accurate measurment of the time made to generate the page
	cp.FooterText = "Alexander Rødseth, " + strconv.Itoa(y)

	cp.Url = "/" // Is replaced when the contentpage is published

	cp.ColorScheme = NewJokulhlaupColorScheme()

	// Behind the text
	//cp.BgImageURL = "/img/nasty_fabric.png"
	//cp.BgImageURL = "/img/cloth_alike.png"
	cp.BgImageURL = "/img/strange_bullseyes.png"
	//cp.BgImageURL = "/img/rough_diagonal.png"
	cp.StretchBackground = false

	// Behind the menu
	//cp.BackgroundTextureURL = "/img/bg2.png"
	//cp.BackgroundTextureURL = "/img/simple_dashed.png"
	//cp.BackgroundTextureURL = "/img/grey.png"
	//cp.BackgroundTextureURL = "/img/pw_maze_black.png"
	//cp.BackgroundTextureURL = "/img/black_twill.png"
	cp.BackgroundTextureURL = "/img/dark_wood.png"
	//cp.BackgroundTextureURL = "/img/hixs_pattern_evolution.png"
	//ps_neutral.png"

	cp.SearchBox = false

	return cp
}

// Returns a JokulhlaupBaseCP with the contentTitle set
func JokulhlaupBaseTitleCP(contentTitle string, userState *UserState) *ContentPage {
	cp := JokulhlaupBaseCP(userState)
	cp.ContentTitle = contentTitle
	return cp
}

func OverviewCP(userState *UserState, url string) *ContentPage {
	cp := JokulhlaupBaseCP(userState)
	cp.ContentTitle = `<a href="http://snl.no/j%C3%B8kulhlaup">About</a>`
	cp.ContentHTML = `<img src="http://upload.wikimedia.org/wikipedia/commons/b/b3/Hubbard_Glacier_August_14.2002.jpg">`
	cp.Url = url
	return cp
}

func TextCP(userState *UserState, url string) *ContentPage {
	apc := JokulhlaupBaseCP(userState)
	apc.ContentTitle = "Text"
	apc.ContentHTML = `<p id='textparagraph'>Hi<br/>there<br/></p>`
	apc.Url = url
	return apc
}

// This is where the possibilities for the menu are listed
func Cps2MenuEntries(cps []ContentPage) MenuEntries {
	links := []string{
		"About:/",
		"Log in:/login",
		"Register:/register",
		"Log out:/logout",
		"Admin:/admin",
		"Wiki:/wiki",
	}
	return Links2menuEntries(links)
}

// Routing for the archlinux.no webpage
// Admin, search and user management is already provided
func ServeJokulhlaup(userState *UserState, jquerypath string) MenuEntries {
	cps := []ContentPage{
		*OverviewCP(userState, "/"),
		*TextCP(userState, "/text"),
		*LoginCP(JokulhlaupBaseCP, userState, "/login"),
		*RegisterCP(JokulhlaupBaseCP, userState, "/register"),
	}

	menuEntries := Cps2MenuEntries(cps)

	// template content generator
	tvgf := DynamicMenuFactoryGenerator(menuEntries)

	//ServeSearchPages(JokulhlaupBaseCP, userState, cps, JokulhlaupBaseCP(userState).ColorScheme, tvgf(userState))
	ServeSite(JokulhlaupBaseCP, userState, cps, tvgf, jquerypath)

	return menuEntries
}

func NewJokulhlaupColorScheme() *ColorScheme {
	var cs ColorScheme
	cs.Darkgray = "#202020"
	cs.Nicecolor = "#d80000"   // bright orange!
	cs.Menu_link = "#c0c0c0"   // light gray
	cs.Menu_hover = "#efefe0"  // light gray, somewhat yellow
	cs.Menu_active = "#ffffff" // white
	cs.Default_background = "#000030"
	cs.TitleText = "#e0e0e0" // The first word of the title text
	return &cs
}
