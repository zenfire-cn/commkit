package tool

import (
	"os"
	"path/filepath"
	"strings"
)

/**
 * @description: 获取 项目根目录或可执行文件 + 该文件名 的绝对路径
 *   编译后 os.Executable可以得到当前可执行文件的绝对路径，如果是编译后运行，根据该绝对路径判断配置文件是否存在
 * 	 开发时（未手动编译，go run运行时），os.Executable 只能获取到临时编译路径，所以用 os.Getwd 加上文件名即可得到配置文件
 *	 但是在某些情况下，例如单元测试的时候，程序在某个包下运行，os.Getwd 获取到的是对应包的路径而不是项目的根路径，向上遍历程序运行的路径，得到正确的配置文件路径
 * @author: Lorin
 * @time: 2020/11/20 上午11:03
 */
func ExePath(fileName string) string {
	sp := string(os.PathSeparator)
	// 编译后 os.Executable可以得到当前可执行文件的绝对路径
	executablePath, _ := os.Executable()
	executablePath, _ = filepath.EvalSymlinks(filepath.Dir(executablePath))
	executablePath += sp + fileName

	if pathExists(executablePath) {
		// 如果是编译后运行，根据该绝对路径判断配置文件是否存在，存在则初始化
		return executablePath
	} else {
		splits := strings.Split(executablePath, sp)
		for i := len(splits); i > 0; i-- {
			path := strings.Join(splits[0:i], sp) + sp + fileName
			if pathExists(path) {
				return path
			}
		}
		pwd, _ := os.Getwd()
		splits = strings.Split(pwd, sp)
		for i := len(splits); i > 0; i-- {
			path := strings.Join(splits[0:i], sp) + sp + fileName
			if pathExists(path) {
				return path
			}
		}
	}
	return ""
}

func pathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}
