// go 标准库 path/filepath
package main

import (
	"fmt"
	"os"
	"path/filepath"
)

// 打印路径名称
func pathName(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	} else {
		fmt.Println(path)
	}
	return nil
}

func main() {
	// 预定义变量
	fmt.Println(string(filepath.Separator), string(filepath.ListSeparator))

	// 返回path 相对当前路径的绝对路径
	dir := "~/gocode/src/go_note/package/filepath"
	real_dir, err := filepath.Abs(dir)
	fmt.Println(real_dir, err)

	// 返回path 的最短路径
	dir = "/usr/../etc/../tmp"
	clear_dir := filepath.Clean(dir)
	fmt.Println(clear_dir) // /tmp

	// 返回targpath 相对 basepath路径
	basepath, targpath := "/usr/local", "/usr/local/go/bin"
	rel_dir, err := filepath.Rel(basepath, targpath)
	fmt.Println(rel_dir, err) // go/bin <nil>

	// 返回软链指向的路径
	symlink := "/usr/local"
	real_dir, err = filepath.EvalSymlinks(symlink)
	fmt.Println(real_dir, err) // /usr/local <nil>

	// 返回路径最前面的卷名
	root := "/usr/local/go"
	vol := filepath.VolumeName(root)
	fmt.Println(vol) // ''

	// 路径分隔符替换为 `/`
	slash_dir := filepath.ToSlash(root)
	fmt.Println(slash_dir) // /usr/local/go

	//  `/` 替换为路径分隔符
	from_slash := filepath.FromSlash(slash_dir)
	fmt.Println(from_slash) // /usr/local/go

	// 分隔环境变量里面的路径
	env_path := "/usr/local/bin:/usr/bin:/bin:/usr/sbin:/sbin:/opt/X11/bin:/usr/local/go/bin"
	env_slice := filepath.SplitList(env_path)
	for k, v := range env_slice {
		fmt.Println(k, v)
	}
	// 0 /usr/local/bin
	// 1 /usr/bin
	// 2 /bin
	// 3 /usr/sbin
	// 4 /sbin
	// 5 /opt/X11/bin
	// 6 /usr/local/go/bin

	// 遍历 root 目录下的文件树，并调用 walkFn
	rootDir, err := os.Getwd()
	err = filepath.Walk(rootDir, pathName)
	fmt.Println(err)
}
