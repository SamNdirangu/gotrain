package contextTrain

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func Context() {
	example1()
}

func example1() {
	//This example shows the working via http request

	timeoutContext, cancel := context.WithTimeout(context.Background(), time.Millisecond*100)
	defer cancel()

	req, err := http.NewRequestWithContext(timeoutContext, http.MethodGet, "https://miro.medium.com/max/500/0*M36jfoWgm1tj8KD-", nil)
	if err != nil {
		panic(err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	//get data from http reponse
	imageData, err := ioutil.ReadAll(res.Body)
	fmt.Printf("Downloaded image of size %d\n", len(imageData))
}
