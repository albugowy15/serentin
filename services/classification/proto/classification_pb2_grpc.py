# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

from proto import classification_pb2 as classification__pb2

class ClassificationStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.PredictStressLevel = channel.unary_unary(
                '/Classification/PredictStressLevel',
                request_serializer=classification__pb2.StressLevelRequest.SerializeToString,
                response_deserializer=classification__pb2.StressLevelResponse.FromString,
                )


class ClassificationServicer(object):
    """Missing associated documentation comment in .proto file."""

    def PredictStressLevel(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_ClassificationServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'PredictStressLevel': grpc.unary_unary_rpc_method_handler(
                    servicer.PredictStressLevel,
                    request_deserializer=classification__pb2.StressLevelRequest.FromString,
                    response_serializer=classification__pb2.StressLevelResponse.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'Classification', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class Classification(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def PredictStressLevel(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/Classification/PredictStressLevel',
            classification__pb2.StressLevelRequest.SerializeToString,
            classification__pb2.StressLevelResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)