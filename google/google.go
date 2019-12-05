package google

import(
	"net/http"
	"fmt"

	"github.com/Leangeful/FicLink/config"

	"google.golang.org/api/googleapi/transport"
	customsearch "google.golang.org/api/customsearch/v1"
)

//Search does a google customsearch
func Search(query string) string{
	client := &http.Client{Transport: &transport.APIKey{Key: config.Cfg.APIKey}}

	svc, err := customsearch.New(client)
	if err != nil {
		fmt.Println("Could not create customsearch client")
		return ""
	}

	resp, err := svc.Cse.Siterestrict.List(query).Num(1).Fields("items/link").Cx(config.Cfg.CX).Do()
	if err != nil {
		fmt.Println("Could not recieve customsearch results")
		return ""
	}
	return resp.Items[0].Link
}