package main

import (
	paper "Aphtap/papermc-downloader/paper"
	"fmt"
)

func main() {
	latestVersion := paper.GetLatestVersion()
	fmt.Println(latestVersion)
	latestBuild := paper.GetLatestBuild(latestVersion)
	fmt.Println(latestBuild)
	name := paper.GetLatestDownload(latestVersion, latestBuild)
	fmt.Println(name)
	paper.Get(latestVersion, latestBuild, name)
}
