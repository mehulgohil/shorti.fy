package cmd

import (
	"fmt"
	"github.com/mehulgohil/shorti.fy/writer/config"
	"github.com/mehulgohil/shorti.fy/writer/pkg/algorithm/encoding"
	"github.com/mehulgohil/shorti.fy/writer/pkg/algorithm/hashing"
	"github.com/mehulgohil/shorti.fy/writer/services"

	"github.com/spf13/cobra"
)

// writeCmd represents the write command
var writeCmd = &cobra.Command{
	Use:   "write",
	Short: "write command helps you to provide the URL",
	Long:  `write command facilitates the option to provide the longs URL and shorten it`,
	Run: func(cmd *cobra.Command, args []string) {
		getURL, err := cmd.Flags().GetString("url")
		if err != nil {
			config.ZapLogger.Error(err.Error())
			return
		}

		getEmail, err := cmd.Flags().GetString("mail")
		if err != nil {
			config.ZapLogger.Error(err.Error())
			return
		}

		writerService := &services.ShortifyWriterService{
			IEncodingAlgorithm: &encoding.Base62Algorithm{},                          //injecting base62 as the encoding algorithm
			IHashingAlgorithm:  &hashing.MD5Hash{},                                   //injecting md5 as hashing algorithm
			IDataAccessLayer:   config.DynamoDB().(*config.DBClientHandler).DBClient, //injecting db client
			EnvVariables:       config.EnvVariables,                                  //injecting env variables
		}

		shortURL, err := writerService.Writer(getURL, getEmail)
		if err != nil {
			config.ZapLogger.Error(err.Error())
			return
		}

		fmt.Println(shortURL)
	},
}

func init() {
	rootCmd.AddCommand(writeCmd)
	writeCmd.Flags().StringP("url", "u", "", "provide long url")
	writeCmd.Flags().StringP("mail", "m", "", "provide user email")
	_ = writeCmd.MarkFlagRequired("url")
	_ = writeCmd.MarkFlagRequired("mail")
}
