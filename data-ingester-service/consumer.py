from PIL import Image
from io import BytesIO
from kafka_config import read_ccloud_config
from confluent_kafka import Consumer
import io
props = read_ccloud_config("client.properties")
props["group.id"] = "python-group-1"
props["auto.offset.reset"] = "earliest"

consumer = Consumer(props)
consumer.subscribe(["topic_0"])
def consume():
    try:
        while True:
            msg = consumer.poll(1.0)
            print(msg, type(msg))
            if msg and msg.error() is None:
                key = msg.key()
                data = msg.value()
                image_stream = io.BytesIO(data)

                # Open and display the image using Pillow
                img = Image.open(image_stream)
                img.show()
                # print("key = {key:12} value = {value:12}".format(key=msg.key(), value=msg.value().decode('utf-8')))
    except KeyboardInterrupt:
            pass
    finally:
            consumer.close()
# def consume():
#     print('Consuming')
#     consumer = KafkaConsumer('CO2_Emissions',bootstrap_servers=['localhost:9093'],api_version=(0,10,1))
#     for msg in consumer:
#         stream = BytesIO(msg.value)
#         image = Image.open(stream).convert("RGBA")
#         print(image.shape)
#         stream.close()
#         image.show()
        
#     # consumer.close()
#     return
if __name__=='__main__':
     consume()
