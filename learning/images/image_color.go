// GO语言中的标准的调色板(image/color包);
// 可将任意Color接口转换为采用自身色彩模型的Color接口, 转换可能会丢失色彩信息;
// var (
//    RGBAModel    Model = ModelFunc(rgbaModel)
//    RGBA64Model  Model = ModelFunc(rgba64Model)
//    NRGBAModel   Model = ModelFunc(nrgbaModel)
//    NRGBA64Model Model = ModelFunc(nrgba64Model)
//    AlphaModel   Model = ModelFunc(alphaModel)
//    Alpha16Model Model = ModelFunc(alpha16Model)
//    GrayModel    Model = ModelFunc(grayModel)
//    Gray16Model  Model = ModelFunc(gray16Model)
// )
// 将RGB三原色转换为Y'CbCr三原色;
// func RGBToYCbCr(r, g, b uint8) (uint8, uint8, uint8);
// 将Y'CbCr三原色转换为RGB三原色;
// func YCbCrToRGB(y, cb, cr uint8) (uint8, uint8, uint8);
package images
