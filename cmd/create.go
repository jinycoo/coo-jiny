/**------------------------------------------------------------**
 * @filename cmd/create.go
 * @author   jiny - caojingyin@jinycoo.com
 * @version  1.0.0
 * @date     2017/07/02 12:12
 * @desc     cmd-create - summary
 **------------------------------------------------------------**/

package cmd

import (
	"fmt"
	"log"
	"os"
	"os/user"
	"path"
	"strings"
	"time"

	"github.com/jinycoo/jiny/project"
	"github.com/spf13/cobra"
)

const AppPath = "app"

func addCreate(cmd *cobra.Command) {
	create := &cobra.Command{
		Use:   "new PROJECT_NAME",
		Short: "创建新项目",
		Long:  `快速创建基于Jinygo的Golang项目，你只需要关注业务实现就好，其他一切给你搞定！`,
		Example: `
  # 创建新项目 jiny new project_name -o owner -m module -p project_path -d company.com
  # project_name 最好为单个有意义的单词：
  #   jiny new member -o jinycoo -m service -p $HOME -d company.com
  # 其中有三个参数，-o 和 -m
  # -o author         项目创建人 注册创建人姓名 -o 后最好是自己姓名全拼，有利于跟公司邮箱一致 Default: os current username
  # -m module         项目类型模块名称 Default：service
  # -p project_path   项目父级目录 Default：pwd
  # -d domain         项目主域名 Default: project_path 中最后一层文件夹（小写）名称
  # 依据规则可选择：： 后台 = admin  接口 = interface  服务 = service`,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				log.Fatalf("required project name")
			}
			project.P.Name = args[0]
			project.P.Date = time.Now().Format("2006-01-02 15:04")

			if len(project.P.Owner) == 0 {
				u, _ := user.Current()
				project.P.Owner = u.Username
			}

			if len(project.P.RootPath) == 0 {
				pwd, _ := os.Getwd()
				project.P.RootPath = pwd
			}
			if len(project.P.Domain) == 0 {
				project.P.Domain = strings.ToLower(path.Base(project.P.RootPath))
			}

			project.P.Path = path.Join(project.P.RootPath, AppPath, project.P.Module, project.P.Name)

			// creata a project
			if err := project.Create(); err != nil {
				log.Fatalf("create new project err (%v)", err)
			}
			fmt.Printf("Project: %s\n", project.P.Name)
			fmt.Printf("Owner: %s\n", project.P.Owner)
			fmt.Printf("Module Name: %s\n", project.P.Module)
			fmt.Printf("WithGRPC: %t\n", project.P.WithGRPC)
			fmt.Printf("Directory: %s\n\n", project.P.Path)
			fmt.Println("The application has been created.")
		},
	}
	AddProjectCOptions(create, &project.P)
	cmd.AddCommand(create)
}

func AddProjectCOptions(cmd *cobra.Command, p *project.Project) {
	cmd.Flags().StringVarP(&p.Module, "module", "m", "service", "project type name for project")
	cmd.Flags().StringVarP(&p.Owner, "owner", "o", "", "project creator for create project")
	cmd.Flags().StringVarP(&p.RootPath, "path", "p", "", "project root path for create project")
	cmd.Flags().StringVarP(&p.Domain, "domain", "d", "", "domain for project")
}
