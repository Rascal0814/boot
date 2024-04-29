package project

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/AlecAivazis/survey/v2"
	"github.com/fatih/color"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var DefaultRepo = "https://github.com/Rascal0814/boot-template.git"

// options 创建新项目所需要使用的参数
var options struct {
	repo       string        // 模板仓库的地址
	module     string        // Go项目模板
	timeout    time.Duration // 创建项目超时时间
	withoutGit bool          // 是否要生成Git仓库
}

// CommandNew 创建一个新项目
func CommandNew() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "new",
		Short: "create a service project using the repository template",
		PreRun: func(cmd *cobra.Command, args []string) {
			if len(options.repo) == 0 {
				options.repo = DefaultRepo
			}
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			_, cancel := context.WithTimeout(context.Background(), options.timeout)
			defer cancel()

			var projectName string
			if len(args) == 0 {
				prompt := &survey.Input{
					Message: "What is project name ?",
				}

				err := survey.AskOne(prompt, &projectName)
				if err != nil || len(projectName) == 0 {
					return nil
				}
			} else {
				projectName = args[0]
			}

			if len(options.module) == 0 {
				prompt := &survey.Input{
					Message: "What is golang module name (without project name) ?",
				}

				_ = survey.AskOne(prompt, &options.module)
			}

			if len(options.module) != 0 {
				options.module = strings.TrimRight(options.module, "/")
				options.module += "/"
			}

			return createProject(context.Background(), projectName)
		},
	}

	cmd.PersistentFlags().StringVar(&options.repo, "repo", "", "the address of the service template")
	cmd.PersistentFlags().StringVar(&options.module, "module", "", "the name of the golang project")
	cmd.PersistentFlags().BoolVar(&options.withoutGit, "without-git", true, "without git-vcs init")

	return cmd

}

func createProject(ctx context.Context, name string) error {
	projectDir, err := filepath.Abs(name)
	if err != nil {
		return err
	}

	stat, err := os.Stat(projectDir)
	if err != nil && !os.IsNotExist(err) {
		return err
	} else if stat != nil && !stat.IsDir() {
		return errors.Errorf("the project(%s) already exists", name)
	} else if stat != nil && stat.IsDir() {
		dir, err := os.Open(projectDir)
		if err != nil {
			return err
		}
		defer func() { _ = dir.Close() }()

		fis, err := dir.Readdir(1)
		if err != nil {
			return err
		}

		if len(fis) != 0 {
			return errors.Errorf("the project(%s) not empty", name)
		}
	} else if stat == nil || os.IsNotExist(err) {
		if err = os.MkdirAll(projectDir, fs.FileMode(0755)); err != nil {
			return err
		}
	}

	defer func() {
		if err != nil {
			_ = os.RemoveAll(projectDir)
		}
	}()

	fmt.Printf("🚀 Creating service %s, please wait a moment.\n", name)
	for file := range clone(ctx, options.repo) {
		if file.Error != nil {
			return file.Error
		}

		err = renderAndWrite(filepath.Join(projectDir, file.RelPath), file.Content, buildParams(name))
		if err != nil {
			return err
		}
	}

	if !options.withoutGit {
		if err = initRepo(projectDir); err != nil {
			return err
		}
	}

	fmt.Printf("\n🍺 Project creation succeeded %s\n", color.GreenString(name))
	return nil
}

// renderAndWrite 渲染并生成对应的项目目录结构
func renderAndWrite(dst string, src io.ReadCloser, params map[string]string) error {
	// do not forget close file
	defer func() { _ = src.Close() }()

	data, err := ioutil.ReadAll(src)
	if err != nil {
		return err
	}

	dst = string(render([]byte(dst), params))
	if err = os.MkdirAll(filepath.Dir(dst), fs.FileMode(0755)); err != nil {
		return err
	}

	fp, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer func() { _ = fp.Close() }()

	if isBinary(data) {
		_, err = fp.Write(data)
		return err
	}

	_, err = fp.Write(render(data, params))
	return err
}

// isBinary 检查文件是否是二进制文件
func isBinary(data []byte) bool {
	for _, b := range data {
		if b == 0 {
			return true
		}
	}
	return false
}

// render 渲染特定的内容
func render(data []byte, params map[string]string) []byte {
	for k, v := range params {
		data = bytes.ReplaceAll(data, []byte(k), []byte(v))
	}
	return data
}

// buildParams 构建模板参数
func buildParams(name string) map[string]string {
	return map[string]string{
		"hml":      name,
		"example/": options.module,
	}
}
