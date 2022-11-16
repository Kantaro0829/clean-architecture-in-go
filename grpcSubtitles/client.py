import grpc
import chat_pb2
import chat_pb2_grpc

def run():
    with grpc.insecure_channel('localhost:9000') as channel:
        #stub = hello_pb2_grpc.HelloWorldStub(channel)
        stub = chat_pb2_grpc.ChatServiceStub(channel)
        #response = stub.SayHello(hello_pb2.HelloRequest(name='Yamada'))
        response = stub.SayHello(chat_pb2.Message(body="kantaro"))
    print("RECV: %s" % response.body)

if __name__ == '__main__':
    run()