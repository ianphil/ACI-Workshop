using System;
using RabbitMQ.Client;
using System.Text;
using System.Threading;

namespace send
{
    class Send
    {
        public static int Fibonacci(int n)
        {
            int a = 0;
            int b = 1;
            
            for (int i = 0; i < n; i++)
            {
                int temp = a;
                a = b;
                b = temp + b;
            }
            return a;
        }

        static void Main(string[] args)
        {
            var factory = new ConnectionFactory { HostName = "localhost" };
            // var factory = new ConnectionFactory { HostName = "some-rabbit" };
            using (var connection = factory.CreateConnection())
            {
                using (var channel = connection.CreateModel())
                {
                    channel.ExchangeDeclare(exchange: "events", type: "fanout");
                    
                    for (int i = 0; i < 150; i++)
                    {
                        string message = Fibonacci(i).ToString();
                        var body = Encoding.UTF8.GetBytes(message);

                        channel.BasicPublish(exchange: "events",
                                            routingKey: "",
                                            basicProperties: null,
                                            body: body);
                        
                        Console.WriteLine(" [x] sent {0}", message);
                        Thread.Sleep(1000);
                    }
                }
            }

            Console.WriteLine(" Press [enter] to exit.");
            Console.ReadLine();
        }
    }
}
