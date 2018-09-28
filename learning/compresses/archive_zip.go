// GO语言中的zip档案文件的读写服务(archive/zip包);
// 本包不支持跨硬盘的压缩;
// 关于ZIP64:
// 为了向下兼容, FileHeader同时拥有32位和64位的Size字段;
// 64位字段总是包含正确的值, 对普通格式的档案未见它们的值是相同的;
// 对zip64格式的档案文件32位字段将是0xffffffff, 必须使用64位字段;
package compresses

import "fmt"
import "os"
import "io"
import "path/filepath"
import "archive/zip"
import "github.com/kingreatwill/go_study/learning/common"

// zip压缩;
func zip01() {
	common.Println("archive_zip.go")

	// FileWriter;
	fileWriter, err := os.Create("./doc/test/zip01.zip")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer fileWriter.Close()
	// ZipWriter;
	zipWriter := zip.NewWriter(fileWriter)
	defer zipWriter.Close()
	// 遍历目录下的文件夹/文件;
	root := "./doc/temp"
	err = filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		// 文件夹?
		if info.IsDir() {
			return nil
		}
		// FileHeader;
		header, err := zip.FileInfoHeader(info)
		if err != nil {
			fmt.Println(err)
			return nil
		}
		header.Name, err = filepath.Rel(filepath.Dir(root), path)
		header.Method = zip.Deflate
		// header.Method = zip.Store
		fmt.Println("header.Name:", header.Name)
		if err != nil {
			fmt.Println(err)
			return nil
		}
		// 元数据(文件);
		metaWriter, err := zipWriter.CreateHeader(header)
		if err != nil {
			fmt.Println(err)
			return nil
		}
		// 打开文件;
		file, err := os.Open(path)
		if err != nil {
			fmt.Println(err)
			return nil
		}
		defer file.Close()
		// 拷贝;
		_, err = io.Copy(metaWriter, file)
		if err != nil {
			fmt.Println(err)
			return nil
		}
		return nil
	})
	if err != nil {
		fmt.Println(err)
		return
	}
}

// zip解压;
func zip02() {
	// ReadCloser;
	ziprc, err := zip.OpenReader("./doc/test/zip01.zip")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer ziprc.Close()
	// 遍历元数据(文件);
	for _, file := range ziprc.File {
		// 创建目录;
		path := filepath.Dir("./doc/test/zip01/" + file.Name)
		err = os.MkdirAll(path, os.ModePerm)
		if err != nil {
			fmt.Println(err)
			return
		}
		// 创建文件;
		fileWriter, err := os.Create("./doc/test/zip01/" + file.Name)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer fileWriter.Close()
		// Multiple Files;
		rc, err := file.Open()
		if err != nil {
			fmt.Println(err)
			return
		}
		defer rc.Close()
		// 拷贝;
		_, err = io.Copy(fileWriter, rc)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}
