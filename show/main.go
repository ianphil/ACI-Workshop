package main

import (
	"flag"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/streadway/amqp"
)

var addr = flag.String("addr", ":8080", "http service address")

var upgrader = websocket.Upgrader{} // use default options

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func echo(w http.ResponseWriter, r *http.Request) {
	conn, err := amqp.Dial("amqp://guest:guest@some-rabbit:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	err = ch.ExchangeDeclare(
		"events", // name
		"fanout", // type
		false,    // durable
		false,    // auto-deleted
		false,    // internal
		false,    // no-wait
		nil,      // arguments
	)
	failOnError(err, "Failed to declare an exchange")

	q, err := ch.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when usused
		true,  // exclusive
		false, // no-wait
		nil,   // arguments
	)
	failOnError(err, "Failed to declare a queue")

	err = ch.QueueBind(
		q.Name,   // queue name
		"",       // routing key
		"events", // exchange
		false,
		nil)
	failOnError(err, "Failed to bind a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
			err = c.WriteMessage(1, d.Body)
			if err != nil {
				log.Println("write:", err)
				break
			}
		}
	}()

	<-forever
}

func home(w http.ResponseWriter, r *http.Request) {
	homeTemplate.Execute(w, "ws://"+r.Host+"/echo")
}

func main() {
	flag.Parse()
	log.SetFlags(0)
	http.HandleFunc("/echo", echo)
	http.HandleFunc("/", home)
	log.Fatal(http.ListenAndServe(*addr, nil))
}

var homeTemplate = template.Must(template.New("").Parse(`
<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8">
<script>  
window.addEventListener("load", function(evt) {

    var output = document.getElementById("output");
    var ws;

    var print = function(message) {
        var d = document.createElement("div");
        d.innerHTML = message;
        output.appendChild(d);
	};
	
	if (ws) {
		return false;
	}
	ws = new WebSocket("{{.}}");
	ws.onopen = function(evt) {
		print("OPEN");
	}
	ws.onclose = function(evt) {
		print("CLOSE");
		ws = null;
	}
	ws.onmessage = function(evt) {
		print("RESPONSE: " + evt.data);
	}
	ws.onerror = function(evt) {
		print("ERROR: " + evt.data);
	}

});
</script>
</head>
<body>
<table>
<tr><td valign="top" width="50%">
<div id="output"></div>
</td></tr></table>
</body>
</html>
`))
