#!/usr/bin/env python

import pika
from pymongo import MongoClient

connection = pika.BlockingConnection(pika.ConnectionParameters('localhost'))
channel = connection.channel()
channel.queue_declare(queue='hello')

client = MongoClient('mongodb://localhost:27017/')
db = client.offsite
col = db.events

def callback(ch, method, properties, body):
    print(" [x] Received %r" % body)
    col.insert_one({"text": body.decode('ascii')})
    

channel.basic_consume(callback,
                      queue='hello',
                      no_ack=True)

print(' [*] Waiting for messages. To exit press CTRL+C')
channel.start_consuming()