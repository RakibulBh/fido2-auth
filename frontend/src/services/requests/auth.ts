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
  try {
    const response = await fetch(
      `${import.meta.env.VITE_API_URL}/auth/begin-register`,
      {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ email: email }),
      }
    );

    const data: registerResponse = await response.json();

    if (data.error || !data.data?.options) {
      console.error("Registration error:", data.message);
      return;
    }

    const options: globalThis.PublicKeyCredentialCreationOptions =
      data.data.options.publicKey;

    options.challenge = base64ToArrayBuffer(String(options.challenge));
    options.user.id = base64ToArrayBuffer(String(options.user.id));

    // Prompt user to authenticate
    const credential = (await navigator.credentials.create({
      publicKey: options,
    })) as CredentialType;

    if (!credential) {
      console.error("No credential returned");
      toast("Authentication failed");
      return;
    }

    // Log credential for debugging
    console.log("Credential:", {
      id: credential.id,
      rawId: credential.rawId,
      response: {
        attestationObject: credential.response.attestationObject,
        clientDataJSON: credential.response.clientDataJSON,
      },
      type: credential.type,
    });

    const payload = {
      id: credential.id,
      rawId: arrayBufferToBase64(credential.rawId),
      response: {
        attestationObject: arrayBufferToBase64(
          credential.response.attestationObject
        ),
        clientDataJSON: arrayBufferToBase64(credential.response.clientDataJSON),
      },
      type: credential.type,
      sessionKey: data.data.sessionKey,
    };

    const finishResponse = await fetch(
      `${import.meta.env.VITE_API_URL}/auth/complete-register`,
      {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(payload),
      }
    );

    const finishData = await finishResponse.json();

    if (finishData.error) {
      console.error("Finish registration error:", finishData.message);
      toast(finishData.message);
      return;
    }

    toast("Registration successful!");
  } catch (error) {
    console.error("Registration error:", error);
    toast("An error occurred during registration");
  }
};

function base64ToArrayBuffer(base64: string) {
  const base64Standard = base64.replace(/-/g, "+").replace(/_/g, "/");
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
  return window
    .btoa(binary)
    .replace(/\+/g, "-")
    .replace(/\//g, "_")
    .replace(/=/g, "");
}
