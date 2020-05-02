// package: 
// file: mood.proto

import * as mood_pb from "./mood_pb";
import {grpc} from "@improbable-eng/grpc-web";

type MoodAddEntry = {
  readonly methodName: string;
  readonly service: typeof Mood;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof mood_pb.AddEntryRequest;
  readonly responseType: typeof mood_pb.AddEntryResponse;
};

type MoodGetMoodFromEntry = {
  readonly methodName: string;
  readonly service: typeof Mood;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof mood_pb.GetMoodFromEntryRequest;
  readonly responseType: typeof mood_pb.GetMoodFromEntryResponse;
};

type MoodGetMood = {
  readonly methodName: string;
  readonly service: typeof Mood;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof mood_pb.GetMoodRequest;
  readonly responseType: typeof mood_pb.GetMoodResponse;
};

type MoodCreateMood = {
  readonly methodName: string;
  readonly service: typeof Mood;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof mood_pb.CreateMoodRequest;
  readonly responseType: typeof mood_pb.CreateMoodResponse;
};

export class Mood {
  static readonly serviceName: string;
  static readonly AddEntry: MoodAddEntry;
  static readonly GetMoodFromEntry: MoodGetMoodFromEntry;
  static readonly GetMood: MoodGetMood;
  static readonly CreateMood: MoodCreateMood;
}

export type ServiceError = { message: string, code: number; metadata: grpc.Metadata }
export type Status = { details: string, code: number; metadata: grpc.Metadata }

interface UnaryResponse {
  cancel(): void;
}
interface ResponseStream<T> {
  cancel(): void;
  on(type: 'data', handler: (message: T) => void): ResponseStream<T>;
  on(type: 'end', handler: (status?: Status) => void): ResponseStream<T>;
  on(type: 'status', handler: (status: Status) => void): ResponseStream<T>;
}
interface RequestStream<T> {
  write(message: T): RequestStream<T>;
  end(): void;
  cancel(): void;
  on(type: 'end', handler: (status?: Status) => void): RequestStream<T>;
  on(type: 'status', handler: (status: Status) => void): RequestStream<T>;
}
interface BidirectionalStream<ReqT, ResT> {
  write(message: ReqT): BidirectionalStream<ReqT, ResT>;
  end(): void;
  cancel(): void;
  on(type: 'data', handler: (message: ResT) => void): BidirectionalStream<ReqT, ResT>;
  on(type: 'end', handler: (status?: Status) => void): BidirectionalStream<ReqT, ResT>;
  on(type: 'status', handler: (status: Status) => void): BidirectionalStream<ReqT, ResT>;
}

export class MoodClient {
  readonly serviceHost: string;

  constructor(serviceHost: string, options?: grpc.RpcOptions);
  addEntry(
    requestMessage: mood_pb.AddEntryRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: mood_pb.AddEntryResponse|null) => void
  ): UnaryResponse;
  addEntry(
    requestMessage: mood_pb.AddEntryRequest,
    callback: (error: ServiceError|null, responseMessage: mood_pb.AddEntryResponse|null) => void
  ): UnaryResponse;
  getMoodFromEntry(
    requestMessage: mood_pb.GetMoodFromEntryRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: mood_pb.GetMoodFromEntryResponse|null) => void
  ): UnaryResponse;
  getMoodFromEntry(
    requestMessage: mood_pb.GetMoodFromEntryRequest,
    callback: (error: ServiceError|null, responseMessage: mood_pb.GetMoodFromEntryResponse|null) => void
  ): UnaryResponse;
  getMood(
    requestMessage: mood_pb.GetMoodRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: mood_pb.GetMoodResponse|null) => void
  ): UnaryResponse;
  getMood(
    requestMessage: mood_pb.GetMoodRequest,
    callback: (error: ServiceError|null, responseMessage: mood_pb.GetMoodResponse|null) => void
  ): UnaryResponse;
  createMood(
    requestMessage: mood_pb.CreateMoodRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: mood_pb.CreateMoodResponse|null) => void
  ): UnaryResponse;
  createMood(
    requestMessage: mood_pb.CreateMoodRequest,
    callback: (error: ServiceError|null, responseMessage: mood_pb.CreateMoodResponse|null) => void
  ): UnaryResponse;
}

