// GO语言中的基本的2D图片库(image包);
package images

import "fmt"
import "os"
import "strings"
import "image"
import "image/color"
import "image/color/palette"
import "image/draw"
import "image/jpeg"
import "github.com/kingreatwill/go_study/learning/common"

// Alpha/Alpha16;
func image01() {
	common.Println("image.go")

	// 一幅内存中的图像(8位Alpah通道, Alpha通道为透明度);
	// func NewAlpha(r Rectangle) *Alpha;
	// func (p *Alpha) At(x, y int) color.Color;
	// func (p *Alpha) Bounds() Rectangle;
	// func (p *Alpha) ColorModel() color.Model;
	// func (p *Alpha) Opaque() bool;
	// func (p *Alpha) PixOffset(x, y int) int;
	// func (p *Alpha) Set(x, y int, c color.Color);
	// func (p *Alpha) SetAlpha(x, y int, c color.Alpha);
	// func (p *Alpha) SubImage(r Rectangle) Image;

	// 一幅内存中的图像(16位Alpah, Alpha通道为透明度);
	// func NewAlpha16(r Rectangle) *Alpha16;
	// func (p *Alpha16) At(x, y int) color.Color;
	// func (p *Alpha16) Bounds() Rectangle;
	// func (p *Alpha16) ColorModel() color.Model;
	// func (p *Alpha16) Opaque() bool;
	// func (p *Alpha16) PixOffset(x, y int) int;
	// func (p *Alpha16) Set(x, y int, c color.Color);
	// func (p *Alpha16) SetAlpha16(x, y int, c color.Alpha16);
	// func (p *Alpha16) SubImage(r Rectangle) Image;

	// 创建文件(图像输出端);
	root, _ := os.Getwd()
	_path := strings.Join([]string{root, "doc", "temp", "image01.jpeg"}, string(os.PathSeparator))
	file, err := os.Create(_path)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	// 创建一幅内存中的图像(指定宽高);
	alpha := image.NewAlpha(image.Rect(0, 0, 400, 200))
	for x := 0; x < 400; x++ {
		for y := 0; y < 200; y++ {
			// alpha.Set(x, y, color.Alpha{uint8(x % 200)})
			alpha.SetAlpha(x, y, color.Alpha{uint8(x % 200)})
		}
	}
	// 将内存中的图像存入文件;
	jpeg.Encode(file, alpha, nil)
}

// (常用)RGBA/RGBA64/NRGBA/NRGBA64;
func image02() {
	// (常用)一幅内存中的图像(32位RGBA, 传统的预乘了alpha通道的32位RGB色彩, Red/Green/Blue/Alpha各8位);
	// func NewRGBA(r Rectangle) *RGBA;
	// func (p *RGBA) At(x, y int) color.Color;
	// func (p *RGBA) Bounds() Rectangle;
	// func (p *RGBA) ColorModel() color.Model;
	// func (p *RGBA) Opaque() bool;
	// func (p *RGBA) PixOffset(x, y int) int;
	// func (p *RGBA) Set(x, y int, c color.Color);
	// func (p *RGBA) SetRGBA(x, y int, c color.RGBA);
	// func (p *RGBA) SubImage(r Rectangle) Image;

	// 一幅内存中的图像(64位RGBA, 传统的预乘了alpha通道的64位RGB色彩, Red/Green/Blue/Alpha各16位);
	// func NewRGBA64(r Rectangle) *RGBA64;
	// func (p *RGBA64) At(x, y int) color.Color;
	// func (p *RGBA64) Bounds() Rectangle;
	// func (p *RGBA64) ColorModel() color.Model;
	// func (p *RGBA64) Opaque() bool;
	// func (p *RGBA64) PixOffset(x, y int) int;
	// func (p *RGBA64) Set(x, y int, c color.Color);
	// func (p *RGBA64) SetRGBA64(x, y int, c color.RGBA64);
	// func (p *RGBA64) SubImage(r Rectangle) Image;

	// 一幅内存中的图像(32位NRGBA, 没有预乘alpha通道的32位RGB色彩, Red/Green/Blue/Alpha各8位);
	// func NewNRGBA(r Rectangle) *NRGBA;
	// func (p *NRGBA) At(x, y int) color.Color;
	// func (p *NRGBA) Bounds() Rectangle;
	// func (p *NRGBA) ColorModel() color.Model;
	// func (p *NRGBA) Opaque() bool;
	// func (p *NRGBA) PixOffset(x, y int) int;
	// func (p *NRGBA) Set(x, y int, c color.Color);
	// func (p *NRGBA) SetNRGBA(x, y int, c color.NRGBA);
	// func (p *NRGBA) SubImage(r Rectangle) Image;

	// 一幅内存中的图像(64位NRGBA, 没有预乘alpha通道的64位RGB色彩, Red/Green/Blue/Alpha各16位);
	// func NewNRGBA64(r Rectangle) *NRGBA64;
	// func (p *NRGBA64) At(x, y int) color.Color;
	// func (p *NRGBA64) Bounds() Rectangle;
	// func (p *NRGBA64) ColorModel() color.Model;
	// func (p *NRGBA64) Opaque() bool;
	// func (p *NRGBA64) PixOffset(x, y int) int;
	// func (p *NRGBA64) Set(x, y int, c color.Color);
	// func (p *NRGBA64) SetNRGBA64(x, y int, c color.NRGBA64);
	// func (p *NRGBA64) SubImage(r Rectangle) Image;

	// 创建文件(图像输出端);
	root, _ := os.Getwd()
	_path := strings.Join([]string{root, "doc", "temp", "image02.jpeg"}, string(os.PathSeparator))
	file, err := os.Create(_path)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	// 创建一幅内存中的图像(指定宽高);
	rgba := image.NewRGBA(image.Rect(0, 0, 400, 200))
	for x := 0; x < 400; x++ {
		for y := 0; y < 200; y++ {
			// rgba.Set(x, y, color.NRGBA{uint8(x % 200), uint8(y % 200), 0, 200})
			rgba.SetRGBA(x, y, color.RGBA{uint8(x % 200), uint8(y % 200), 0, 200})
		}
	}
	// 将内存中的图像存入文件;
	jpeg.Encode(file, rgba, nil)
}

// Gray/Gray16;
func image03() {
	// 一幅内存中的图像(8位Gray, 8位的灰度色彩);
	// func NewGray(r Rectangle) *Gray;
	// func (p *Gray) At(x, y int) color.Color;
	// func (p *Gray) Bounds() Rectangle;
	// func (p *Gray) ColorModel() color.Model;
	// func (p *Gray) Opaque() bool;
	// func (p *Gray) PixOffset(x, y int) int;
	// func (p *Gray) Set(x, y int, c color.Color);
	// func (p *Gray) SetGray(x, y int, c color.Gray);
	// func (p *Gray) SubImage(r Rectangle) Image;

	// 一幅内存中的图像(16位Gray, 16位的灰度色彩);
	// func NewGray16(r Rectangle) *Gray16;
	// func (p *Gray16) At(x, y int) color.Color;
	// func (p *Gray16) Bounds() Rectangle;
	// func (p *Gray16) ColorModel() color.Model;
	// func (p *Gray16) Opaque() bool;
	// func (p *Gray16) PixOffset(x, y int) int;
	// func (p *Gray16) Set(x, y int, c color.Color);
	// func (p *Gray16) SetGray16(x, y int, c color.Gray16);
	// func (p *Gray16) SubImage(r Rectangle) Image;

	// 创建文件(图像输出端);
	root, _ := os.Getwd()
	_path := strings.Join([]string{root, "doc", "temp", "image03.jpeg"}, string(os.PathSeparator))
	file, err := os.Create(_path)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	// 创建一幅内存中的图像(指定宽高);
	gray := image.NewGray(image.Rect(0, 0, 400, 200))
	for x := 0; x < 400; x++ {
		for y := 0; y < 200; y++ {
			// gray.Set(x, y, color.Gray{uint8(x % 200)})
			gray.SetGray(x, y, color.Gray{uint8(x % 200)})
		}
	}
	// 将内存中的图像存入文件;
	jpeg.Encode(file, gray, nil)
}

// (常用)Paletted;
func image04() {
	// (常用)一幅采用uint8类型索引调色板的内存中的图像;
	// func NewPaletted(r Rectangle, p color.Palette) *Paletted;
	// func (p *Paletted) At(x, y int) color.Color;
	// func (p *Paletted) Bounds() Rectangle;
	// func (p *Paletted) ColorIndexAt(x, y int) uint8;
	// func (p *Paletted) ColorModel() color.Model;
	// func (p *Paletted) Opaque() bool;
	// func (p *Paletted) PixOffset(x, y int) int;
	// func (p *Paletted) Set(x, y int, c color.Color);
	// func (p *Paletted) SetColorIndex(x, y int, index uint8);
	// func (p *Paletted) SubImage(r Rectangle) Image;

	// 创建文件(图像输出端);
	root, _ := os.Getwd()
	_path := strings.Join([]string{root, "doc", "temp", "image04.jpeg"}, string(os.PathSeparator))
	file, err := os.Create(_path)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	// 创建一幅内存中的图像(指定宽高与色彩);
	paletted := image.NewPaletted(image.Rect(0, 0, 400, 200), palette.Plan9)
	for x := 0; x < 400; x++ {
		for y := 0; y < 200; y++ {
			//paletted.Set(x, y, palette.Plan9[uint8(x%200)])
			paletted.SetColorIndex(x, y, uint8(x%200))
		}
	}
	// 将内存中的图像存入文件;
	jpeg.Encode(file, paletted, nil)
}

// YCbCr;
func image05() {
	// YCbCr: 其中Y是指亮度分量, Cb指蓝色色度分量, 而Cr指红色色度分量;
	// YCbCr模型来源于YUV模型, 应用于数字视频(ITU-R BT.601 recommendation);
	// YUV色彩模型来源于RGB模型, 该模型的特点是将亮度和色度分离开, 从而适合于图像处理领域;
	// 为了使用人的视角特性以降低数据量, 通常把RGB空间表示的彩色图像变换到其他彩色空间;
	// 目前采用的彩色空间变换有三种: YIQ/YUV/YCrCb, 每一种彩色空间都产生一种亮度分量信号和两种色度分量信号, 而每一种变换使用的参数都是为了适应某种类型的显示设备;
	// 其中YIQ适用于NTSC彩色电视制式, YUV适用于PAL和SECAM彩色电视制式, 而YCrCb适用于计算机用的显示器;
	// RGB、YUV和YCbCr三种颜色, 可参考: https://blog.csdn.net/u010186001/article/details/52800250

	// 具有指定宽度、高度和二次采样率的YcbCr的内存中的图像;
	// func NewYCbCr(r Rectangle, subsampleRatio YCbCrSubsampleRatio) *YCbCr;
	// func (p *YCbCr) At(x, y int) color.Color;
	// func (p *YCbCr) Bounds() Rectangle;
	// func (p *YCbCr) COffset(x, y int) int;
	// func (p *YCbCr) ColorModel() color.Model;
	// func (p *YCbCr) Opaque() bool;
	// func (p *YCbCr) SubImage(r Rectangle) Image;
	// func (p *YCbCr) YOffset(x, y int) int;
}

// (常用)Uniform;
func image06() {
	// (常用)一幅面积无限大的具有同一色彩的图像;
	// func NewUniform(c color.Color) *Uniform;
	// func (c *Uniform) At(x, y int) color.Color;
	// func (c *Uniform) Bounds() Rectangle;
	// func (c *Uniform) ColorModel() color.Model;
	// func (c *Uniform) Convert(color.Color) color.Color;
	// func (c *Uniform) Opaque() bool;
	// func (c *Uniform) RGBA() (r, g, b, a uint32);

	// 创建文件(图像输出端);
	root, _ := os.Getwd()
	_path := strings.Join([]string{root, "doc", "temp", "image06.jpeg"}, string(os.PathSeparator))
	file, err := os.Create(_path)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	// 创建一幅内存中的图像(指定宽高);
	rgba := image.NewRGBA(image.Rect(0, 0, 400, 200))
	// 底色填充;
	draw.Draw(rgba, rgba.Bounds(), &image.Uniform{color.RGBA{0, 0, 255, 255}}, image.ZP, draw.Src)
	// 将内存中的图像存入文件;
	jpeg.Encode(file, rgba, nil)
}

// Point/Rectangle;
func image07() {
	// Point(X, Y坐标对);
	// func Pt(X, Y int) Point;
	// func (p Point) Eq(q Point) bool;
	// func (p Point) Add(q Point) Point;
	// func (p Point) Sub(q Point) Point;
	// func (p Point) Mul(k int) Point;
	// func (p Point) Div(k int) Point;
	// func (p Point) In(r Rectangle) bool;
	// func (p Point) Mod(r Rectangle) Point;
	// func (p Point) String() string;

	// pt := image.Point{X: 5, Y: 5}
	pt := image.Pt(5, 5)
	fmt.Println(pt.String()) // (5,5);

	// Rectangle(矩形);
	// func Rect(x0, y0, x1, y1 int) Rectangle;
	// func (r Rectangle) Canon() Rectangle;
	// func (r Rectangle) Dx() int;
	// func (r Rectangle) Dy() int;
	// func (r Rectangle) Size() Point;
	// func (r Rectangle) Empty() bool;
	// func (r Rectangle) Eq(s Rectangle) bool;
	// func (r Rectangle) In(s Rectangle) bool;
	// func (r Rectangle) Overlaps(s Rectangle) bool;
	// func (r Rectangle) Add(p Point) Rectangle;
	// func (r Rectangle) Sub(p Point) Rectangle;
	// func (r Rectangle) Intersect(s Rectangle) Rectangle;
	// func (r Rectangle) Union(s Rectangle) Rectangle;
	// func (r Rectangle) Inset(n int) Rectangle;
	// func (r Rectangle) String() string;
	// pt := image.Point{X: 5, Y: 5}

	// rt := image.Rectangle{image.Point{X: 5, Y: 5}, image.Point{X: 10, Y: 10}}
	rt := image.Rect(5, 5, 10, 10)
	fmt.Println(rt.String()) // (5,5)-(10,5);
}
