package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"ImageConverter/tools"

	"github.com/schollz/progressbar/v3"
)


var path string
var te string
var de string

var imageInfoList []tools.ImageInfo

func init(){
	flag.StringVar(&path, "path", "./", "filePath" )
	flag.StringVar(&te, "te", ".png", "target ext" )
	flag.StringVar(&de, "de", ".jpg", "distinction ext" )
}

func main(){

	// filepathの取得
	flag.Parse()
	//fmt.Println(path)
	//path = "./"
	//te = ".png"
	//de = ".jpg"

	// filepathがディレクトリだったら以下すべて、
	if filepath.Ext(path) == ""{
		// 対象Imageの情報を配列に格納
		imageInfoList = tools.GetAllImagesInfo(path, te, imageInfoList)

	}else{
		imageInfoList = tools.GetOneImageInfo(path, imageInfoList)
	}




	// 該当ファイルがなければ終了
	if len(imageInfoList) == 0 {
		fmt.Println("該当ファイルが見つかりません")
		return
	}
	// 検索語対話型
	fmt.Println(strconv.Itoa(len(imageInfoList)) + "個のファイルが見つかりました。 \n変換しますか? yes/no")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if scanner.Text() == "yes" {
			// progress Barの表示
			bar := progressbar.Default(int64(len(imageInfoList)))
			fmt.Println("convert")
			for _, imageInfo := range imageInfoList{
				bar.Add(1)
				tools.Convert(imageInfo, de)
			}
			break
		}else if scanner.Text() == "no"{
			fmt.Println("キャンセルしました。")
			break
		}else{
			fmt.Println("\"yes\"か\"no\"を入力してください。")
			continue
		}
	}

}
