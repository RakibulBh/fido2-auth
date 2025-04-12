import { toast } from "react-toastify";
import { CredentialType } from "../../types";

interface response {
  data?: any;
  message: string;
  error: boolean;
}

interface registerResponse {
  data?: {
    options: any;
    sessionKey: string;
  };
  message: string;
  error: boolean;
}
export const registerUser = async ({ email }: { email: string }) => {
  const response = await fetch(
    `${import.meta.env.VITE_API_URL}/auth/register`,
    {
      method: "POST",
      body: JSON.stringify({ email: email }),
    }
  );

  const data: registerResponse = await response.json();
  toast(data.message);

  if (data.error || !data.data?.options) {
    return;
  }

  const options: globalThis.PublicKeyCredentialCreationOptions =
    data.data.options.publicKey;

  options.challenge = base64ToArrayBuffer(String(options.challenge));
  options.user.id = base64ToArrayBuffer(String(options.user.id));

  // Prompt user to authenticate
  const credential = (await navigator.credentials.create({
    publicKey: options,
  })) as unknown as CredentialType;

  if (!credential) {
    return;
  }

  const credentialData = {
    id: credential.id,
    rawId: arrayBufferToBase64(credential.rawId),
    response: {
      attestationObject: arrayBufferToBase64(
        credential.response.clientDataJSON
      ),
      clientDataJSON: arrayBufferToBase64(credential.response.clientDataJSON),
    },
    type: credential.type,
  };

  console.log(credentialData);

  const finishResponse = await fetch("/complete-registration", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ credentialData, sessionKey: data.data.sessionKey }),
  });

  const finishData = await finishResponse.json();

  if (finishData.error) {
    toast(finishData.message);
    return;
  }

  toast("Reached the end!");
};

// Helper functions
function base64ToArrayBuffer(base64: string) {
  // Convert base64url to base64 by replacing - with + and _ with /
  const base64Standard = base64.replace(/-/g, "+").replace(/_/g, "/");
  // Add padding if needed
  const paddedBase64 = base64Standard.padEnd(
    base64Standard.length + ((4 - (base64Standard.length % 4)) % 4),
    "="
  );

  const binary = window.atob(paddedBase64);
  const len = binary.length;
  const bytes = new Uint8Array(len);
  for (let i = 0; i < len; i++) {
    bytes[i] = binary.charCodeAt(i);
  }
  return bytes.buffer;
}

function arrayBufferToBase64(buffer: ArrayBuffer) {
  let binary = "";
  const bytes = new Uint8Array(buffer);
  for (let i = 0; i < bytes.byteLength; i++) {
    binary += String.fromCharCode(bytes[i]);
  }
  return window.btoa(binary);
}
