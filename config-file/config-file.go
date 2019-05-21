package main

import (
	"flag"
	"fmt"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

func init() {

	// import environment variables
	viper.SetEnvPrefix("SNITCHIT")
	viper.BindEnv("config")

	// define flags
	flag.String("config", "config.yaml", "Configuration file, default = config.yaml")

	// initialise flags
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)

	// display configfile details
	fmt.Println("   CLI:", viper.GetString("config"))
	dir, file := filepath.Split(viper.GetString("config"))
	fmt.Println("   DIR:", dir)
	fmt.Println("  FILE:", file)

	// setup configuration file
	viper.SetConfigType("yaml")
	config := strings.TrimSuffix(file, ".yml")
	fmt.Println("STRIP1:", config)
	config = strings.TrimSuffix(config, ".yaml")
	fmt.Println("STRIP2:", config)
	//viper.AddConfigPath(*configPath)
	viper.AddConfigPath(dir)

	// read in configuration file
	viper.SetConfigName(config)
	err := viper.ReadInConfig()

	if err != nil {
		fmt.Printf(" ERROR: Cannot read configuration: %s%s\n", dir, config)
	}

}

func main() {
	fmt.Println("-----------------------")
	displayConfig()
	os.Exit(0)
}

func displayConfig() {
	fmt.Printf("CONFIG: file: %s\n", viper.ConfigFileUsed())
	allmysettings := viper.AllSettings()
	var keys []string
	for k := range allmysettings {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		//fmt.Println("CONFIG:", k, ":", allmysettings[k])
		fmt.Printf("CONFIG: %s: %s\n", k, allmysettings[k])
	}
}
