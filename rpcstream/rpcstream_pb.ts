// @generated by protoc-gen-es v1.8.0 with parameter "target=ts"
// @generated from file github.com/aperturerobotics/starpc/rpcstream/rpcstream.proto (package rpcstream, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type {
  BinaryReadOptions,
  FieldList,
  JsonReadOptions,
  JsonValue,
  PartialMessage,
  PlainMessage,
} from '@bufbuild/protobuf'
import { Message, proto3 } from '@bufbuild/protobuf'

/**
 * RpcStreamPacket is a packet encapsulating data for a RPC stream.
 *
 * @generated from message rpcstream.RpcStreamPacket
 */
export class RpcStreamPacket extends Message<RpcStreamPacket> {
  /**
   * @generated from oneof rpcstream.RpcStreamPacket.body
   */
  body:
    | {
        /**
         * Init is the first packet in the stream.
         * Sent by the initiator.
         *
         * @generated from field: rpcstream.RpcStreamInit init = 1;
         */
        value: RpcStreamInit
        case: 'init'
      }
    | {
        /**
         * Ack is sent in response to Init.
         * Sent by the server.
         *
         * @generated from field: rpcstream.RpcAck ack = 2;
         */
        value: RpcAck
        case: 'ack'
      }
    | {
        /**
         * Data is the encapsulated data packet.
         *
         * @generated from field: bytes data = 3;
         */
        value: Uint8Array
        case: 'data'
      }
    | { case: undefined; value?: undefined } = { case: undefined }

  constructor(data?: PartialMessage<RpcStreamPacket>) {
    super()
    proto3.util.initPartial(data, this)
  }

  static readonly runtime: typeof proto3 = proto3
  static readonly typeName = 'rpcstream.RpcStreamPacket'
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: 'init', kind: 'message', T: RpcStreamInit, oneof: 'body' },
    { no: 2, name: 'ack', kind: 'message', T: RpcAck, oneof: 'body' },
    {
      no: 3,
      name: 'data',
      kind: 'scalar',
      T: 12 /* ScalarType.BYTES */,
      oneof: 'body',
    },
  ])

  static fromBinary(
    bytes: Uint8Array,
    options?: Partial<BinaryReadOptions>,
  ): RpcStreamPacket {
    return new RpcStreamPacket().fromBinary(bytes, options)
  }

  static fromJson(
    jsonValue: JsonValue,
    options?: Partial<JsonReadOptions>,
  ): RpcStreamPacket {
    return new RpcStreamPacket().fromJson(jsonValue, options)
  }

  static fromJsonString(
    jsonString: string,
    options?: Partial<JsonReadOptions>,
  ): RpcStreamPacket {
    return new RpcStreamPacket().fromJsonString(jsonString, options)
  }

  static equals(
    a: RpcStreamPacket | PlainMessage<RpcStreamPacket> | undefined,
    b: RpcStreamPacket | PlainMessage<RpcStreamPacket> | undefined,
  ): boolean {
    return proto3.util.equals(RpcStreamPacket, a, b)
  }
}

/**
 * RpcStreamInit is the first message in a RPC stream.
 *
 * @generated from message rpcstream.RpcStreamInit
 */
export class RpcStreamInit extends Message<RpcStreamInit> {
  /**
   * ComponentId is the identifier of the component making the request.
   *
   * @generated from field: string component_id = 1;
   */
  componentId = ''

  constructor(data?: PartialMessage<RpcStreamInit>) {
    super()
    proto3.util.initPartial(data, this)
  }

  static readonly runtime: typeof proto3 = proto3
  static readonly typeName = 'rpcstream.RpcStreamInit'
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    {
      no: 1,
      name: 'component_id',
      kind: 'scalar',
      T: 9 /* ScalarType.STRING */,
    },
  ])

  static fromBinary(
    bytes: Uint8Array,
    options?: Partial<BinaryReadOptions>,
  ): RpcStreamInit {
    return new RpcStreamInit().fromBinary(bytes, options)
  }

  static fromJson(
    jsonValue: JsonValue,
    options?: Partial<JsonReadOptions>,
  ): RpcStreamInit {
    return new RpcStreamInit().fromJson(jsonValue, options)
  }

  static fromJsonString(
    jsonString: string,
    options?: Partial<JsonReadOptions>,
  ): RpcStreamInit {
    return new RpcStreamInit().fromJsonString(jsonString, options)
  }

  static equals(
    a: RpcStreamInit | PlainMessage<RpcStreamInit> | undefined,
    b: RpcStreamInit | PlainMessage<RpcStreamInit> | undefined,
  ): boolean {
    return proto3.util.equals(RpcStreamInit, a, b)
  }
}

/**
 * RpcAck is the ack message in a RPC stream.
 *
 * @generated from message rpcstream.RpcAck
 */
export class RpcAck extends Message<RpcAck> {
  /**
   * Error indicates there was some error setting up the stream.
   *
   * @generated from field: string error = 1;
   */
  error = ''

  constructor(data?: PartialMessage<RpcAck>) {
    super()
    proto3.util.initPartial(data, this)
  }

  static readonly runtime: typeof proto3 = proto3
  static readonly typeName = 'rpcstream.RpcAck'
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: 'error', kind: 'scalar', T: 9 /* ScalarType.STRING */ },
  ])

  static fromBinary(
    bytes: Uint8Array,
    options?: Partial<BinaryReadOptions>,
  ): RpcAck {
    return new RpcAck().fromBinary(bytes, options)
  }

  static fromJson(
    jsonValue: JsonValue,
    options?: Partial<JsonReadOptions>,
  ): RpcAck {
    return new RpcAck().fromJson(jsonValue, options)
  }

  static fromJsonString(
    jsonString: string,
    options?: Partial<JsonReadOptions>,
  ): RpcAck {
    return new RpcAck().fromJsonString(jsonString, options)
  }

  static equals(
    a: RpcAck | PlainMessage<RpcAck> | undefined,
    b: RpcAck | PlainMessage<RpcAck> | undefined,
  ): boolean {
    return proto3.util.equals(RpcAck, a, b)
  }
}