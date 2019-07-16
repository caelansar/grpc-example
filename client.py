#! /usr/bin/env python3
# -*- coding: utf-8 -*-

import grpc

import hello_pb2
import hello_pb2_grpc


def run():
    with grpc.insecure_channel('localhost:50051') as channel:  # create channel
        stub = hello_pb2_grpc.GreeterStub(channel)  # create stub to call service methods
        response = stub.SayHello(hello_pb2.HelloRequest(name='niconiconi'))  # call service methods
        print("Greeter client received: " + response.message)


if __name__ == '__main__':
    run()
