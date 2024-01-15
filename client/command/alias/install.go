package alias

/*
	Sliver Implant Framework
	Copyright (C) 2021  Bishop Fox

	This program is free software: you can redistribute it and/or modify
	it under the terms of the GNU General Public License as published by
	the Free Software Foundation, either version 3 of the License, or
	(at your option) any later version.

	This program is distributed in the hope that it will be useful,
	but WITHOUT ANY WARRANTY; without even the implied warranty of
	MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
	GNU General Public License for more details.

	You should have received a copy of the GNU General Public License
	along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

import (
	"fmt"
	"github.com/chainreactors/grumble"
	"github.com/chainreactors/malice-network/client/assets"
	"github.com/chainreactors/malice-network/client/console"
	"github.com/chainreactors/malice-network/client/utils"
	"os"
	"path/filepath"
	"strings"
)

// AliasesInstallCmd - Install an alias
func AliasesInstallCmd(ctx *grumble.Context, con *console.Console) {
	aliasLocalPath := ctx.Args.String("path")
	fi, err := os.Stat(aliasLocalPath)
	if os.IsNotExist(err) {
		console.Log.Errorf("alias path '%s' does not exist", aliasLocalPath)
		return
	}
	if !fi.IsDir() {
		InstallFromFile(aliasLocalPath, false, con)
	} else {
		installFromDir(aliasLocalPath, con)
	}
}

// Install an extension from a directory
func installFromDir(aliasLocalPath string, con *console.Console) {
	manifestData, err := os.ReadFile(filepath.Join(aliasLocalPath, ManifestFileName))
	if err != nil {
		console.Log.Errorf("Error reading %s: %s", ManifestFileName, err)
		return
	}
	manifest, err := ParseAliasManifest(manifestData)
	if err != nil {
		console.Log.Errorf("Error parsing %s: %s", ManifestFileName, err)
		return
	}
	installPath := filepath.Join(assets.GetAliasesDir(), filepath.Base(manifest.CommandName))
	if _, err := os.Stat(installPath); !os.IsNotExist(err) {
		console.Log.Infof("Alias '%s' already exists", manifest.CommandName)
		//confirm := false
		// todo rewrite to tea
		//prompt := &survey.Confirm{Message: "Overwrite current install?"}
		//survey.AskOne(prompt, &confirm)
		//if !confirm {
		//	return
		//}
		forceRemoveAll(installPath)
	}

	console.Log.Infof("Installing alias '%s' (%s) ... ", manifest.Name, manifest.Version)
	err = os.MkdirAll(installPath, 0o700)
	if err != nil {
		console.Log.Errorf("Error creating alias directory: %s\n", err)
		return
	}
	err = os.WriteFile(filepath.Join(installPath, ManifestFileName), manifestData, 0o600)
	if err != nil {
		console.Log.Errorf("Failed to write %s: %s\n", ManifestFileName, err)
		forceRemoveAll(installPath)
		return
	}

	for _, cmdFile := range manifest.Files {
		if cmdFile.Path != "" {
			src := filepath.Join(aliasLocalPath, utils.ResolvePath(cmdFile.Path))
			dst := filepath.Join(installPath, utils.ResolvePath(cmdFile.Path))
			err := utils.CopyFile(src, dst)
			if err != nil {
				console.Log.Errorf("Error copying file '%s' -> '%s': %s\n", src, dst, err)
				forceRemoveAll(installPath)
				return
			}
		}
	}

	console.Log.Infof("done!\n")
}

// Install an extension from a .tar.gz file
func InstallFromFile(aliasGzFilePath string, autoOverwrite bool, con *console.Console) *string {
	manifestData, err := utils.ReadFileFromTarGz(aliasGzFilePath, fmt.Sprintf("./%s", ManifestFileName))
	if err != nil {
		console.Log.Errorf("Failed to read %s from '%s': %s\n", ManifestFileName, aliasGzFilePath, err)
		return nil
	}
	manifest, err := ParseAliasManifest(manifestData)
	if err != nil {
		console.Log.Errorf("Failed to parse %s: %s\n", ManifestFileName, err)
		return nil
	}
	installPath := filepath.Join(assets.GetAliasesDir(), filepath.Base(manifest.CommandName))
	if _, err := os.Stat(installPath); !os.IsNotExist(err) {
		if !autoOverwrite {
			console.Log.Infof("Alias '%s' already exists\n", manifest.CommandName)
			//confirm := false
			//prompt := &survey.Confirm{Message: "Overwrite current install?"}
			//survey.AskOne(prompt, &confirm)
			//if !confirm {
			//	return nil
			//}
		}
		forceRemoveAll(installPath)
	}

	console.Log.Infof("Installing alias '%s' (%s) ... ", manifest.Name, manifest.Version)
	err = os.MkdirAll(installPath, 0o700)
	if err != nil {
		console.Log.Errorf("Failed to create alias directory: %s\n", err)
		return nil
	}
	err = os.WriteFile(filepath.Join(installPath, ManifestFileName), manifestData, 0o600)
	if err != nil {
		console.Log.Errorf("Failed to write %s: %s\n", ManifestFileName, err)
		forceRemoveAll(installPath)
		return nil
	}
	for _, aliasFile := range manifest.Files {
		if aliasFile.Path != "" {
			err := installArtifact(aliasGzFilePath, installPath, aliasFile.Path, con)
			if err != nil {
				console.Log.Errorf("Failed to install file: %s\n", err)
				forceRemoveAll(installPath)
				return nil
			}
		}
	}
	console.Log.Console("done!\n")
	return &installPath
}

func installArtifact(aliasGzFilePath string, installPath, artifactPath string, con *console.Console) error {
	data, err := utils.ReadFileFromTarGz(aliasGzFilePath, fmt.Sprintf("./%s", strings.TrimPrefix(artifactPath, string(os.PathSeparator))))
	if err != nil {
		return err
	}
	if len(data) == 0 {
		return fmt.Errorf("empty file %s", artifactPath)
	}
	localArtifactPath := filepath.Join(installPath, utils.ResolvePath(artifactPath))
	artifactDir := filepath.Dir(localArtifactPath)
	if _, err := os.Stat(artifactDir); os.IsNotExist(err) {
		os.MkdirAll(artifactDir, 0o700)
	}
	err = os.WriteFile(localArtifactPath, data, 0o600)
	if err != nil {
		return err
	}
	return nil
}

func forceRemoveAll(rootPath string) {
	utils.ChmodR(rootPath, 0o600, 0o700)
	os.RemoveAll(rootPath)
}