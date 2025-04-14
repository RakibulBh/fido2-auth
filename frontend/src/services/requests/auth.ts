import { toast } from "react-toastify";
import { CredentialType } from "../../types";
import { arrayBufferToBase64, base64ToArrayBuffer } from "../../lib/utils";

// interface response {
//   data?: any;
//   message: string;
//   error: boolean;
// }

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
    const response = await fetch(`${import.meta.env.VITE_API_URL}/gateway`, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({
        service: "begin-register",
        auth_payload: {
          email: email,
        },
      }),
    });

    const data: registerResponse = await response.json();

    console.log(data);

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
      `${import.meta.env.VITE_API_URL}/gateway`,
      {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({
          service: "complete-register",
          complete_register_payload: payload,
        }),
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

export const Login = async ({ email }: { email: string }) => {
  try {
    const response = await fetch(`${import.meta.env.VITE_API_URL}/gateway`, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({
        service: "begin-login",
        auth_payload: {
          email: email,
        },
      }),
    });

    const data = await response.json();

    console.log(data);
  } catch (e) {
    console.log(e);
  }
};
