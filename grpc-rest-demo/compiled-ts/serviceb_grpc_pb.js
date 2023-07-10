// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('@grpc/grpc-js');
var serviceb_pb = require('./serviceb_pb.js');

function serialize_dev_jponc_grpc_rest_serviceb_v1_PongRequest(arg) {
  if (!(arg instanceof serviceb_pb.PongRequest)) {
    throw new Error('Expected argument of type dev.jponc.grpc_rest.serviceb.v1.PongRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_dev_jponc_grpc_rest_serviceb_v1_PongRequest(buffer_arg) {
  return serviceb_pb.PongRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_dev_jponc_grpc_rest_serviceb_v1_PongResponse(arg) {
  if (!(arg instanceof serviceb_pb.PongResponse)) {
    throw new Error('Expected argument of type dev.jponc.grpc_rest.serviceb.v1.PongResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_dev_jponc_grpc_rest_serviceb_v1_PongResponse(buffer_arg) {
  return serviceb_pb.PongResponse.deserializeBinary(new Uint8Array(buffer_arg));
}


var ServiceBService = exports.ServiceBService = {
  pong: {
    path: '/dev.jponc.grpc_rest.serviceb.v1.ServiceB/Pong',
    requestStream: false,
    responseStream: false,
    requestType: serviceb_pb.PongRequest,
    responseType: serviceb_pb.PongResponse,
    requestSerialize: serialize_dev_jponc_grpc_rest_serviceb_v1_PongRequest,
    requestDeserialize: deserialize_dev_jponc_grpc_rest_serviceb_v1_PongRequest,
    responseSerialize: serialize_dev_jponc_grpc_rest_serviceb_v1_PongResponse,
    responseDeserialize: deserialize_dev_jponc_grpc_rest_serviceb_v1_PongResponse,
  },
};

exports.ServiceBClient = grpc.makeGenericClientConstructor(ServiceBService);
