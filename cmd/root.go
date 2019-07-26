package cmd

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/danielkvist/beagle/client"
	"github.com/danielkvist/beagle/logger"
	"github.com/danielkvist/beagle/sites"

	"github.com/spf13/cobra"
)

var (
	agent      string
	csvFile    string
	debug      bool
	disclaimer bool
	goroutines int
	proxy      string
	timeout    time.Duration
	user       string
	verbose    bool
)

func init() {
	rootCmd.PersistentFlags().StringVarP(&agent, "agent", "a", "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:67.0) Gecko/20100101 Firefox/67.0", "user agent")
	rootCmd.PersistentFlags().StringVar(&csvFile, "csv", "./urls.csv", ".csv file with the URLs to parse and check")
	rootCmd.PersistentFlags().BoolVar(&debug, "debug", false, "prints error messages")
	rootCmd.PersistentFlags().BoolVar(&disclaimer, "disclaimer", true, "disables disclaimer")
	rootCmd.PersistentFlags().IntVarP(&goroutines, "goroutines", "g", 1, "number of goroutines")
	rootCmd.PersistentFlags().StringVarP(&proxy, "proxy", "p", "", "proxy URL")
	rootCmd.PersistentFlags().DurationVarP(&timeout, "timeout", "t", 3*time.Second, "max time to wait for a response")
	rootCmd.PersistentFlags().StringVarP(&user, "user", "u", "me", "username you want to search for")
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "enables verbose mode")
}

var rootCmd = &cobra.Command{
	Use:   "beagle",
	Short: "",
	RunE: func(cmd *cobra.Command, args []string) error {
		if disclaimer {
			printDisclaimer()
		}

		siteList, err := sites.Parse(csvFile)
		if err != nil {
			return err
		}

		if len(siteList) == 0 {
			return fmt.Errorf(".csv file %q is empty or does not contains valid URLs", csvFile)
		}

		c, err := client.New(client.WithTimeout(timeout), client.WithProxy(proxy))
		if err != nil {
			return fmt.Errorf("while creating a new http.Client to make the requests: %v", err)
		}

		l := logger.New(os.Stdout, goroutines)
		sema := make(chan struct{}, goroutines)
		var wg sync.WaitGroup

		for _, s := range siteList {
			wg.Add(1)
			sema <- struct{}{}

			go func(site *sites.Site) {
				defer func() {
					<-sema
					wg.Done()
				}()

				site.ReplaceURL("$", user)
				_, statusCode, err := check(c, site.URL, agent)
				if err != nil && debug {
					log.Printf("while checking %q (%q): %v", site.Name, site.URL, err)
				}

				if !verbose && statusCode != http.StatusOK {
					return
				}

				l.Println(formatMsg(site.Name, site.URL, statusCode))
			}(s)
		}

		wg.Wait()
		l.Stop()

		return nil
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func printDisclaimer() {
	beagle := `	    __
 \,--------/_/'--o  	Use beagle with
 /_    ___    /~"   	responsibility.
  /_/_/  /_/_/
^^^^^^^^^^^^^^^^^^
`

	fmt.Println(beagle)
}

func check(c *http.Client, url string, agent string) (string, int, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return "", 0, err
	}
	req.Header.Set("User-Agent", agent)

	resp, err := c.Do(req)
	if err != nil {
		return "", 0, err
	}

	return resp.Status, resp.StatusCode, nil
}

func formatMsg(name string, url string, status int) string {
	if status != http.StatusOK {
		return fmt.Sprintf("NO %s %s", name, url)
	}
	return fmt.Sprintf("OK %s %s", name, url)
}
