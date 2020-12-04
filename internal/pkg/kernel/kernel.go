package kernel

import (
	"fmt"
	"github.com/hpcng/warewulf/internal/pkg/config"
	"github.com/hpcng/warewulf/internal/pkg/errors"
	"github.com/hpcng/warewulf/internal/pkg/util"
	"github.com/hpcng/warewulf/internal/pkg/wwlog"
	"os"
	"os/exec"
	"path"
)

func Build(kernelVersion string) error {

	kernelImage := "/boot/vmlinuz-" + kernelVersion
	kernelDrivers := "/lib/modules/" + kernelVersion
	kernelDestination := config.KernelImage(kernelVersion)
	driversDestination := config.KmodsImage(kernelVersion)

	// Create the destination paths just in case it doesn't exist
	os.MkdirAll(path.Dir(kernelDestination), 0755)
	os.MkdirAll(path.Dir(driversDestination), 0755)

	if util.IsFile(kernelImage) == false {
		return errors.New("Could not locate kernel image: " + kernelImage)
	}

	if util.IsDir(kernelDrivers) == false {
		return errors.New("Could not locate kernel drivers: " + kernelDrivers)
	}

	if _, err := os.Stat(kernelImage); err == nil {
		err := util.CopyFile(kernelImage, kernelDestination)
		if err != nil {
			return err
		}
		wwlog.Printf(wwlog.INFO, "%-45s: Done\n", "vmlinuz-"+kernelVersion)
	}

	if _, err := os.Stat(kernelDrivers); err == nil {
		cmd := fmt.Sprintf("cd /; find .%s | cpio --quiet -o -H newc -F \"%s\"", kernelDrivers, driversDestination)
		err := exec.Command("/bin/sh", "-c", cmd).Run()
		if err != nil {
			return err
		}
		wwlog.Printf(wwlog.INFO, "%-45s: Done\n", "kmods-"+kernelVersion+".img")
	}

	return nil
}
