# LocalizableStrings2Enum 🇺🇸

## About this tool

`Localizable.strings`ファイルからSwiftのEnum定義ファイルを作成します。

## Build 

```
$ go build localizableStrings2Enum.go
```

## Usage

カレントディレクトリに`Localize.swift`を作成
```
$ localizableStrings2Enum -l <Localizable.stringsファイル>
```

指定したディレクトリに`Localize.swift`を作成

```
$ localizableStrings2Enum -l <Localizable.stringsファイル> -o <ファイルの出力先>
```

## Options
```
Application Options:
  -l, --path=FILE_PATH        enumに変換するLocalizeStringsファイルへのパス (default: ./Localizable.strings)
  -o, --output=OUTPUT_PATH    ファイルの出力先 (default: ./)
  -n, --name=FILE_NAME        基本的にはLocalize.swift (default: Localize.swift)
  -e, --ename=ENUM_NAME       Enumの名前 (default: LocalizableStrings)

Help Options:
  -h, --help                  Show this help message
```


## Issue 😫

- [ ] `NSLocalizeStrings`の`Extension`もこのツール上で書く
- [ ] 前回との差分比較を行う