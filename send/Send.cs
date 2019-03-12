using System;
using RabbitMQ.Client;
using System.Text;

namespace send
{
    class Send
    {
        static void Main(string[] args)
        {
            var factory = new ConnectionFactory { HostName = "localhost" };
            using (var connection = factory.CreateConnection())
            {
                using (var channel = connection.CreateModel())
                {
                    channel.ExchangeDeclare(exchange: "events", type: "fanout");
                    
                    // TODO: Send Fibonacci sequence 
                    string message = "Hello world!";
                    var body = Encoding.UTF8.GetBytes(message);

                    channel.BasicPublish(exchange: "events",
                                        routingKey: "",
                                        basicProperties: null,
                                        body: body);
                    
                    Console.WriteLine(" [x] sent {0}", message);
                }
            }

            Console.WriteLine(" Press [enter] to exit.");
            Console.ReadLine();
        }
    }
}
