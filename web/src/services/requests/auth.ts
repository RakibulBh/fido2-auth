import { toast } from "react-toastify";

interface response {
  data?: any;
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

  const data = await response.json();
  toast(data.message);

  console.log(data);
};
