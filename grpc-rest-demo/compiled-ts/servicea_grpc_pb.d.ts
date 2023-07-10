// package: dev.jponc.grpc_rest.servicea.v1
// file: servicea.proto

/* tslint:disable */
/* eslint-disable */

import * as grpc from "@grpc/grpc-js";
import * as servicea_pb from "./servicea_pb";

interface IServiceAService extends grpc.ServiceDefinition<grpc.UntypedServiceImplementation> {
    ping: IServiceAService_IPing;
}

interface IServiceAService_IPing extends grpc.MethodDefinition<servicea_pb.PingRequest, servicea_pb.PingResponse> {
    path: "/dev.jponc.grpc_rest.servicea.v1.ServiceA/Ping";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<servicea_pb.PingRequest>;
    requestDeserialize: grpc.deserialize<servicea_pb.PingRequest>;
    responseSerialize: grpc.serialize<servicea_pb.PingResponse>;
    responseDeserialize: grpc.deserialize<servicea_pb.PingResponse>;
}

export const ServiceAService: IServiceAService;

export interface IServiceAServer extends grpc.UntypedServiceImplementation {
    ping: grpc.handleUnaryCall<servicea_pb.PingRequest, servicea_pb.PingResponse>;
}

export interface IServiceAClient {
    ping(request: servicea_pb.PingRequest, callback: (error: grpc.ServiceError | null, response: servicea_pb.PingResponse) => void): grpc.ClientUnaryCall;
    ping(request: servicea_pb.PingRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: servicea_pb.PingResponse) => void): grpc.ClientUnaryCall;
    ping(request: servicea_pb.PingRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: servicea_pb.PingResponse) => void): grpc.ClientUnaryCall;
}

export class ServiceAClient extends grpc.Client implements IServiceAClient {
    constructor(address: string, credentials: grpc.ChannelCredentials, options?: Partial<grpc.ClientOptions>);
    public ping(request: servicea_pb.PingRequest, callback: (error: grpc.ServiceError | null, response: servicea_pb.PingResponse) => void): grpc.ClientUnaryCall;
    public ping(request: servicea_pb.PingRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: servicea_pb.PingResponse) => void): grpc.ClientUnaryCall;
    public ping(request: servicea_pb.PingRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: servicea_pb.PingResponse) => void): grpc.ClientUnaryCall;
}
