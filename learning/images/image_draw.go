// GO语言中的基本的2D图片库(image/draw包);
package images

import "fmt"
import "os"
import "strings"
import "image"
import "image/draw"
import "image/jpeg"
import "github.com/kingreatwill/go_study/learning/common"

// draw01;
func draw01() {
	common.Println("image_draw.go")

	// Op是Porter-Duff合成的操作符;
	// type Op int;
	// const (
	//    // 源图像透过遮罩后，覆盖在目标图像上（类似图层）
	//    Over Op  = iota
	//    // 源图像透过遮罩后，替换掉目标图像
	//    Src
	// )

	// 使用nil的mask参数调用DrawMask函数;
	// func Draw(dst Image, r image.Rectangle, src image.Image, sp image.Point, op Op);
	// 对齐目标图像dst的矩形r左上角、源图像src的sp点、遮罩mask的mp点, 根据op修改dst的r矩形区域内的内容, mask如果为nil则视为完全透明;
	// func DrawMask(dst Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point, op Op);

	// 打开文件(原图);
	root, _ := os.Getwd()
	src, err1 := os.Open(strings.Join([]string{root, "doc", "temp", "image.draw.src.jpg"}, string(os.PathSeparator)))
	if err1 != nil {
		fmt.Println(err1)
		return
	}
	defer src.Close()
	// 打开文件(LOGO);
	logo, err2 := os.Open(strings.Join([]string{root, "doc", "temp", "image.draw.logo.jpg"}, string(os.PathSeparator)))
	if err2 != nil {
		fmt.Println(err2)
		return
	}
	defer logo.Close()
	// 创建文件(图像输出端);
	dst, err3 := os.Create(strings.Join([]string{root, "doc", "temp", "image.draw.dst.jpg"}, string(os.PathSeparator)))
	if err3 != nil {
		fmt.Println(err3)
		return
	}
	defer dst.Close()
	// 文件转图像;
	imgsrc, _ := jpeg.Decode(src)
	imglogo, _ := jpeg.Decode(logo)
	imgdst := image.NewNRGBA(imgsrc.Bounds())
	// 绘画(原图);
	draw.Draw(imgdst, imgdst.Bounds(), imgsrc, image.ZP, draw.Over)
	// 绘画(LOGO);
	offset := image.Pt(imgdst.Bounds().Dx()-imglogo.Bounds().Dx()-10, imgdst.Bounds().Dy()-imglogo.Bounds().Dy()-10)
	draw.Draw(imgdst, imgdst.Bounds().Add(offset), imglogo, image.ZP, draw.Over)
	// 将内存中的图像存入文件;
	jpeg.Encode(dst, imgdst, &jpeg.Options{100})
}
