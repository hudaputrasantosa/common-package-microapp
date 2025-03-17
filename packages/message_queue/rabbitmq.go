package messagequeue

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

type ConfigConnection struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// RabbitMQConfig struct untuk menyimpan koneksi
type RabbitMQConfig struct {
	Connection *amqp.Connection
	Channel    *amqp.Channel
}

// NewRabbitMQConnection membuat koneksi ke RabbitMQ
func NewRabbitMQConnection(cc ConfigConnection) (*RabbitMQConfig, error) {
	connAddr := fmt.Sprintf(
		"amqp://%s:%s@%s:%s/",
		cc.Username,
		cc.Password,
		cc.Host,
		cc.Port,
	)

	conn, err := amqp.Dial(connAddr)
	if err != nil {
		log.Fatalf("Gagal terhubung ke RabbitMQ: %v", err)
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Gagal membuka channel: %v", err)
		conn.Close()
		return nil, err
	}

	return &RabbitMQConfig{Connection: conn, Channel: ch}, nil
}

// Close menutup koneksi dan channel
func (r *RabbitMQConfig) Close() {
	r.Channel.Close()
	r.Connection.Close()
}
