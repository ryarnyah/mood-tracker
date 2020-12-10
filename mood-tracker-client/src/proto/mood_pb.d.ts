// package: 
// file: mood.proto

import * as jspb from "google-protobuf";
import * as google_protobuf_timestamp_pb from "google-protobuf/google/protobuf/timestamp_pb";

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

export class EntryWithDate extends jspb.Message {
  getRecord(): number;
  setRecord(value: number): void;

  getComment(): string;
  setComment(value: string): void;

  hasRecordEntry(): boolean;
  clearRecordEntry(): void;
  getRecordEntry(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setRecordEntry(value?: google_protobuf_timestamp_pb.Timestamp): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): EntryWithDate.AsObject;
  static toObject(includeInstance: boolean, msg: EntryWithDate): EntryWithDate.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: EntryWithDate, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): EntryWithDate;
  static deserializeBinaryFromReader(message: EntryWithDate, reader: jspb.BinaryReader): EntryWithDate;
}

export namespace EntryWithDate {
  export type AsObject = {
    record: number,
    comment: string,
    recordEntry?: google_protobuf_timestamp_pb.Timestamp.AsObject,
  }
}

export class AddEntryRequest extends jspb.Message {
  hasEntry(): boolean;
  clearEntry(): void;
  getEntry(): Entry | undefined;
  setEntry(value?: Entry): void;

  getMoodId(): string;
  setMoodId(value: string): void;

  getEntryId(): string;
  setEntryId(value: string): void;

  getEntrySignature(): string;
  setEntrySignature(value: string): void;

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
    moodId: string,
    entryId: string,
    entrySignature: string,
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
  getMoodId(): string;
  setMoodId(value: string): void;

  getEntryId(): string;
  setEntryId(value: string): void;

  getEntrySignature(): string;
  setEntrySignature(value: string): void;

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
    moodId: string,
    entryId: string,
    entrySignature: string,
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
  getMoodId(): string;
  setMoodId(value: string): void;

  getMoodSignature(): string;
  setMoodSignature(value: string): void;

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
    moodId: string,
    moodSignature: string,
  }
}

export class RecordStat extends jspb.Message {
  hasRecordEntry(): boolean;
  clearRecordEntry(): void;
  getRecordEntry(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setRecordEntry(value?: google_protobuf_timestamp_pb.Timestamp): void;

  getCount(): number;
  setCount(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): RecordStat.AsObject;
  static toObject(includeInstance: boolean, msg: RecordStat): RecordStat.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: RecordStat, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): RecordStat;
  static deserializeBinaryFromReader(message: RecordStat, reader: jspb.BinaryReader): RecordStat;
}

export namespace RecordStat {
  export type AsObject = {
    recordEntry?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    count: number,
  }
}

export class MoodStat extends jspb.Message {
  getRecord(): number;
  setRecord(value: number): void;

  clearRecordStatsList(): void;
  getRecordStatsList(): Array<RecordStat>;
  setRecordStatsList(value: Array<RecordStat>): void;
  addRecordStats(value?: RecordStat, index?: number): RecordStat;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): MoodStat.AsObject;
  static toObject(includeInstance: boolean, msg: MoodStat): MoodStat.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: MoodStat, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): MoodStat;
  static deserializeBinaryFromReader(message: MoodStat, reader: jspb.BinaryReader): MoodStat;
}

export namespace MoodStat {
  export type AsObject = {
    record: number,
    recordStatsList: Array<RecordStat.AsObject>,
  }
}

export class GetMoodResponse extends jspb.Message {
  getTitle(): string;
  setTitle(value: string): void;

  getContent(): string;
  setContent(value: string): void;

  clearEntriesList(): void;
  getEntriesList(): Array<EntryWithDate>;
  setEntriesList(value: Array<EntryWithDate>): void;
  addEntries(value?: EntryWithDate, index?: number): EntryWithDate;

  clearStatsList(): void;
  getStatsList(): Array<MoodStat>;
  setStatsList(value: Array<MoodStat>): void;
  addStats(value?: MoodStat, index?: number): MoodStat;

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
    title: string,
    content: string,
    entriesList: Array<EntryWithDate.AsObject>,
    statsList: Array<MoodStat.AsObject>,
  }
}

export class CreateMoodRequest extends jspb.Message {
  getTitle(): string;
  setTitle(value: string): void;

  getContent(): string;
  setContent(value: string): void;

  getNumberOfRecordsNeeded(): number;
  setNumberOfRecordsNeeded(value: number): void;

  clearEmailsList(): void;
  getEmailsList(): Array<string>;
  setEmailsList(value: Array<string>): void;
  addEmails(value: string, index?: number): string;

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
    emailsList: Array<string>,
  }
}

export class EntrySigned extends jspb.Message {
  getEntryId(): string;
  setEntryId(value: string): void;

  getEntrySignature(): string;
  setEntrySignature(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): EntrySigned.AsObject;
  static toObject(includeInstance: boolean, msg: EntrySigned): EntrySigned.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: EntrySigned, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): EntrySigned;
  static deserializeBinaryFromReader(message: EntrySigned, reader: jspb.BinaryReader): EntrySigned;
}

export namespace EntrySigned {
  export type AsObject = {
    entryId: string,
    entrySignature: string,
  }
}

export class CreateMoodResponse extends jspb.Message {
  getMoodId(): string;
  setMoodId(value: string): void;

  getMoodSignature(): string;
  setMoodSignature(value: string): void;

  clearEntriesIdsList(): void;
  getEntriesIdsList(): Array<EntrySigned>;
  setEntriesIdsList(value: Array<EntrySigned>): void;
  addEntriesIds(value?: EntrySigned, index?: number): EntrySigned;

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
    moodId: string,
    moodSignature: string,
    entriesIdsList: Array<EntrySigned.AsObject>,
  }
}

