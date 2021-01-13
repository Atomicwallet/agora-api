# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

from account.v4 import account_service_pb2 as account_dot_v4_dot_account__service__pb2


class AccountStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.CreateAccount = channel.unary_unary(
                '/kin.agora.account.v4.Account/CreateAccount',
                request_serializer=account_dot_v4_dot_account__service__pb2.CreateAccountRequest.SerializeToString,
                response_deserializer=account_dot_v4_dot_account__service__pb2.CreateAccountResponse.FromString,
                )
        self.GetAccountInfo = channel.unary_unary(
                '/kin.agora.account.v4.Account/GetAccountInfo',
                request_serializer=account_dot_v4_dot_account__service__pb2.GetAccountInfoRequest.SerializeToString,
                response_deserializer=account_dot_v4_dot_account__service__pb2.GetAccountInfoResponse.FromString,
                )
        self.ResolveTokenAccounts = channel.unary_unary(
                '/kin.agora.account.v4.Account/ResolveTokenAccounts',
                request_serializer=account_dot_v4_dot_account__service__pb2.ResolveTokenAccountsRequest.SerializeToString,
                response_deserializer=account_dot_v4_dot_account__service__pb2.ResolveTokenAccountsResponse.FromString,
                )
        self.GetEvents = channel.unary_stream(
                '/kin.agora.account.v4.Account/GetEvents',
                request_serializer=account_dot_v4_dot_account__service__pb2.GetEventsRequest.SerializeToString,
                response_deserializer=account_dot_v4_dot_account__service__pb2.Events.FromString,
                )


class AccountServicer(object):
    """Missing associated documentation comment in .proto file."""

    def CreateAccount(self, request, context):
        """CreateAccount creates a kin token account.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetAccountInfo(self, request, context):
        """GetAccountInfo returns the information of a specified account.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def ResolveTokenAccounts(self, request, context):
        """ResolveTokenAccounts resolves a set of Token Accounts owned by the specified account ID.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetEvents(self, request, context):
        """GetEvents returns a stream of events related to the specified account.

        Note: Only events occurring after stream initiation will be returned.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_AccountServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'CreateAccount': grpc.unary_unary_rpc_method_handler(
                    servicer.CreateAccount,
                    request_deserializer=account_dot_v4_dot_account__service__pb2.CreateAccountRequest.FromString,
                    response_serializer=account_dot_v4_dot_account__service__pb2.CreateAccountResponse.SerializeToString,
            ),
            'GetAccountInfo': grpc.unary_unary_rpc_method_handler(
                    servicer.GetAccountInfo,
                    request_deserializer=account_dot_v4_dot_account__service__pb2.GetAccountInfoRequest.FromString,
                    response_serializer=account_dot_v4_dot_account__service__pb2.GetAccountInfoResponse.SerializeToString,
            ),
            'ResolveTokenAccounts': grpc.unary_unary_rpc_method_handler(
                    servicer.ResolveTokenAccounts,
                    request_deserializer=account_dot_v4_dot_account__service__pb2.ResolveTokenAccountsRequest.FromString,
                    response_serializer=account_dot_v4_dot_account__service__pb2.ResolveTokenAccountsResponse.SerializeToString,
            ),
            'GetEvents': grpc.unary_stream_rpc_method_handler(
                    servicer.GetEvents,
                    request_deserializer=account_dot_v4_dot_account__service__pb2.GetEventsRequest.FromString,
                    response_serializer=account_dot_v4_dot_account__service__pb2.Events.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'kin.agora.account.v4.Account', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class Account(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def CreateAccount(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/kin.agora.account.v4.Account/CreateAccount',
            account_dot_v4_dot_account__service__pb2.CreateAccountRequest.SerializeToString,
            account_dot_v4_dot_account__service__pb2.CreateAccountResponse.FromString,
            options, channel_credentials,
            call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def GetAccountInfo(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/kin.agora.account.v4.Account/GetAccountInfo',
            account_dot_v4_dot_account__service__pb2.GetAccountInfoRequest.SerializeToString,
            account_dot_v4_dot_account__service__pb2.GetAccountInfoResponse.FromString,
            options, channel_credentials,
            call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def ResolveTokenAccounts(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/kin.agora.account.v4.Account/ResolveTokenAccounts',
            account_dot_v4_dot_account__service__pb2.ResolveTokenAccountsRequest.SerializeToString,
            account_dot_v4_dot_account__service__pb2.ResolveTokenAccountsResponse.FromString,
            options, channel_credentials,
            call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def GetEvents(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_stream(request, target, '/kin.agora.account.v4.Account/GetEvents',
            account_dot_v4_dot_account__service__pb2.GetEventsRequest.SerializeToString,
            account_dot_v4_dot_account__service__pb2.Events.FromString,
            options, channel_credentials,
            call_credentials, compression, wait_for_ready, timeout, metadata)
