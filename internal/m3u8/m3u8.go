package m3u8

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/oopsguy/m3u8/dl"
)

//转化为Mp4
func VideoSwitch(url string, output string, chanSize int) error {
	if url == "" {
		panicParameter("u")
	}

	if output == "" {
		panicParameter("o")
	}

	if chanSize <= 0 {
		panic("parameter 'c' must be greater than 0")
	}

	fileName, err := switchToTs(url, output, chanSize)
	if err != nil {
		return err
	}

	err = switchToMp4(fileName, output)
	if err != nil {
		fmt.Println(err)
	}
	return nil
}

func panicParameter(name string) {
	panic("parameter '" + name + "' is required")
}

//转换为ts
func switchToTs(url string, output string, chanSize int) (string, error) {
	fileName := output + "/main.ts"
	downloader, err := dl.NewTask(output, url)
	if err != nil {
		return fileName, err
	}

	if err := downloader.Start(chanSize); err != nil {
		return fileName, err
	}

	return fileName, nil
}

//转换为Mp4
func switchToMp4(fileName string, dir string) error {
	binary, lookErr := exec.LookPath("ffmpeg")
	if lookErr != nil {
		panic(lookErr)
	}

	mp4FileName := dir + "/" + time.Now().Format("2006-01-02150405") + ".mp4"
	args := []string{
		"-i",
		fileName,
		"-acodec",
		"copy",
		"-vcodec",
		"copy",
		"-absf",
		"aac_adtstoasc",
		mp4FileName,
	}

	cmd := exec.Command(binary, args...)
	_, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
		return err
	}

	err = os.Remove(fileName)
	if err != nil {
		log.Println(err)
	}
	log.Println("Success!")
	return nil
}
