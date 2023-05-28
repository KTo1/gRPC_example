import sys
sys.path.append('proto/reverse')

import grpc
import proto.reverse.reverse_pb2 as pb2
import proto.reverse.reverse_pb2_grpc as pb2_grpc


class ReverseMessageService:
    def __init__(self):
        self.channel = grpc.insecure_channel("192.168.25.101:5300")
        self.stub = pb2_grpc.ReverseStub(self.channel)

    def do(self, message: str):
        request = pb2.Request(message=message)
        response = self.stub.Do(request)
        return response


if __name__ == '__main__':
    rms = ReverseMessageService()
    for i in range(10):
        print(rms.do('Привет'))
