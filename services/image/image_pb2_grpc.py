# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

import image_pb2 as image__pb2


class UploadImageStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.Upload = channel.stream_unary(
                '/image.UploadImage/Upload',
                request_serializer=image__pb2.UploadRequest.SerializeToString,
                response_deserializer=image__pb2.UploadResponse.FromString,
                )


class UploadImageServicer(object):
    """Missing associated documentation comment in .proto file."""

    def Upload(self, request_iterator, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_UploadImageServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'Upload': grpc.stream_unary_rpc_method_handler(
                    servicer.Upload,
                    request_deserializer=image__pb2.UploadRequest.FromString,
                    response_serializer=image__pb2.UploadResponse.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'image.UploadImage', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class UploadImage(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def Upload(request_iterator,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.stream_unary(request_iterator, target, '/image.UploadImage/Upload',
            image__pb2.UploadRequest.SerializeToString,
            image__pb2.UploadResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)