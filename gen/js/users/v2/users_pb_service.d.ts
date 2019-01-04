// package: gomesh.users.v2
// file: users/v2/users.proto

import * as users_v2_users_pb from "../../users/v2/users_pb";
import * as google_protobuf_empty_pb from "google-protobuf/google/protobuf/empty_pb";
import {grpc} from "grpc-web-client";

type UsersGet = {
  readonly methodName: string;
  readonly service: typeof Users;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof users_v2_users_pb.GetRequest;
  readonly responseType: typeof users_v2_users_pb.User;
};

type UsersCreate = {
  readonly methodName: string;
  readonly service: typeof Users;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof users_v2_users_pb.CreateRequest;
  readonly responseType: typeof users_v2_users_pb.User;
};

type UsersUpdate = {
  readonly methodName: string;
  readonly service: typeof Users;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof users_v2_users_pb.UpdateRequest;
  readonly responseType: typeof users_v2_users_pb.User;
};

type UsersDelete = {
  readonly methodName: string;
  readonly service: typeof Users;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof users_v2_users_pb.DeleteRequest;
  readonly responseType: typeof google_protobuf_empty_pb.Empty;
};

type UsersList = {
  readonly methodName: string;
  readonly service: typeof Users;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof users_v2_users_pb.ListRequest;
  readonly responseType: typeof users_v2_users_pb.ListResponse;
};

type UsersBatchGet = {
  readonly methodName: string;
  readonly service: typeof Users;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof users_v2_users_pb.BatchGetRequest;
  readonly responseType: typeof users_v2_users_pb.BatchGetResponse;
};

export class Users {
  static readonly serviceName: string;
  static readonly Get: UsersGet;
  static readonly Create: UsersCreate;
  static readonly Update: UsersUpdate;
  static readonly Delete: UsersDelete;
  static readonly List: UsersList;
  static readonly BatchGet: UsersBatchGet;
}

export type ServiceError = { message: string, code: number; metadata: grpc.Metadata }
export type Status = { details: string, code: number; metadata: grpc.Metadata }

interface UnaryResponse {
  cancel(): void;
}
interface ResponseStream<T> {
  cancel(): void;
  on(type: 'data', handler: (message: T) => void): ResponseStream<T>;
  on(type: 'end', handler: () => void): ResponseStream<T>;
  on(type: 'status', handler: (status: Status) => void): ResponseStream<T>;
}
interface RequestStream<T> {
  write(message: T): RequestStream<T>;
  end(): void;
  cancel(): void;
  on(type: 'end', handler: () => void): RequestStream<T>;
  on(type: 'status', handler: (status: Status) => void): RequestStream<T>;
}
interface BidirectionalStream<ReqT, ResT> {
  write(message: ReqT): BidirectionalStream<ReqT, ResT>;
  end(): void;
  cancel(): void;
  on(type: 'data', handler: (message: ResT) => void): BidirectionalStream<ReqT, ResT>;
  on(type: 'end', handler: () => void): BidirectionalStream<ReqT, ResT>;
  on(type: 'status', handler: (status: Status) => void): BidirectionalStream<ReqT, ResT>;
}

export class UsersClient {
  readonly serviceHost: string;

  constructor(serviceHost: string, options?: grpc.RpcOptions);
  get(
    requestMessage: users_v2_users_pb.GetRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: users_v2_users_pb.User|null) => void
  ): UnaryResponse;
  get(
    requestMessage: users_v2_users_pb.GetRequest,
    callback: (error: ServiceError|null, responseMessage: users_v2_users_pb.User|null) => void
  ): UnaryResponse;
  create(
    requestMessage: users_v2_users_pb.CreateRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: users_v2_users_pb.User|null) => void
  ): UnaryResponse;
  create(
    requestMessage: users_v2_users_pb.CreateRequest,
    callback: (error: ServiceError|null, responseMessage: users_v2_users_pb.User|null) => void
  ): UnaryResponse;
  update(
    requestMessage: users_v2_users_pb.UpdateRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: users_v2_users_pb.User|null) => void
  ): UnaryResponse;
  update(
    requestMessage: users_v2_users_pb.UpdateRequest,
    callback: (error: ServiceError|null, responseMessage: users_v2_users_pb.User|null) => void
  ): UnaryResponse;
  delete(
    requestMessage: users_v2_users_pb.DeleteRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: google_protobuf_empty_pb.Empty|null) => void
  ): UnaryResponse;
  delete(
    requestMessage: users_v2_users_pb.DeleteRequest,
    callback: (error: ServiceError|null, responseMessage: google_protobuf_empty_pb.Empty|null) => void
  ): UnaryResponse;
  list(
    requestMessage: users_v2_users_pb.ListRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: users_v2_users_pb.ListResponse|null) => void
  ): UnaryResponse;
  list(
    requestMessage: users_v2_users_pb.ListRequest,
    callback: (error: ServiceError|null, responseMessage: users_v2_users_pb.ListResponse|null) => void
  ): UnaryResponse;
  batchGet(
    requestMessage: users_v2_users_pb.BatchGetRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: users_v2_users_pb.BatchGetResponse|null) => void
  ): UnaryResponse;
  batchGet(
    requestMessage: users_v2_users_pb.BatchGetRequest,
    callback: (error: ServiceError|null, responseMessage: users_v2_users_pb.BatchGetResponse|null) => void
  ): UnaryResponse;
}

