package image

import (
	"fmt"
	"testing"

	"gocv.io/x/gocv"
)

func TestSave2Jpg(t *testing.T) {

	srcFile := "/Users/apple/workspace_stariverpool/go-image/testdata/opencv-logo.png"

	// 只截图，不保存
	f, err := Image(srcFile)

	if err != nil {
		fmt.Println(err)
		return
	}

	thFile := "/Users/apple/workspace_stariverpool/go-image/testdata/output/opencv-logo1.jpg"
	coverFile := "/Users/apple/workspace_stariverpool/go-image/testdata/output/opencv-logo2.jpg"

	// 保存到 指定目录
	if err = Save2Jpg(f, thFile, coverFile); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("保存成功")

}

func TestImageAndSave(t *testing.T) {

	//srcFile := "//Users/apple/视频/科技地球粒子mkv.mkv"
	//srcFile := "/Users/apple/workspace_stariverpool/go-image/testdata/1.png"

	//srcFile := "/Users/apple/workspace_stariverpool/go-image/testdata/2.png"
	//srcFile := "/Users/apple/workspace_stariverpool/go-image/testdata/opencv-logo.png"

	//srcFile := "/Users/apple/workspace_stariverpool/go-image/testdata/mimetype/gif.gif"
	srcFile := "/Users/apple/workspace_stariverpool/go-image/testdata/bad_gif/upload.gif"

	//srcFile := "/Users/apple/Desktop/bafybeiezq6jb7mptlqxtzvbwj2n257h6jbse24vxr66sc4gcd4na6kvpau.heic"
	//srcFile := "/Users/apple/Desktop/压缩后/五牛图-166x25.jpg"
	//srcFile := "/Users/apple/workspace_stariverpool/go-image/testdata/清明上河图.jpg"
	//srcFile := "/Users/apple/Desktop/压缩后/清明上河图-516x25.jpg"
	//srcFile := "/Users/apple/Desktop/压缩后/清明上河图-516x25.jpg"
	//srcFile := "/Users/apple/workspace_stariverpool/go-image/testdata/opencv-logo.png"
	//srcFile := "/Users/apple/workspace_stariverpool/go-image/testdata/1.png"
	//srcFile := "/Users/apple/workspace_stariverpool/go-image/testdata/2.png"
	//
	f, err := ImageAndSave(srcFile, "/Users/apple/workspace_stariverpool/go-image/testdata/output6/")

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(f.ThumbnailImgPath, f.CoverImgPath)

}

func TestImageResize(t *testing.T) {
	//srcFile := "/Users/apple/workspace_stariverpool/go-image/testdata/opencv-logo.png"
	srcFile := "/Users/apple/workspace_stariverpool/go-image/testdata/1.png"
	//srcFile := "/Users/apple/workspace_stariverpool/go-image/testdata/2.png"

	srcMat := gocv.IMRead(srcFile, gocv.IMReadColor)
	defer srcMat.Close()

	//desMat := gocv.IMRead(srcFile, gocv.IMReadColor)

	dst := gocv.NewMat()
	defer dst.Close()

	//dst := srcMat

	Resize(srcMat, dst)

	fmt.Println(dst.Size()) // []
	fmt.Println(len(srcMat.ToBytes()), len(dst.ToBytes()))

	if ok := gocv.IMWrite("/Users/apple/workspace_stariverpool/go-image/testdata/test_r.jpg", dst); !ok {
		fmt.Println("缩略图 保存失败")
	}

}

func TestResizeROI(t *testing.T) {
	//srcFile := "/Users/apple/workspace_stariverpool/go-image/testdata/opencv-logo.png"
	srcFile := "/Users/apple/workspace_stariverpool/go-image/testdata/output4/1.png"
	//srcFile := "/Users/apple/workspace_stariverpool/go-image/testdata/output4/2.png"

	srcMat := gocv.IMRead(srcFile, gocv.IMReadColor)
	defer srcMat.Close()

	//desMat := gocv.IMRead(srcFile, gocv.IMReadColor)

	dst := gocv.NewMat()
	defer dst.Close()

	//dst := srcMat
	fmt.Println("1111111", dst.Size()) // []
	ResizeROI(&srcMat, &dst)
	fmt.Println("222222", dst.Size()) // []

	fmt.Println(len(srcMat.ToBytes()), len(dst.ToBytes()))

	if !dst.Empty() { // 添加一个判断才行
		if ok := gocv.IMWrite("/Users/apple/workspace_stariverpool/go-image/testdata/test_r.jpg", dst); !ok {
			fmt.Println("缩略图 保存失败")
		} else {
			fmt.Println("保存成功")
		}
	}

}
