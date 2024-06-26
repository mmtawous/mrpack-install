package server

import (
	"errors"
	"github.com/nothub/mrpack-install/web"
	"log"
	"strconv"
)

type PaperInstaller struct {
	MinecraftVersion string
}

func (inst *PaperInstaller) Install(serverDir string, serverFile string) error {
	var response struct {
		Builds []struct {
			Id        int    `json:"build"`
			Channel   string `json:"channel"`
			Downloads struct {
				Application struct {
					Name   string `json:"name"`
					Sha256 string `json:"sha256"`
				} `json:"application"`
			} `json:"downloads"`
		} `json:"builds"`
	}

	err := web.DefaultClient.GetJson("https://api.papermc.io/v2/projects/paper/versions/"+
		inst.MinecraftVersion+"/builds", &response, nil)
	if err != nil {
		return err
	}

	for i := range response.Builds {
		i = len(response.Builds) - 1 - i
		if response.Builds[i].Channel == "default" {
			u := "https://api.papermc.io/v2/projects/paper/versions/" + inst.MinecraftVersion +
				"/builds/" + strconv.Itoa(response.Builds[i].Id) +
				"/downloads/" + response.Builds[i].Downloads.Application.Name
			file, err := web.DefaultClient.DownloadFile(u, serverDir, serverFile)
			if err != nil {
				return err
			}

			log.Println("Server jar downloaded to:", file)
			return nil
		}
	}

	return errors.New("no stable paper release found")
}
