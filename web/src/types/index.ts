export type CredentialType = {
  authenticatorAttachment: "platform";
  id: string;
  rawId: ArrayBuffer;
  response: {
    attestationObject: ArrayBuffer;
    clientDataJSON: ArrayBuffer;
  };
  type: "public-key";
};
