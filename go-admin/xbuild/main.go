package main

import (
	"archive/zip"
	"fmt"
	"gopkg.in/yaml.v2"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

var (
	BackFold  string
	FrontFold string

	Platform   string
	BuildFile  string
	FileName   string
	ApiVersion string
	// GoBuildTargetFileName 打包后结果信息
	GoBuildTargetFileName string
	GoBuildFold           string
)

type Config struct {
	BackFold  string `yaml:"backFold"`
	FrontFold string `yaml:"frontFold"`

	Platform        string `yaml:"platform"`
	GoBuildFileName string `yaml:"goBuildFileName"`
	ApiVersion      string `yaml:"apiVersion"`
}

func main() {
	readConfig()
	clearFile(".\\" + ApiVersion)
	clearFile(".\\" + ApiVersion + "\\config")
	clearFile(".\\" + ApiVersion + "\\public")
	clearFile(".\\" + ApiVersion + "\\front")
	setGoEnv("GOOS", Platform)
	buildMainGo(BuildFile, FileName, ApiVersion, Platform)
	buildFront()
	scpBuildFile()
	deleteFile()
	zipFormat := ".zip"
	if Platform == "linux" {
		zipFormat = ".tar"
	}
	split := strings.Split(GoBuildTargetFileName, ",")
	zipFolder(".\\"+ApiVersion, split[0]+"_"+Platform+zipFormat)
}

func buildFront() {
	executeStringCommandInDir("../../soybean-admin/", "pnpm", "i")
	executeStringCommandInDir("../../soybean-admin/", "pnpm", "build")
}

func setGoEnv(key, value string) {
	fmt.Printf("设置目标平台%s...\n", value)
	cmd := exec.Command("go", "env", "-w", fmt.Sprintf("%s=%s", key, value))
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Println(fmt.Errorf("无法设置环境变量 %s: %w", key, err))
		return
	}
	fmt.Printf("目标平台设置成功！%s\n", value)
}

func buildMainGo(file string, buildName string, version string, platform string) {
	GoBuildTargetFileName = buildName + "_" + version
	if platform == "windows" {
		GoBuildTargetFileName = GoBuildTargetFileName + ".exe"
	}
	fmt.Printf("开始打包%s...\n", file)

	goFile := "main.go"
	executeStringCommandInDir(BackFold, "go", "work", "sync")
	executeStringCommandInDir(BackFold, "go", "mod", "tidy")
	executeStringCommandInDir(BackFold, "go", "build", "-o", GoBuildTargetFileName, goFile)
	fmt.Printf("打包成功:%s\n", GoBuildTargetFileName)
}

func executeStringCommandInDir(dir, command string, args ...string) {
	// 创建命令对象
	cmd := exec.Command(command, args...)

	// 设置命令的工作目录
	cmd.Dir = dir

	// 将命令的输出连接到当前进程的标准输出和标准错误输出
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// 执行命令
	err := cmd.Run()
	if err != nil {
		fmt.Println(fmt.Errorf("命令执行失败: %w", err))
	}
}

func clearFile(folderPath string) {

	// 检测文件夹是否存在
	_, err := os.Stat(folderPath)
	if err == nil {
		// 文件夹存在，删除文件夹内容
		err := removeAllContents(folderPath)
		if err != nil {
			fmt.Println("清除文件夹内容失败:", err)
			return
		}
		fmt.Println("文件夹内容已清除")
	} else if os.IsNotExist(err) {
		// 文件夹不存在，创建文件夹
		err := os.MkdirAll(folderPath, os.ModePerm)
		if err != nil {
			fmt.Println("创建文件夹失败:", err)
			return
		}
		fmt.Println("文件夹已创建")
	} else {
		// 其他错误，无法确定文件夹是否存在
		fmt.Println("无法确定文件夹是否存在:", err)
		return
	}
}

func scpBuildFile() {
	scpFile(BackFold+GoBuildTargetFileName, ".\\"+ApiVersion)
	copyFolder(GoBuildFold+GoBuildTargetFileName, ".\\"+ApiVersion)
	copyFolder(FrontFold+"\\dist\\", ".\\"+ApiVersion+"\\front")
	copyFolder(BackFold+"\\config\\", ".\\"+ApiVersion+"\\config")
	copyFolder(BackFold+"\\public\\", ".\\"+ApiVersion+"\\public")
}

func deleteFile() {
	DeleteIfExists(BackFold + "\\" + GoBuildTargetFileName)
}

func readConfig() *Config {
	// 读取 YAML 文件内容
	yamlFile, err := os.ReadFile("config.yaml")
	if err != nil {
		log.Fatalf("无法读取 YAML 文件: %v", err)
	}

	// 解析 YAML 文件
	var config Config
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		log.Fatalf("解析 YAML 文件失败: %v", err)
	}

	// 打印配置信息
	fmt.Println("Name:", config)
	BackFold = config.BackFold
	FrontFold = config.FrontFold
	Platform = config.Platform
	FileName = config.GoBuildFileName
	ApiVersion = config.ApiVersion
	return &config

}

// 复制文件夹
func copyFolder(source, destination string) error {
	// 创建目标文件夹
	err := os.MkdirAll(destination, os.ModePerm)
	if err != nil {
		return err
	}

	// 遍历源文件夹中的文件和子文件夹
	err = filepath.Walk(source, func(sourcePath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// 构建目标路径
		relativePath, err := filepath.Rel(source, sourcePath)
		if err != nil {
			return err
		}
		destinationPath := filepath.Join(destination, relativePath)

		if info.IsDir() {
			// 如果是文件夹，创建对应的目标文件夹
			err = os.MkdirAll(destinationPath, os.ModePerm)
			if err != nil {
				return err
			}
		} else {
			// 如果是文件，复制文件内容到目标文件
			err = copyFile(sourcePath, destinationPath)
			if err != nil {
				return err
			}
		}

		return nil
	})

	return err
}

// 复制文件内容
func copyFile(source, destination string) error {
	sourceFile, err := os.Open(source)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destFile, err := os.Create(destination)
	if err != nil {
		return err
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, sourceFile)
	if err != nil {
		return err
	}

	return nil
}

func scpFile(sourceFile string, destinationDir string) {
	// 获取源文件的绝对路径
	sourcePath, err := filepath.Abs(sourceFile)
	if err != nil {
		fmt.Println("获取源文件路径失败:", err)
		return
	}

	// 获取目标目录的绝对路径
	destinationPath, err := filepath.Abs(filepath.Join(destinationDir, filepath.Base(sourceFile)))
	if err != nil {
		fmt.Println("获取目标目录路径失败:", err)
		return
	}

	// 剪切文件
	err = os.Rename(sourcePath, destinationPath)
	if err != nil {
		fmt.Println("剪切文件失败:", err)
		return
	}
	fmt.Println("文件剪切成功")
}

// 递归删除文件夹中的所有内容（包括文件和子文件夹）
func removeAllContents(path string) error {
	err := filepath.Walk(path, func(subPath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if subPath != path {
			err = os.RemoveAll(subPath)
			if err != nil {
				return err
			}
		}
		return nil
	})

	return err
}

// 压缩文件夹
func zipFolder(sourceFolder, destinationFile string) error {
	// 创建目标压缩文件
	zipFile, err := os.Create(destinationFile)
	if err != nil {
		return err
	}
	defer zipFile.Close()

	// 创建 zip.Writer
	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	// 遍历源文件夹中的文件和子文件夹
	err = filepath.Walk(sourceFolder, func(sourcePath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// 构建 zip 文件内部的相对路径
		relativePath, err := filepath.Rel(sourceFolder, sourcePath)
		if err != nil {
			return err
		}

		// 如果是文件夹，创建对应的 zip 文件夹
		if info.IsDir() {
			_, err = zipWriter.CreateHeader(&zip.FileHeader{
				Name:     relativePath + "/",
				Method:   zip.Store, // 存储方式，不进行压缩
				Modified: info.ModTime(),
			})
			if err != nil {
				return err
			}
		} else {
			// 如果是文件，创建对应的 zip 文件并写入内容
			zipFile, err := zipWriter.CreateHeader(&zip.FileHeader{
				Name:     relativePath,
				Method:   zip.Deflate, // 压缩方式
				Modified: info.ModTime(),
			})
			if err != nil {
				return err
			}

			// 打开源文件
			sourceFile, err := os.Open(sourcePath)
			if err != nil {
				return err
			}
			defer sourceFile.Close()

			// 将源文件内容复制到 zip 文件
			_, err = io.Copy(zipFile, sourceFile)
			if err != nil {
				return err
			}
		}

		return nil
	})

	return err
}

// DeleteIfExists 删除文件或文件夹（如果存在）
func DeleteIfExists(path string) error {
	// 检查文件或文件夹是否存在
	_, err := os.Stat(path)
	if err != nil {
		// 文件或文件夹不存在，无需删除
		if os.IsNotExist(err) {
			return nil
		}
		// 发生其他错误
		return err
	}

	// 删除文件或文件夹
	err = os.RemoveAll(path)
	if err != nil {
		return err
	}

	return nil
}
