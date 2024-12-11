package cli

import (
	"fmt"
	"os"
	"os/exec"
	"os/signal"

	"github.com/spf13/cobra"
)

func (r *Operations) ServeCmd() *cobra.Command {
	var port int
	var host string
	var watch bool
	var local bool
	var module string

	cmd := &cobra.Command{
		Use:   "serve",
		Args:  cobra.MaximumNArgs(1),
		Short: "Serve a beamlit project",
		Long:  "Serve a beamlit project",
		Example: `
			bl serve --local --port 1338
		`,
		Run: func(cmd *cobra.Command, args []string) {
			uvicorn := exec.Command(
				"uvicorn",
				"beamlit.serve.app:app",
				"--reload",
				"--port",
				fmt.Sprintf("%d", port),
			)
			uvicornEnvironment := "production"
			if local {
				uvicornEnvironment = "local"
			} else if watch {
				uvicornEnvironment = "development"
			}

			uvicorn.Stdout = os.Stdout
			uvicorn.Stderr = os.Stderr
			uvicorn.Env = append(uvicorn.Env, fmt.Sprintf("BL_ENVIRONMENT=%s", uvicornEnvironment))
			uvicorn.Env = append(uvicorn.Env, fmt.Sprintf("BL_WORKSPACE=%s", workspace))
			uvicorn.Env = append(uvicorn.Env, fmt.Sprintf("BL_SERVER_PORT=%d", port))
			uvicorn.Env = append(uvicorn.Env, fmt.Sprintf("BL_SERVER_HOST=%s", host))
			uvicorn.Env = append(uvicorn.Env, fmt.Sprintf("BL_SERVER_MODULE=%s", module))

			err := uvicorn.Start()
			if err != nil {
				fmt.Printf("Error starting server: %v\n", err)
				os.Exit(1)
			}

			// Handle graceful shutdown on interrupt
			c := make(chan os.Signal, 1)
			signal.Notify(c, os.Interrupt)
			go func() {
				<-c
				fmt.Println("\nShutting down server...")
				if err := uvicorn.Process.Kill(); err != nil {
					fmt.Printf("Error killing server process: %v\n", err)
				}
				os.Exit(0)
			}()

			if err := uvicorn.Wait(); err != nil {
				fmt.Printf("Server error: %v\n", err)
				os.Exit(1)
			}
		},
	}

	cmd.Flags().IntVarP(&port, "port", "p", 1338, "Bind socket to this host")
	cmd.Flags().StringVarP(&host, "host", "H", "0.0.0.0", "Bind socket to this port. If 0, an available port will be picked")
	cmd.Flags().StringVarP(&module, "module", "m", "agent.main", "Module to serve, can be an agent or a function")
	cmd.Flags().BoolVarP(&watch, "watch", "W", false, "Watch for changes in the project, save changes to beamlit development environment and execute on it")
	cmd.Flags().BoolVarP(&local, "local", "l", false, "Serve the project locally, without using the cloud")
	return cmd
}
