import { FormEvent, useState } from "react";
import Input from "./components/Input";

function App() {
  const [screen, setScreen] = useState<number>(0);
  return (
    <main className="h-screen bg-[#141e1b] flex items-center justify-center bg-gradient-to-b from-[#141e1b] to-[#1a2421]">
      {screen == 0 ? <Register /> : <Login />}
    </main>
  );
}

const Register = () => {
  const onSubmit = (e: FormEvent) => {
    e.preventDefault();
    console.log();
  };

  return (
    <div className="bg-[#1a2421]/90 w-[60rem] xl:w-[80rem] h-[40rem] xl:h-[50rem] rounded-2xl flex overflow-hidden shadow-2xl border border-[#3c5045]/30">
      <div className="w-96 flex flex-col items-center pt-24 pb-16">
        {/* Register title */}
        <div className="text-center">
          <h1 className="font-bold text-2xl text-[#d3e4cd]">
            Sign up for free
          </h1>
          <p className="text-[#a0b9a6] text-sm mt-3">
            FIDO2 Authentication Demo
          </p>
        </div>

        {/* Form */}
        <form onSubmit={onSubmit} className="flex flex-col gap-4 w-72 mt-16">
          <Input
            id="email"
            label="Your email"
            type="email"
            placeholder="Enter your email"
          />
          <button
            type="submit"
            className="p-3 bg-[#3c5045] hover:bg-[#4a6154] transition-colors text-[#d3e4cd] text-sm rounded-xl mt-2 font-medium shadow-md"
          >
            Sign up with FIDO2
          </button>
        </form>
        {/* Dont have account */}
        <div className="text-center mt-4">
          <p className="text-[#a0b9a6]/70 text-xs">
            Already have an account?
            <button
              className="text-[#8fb996] ml-1 hover:underline"
              onClick={(e) => {
                e.preventDefault();
              }}
            >
              Log in
            </button>
          </p>
        </div>
      </div>

      {/* Image */}
      <div className="flex-1 relative">
        <div className="absolute inset-0 bg-gradient-to-r from-[#1a2421] to-transparent z-10"></div>
        <img
          src="/green-forest.jpg"
          className="absolute w-full h-full object-cover"
          alt="Forest background"
        />
      </div>
    </div>
  );
};

const Login = () => {
  return <div>Login</div>;
};

export default App;
