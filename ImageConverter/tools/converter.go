package tools

import (
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"os"
	"path/filepath"
)

type ImageInfo struct{
	fileName string
	path string
	ext string
}

func GetOneImageInfo(path string, imageInfoList []ImageInfo) []ImageInfo{
	file, err := os.Open(path)

	if err != nil{
		panic(err)
	}
	defer file.Close()

	var tmpImageInfo ImageInfo
	tmpImageInfo.ext = filepath.Ext(path)
	tmpImageInfo.path = filepath.Dir(path)

	tmpImageInfo.fileName = filepath.Base(path)[:len(filepath.Base(path))-len(tmpImageInfo.ext)]
	imageInfoList = append(imageInfoList, tmpImageInfo)

	fmt.Println(tmpImageInfo.fileName)

	return imageInfoList
}

func GetAllImagesInfo(path string, ext string, imageInfoList []ImageInfo) []ImageInfo{
	// 全てのFilePathの取得
	files, err := ioutil.ReadDir(path)

	if err != nil{
		panic(err)
	}

	// fileの検索
	for _, file := range files {
		if filepath.Ext(file.Name()) == ext {
			var tmpImageInfo ImageInfo
			tmpImageInfo.ext = ext
			tmpImageInfo.path = path
			tmpImageInfo.fileName = file.Name()[:len(file.Name())-len(ext)]
			imageInfoList = append(imageInfoList, tmpImageInfo)
			fmt.Println(filepath.Join(path, file.Name()))
		}else if file.IsDir(){
			imageInfoList = GetAllImagesInfo(filepath.Join(path, file.Name()), ext, imageInfoList)
		}
	}
	return imageInfoList
}

func Convert(imageInfo ImageInfo, destinationExt string){
	// 元ファイルを開く
	sf, err := os.Open(filepath.Join(imageInfo.path, imageInfo.fileName + imageInfo.ext))
	if err != nil {
		panic(err)
	}
	defer sf.Close()

	// 空のファイルを作成
	outputImage, err := os.Create(filepath.Join(imageInfo.path, imageInfo.fileName + destinationExt))
	if err != nil {
		panic(err)
	}
	defer outputImage.Close()

	// Imageファイルに変換
	var sourceImage image.Image
	switch imageInfo.ext {
	case ".jpg", "jpeg":
		sourceImage, err = jpeg.Decode(sf)
		if err != nil {
			panic(err)
		}
	case ".png":
		sourceImage, err = png.Decode(sf)
		if err != nil {
			panic(err)
		}
	case ".gif":
		sourceImage, err = gif.Decode(sf)
		if err != nil {
			panic(err)
		}
	}

	// 対象のファイルに変換
	switch destinationExt {
	case ".jpg", "jpeg":
		err := jpeg.Encode(outputImage, sourceImage, nil)
		if err != nil {
			panic(err)
		}
	case ".png":
		err := png.Encode(outputImage, sourceImage)
		if err != nil {
			panic(err)
		}
	case ".gif":
		err := gif.Encode(outputImage, sourceImage, nil)
		if err != nil {
			panic(err)
		}
	}
}