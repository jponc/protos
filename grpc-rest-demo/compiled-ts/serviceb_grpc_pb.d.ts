// package: dev.jponc.grpc_rest.serviceb.v1
// file: serviceb.proto

/* tslint:disable */
/* eslint-disable */

import * as grpc from "@grpc/grpc-js";
import * as serviceb_pb from "./serviceb_pb";

interface IServiceBService extends grpc.ServiceDefinition<grpc.UntypedServiceImplementation> {
    pong: IServiceBService_IPong;
}

interface IServiceBService_IPong extends grpc.MethodDefinition<serviceb_pb.PongRequest, serviceb_pb.PongResponse> {
    path: "/dev.jponc.grpc_rest.serviceb.v1.ServiceB/Pong";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<serviceb_pb.PongRequest>;
    requestDeserialize: grpc.deserialize<serviceb_pb.PongRequest>;
    responseSerialize: grpc.serialize<serviceb_pb.PongResponse>;
    responseDeserialize: grpc.deserialize<serviceb_pb.PongResponse>;
}

export const ServiceBService: IServiceBService;

export interface IServiceBServer extends grpc.UntypedServiceImplementation {
    pong: grpc.handleUnaryCall<serviceb_pb.PongRequest, serviceb_pb.PongResponse>;
}

export interface IServiceBClient {
    pong(request: serviceb_pb.PongRequest, callback: (error: grpc.ServiceError | null, response: serviceb_pb.PongResponse) => void): grpc.ClientUnaryCall;
    pong(request: serviceb_pb.PongRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: serviceb_pb.PongResponse) => void): grpc.ClientUnaryCall;
    pong(request: serviceb_pb.PongRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: serviceb_pb.PongResponse) => void): grpc.ClientUnaryCall;
}

export class ServiceBClient extends grpc.Client implements IServiceBClient {
    constructor(address: string, credentials: grpc.ChannelCredentials, options?: Partial<grpc.ClientOptions>);
    public pong(request: serviceb_pb.PongRequest, callback: (error: grpc.ServiceError | null, response: serviceb_pb.PongResponse) => void): grpc.ClientUnaryCall;
    public pong(request: serviceb_pb.PongRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: serviceb_pb.PongResponse) => void): grpc.ClientUnaryCall;
    public pong(request: serviceb_pb.PongRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: serviceb_pb.PongResponse) => void): grpc.ClientUnaryCall;
}
