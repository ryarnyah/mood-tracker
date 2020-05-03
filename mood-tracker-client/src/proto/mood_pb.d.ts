// package: 
// file: mood.proto

import * as jspb from "google-protobuf";

export class Entry extends jspb.Message {
  getRecord(): number;
  setRecord(value: number): void;

  getComment(): string;
  setComment(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Entry.AsObject;
  static toObject(includeInstance: boolean, msg: Entry): Entry.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Entry, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Entry;
  static deserializeBinaryFromReader(message: Entry, reader: jspb.BinaryReader): Entry;
}

export namespace Entry {
  export type AsObject = {
    record: number,
    comment: string,
  }
}

export class AddEntryRequest extends jspb.Message {
  hasEntry(): boolean;
  clearEntry(): void;
  getEntry(): Entry | undefined;
  setEntry(value?: Entry): void;

  getMoodId(): number;
  setMoodId(value: number): void;

  getEntryAccessCode(): string;
  setEntryAccessCode(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): AddEntryRequest.AsObject;
  static toObject(includeInstance: boolean, msg: AddEntryRequest): AddEntryRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: AddEntryRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): AddEntryRequest;
  static deserializeBinaryFromReader(message: AddEntryRequest, reader: jspb.BinaryReader): AddEntryRequest;
}

export namespace AddEntryRequest {
  export type AsObject = {
    entry?: Entry.AsObject,
    moodId: number,
    entryAccessCode: string,
  }
}

export class AddEntryResponse extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): AddEntryResponse.AsObject;
  static toObject(includeInstance: boolean, msg: AddEntryResponse): AddEntryResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: AddEntryResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): AddEntryResponse;
  static deserializeBinaryFromReader(message: AddEntryResponse, reader: jspb.BinaryReader): AddEntryResponse;
}

export namespace AddEntryResponse {
  export type AsObject = {
  }
}

export class GetMoodFromEntryRequest extends jspb.Message {
  getMoodId(): number;
  setMoodId(value: number): void;

  getEntryAccessCode(): string;
  setEntryAccessCode(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetMoodFromEntryRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetMoodFromEntryRequest): GetMoodFromEntryRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: GetMoodFromEntryRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetMoodFromEntryRequest;
  static deserializeBinaryFromReader(message: GetMoodFromEntryRequest, reader: jspb.BinaryReader): GetMoodFromEntryRequest;
}

export namespace GetMoodFromEntryRequest {
  export type AsObject = {
    moodId: number,
    entryAccessCode: string,
  }
}

export class GetMoodFromEntryResponse extends jspb.Message {
  getTitle(): string;
  setTitle(value: string): void;

  getContent(): string;
  setContent(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetMoodFromEntryResponse.AsObject;
  static toObject(includeInstance: boolean, msg: GetMoodFromEntryResponse): GetMoodFromEntryResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: GetMoodFromEntryResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetMoodFromEntryResponse;
  static deserializeBinaryFromReader(message: GetMoodFromEntryResponse, reader: jspb.BinaryReader): GetMoodFromEntryResponse;
}

export namespace GetMoodFromEntryResponse {
  export type AsObject = {
    title: string,
    content: string,
  }
}

export class GetMoodRequest extends jspb.Message {
  getMoodId(): number;
  setMoodId(value: number): void;

  getMoodAccessCode(): string;
  setMoodAccessCode(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetMoodRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetMoodRequest): GetMoodRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: GetMoodRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetMoodRequest;
  static deserializeBinaryFromReader(message: GetMoodRequest, reader: jspb.BinaryReader): GetMoodRequest;
}

export namespace GetMoodRequest {
  export type AsObject = {
    moodId: number,
    moodAccessCode: string,
  }
}

export class GetMoodResponse extends jspb.Message {
  clearEntriesList(): void;
  getEntriesList(): Array<Entry>;
  setEntriesList(value: Array<Entry>): void;
  addEntries(value?: Entry, index?: number): Entry;

  getStatsMap(): jspb.Map<number, number>;
  clearStatsMap(): void;
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetMoodResponse.AsObject;
  static toObject(includeInstance: boolean, msg: GetMoodResponse): GetMoodResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: GetMoodResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetMoodResponse;
  static deserializeBinaryFromReader(message: GetMoodResponse, reader: jspb.BinaryReader): GetMoodResponse;
}

export namespace GetMoodResponse {
  export type AsObject = {
    entriesList: Array<Entry.AsObject>,
    statsMap: Array<[number, number]>,
  }
}

export class CreateMoodRequest extends jspb.Message {
  getTitle(): string;
  setTitle(value: string): void;

  getContent(): string;
  setContent(value: string): void;

  getNumberOfRecordsNeeded(): number;
  setNumberOfRecordsNeeded(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateMoodRequest.AsObject;
  static toObject(includeInstance: boolean, msg: CreateMoodRequest): CreateMoodRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: CreateMoodRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateMoodRequest;
  static deserializeBinaryFromReader(message: CreateMoodRequest, reader: jspb.BinaryReader): CreateMoodRequest;
}

export namespace CreateMoodRequest {
  export type AsObject = {
    title: string,
    content: string,
    numberOfRecordsNeeded: number,
  }
}

export class CreateMoodResponse extends jspb.Message {
  getMoodId(): number;
  setMoodId(value: number): void;

  getMoodAccessCode(): string;
  setMoodAccessCode(value: string): void;

  clearEntriesAccessCodesList(): void;
  getEntriesAccessCodesList(): Array<string>;
  setEntriesAccessCodesList(value: Array<string>): void;
  addEntriesAccessCodes(value: string, index?: number): string;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateMoodResponse.AsObject;
  static toObject(includeInstance: boolean, msg: CreateMoodResponse): CreateMoodResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: CreateMoodResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateMoodResponse;
  static deserializeBinaryFromReader(message: CreateMoodResponse, reader: jspb.BinaryReader): CreateMoodResponse;
}

export namespace CreateMoodResponse {
  export type AsObject = {
    moodId: number,
    moodAccessCode: string,
    entriesAccessCodesList: Array<string>,
  }
}

