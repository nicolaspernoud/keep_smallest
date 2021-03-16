package main

import (
	"errors"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

var resourceIconPng = &fyne.StaticResource{
	StaticName: "Icon.png",
	StaticContent: []byte(
		"\x89PNG\r\n\x1a\n\x00\x00\x00\rIHDR\x00\x00\x00\x80\x00\x00\x00\x80\b\x06\x00\x00\x00\xc3>a\xcb\x00\x00\x00\tpHYs\x00\x00\x0fD\x00\x00\x0fD\x01\x17\xad\xe3\v\x00\x00\x00\x19tEXtSoftware\x00www.inkscape.org\x9b\xee<\x1a\x00\x00\b\xdfIDATx\x9c\xed\xdd}P\x14\xe7\x1d\a\xf0\xef\xb3@\x8e\x03\xa2ƓP\x13\xd2&\xf8\x82:!\x96\x90D\x11\x9d\x1cISy1\xb53\xcaiՀiL1ք\xb4\x99&\xadi&Mk\xfeʴ\x8dNb\"\x194\xa1\xd2*\x87M\xad\x06Q\xab\x90\x88\x9c/\xe3\xa00b$\xa2\xd5\"FyQy\x17\xee\xf6\xe9\x1fx\x0ex{\x06\xeed\x17\x9e\xe7\xf7\xf9Kv\x97\xe5\xeb\xecwn\x9f\xdb}n\x8f\xc1OVkq`\xb7\xb9\xe91\x0e\xf5I\xce0\x95qL\xe4\xc0\xfd\x00B\x01\xdc\xed\xef\xfe\t\x00\xa0\x05@\x1b\x80Z\x80Us\xf0\xe3`\xec\xcb\xe0v\xcbђ\x92D\xa7?;f\xbe\xfeb|\xaa}\x06T\x96\x01p\x1b\x80Q\xfe\x84 >\xbb\x02\x86|\xa6\xf2\xcfʊl\x0e_v0\xe0\x02$$\x15\xa4\xa8\xe0o\x82a\x86/\u007f\x90\f\x16~\x00\f\xab\x1d\x85\xb6\xa2\x81\xfcV\xbf\v0#y\xf38\x15\xcaZ\x06\x962\xf0pDG{\xe1b+\x1d\xbb\xe7\u007fݟ\x8d\xfbU\x80\xf8d{:\x80u\xe89\xaf\x93\xa1\xaf\x13\x8c\xff\xd6Qh[\xf3]\x1b\u07b6\x00qq\xeb\x83L\x11\x96l\xce\xf9\xd2;\x16\x8d臱\x1cS\xbbe\xf9\xed\x06\x8a^\v\x10\xf7\xec\xf6\x90 g\x87\x9d^\xf2\x87\xbd\x1d\b\xe36\x87\xdd֡\xb5R\xd1Z\x18\x17\xb7>\x88\x0e\xbe0栍٭\xd6\xe2@\xad\x95\x9a\x050EX\xb2\xe9\xe0\v\x84#\xb5+\xb8\xe1#\xadU\x1e\x05\x98\x91\\\xf0\"\x9d\xf3\xc5\xc3\x19\x96MO\xca_z\xeb\xf2>c\x80is\n&(.~\x1c\x80Y\xaf`DWm\x01\x8a\xf2H\xe9\x17\xf3θ\x17\xf4y\x05`.\xf5}\xd0\xc1\x17Y\xa8˥~\xd8{\xc1\xcd\x02$$\x15\xa4\xd0y_\x02\fI\xf1I\x05\xb3\xdd?\xde,\x80\xca\xd4U\xc6$\"\xbac\xfcm\xf7?\x15\xe0ƍ\x1d\xb0\x04\xe3\x12\x11\x9d\xc5OO)\x98\x06\xb8_\x01T\x96ah\x1cb\x00\x9e\x0e\x00\xccj-\x0e\xbcnn\xb8\f\xe0\x1e\x83\x13\x11}5F\x86\xf1\b\xa5\xdb\xdc\xf4\x18\xe8\xe0\xcb\xc8Rצ\xc4**wY\x8dNB\x8c\xa1\xaaj\xa2\x02\xa6\xc4\x18\x1d\x84\x18E\x89Q\x18x\xb4\xd11\x88A\x14\x1e\xadp\xe0>\xa3s\x10\x83pD*\x00F\x18\x9d\x83\x18\xe6n\x05@\x88\xd1)\x88a\xc2\x14\xf815\x9c\f{LsB\b\x91\a\x15@rT\x00\xc9Q\x01$G\x05\x90\x1c\x15@rT\x00\xc9Q\x01$G\x05\x90\x9c\xe6ǅ\xc8\xd0b\x19\x1d\x8cȱa\x9a\xeb*\xaa\x1a\xc0\xb9\xef\xfb\xa6\x02\f\x03O͊į2\u007f\xa8\xb9\xeeɹ\xffDw\xb7\xea\xf3\xbe\xe9\x14 9*\x80\xe4\xa8\x00\x92\xa3\x02H\x8e\n 9*\x80\xe4|z\x1b8q\xdc(\xc4Ƅ\xf7Y\xb6\xff`\x1d\xea\xbem\xbb#\xa1\x88~|*@lL8\xb2~1\xb5ϲ\v\x17[\xa9\x00\xc3\x10\x9d\x02$G\x05\x90\x9c\x10\x05\b\r\t2:°%D\x01\xde_=\vKҢ\xa1(4\xc3}\xa0\x84(\x80\xc9\x14\x80\x15\xcf\xc7\xe0ӵ?\u0094\xe8\xd1F\xc7\x19V\x84(\x80\xdb\xf8\xa8\x91\xc8\xfeK\"\xdex\xf9Q:-\xf4\x93P\x05\x00\x00\x851\xccM\x8e\xc2\xe6\xec\xd9H~\xfa\aF\xc7\x19\xf2\x84+\x80\x9bet0\xdez\xedq\xbc\xf7\x87\x04D\x84\xd3\xc7\x1f\xbd\x11\xa2\x00\x9b?\xafƵ\xe6.\xcdu\tO\x8cE\xde\xc7?\x86m\xee\x04\x1a$j\x10\xa2\x00\x85\xff9\a۲\"ط\x9d\x86\xaazΏ\n1\a\xe2\xd5̩ظ\xf6i\x1a$\xdeB\x88\x02\x00@Kk\x17\xfe\xba\xfe\x18~\x9e\xb5\x17\x15U\x8d\x9a\xdbL\x88\x1aus\x90\x18b\xa6\xd9p\x80@\x05p\xab\xae\xb9\x8a\x97~S\x8c?\xfd\xf9\b\x9a\xaevz\xac\xef=H|jV\xa4\x01\t\x87\x16\xe1\n\x00\x00\x9c\x03;\xf7\x9e\xc3\xc2\x17wa\x93\xfd\x14\xba\x9d\x9e\x93&\xc7X\xccX\xfd\xbb\xe9\xd2\x0f\x12\x85,\x80[k[7\xd6m\xac\xc4s+\xf6\xe0p\xf9%\xcdm\x12\x9e\x18\x8b\u007fdϖ\xf6J\xa2\xd0\x05p;_ۂW\xdf\u070f\xd7\xdf9\x80K\xf5\xed\x1e\xeb\x83o\\I\x94q\x90(E\x01\xdcJ\x0f]Ģ\xcc\xdd\xc8ɫҜK/\xe3 Q\xaa\x02\x00@G\xa7\x139yUX\xf2\xd2n\x1c<\xfa\xad\xc7\xfaރ\xc4ę\xe2\x0f\x12\xa5+\x80\xdb\xff\xeaZ\xf1\xeb\xb7J\xf1\xfa;\ap\xf1\x92\xe7L\xa61\x163\xde]%\xfe Q\xda\x02\xb8\x95\x1e\xba\x88\xc5\xcb{N\v]].\x8f\xf5\xa2\x0f\x12\xa5/\x00\x00t^w\xf5\x9c\x16V\xec\x81\xe3\x88\xe7i\xc1=H|a\xf1\x14\x03\xd2\r.*@/\xb5u\xadx\xed\xedR\xbc\xf1\xc72ttx~\xdbj\xb0)\xc0\x80T\x83\x8b\np\x8b\xc7c\xef\xc5\xf3\x8b&\xc3,ɻ\x009\xfe\x97\xfd\x10\x1b\x13\x8ě\x87\xf1\xc8\x14\x8b\xd7mZۺuL\xa4\x0f\xe9\v0\xfe\xa1\x91X\xbap\xf2m\xef\v\x9c\xbfЂ\xec\xdc\x13(.\xad\xd51\x99>\xa4-\xc0\x83\x0f\x8c\xc0\xb2%S\x9083\x12\xcc\xcb\xe0\xfeZs\x17\xf2\xb6\x9e\u0096\u007f}\xe3\xd7C\x18\x862\xe9\n\xf0\xbd{C\x90\xb1`\x12\x9e\x9d\xfd\x90\u05f7u\u05fb\\\xb0o;\x8d\xdc\xfc\xaf\x85|\xd9\xefM\x9a\x02\x84[\xccX2?\x1a?M\x89BP\x90\xf6\xd8W\xe5\x1c%\xa5\x17\xf0\xe1\x86J͋C\"\x12\xbe\x00#G܅\xc5\xf3\xa2\x916w<Lwy\u007f\x1bw\xa4\xfc2>ȩ\xc07g\xae\xea\x98\xcex\xc2\x16\xc0l\x0eļ9\xe3\x90n\x9b\x84\xb0P\xefS\xc4Ϟoƺ\r\x958p\xf8\xa2\x8e\xe9\x86\x0e\xe1\n\x10l\n\xc0O\x92\xa2\x90\xb1`\x12\xee\x19e\xf2\xba]}c\a6\xfc\xfd$\xb6\xef:\xab9\x8fP\x16\xc2\x14 0PA\xea3\x0f\xe2\x85E\x931\xc6b\xf6\xba]G\xa7\x13[\xb7\xd7`\xe3據W\xfbd#D\x01\x9e\xb1>\x80\xe5\x19\x0fclD\xa8\xd7m\x9cN\x15ۊ\xce\"'\xaf\nW\xaf]\xd71\xdd\xd0&D\x01\x9eK\x9btۃ\u007f\xa4\xfc2\xd6d\x1fÙs\xcd:\xa6\x1a\x1e\x84(\x807'\xab\xaf\xe0\x83\x9c\n\x94W\xd6\x1b\x1de\xc8\x12\xb2\x00\x97\xeaۑ\x9d{\x02E\xfb\xce\xf9\xf5\x1c]\x19\bU\x80\xe6\x96.l*\x10\xfb\xd2\xed\x9d&D\x01\x9cN\x15\xf6m\xa7\xf1ɦ\x13\xc2_\xba\xbdӄ(\xc0+\xab\xbe\xa2\x03\xef#!&\x84\xd0\xc1\xf7\x9d\x10\x05 \xbe\xa3\x02HΧ1@ye=\xd6d\x1fﳌ.\xb2\fO>\x15\xa0\xba\xe6*\xaak\xe4\xbam**:\x05H\x8e\n 9*\x80\xe4\xa8\x00\x92\xa3\x02H\x8e\n 9!\xee\x05\x88n\xdf\xfeZT\x9f\xd6~\xdb\xed\xd4x\x00\xd6@P\x01\x86\x81ƦN46y>\xf2\xeeN\xa0S\x80\xe4\xa8\x00\x92\xa3\x02H\x8e\n 9*\x80\xe4\xa8\x00\x92\xa3\x02H\x8e\n 9*\x80\xe4\x14\x00\xf4\xd9\x19yq\x05\x80\xe7\xf3Ӊ,Z\x15\x004\x9bS^-\n\x03\xea\x8cNA\f\xc2P\xabp\xb0SF\xe7 \x06Q\xd9)\x85\x81W\x18\x9d\x83\x18E\xadTT\xc6J\x8c\x8eA\x8c\xa1\x06`\x9f\x12\xdcn9\n\xe0\x8a\xd1a\x88\xbe\x18\xd0\xf0\xfd\x10\x1cSJJ\x12\x9d`\xc87:\x10\xd1\x17\xe7\xd8b\xb7\xdb\\\n\x000\x95\u007fft \xa23\xces\x81\x1b\x97\x82ˊl\x0e\x06\x94\x1a\x9b\x88\xe8\xe8K\xc7.\xdba\xa0\u05fd\x00\xce\xf8\xbb\xc6\xe5!zR{\x1d\xeb\x9b\x05p\x14ڊ\xc0\xf0\x851\x91\x88^\x18\xf0\xefC\x85\xb6=\xee\x9f\xfb\xdc\rdܕ\x05\xba7 \xb2V\x95;\xb3z/\xe8S\x80\xb2\x9d\vk\x00\xf6\xb2\xbe\x99\x88\x8e~y\xb0\xe8g\xff\xed\xbd\xc0c>\x80c\xe7\xfc\r`,G\xb7HD\x1f\f\xeb\x1d;\xd3ro]\xac9!$2T\xcd\x04\xf0\xf9\xa0\x87\"z\xd9aj\x1f\xb3Rk\x85f\x01\xecv\x9b\va|1\x80\x1d\x83\x1a\x8b\xe8\x80oG\x18\xb7\x95\x94$j>\x1b\xff\xb6_\x86k\xb5\x16\av\x057|\xc4\x19\x96\rN82\xa8\x18֛\xdaǬ\xf4v\xf0{6\xe9\x87\xf8d{:\x80u\x00\xbc?\x93\x9d\f%\x1d\f,\xabl\xe7\xfcO\xbek\xc3~\u007f\x1d\xf6\xccԭQ.\xae\xae\x05G\xaa\u007f\xd9\xc8 \xdb\xe1r\xf1W\x0eﶝ\xed\xcf\xc6\x03\xfe>\xf4\xf8\xa4\x82ٌ\xf1\xdfs`\xe6\xc0\xb3\x91\xc1ÿR\x19V\xf7\xbe\xc8\xd3\x1f\x03.\x80\xdb\xf4\x94\x82i\x00Og\x1c\v\x00x\xff\xc2]2h\x18\xd0\xc09\xb6p\x85\xfd\xed`\xe1\xfcC>\xee\xc3?ii\xf9\x01\x17Z\x02\x1e\xe5pY\x19cS9\xf8D\x80\xdd\x0f \f\xc0\b\u007f\xf7O\x00\xf4L\xdcm\x05C-\a\xaa\x19\xc7q\x85\xb1\x92\xfbB\xd5r\xbb\xdd\xe6\xf2g\xc7\xff\aČ\xa5\xa1W@\xf2\xb2\x00\x00\x00\x00IEND\xaeB`\x82"),
}

func main() {
	a := app.New()
	w := a.NewWindow("Keep Smallest")
	w.SetIcon(resourceIconPng)

	var sourceDir fyne.ListableURI
	var destDir fyne.ListableURI

	// Usage info
	usageLabel := widget.NewLabel("If a file exists with the same name in both directories, it will replace the file in destination directory with the one in source directory if it is smaller.")
	usageLabel.Wrapping = fyne.TextWrapWord
	usageCard := widget.NewCard("Keep Smallest", "Keep Smallest compare files by names in source directory and destination directory.", usageLabel)

	// Compare button
	compareBtn := widget.NewButton("Compare", func() {
		list, err := KeepSmallest(sourceDir.Path(), destDir.Path())
		if err != nil {
			dialog.ShowError(err, w)
			return
		}
		if list == "" {
			dialog.ShowInformation("Compare successfull !", "No files altered.", w)
			return

		}
		dialog.ShowInformation("Compare successfull !", "Altered files : "+list, w)

	})
	compareBtn.Disable()
	compareBtnCard := widget.NewCard("Start comparison", "", compareBtn)

	// Path buttons
	sourceDirlabel := widget.NewLabel("Select a source directory")
	destDirlabel := widget.NewLabel("Select a destination directory")
	sourceDirBtn := widget.NewButton("Select source directory", func() {
		dialog.ShowFolderOpen(func(dir fyne.ListableURI, err error) {
			if err != nil {
				dialog.ShowError(err, w)
				return
			}
			if dir != nil {
				sourceDir = dir
				sourceDirlabel.SetText("source directory : " + sourceDir.Path())
				enableCompareButton(sourceDir, destDir, compareBtn)
			}
		}, w)
	})
	destDirBtn := widget.NewButton("Select destination directory", func() {
		dialog.ShowFolderOpen(func(dir fyne.ListableURI, err error) {
			if err != nil {
				dialog.ShowError(err, w)
				return
			}
			if dir != nil {
				destDir = dir
				destDirlabel.SetText("destination directory : " + destDir.Path())
				enableCompareButton(sourceDir, destDir, compareBtn)
			}
		}, w)
	})
	pathsCard := widget.NewCard("Choose directories", "", container.NewVBox(sourceDirlabel, sourceDirBtn, destDirlabel, destDirBtn))

	w.SetContent(container.NewVBox(
		usageCard,
		pathsCard,
		compareBtnCard,
	))

	w.ShowAndRun()
}

func enableCompareButton(sourceDir, destDir fyne.ListableURI, compareBtn *widget.Button) {
	if sourceDir != nil && destDir != nil {
		compareBtn.Enable()
	}
}

// KeepSmallest compare files by names in sourceDir and destinationDir.
// If a file exists with the same name in both directories, it will replace the file in destDir with the one in sourceDir if it is smaller.
func KeepSmallest(sourceDir string, destDir string) (alteredFiles string, e error) {
	_, err := os.Stat(sourceDir)
	if err != nil {
		return "", errors.New("source directory must exist")
	}
	_, err = os.Stat(destDir)
	if err != nil {
		return "", errors.New("destination directory must exist")
	}
	if sourceDir == destDir {
		return "", errors.New("destination cannot be the same as source")
	}
	files, err := ioutil.ReadDir(destDir)
	if err != nil {
		return "", err
	}
	// Iterate in destDir
	for _, dfi := range files {
		// Look for a file with the same name in sourceDir and replace in destDir if it is smaller
		if sfi, err := os.Stat(filepath.Join(sourceDir, dfi.Name())); err == nil && sfi.Size() < dfi.Size() {
			sf, err := os.Open(filepath.Join(sourceDir, sfi.Name()))
			if err != nil {
				return "", err
			}
			err = os.Remove(filepath.Join(destDir, dfi.Name()))
			if err != nil {
				return "", err
			}
			df, err := os.OpenFile(filepath.Join(destDir, dfi.Name()), os.O_CREATE|os.O_WRONLY, dfi.Mode())
			if err != nil {
				return "", err
			}
			_, err = io.Copy(df, sf)
			if err != nil {
				return "", err
			}
			df.Close()
			sf.Close()
			alteredFiles += sfi.Name() + " "
		}
	}
	return alteredFiles, nil
}
