// @generated by protoc-gen-es v1.8.0 with parameter "target=ts"
// @generated from file github.com/aperturerobotics/starpc/e2e/mock/mock.proto (package e2e.mock, syntax proto3)
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
 * MockMsg is the mock message body.
 *
 * @generated from message e2e.mock.MockMsg
 */
export class MockMsg extends Message<MockMsg> {
  /**
   * @generated from field: string body = 1;
   */
  body = ''

  constructor(data?: PartialMessage<MockMsg>) {
    super()
    proto3.util.initPartial(data, this)
  }

  static readonly runtime: typeof proto3 = proto3
  static readonly typeName = 'e2e.mock.MockMsg'
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: 'body', kind: 'scalar', T: 9 /* ScalarType.STRING */ },
  ])

  static fromBinary(
    bytes: Uint8Array,
    options?: Partial<BinaryReadOptions>,
  ): MockMsg {
    return new MockMsg().fromBinary(bytes, options)
  }

  static fromJson(
    jsonValue: JsonValue,
    options?: Partial<JsonReadOptions>,
  ): MockMsg {
    return new MockMsg().fromJson(jsonValue, options)
  }

  static fromJsonString(
    jsonString: string,
    options?: Partial<JsonReadOptions>,
  ): MockMsg {
    return new MockMsg().fromJsonString(jsonString, options)
  }

  static equals(
    a: MockMsg | PlainMessage<MockMsg> | undefined,
    b: MockMsg | PlainMessage<MockMsg> | undefined,
  ): boolean {
    return proto3.util.equals(MockMsg, a, b)
  }
}