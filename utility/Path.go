package utility

import (
	"os"
	"path/filepath"
)

/**
 * @description: 获取 项目根目录或可执行文件 + 该文件名 的绝对路径
 *   编译后 os.Executable可以得到当前可执行文件的绝对路径，如果是编译后运行，根据该绝对路径判断配置文件是否存在
 * 	 开发时（未手动编译，go run运行时），os.Executable 只能获取到临时编译路径，所以用 os.Getwd 加上文件名即可得到配置文件
 *	 但是在某些情况下，例如单元测试的时候，程序在某个包下运行，os.Getwd 获取到的是对应包的路径而不是项目的根路径，向上遍历程序运行的路径，得到正确的配置文件路径
 * @author: Lorin
 * @time: 2020/11/20 上午11:03
 */




func GetExePath() string {
	exePath, err := os.Executable()
	if err != nil {
		return ""
	}

	return filepath.Dir(exePath) + string(os.PathSeparator)
}

func GetWorkDir() string {
	path, err := os.Getwd()
	if err != nil {
		return ""
	}

	return path + string(os.PathSeparator)
}

func PathFileExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}

	if os.IsExist(err) {
		return true
	}

	return false
}

// 判断所给路径是否为文件夹
func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

// 从工作路径向上递归查找文件，返回文件绝对路径
func RecursiveFind(fileName string) string {

	return ""
}
