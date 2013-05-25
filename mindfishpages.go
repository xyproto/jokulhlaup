package main

import (
	"strconv"
	"time"

	. "github.com/xyproto/genericsite"
	. "github.com/xyproto/siteengines"
)

// The default settings for Mindfish content pages
func MindfishBaseCP(state *UserState) *ContentPage {
	cp := DefaultCP(state)
	cp.Title = "MindFish Hi"
	cp.Subtitle = "There"

	y := time.Now().Year()

	// TODO: Use templates for the footer, for more accurate measurment of the time made to generate the page
	cp.FooterText = "Alexander Rødseth, " + strconv.Itoa(y)

	cp.Url = "/" // Is replaced when the contentpage is published

	cp.ColorScheme = NewMindfishColorScheme()

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

// Returns a MindfishBaseCP with the contentTitle set
func MindfishBaseTitleCP(contentTitle string, userState *UserState) *ContentPage {
	cp := MindfishBaseCP(userState)
	cp.ContentTitle = contentTitle
	return cp
}

func OverviewCP(userState *UserState, url string) *ContentPage {
	cp := MindfishBaseCP(userState)
	cp.ContentTitle = "About"
	cp.ContentHTML = `Yet to come`
	cp.Url = url
	return cp
}

func TextCP(userState *UserState, url string) *ContentPage {
	apc := MindfishBaseCP(userState)
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
func ServeMindfish(userState *UserState, jquerypath string) MenuEntries {
	cps := []ContentPage{
		*OverviewCP(userState, "/"),
		*TextCP(userState, "/text"),
		*LoginCP(MindfishBaseCP, userState, "/login"),
		*RegisterCP(MindfishBaseCP, userState, "/register"),
	}

	menuEntries := Cps2MenuEntries(cps)

	// template content generator
	tvgf := DynamicMenuFactoryGenerator(menuEntries)

	//ServeSearchPages(MindfishBaseCP, userState, cps, MindfishBaseCP(userState).ColorScheme, tvgf(userState))
	ServeSite(MindfishBaseCP, userState, cps, tvgf, jquerypath)

	return menuEntries
}

func NewMindfishColorScheme() *ColorScheme {
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
