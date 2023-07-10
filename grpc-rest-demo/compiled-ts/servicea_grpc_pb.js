// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('@grpc/grpc-js');
var servicea_pb = require('./servicea_pb.js');

function serialize_dev_jponc_grpc_rest_servicea_v1_PingRequest(arg) {
  if (!(arg instanceof servicea_pb.PingRequest)) {
    throw new Error('Expected argument of type dev.jponc.grpc_rest.servicea.v1.PingRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_dev_jponc_grpc_rest_servicea_v1_PingRequest(buffer_arg) {
  return servicea_pb.PingRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_dev_jponc_grpc_rest_servicea_v1_PingResponse(arg) {
  if (!(arg instanceof servicea_pb.PingResponse)) {
    throw new Error('Expected argument of type dev.jponc.grpc_rest.servicea.v1.PingResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_dev_jponc_grpc_rest_servicea_v1_PingResponse(buffer_arg) {
  return servicea_pb.PingResponse.deserializeBinary(new Uint8Array(buffer_arg));
}


var ServiceAService = exports.ServiceAService = {
  ping: {
    path: '/dev.jponc.grpc_rest.servicea.v1.ServiceA/Ping',
    requestStream: false,
    responseStream: false,
    requestType: servicea_pb.PingRequest,
    responseType: servicea_pb.PingResponse,
    requestSerialize: serialize_dev_jponc_grpc_rest_servicea_v1_PingRequest,
    requestDeserialize: deserialize_dev_jponc_grpc_rest_servicea_v1_PingRequest,
    responseSerialize: serialize_dev_jponc_grpc_rest_servicea_v1_PingResponse,
    responseDeserialize: deserialize_dev_jponc_grpc_rest_servicea_v1_PingResponse,
  },
};

exports.ServiceAClient = grpc.makeGenericClientConstructor(ServiceAService);
