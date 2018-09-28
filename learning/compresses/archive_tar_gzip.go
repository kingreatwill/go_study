// GO语言中的tar格式压缩文件的存取(archive/tar包);
// GO语言中的gzip格式压缩文件的读写(compress/gzip包);
// 本包目标是覆盖大多数tar的变种, 包括GNU和BSD生成的tar文件;
package compresses

import "fmt"
import "os"
import "io"
import "path/filepath"
import "archive/tar"
import "compress/gzip"
import "github.com/kingreatwill/go_study/learning/common"

// tar_gzip压缩;
func tar_gzip_01() {
	common.Println("archive_tar_gzip.go")

	// FileWriter;
	fileWriter, err := os.Create("./doc/test/tar_gzip_01.tar.gz")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer fileWriter.Close()
	// GzipWriter;
	gzipWriter := gzip.NewWriter(fileWriter)
	defer gzipWriter.Close()
	// TarWriter;
	tarWriter := tar.NewWriter(gzipWriter)
	defer tarWriter.Close()
	// 遍历目录下的文件夹/文件;
	root := "./doc/temp"
	err = filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		// 文件夹?
		if info.IsDir() {
			// FileHeader;
			header, err := tar.FileInfoHeader(info, "")
			if err != nil {
				fmt.Println(err)
				return nil
			}
			header.Name, err = filepath.Rel(filepath.Dir(root), path)
			header.Typeflag = tar.TypeDir
			header.Mode = int64(info.Mode())
			fmt.Println("header.Name:", header.Name)
			if err != nil {
				fmt.Println(err)
				return nil
			}
			// 写入;
			err = tarWriter.WriteHeader(header)
			if err != nil {
				fmt.Println(err)
				return nil
			}
		} else {
			// FileHeader;
			header, err := tar.FileInfoHeader(info, "")
			if err != nil {
				fmt.Println(err)
				return nil
			}
			header.Name, err = filepath.Rel(filepath.Dir(root), path)
			header.Mode = int64(info.Mode())
			if err != nil {
				fmt.Println(err)
				return nil
			}
			// 写入;
			err = tarWriter.WriteHeader(header)
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
			_, err = io.Copy(tarWriter, file)
			if err != nil {
				fmt.Println(err)
				return nil
			}
		}
		return nil
	})
	if err != nil {
		fmt.Println(err)
		return
	}
}

// tar_gzip解压;
func tar_gzip_02() {
	// 打开文件;
	file, err := os.Open("./doc/test/tar_gzip_01.tar.gz")
	if err != nil {
		fmt.Println(err)
		return
	}
	// GzipReader;
	gzipReader, err := gzip.NewReader(file)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer gzipReader.Close()
	// tarReader;
	tarReader := tar.NewReader(gzipReader)
	// 遍历;
	for {
		header, err := tarReader.Next()
		if err == io.EOF {
			break
		}
		if header.Typeflag != tar.TypeDir {
			// 创建目录;
			path := filepath.Dir("./doc/test/tar_gzip_01/" + header.Name)
			err = os.MkdirAll(path, os.ModePerm)
			if err != nil {
				fmt.Println(err)
				return
			}
			// 创建文件;
			fileWriter, err := os.Create("./doc/test/tar_gzip_01/" + header.Name)
			if err != nil {
				fmt.Println(err)
				return
			}
			defer fileWriter.Close()
			// 拷贝;
			_, err = io.Copy(fileWriter, tarReader)
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	}
}
