package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)
const banner = `
▓█████▄  ▒█████   ▄████▄   ██ ▄█▀  ██████  ██░ ██  ██▓▓█████  ██▓    ▓█████▄ 
▒██▀ ██▌▒██▒  ██▒▒██▀ ▀█   ██▄█▒ ▒██    ▒ ▓██░ ██▒▓██▒▓█   ▀ ▓██▒    ▒██▀ ██▌
░██   █▌▒██░  ██▒▒▓█    ▄ ▓███▄░ ░ ▓██▄   ▒██▀▀██░▒██▒▒███   ▒██░    ░██   █▌
░▓█▄   ▌▒██   ██░▒▓▓▄ ▄██▒▓██ █▄   ▒   ██▒░▓█ ░██ ░██░▒▓█  ▄ ▒██░    ░▓█▄   ▌
░▒████▓ ░ ████▓▒░▒ ▓███▀ ░▒██▒ █▄▒██████▒▒░▓█▒░██▓░██░░▒████▒░██████▒░▒████▓ 
 ▒▒▓  ▒ ░ ▒░▒░▒░ ░ ░▒ ▒  ░▒ ▒▒ ▓▒▒ ▒▓▒ ▒ ░ ▒ ░░▒░▒░▓  ░░ ▒░ ░░ ▒░▓  ░ ▒▒▓  ▒ 
 ░ ▒  ▒   ░ ▒ ▒░   ░  ▒   ░ ░▒ ▒░░ ░▒  ░ ░ ▒ ░▒░ ░ ▒ ░ ░ ░  ░░ ░ ▒  ░ ░ ▒  ▒ 
 ░ ░  ░ ░ ░ ░ ▒  ░        ░ ░░ ░ ░  ░  ░   ░  ░░ ░ ▒ ░   ░     ░ ░    ░ ░  ░ 
   ░        ░ ░  ░ ░      ░  ░         ░   ░  ░  ░ ░     ░  ░    ░  ░   ░    
 ░               ░                                                    ░      
 🛡️ DockShield - Secure Your Docker Images with DockShield
`
var rootCmd = &cobra.Command{
	Use:   "dockshield",
	Short: "DockShield - Scan Docker images for vulnerabilities using Trivy",
	Long:  "DockShield is a CLI tool that scans Docker images for security vulnerabilities using Trivy.",
}
func Execute() {
	fmt.Println(banner)
	cobra.CheckErr(rootCmd.Execute())
}
func init() {
	rootCmd.AddCommand(scanCmd)
}