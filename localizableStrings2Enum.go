package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"

	"github.com/jessevdk/go-flags"
	
)

var opts struct {
	INPUT_FILE_PATH   	string `short:"l" long:"path" default:"./Localizable.strings" description:"enumに変換するLocalizeStringsファイルへのパス" value-name:"FILE_PATH" required:"true"`
	OUTPUT_PATH 		string `short:"o" long:"output" default:"./" description:"ファイルの出力先" value-name:"OUTPUT_PATH" required:"true"`
	OUTPUT_FILE_NAME	string `short:"n" long:"name" default:"Localize.swift" description:"基本的にはLocalize.swift" value-name:"FILE_NAME" required:"false"`
	ENUM_NAME			string `short:"e" long:"ename" default:"LocalizableStrings" description:"Enumの名前" value-name:"ENUM_NAME" required:"false"`
}


func main() {

	_, err := flags.Parse(&opts)
	CheckIfError(err)

	texts := []string{}

	LocalisableStringsAnalyzer(opts.INPUT_FILE_PATH, &texts)

	filePath := opts.OUTPUT_PATH + opts.OUTPUT_FILE_NAME
	err = os.RemoveAll(filePath)
	CheckIfError(err)

    fp, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	CheckIfError(err)
	
	fmt.Fprintln(fp, "import Foundation")
	fmt.Fprintln(fp, "enum " + opts.ENUM_NAME + ": String {")

	for _, text := range texts {
		if text == "" {
			continue
		}
		// 空白はアンダースコアに置換
		keyword := strings.Replace(text, " ", "_", -1)
		// ハイフンはアンダースコアに置き換え
		keyword =  strings.Replace(text, "-", "_", -1)
		// CamelCase
		keyword = convertToCamelCase(keyword)

		str := fmt.Sprintf("    case %s = \"%s\"", keyword, text)
		fmt.Fprintln(fp,str)
		fmt.Println(str)
	}
	fmt.Fprintln(fp, `    var localized: String {
        NSLocalizedString(self.rawValue, comment: self.rawValue)
    }
    
    func localized(withTableName tableName: String? = nil, bundle: Bundle = Bundle.main, value: String = "") -> String {
        NSLocalizedString(self.rawValue, tableName: tableName, bundle: bundle, value: value, comment: self.rawValue)
    }`)

	fmt.Fprintln(fp, "}")

	fmt.Print("success!")
}

func LocalisableStringsAnalyzer(path string, texts *[]string ){
	// xxx.strings のみを解析
	if !strings.HasSuffix(path, ".strings") {
		return
	}

	// ファイルをOpenする
	file, err := os.Open(path)
	// 読み取り時の例外処理
	if err != nil {
		fmt.Println("error")
	}
	// 関数が終了した際に確実に閉じるようにする
	defer file.Close()

	sc := bufio.NewScanner(file)
	for i := 1; sc.Scan(); i++ {
		if err := sc.Err(); err != nil {
			// エラー処理
			break
		}

		text := sc.Text()
		text = strings.TrimSpace(text)
		index := strings.Index(text, "//")
		if index > 0 {
			text = text[:index]
			text = strings.TrimSpace(text)
		}

		if strings.HasSuffix(text, ";") {
			index := strings.Index(text, "=")
			if index > 0 {
				text = text[:index]
				text = strings.TrimSpace(text)
				text = strings.Trim(text, "\"")
			}

			var result bool = false
			for _, element := range *texts {
				if element == text {
					result = true
					break
				}
			}
			if result == false {
				*texts = append(*texts, text)
			}
		}
	}

}


func output(text string) {
	fmt.Print(text)
}

// 
func convertToCamelCase(text string) string {
	var keyword string
	var foundUnderScore = false
	for i := 0; i < len(text); i++ {
		letter := text[i : i+1]
		if letter == "_" {
			foundUnderScore = true
			continue
		}
		if foundUnderScore {
			foundUnderScore = false
			keyword = keyword + strings.ToUpper(letter)
		} else {
			keyword = keyword + letter
		}
	}
	return keyword
}

func CheckIfError(err error) {
	if err == nil {
		return
	}

	fmt.Printf("\x1b[31;1m%s\x1b[0m\n", fmt.Sprintf("error: %s", err))
	os.Exit(1)
}

func Exists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}