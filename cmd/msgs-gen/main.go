// main package.
package main

import (
	"fmt"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
)

var tplPackage = template.Must(template.New("").Parse(
	`// Package {{ .PkgName }} contains message definitions.
//autogenerated:yes
//nolint:revive
package {{ .PkgName }}
`))

func shellCommand(cmdstr string) error {
	fmt.Fprintf(os.Stderr, "%s\n", cmdstr)
	cmd := exec.Command("sh", "-c", cmdstr)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	return cmd.Run()
}

func writeTemplate(fpath string, tpl *template.Template, args map[string]interface{}) error {
	f, err := os.Create(fpath)
	if err != nil {
		return err
	}
	defer f.Close()

	return tpl.Execute(f, args)
}

func processDir(name string, dir string) error {
	fmt.Fprintf(os.Stderr, "[%s]\n", name)

	os.Mkdir(filepath.Join("pkg", "msgs", name), 0o755)

	err := writeTemplate(
		filepath.Join("pkg", "msgs", name, "package.go"),
		tplPackage,
		map[string]interface{}{
			"PkgName": name,
		})
	if err != nil {
		return err
	}

	err = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		switch {
		case strings.HasSuffix(info.Name(), ".msg"):
			outpath := strings.ToLower(filepath.Join("pkg", "msgs", name, "Msg"+strings.TrimSuffix(info.Name(), ".msg")+".go"))
			err = shellCommand(fmt.Sprintf("go run ./cmd/msg-import --gopackage=%s --rospackage=%s %s > %s",
				name,
				name,
				path,
				outpath))
			if err != nil {
				os.Remove(outpath)
				return err
			}

		case strings.HasSuffix(info.Name(), ".srv"):
			outpath := strings.ToLower(filepath.Join("pkg", "msgs", name, "Srv"+strings.TrimSuffix(info.Name(), ".srv")+".go"))
			err = shellCommand(fmt.Sprintf("go run ./cmd/srv-import --gopackage=%s --rospackage=%s %s > %s",
				name,
				name,
				path,
				outpath))
			if err != nil {
				os.Remove(outpath)
				return err
			}

		case strings.HasSuffix(info.Name(), ".action"):
			outpath := strings.ToLower(filepath.Join("pkg", "msgs", name,
				"Action"+strings.TrimSuffix(info.Name(), ".action")+".go"))
			err = shellCommand(fmt.Sprintf("go run ./cmd/action-import --gopackage=%s --rospackage=%s %s > %s",
				name,
				name,
				path,
				outpath))
			if err != nil {
				os.Remove(outpath)
				return err
			}
		}

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

func processRepo(ur string, branch string) error {
	dir, err := os.MkdirTemp("", "goroslib")
	if err != nil {
		return err
	}
	defer os.RemoveAll(dir)

	_, err = git.PlainClone(dir, false, &git.CloneOptions{
		URL:           ur,
		Depth:         1,
		ReferenceName: plumbing.NewBranchReferenceName(branch),
	})
	if err != nil {
		return err
	}

	// find folders which contain a "msg", "srv" or "action" subfolder
	paths := make(map[string]struct{})
	err = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() &&
			(info.Name() == "msg" || info.Name() == "srv" || info.Name() == "action") {
			paths[filepath.Dir(path)] = struct{}{}
			return nil
		}

		return nil
	})
	if err != nil {
		return err
	}

	u, _ := url.Parse(ur)

	for path := range paths {
		err := processDir(filepath.Base(filepath.Join(u.Path, path[len(dir):])), path)
		if err != nil {
			return err
		}
	}

	return nil
}

func run() error {
	err := shellCommand("rm -rf pkg/msgs/*/")
	if err != nil {
		return err
	}

	done := make(chan error)
	count := 0

	for _, entry := range []struct {
		ur     string
		branch string
	}{
		{
			ur:     "https://github.com/ros/std_msgs",
			branch: "kinetic-devel",
		},
		{
			ur:     "https://github.com/ros/ros_comm_msgs",
			branch: "kinetic-devel",
		},
		{
			ur:     "https://github.com/ros/common_msgs",
			branch: "noetic-devel",
		},
		{
			ur:     "https://github.com/ros-drivers/ackermann_msgs",
			branch: "master",
		},
		{
			ur:     "https://github.com/ros-drivers/audio_common",
			branch: "master",
		},
		{
			ur:     "https://github.com/ros-drivers/velodyne",
			branch: "master",
		},
		{
			ur:     "https://github.com/ros-controls/control_msgs",
			branch: "kinetic-devel",
		},
		{
			ur:     "https://github.com/ros-perception/vision_msgs",
			branch: "noetic-devel",
		},
		{
			ur:     "https://github.com/ros/actionlib",
			branch: "noetic-devel",
		},
		{
			ur:     "https://github.com/mavlink/mavros",
			branch: "master",
		},
		{
			ur:     "https://github.com/ros-geographic-info/geographic_info",
			branch: "master",
		},
		{
			ur:     "https://github.com/ros-geographic-info/unique_identifier",
			branch: "master",
		},
		{
			ur:     "https://github.com/ros/geometry",
			branch: "noetic-devel",
		},
		{
			ur:     "https://github.com/ros/geometry2",
			branch: "noetic-devel",
		},
	} {
		count++
		go func(ur string, branch string) {
			done <- processRepo(ur, branch)
		}(entry.ur, entry.branch)
	}

	for i := 0; i < count; i++ {
		err := <-done
		if err != nil {
			return err
		}
	}

	return nil
}

func main() {
	err := run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERR: %s\n", err)
		os.Exit(1)
	}
}
