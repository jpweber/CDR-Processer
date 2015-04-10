/*
* @Author: jpweber
* @Date:   2015-01-29 15:12:44
* @Last Modified by:   jpweber
* @Last Modified time: 2015-01-30 10:35:48
 */

package Queue

import (
	"fmt"
	"github.com/streadway/amqp"
	// "log"
	"time"
)

func Connect() *amqp.Channel {

	// Connects opens an AMQP connection from the credentials in the URL.
	conn, err := amqp.Dial("amqp://mq01.bluetonecommunications.com:5672/")
	if err != nil {
		// log.Fatalf("connection.open: %s", err)
		fmt.Printf("connection.open: %s", err)
	}

	// This waits for a server acknowledgment which means the sockets will have
	// flushed all outbound publishings prior to returning.  It's important to
	// block on Close to not lose any publishings.
	defer conn.Close()

	c, err := conn.Channel()
	if err != nil {
		// log.Fatalf("channel.open: %s", err)
		fmt.Printf("channel.open: %s", err)
	}

	// We declare our topology on both the publisher and consumer to ensure they
	// are the same.  This is part of AMQP being a programmable messaging model.
	//
	// See the Channel.Consume example for the complimentary declare.
	//ExchangeDeclare(name, kind string, durable, autoDelete, internal, noWait bool, args Table)
	err = c.ExchangeDeclare("gotests2", "topic", false, true, false, false, nil)
	if err != nil {
		// log.Fatalf("exchange.declare: %v", err)
		fmt.Printf("exchange.declare: %v", err)
	}

	// Prepare this message to be persistent.  Your publishing requirements may
	// be different.
	// msg := amqp.Publishing{
	// 	DeliveryMode: amqp.Persistent,
	// 	Timestamp:    time.Now(),
	// 	ContentType:  "text/plain",
	// 	Body:         []byte("Go Go AMQP!"),
	// }

	return c
}

func Publish(c amqp.Channel, payload string) {

	msg := amqp.Publishing{
		DeliveryMode: amqp.Persistent,
		Timestamp:    time.Now(),
		ContentType:  "text/plain",
		Body:         []byte(payload),
	}

	fmt.Println(c)

	// This is not a mandatory delivery, so it will be dropped if there are no
	// queues bound to the logs exchange.
	err := c.Publish("gotests2", "cdr.gotests.foo", false, false, msg)
	if err != nil {
		// Since publish is asynchronous this can happen if the network connection
		// is reset or if the server has run out of resources.
		// log.Fatalf("basic.publish: %v", err)
		fmt.Printf("basic.publish: %v", err)
	}
}
