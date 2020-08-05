# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

from transaction.v3 import transaction_service_pb2 as transaction_dot_v3_dot_transaction__service__pb2


class TransactionStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.GetHistory = channel.unary_unary(
                '/kin.agora.transaction.v3.Transaction/GetHistory',
                request_serializer=transaction_dot_v3_dot_transaction__service__pb2.GetHistoryRequest.SerializeToString,
                response_deserializer=transaction_dot_v3_dot_transaction__service__pb2.GetHistoryResponse.FromString,
                )
        self.SubmitTransaction = channel.unary_unary(
                '/kin.agora.transaction.v3.Transaction/SubmitTransaction',
                request_serializer=transaction_dot_v3_dot_transaction__service__pb2.SubmitTransactionRequest.SerializeToString,
                response_deserializer=transaction_dot_v3_dot_transaction__service__pb2.SubmitTransactionResponse.FromString,
                )
        self.GetTransaction = channel.unary_unary(
                '/kin.agora.transaction.v3.Transaction/GetTransaction',
                request_serializer=transaction_dot_v3_dot_transaction__service__pb2.GetTransactionRequest.SerializeToString,
                response_deserializer=transaction_dot_v3_dot_transaction__service__pb2.GetTransactionResponse.FromString,
                )


class TransactionServicer(object):
    """Missing associated documentation comment in .proto file."""

    def GetHistory(self, request, context):
        """GetHistory returns the transaction history for an account,
        with additional off-chain invoice data, if available.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def SubmitTransaction(self, request, context):
        """SubmitTransaction submits a transaction.

        If the memo does not conform to the Kin binary memo format,
        the transaction is not eligible for whitelisting.

        If the memo _does_ conform to the Kin binary memo format,
        the transaction may be whitelisted depending on app
        configuration.

        See: https://github.com/kinecosystem/agora-api/blob/master/spec/memo.md
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetTransaction(self, request, context):
        """GetTransaction returns a transaction and additional off-chain
        invoice data, if available.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_TransactionServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'GetHistory': grpc.unary_unary_rpc_method_handler(
                    servicer.GetHistory,
                    request_deserializer=transaction_dot_v3_dot_transaction__service__pb2.GetHistoryRequest.FromString,
                    response_serializer=transaction_dot_v3_dot_transaction__service__pb2.GetHistoryResponse.SerializeToString,
            ),
            'SubmitTransaction': grpc.unary_unary_rpc_method_handler(
                    servicer.SubmitTransaction,
                    request_deserializer=transaction_dot_v3_dot_transaction__service__pb2.SubmitTransactionRequest.FromString,
                    response_serializer=transaction_dot_v3_dot_transaction__service__pb2.SubmitTransactionResponse.SerializeToString,
            ),
            'GetTransaction': grpc.unary_unary_rpc_method_handler(
                    servicer.GetTransaction,
                    request_deserializer=transaction_dot_v3_dot_transaction__service__pb2.GetTransactionRequest.FromString,
                    response_serializer=transaction_dot_v3_dot_transaction__service__pb2.GetTransactionResponse.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'kin.agora.transaction.v3.Transaction', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class Transaction(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def GetHistory(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/kin.agora.transaction.v3.Transaction/GetHistory',
            transaction_dot_v3_dot_transaction__service__pb2.GetHistoryRequest.SerializeToString,
            transaction_dot_v3_dot_transaction__service__pb2.GetHistoryResponse.FromString,
            options, channel_credentials,
            call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def SubmitTransaction(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/kin.agora.transaction.v3.Transaction/SubmitTransaction',
            transaction_dot_v3_dot_transaction__service__pb2.SubmitTransactionRequest.SerializeToString,
            transaction_dot_v3_dot_transaction__service__pb2.SubmitTransactionResponse.FromString,
            options, channel_credentials,
            call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def GetTransaction(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/kin.agora.transaction.v3.Transaction/GetTransaction',
            transaction_dot_v3_dot_transaction__service__pb2.GetTransactionRequest.SerializeToString,
            transaction_dot_v3_dot_transaction__service__pb2.GetTransactionResponse.FromString,
            options, channel_credentials,
            call_credentials, compression, wait_for_ready, timeout, metadata)
