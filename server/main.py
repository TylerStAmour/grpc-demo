import logging
import grpc
from concurrent import futures
from proto import demo_pb2
from proto import demo_pb2_grpc


class Paddle(demo_pb2_grpc.DemoServicer):
    def Ping(self, request, context):
        return demo_pb2.PongReply(message="Pong!")


def serve():
    port = "8080"
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    demo_pb2_grpc.add_DemoServicer_to_server(Paddle(), server)
    server.add_insecure_port("[::]:" + port)
    server.start()
    print("Listening on " + port)
    server.wait_for_termination()


if __name__ == "__main__":
    logging.basicConfig()
    serve()
