from kafka import KafkaProducer
from kafka_config import read_ccloud_config
from confluent_kafka import Producer


producer = Producer(read_ccloud_config("client.properties"))


def produce(key,data=""):
    image_data = open('co2.jpg',"rb").read()
    producer.produce("topic_0", key=key, value=image_data)