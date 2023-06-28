package kafka

import (
	"context"

	"github.com/segmentio/kafka-go"
)

const (
	kafka_Topic = "command-log"
	kafka_URL   = "localhost:9092"
)

// kafka writer
func GetKafKaWriter() *kafka.Writer {

	return &kafka.Writer{
		Addr:     kafka.TCP(kafka_URL),
		Topic:    kafka_Topic,
		Balancer: &kafka.LeastBytes{},
	}
}

// kafka Reader
func GetKafKaReader() *kafka.Reader {
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{kafka_URL},
		Topic:     kafka_Topic,
		Partition: 0,
	})

}

func AppendCommandLog(ctx context.Context, kafkawriter *kafka.Writer, key []byte, value []byte) error {
	message := kafka.Message{
		Key:   key,
		Value: value,
	}

	err := kafkawriter.WriteMessages(ctx, message)
	return err
}
