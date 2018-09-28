// GO语言中的gif文件的编码器和解码器(image/gif包);
package images

import "fmt"
import "os"
import "strings"
import "image"
import "image/jpeg"
import "image/color"
import "image/color/palette"
import "github.com/kingreatwill/go_study/learning/common"

// image_gif_jpeg_png;
func image_gif_jpeg_png() {
	common.Println("image_gif_jpeg_png.go")

	// image/gif;
	// 将g中所有的图像按指定的每帧延迟和累计循环时间写入w中;
	// func EncodeAll(w io.Writer, g *GIF) error;
	// 从r中读取一个GIF格式文件(返回值中包含了连续的图帧和时间信息);
	// func DecodeAll(r io.Reader) (*GIF, error);

	// image/jpeg;
	// 将采用JPEG 4:2:0基线格式和指定的编码质量将图像写入w, 如果o为nil将使用DefaultQuality;
	// func Encode(w io.Writer, m image.Image, o *Options) error;
	// 从r读取一幅jpeg格式的图像并解码返回该图像;
	// func Decode(r io.Reader) (image.Image, error);

	// image/png;
	// 将图像m以PNG格式写入w;
	// func Encode(w io.Writer, m image.Image) error;
	// 从r读取一幅png格式的图像并解码返回该图像, 图像的具体类型要看png文件的内容而定;
	// func Decode(r io.Reader) (image.Image, error);

	// 创建文件(图像输出端);
	root, _ := os.Getwd()
	dst, err1 := os.Create(strings.Join([]string{root, "doc", "temp", "image_gif_jpeg_png.jpg"}, string(os.PathSeparator)))
	if err1 != nil {
		fmt.Println(err1)
		return
	}
	defer dst.Close()

	img := image.NewPaletted(image.Rect(0, 0, 110, 110), palette.Plan9)
	for x := 0; x < 100; x++ {
		for y := 0; y < 100; y++ {
			img.Set(50, y, color.RGBA{uint8(x), uint8(y), 255, 255})
		}
	}

	// 图像质量值为100(最高质量);
	err2 := jpeg.Encode(dst, img, &jpeg.Options{100})
	if err2 != nil {
		fmt.Println(err2)
		return
	}
}
