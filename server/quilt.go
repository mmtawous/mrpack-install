package server

import (
	"errors"
	"fmt"
	"github.com/nothub/mrpack-install/files"
	"github.com/nothub/mrpack-install/maven"
	"github.com/nothub/mrpack-install/web"
	"os"
	"os/exec"
	"path"
)

type QuiltInstaller struct {
	MinecraftVersion string
	QuiltVersion     string
}

func (inst *QuiltInstaller) Install(serverDir string, serverFile string) error {
	meta, err := maven.FetchMetadata("https://maven.quiltmc.org/repository/release/org/quiltmc/quilt-installer/maven-metadata.xml")
	if err != nil {
		return err
	}
	quiltInstallerUrl := "https://maven.quiltmc.org/repository/release/org/quiltmc/quilt-installer/" + meta.Versioning.Release + "/quilt-installer-" + meta.Versioning.Release + ".jar"

	installer, err := web.DefaultClient.DownloadFile(quiltInstallerUrl, ".", "")
	if err != nil {
		return err
	}

	cmd := exec.Command("java", "-jar", installer, "install", "server", inst.MinecraftVersion, "--install-dir="+serverDir, "--create-scripts", "--download-server")
	fmt.Println("Executing command:", cmd.String())
	err = cmd.Run()
	if err != nil {
		return err
	}

	if !files.IsFile(path.Join(serverDir, "server.jar")) {
		return errors.New("server.jar not found")
	}

	if serverFile != "" {
		err = os.Rename(path.Join(serverDir, "server.jar"), path.Join(serverDir, serverFile))
		if err != nil {
			return err
		}
	}

	return nil
}
