// @generated by protoc-gen-es-lite unknown with parameter "target=ts,ts_nocheck=false"
// @generated from file github.com/aperturerobotics/starpc/e2e/mock/mock.proto (package e2e.mock, syntax proto3)
/* eslint-disable */

import {
  createMessageType,
  Message,
  MessageType,
  PartialFieldInfo,
} from '@aptre/protobuf-es-lite'

export const protobufPackage = 'e2e.mock'

/**
 * MockMsg is the mock message body.
 *
 * @generated from message e2e.mock.MockMsg
 */
export interface MockMsg extends Message<MockMsg> {
  /**
   * @generated from field: string body = 1;
   */
  body?: string
}

export const MockMsg: MessageType<MockMsg> = createMessageType({
  typeName: 'e2e.mock.MockMsg',
  fields: [
    { no: 1, name: 'body', kind: 'scalar', T: 9 /* ScalarType.STRING */ },
  ] as readonly PartialFieldInfo[],
  packedByDefault: true,
})
