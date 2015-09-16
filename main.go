package main

import(
  "log"
  "net/http"
)

func main() {
	InitDb()
 	router := NewRouter()
	log.Fatal(http.ListenAndServe(":8081", router))

}

/*func main() {
  app := cli.NewApp()
  app.Name = "Todo"
  app.Usage = "Todo json APi in go"
  app.Action = func(c *cli.Context) {
    initApp(c.Args()[0])
  }
  app.Run(os.Args)
}*/