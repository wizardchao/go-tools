package cmd

import (
	"go-tools/internal/m3u8"
	"strings"

	"github.com/spf13/cobra"
)

var (
	url      string
	output   string
	chanSize int
)

var desc = strings.Join([]string{
	"该子命令支持m3u8视频下载为mp4视频",
}, "\n")

var m3u8Cmd = &cobra.Command{
	Use:   "m3u8",
	Short: "下载m3u8视频至本地",
	Long:  desc,
	Run: func(cmd *cobra.Command, args []string) {
		m3u8.VideoSwitch(url, output, chanSize)
	},
}

func init() {
	m3u8Cmd.Flags().StringVarP(&url, "url", "u", "", "请输入链接！")
	m3u8Cmd.Flags().IntVarP(&chanSize, "chanSize", "c", 5, "请输入协程数！")
	m3u8Cmd.Flags().StringVarP(&output, "output", "o", "", "请输入文件地址！")
}
