#!/usr/bin/env python

import pika
from pymongo import MongoClient

connection = pika.BlockingConnection(pika.ConnectionParameters('some-rabbit'))
channel = connection.channel()
channel.exchange_declare(exchange='events', exchange_type='fanout')
result = channel.queue_declare(exclusive=True)
queue_name = result.method.queue
channel.queue_bind(exchange='events', queue=queue_name)

client = MongoClient('mongodb://some-mongo:27017/')
db = client.offsite
col = db.events

def callback(ch, method, properties, body):
    print(" [x] Received %r" % body)
    col.insert_one({"text": body.decode('ascii')})
    

channel.basic_consume(callback,
                      queue=queue_name,
                      no_ack=True)

print(' [*] Waiting for messages. To exit press CTRL+C')
channel.start_consuming()