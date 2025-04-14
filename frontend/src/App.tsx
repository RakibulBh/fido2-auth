import { FormEvent, useState } from "react";
import Input from "./components/Input";
import { Login, registerUser } from "./services/requests/auth";
import { motion, AnimatePresence } from "framer-motion";

function App() {
  const [screen, setScreen] = useState<number>(0);
  return (
    <main className="min-h-screen w-full bg-[#141e1b] flex items-center justify-center bg-gradient-to-b from-[#141e1b] to-[#1a2421] p-4">
      <div className="bg-[#1a2421]/90 w-full max-w-7xl rounded-2xl flex flex-col md:flex-row overflow-hidden shadow-2xl border border-[#3c5045]/30">
        <AnimatePresence mode="wait">
          {screen == 0 ? (
            <motion.div
              key="register"
              initial={{ x: "-100%", opacity: 0 }}
              animate={{ x: 0, opacity: 1 }}
              exit={{ x: "100%", opacity: 0 }}
              transition={{ duration: 0.3 }}
              className="w-full md:w-96"
            >
              <RegisterForm setScreen={setScreen} />
            </motion.div>
          ) : (
            <motion.div
              key="login"
              initial={{ x: "100%", opacity: 0 }}
              animate={{ x: 0, opacity: 1 }}
              exit={{ x: "-100%", opacity: 0 }}
              transition={{ duration: 0.3 }}
              className="w-full md:w-96"
            >
              <LoginForm setScreen={setScreen} />
            </motion.div>
          )}
        </AnimatePresence>

        {/* Static image section */}
        <div className="flex-1 relative min-h-[200px] md:min-h-0">
          <div className="absolute inset-0 bg-gradient-to-r from-[#1a2421] to-transparent z-10 md:block"></div>
          <div className="absolute inset-0 bg-gradient-to-t from-[#1a2421] to-transparent z-10 md:hidden"></div>
          <img
            src="/green-forest.jpg"
            className="absolute w-full h-full object-cover"
            alt="Forest background"
          />
        </div>
      </div>
    </main>
  );
}

const RegisterForm = ({
  setScreen,
}: {
  setScreen: React.Dispatch<React.SetStateAction<number>>;
}) => {
  const [email, setEmail] = useState<string>("");
  const [loading, setLoading] = useState<boolean>(false);

  const onSubmit = (e: FormEvent) => {
    e.preventDefault();
    setLoading(true);
    registerUser({ email });
    setLoading(false);
  };

  return (
    <div className="flex flex-col items-center py-8 md:py-16 px-4">
      {/* Register title */}
      <div className="text-center">
        <h1 className="font-bold text-xl sm:text-2xl text-[#d3e4cd]">
          Sign up for free
        </h1>
        <p className="text-[#a0b9a6] text-xs sm:text-sm mt-2 sm:mt-3">
          FIDO2 Authentication Demo
        </p>
      </div>

      {/* Form */}
      <form
        onSubmit={onSubmit}
        className="flex flex-col gap-3 sm:gap-4 w-full max-w-xs mt-8 sm:mt-16"
      >
        <Input
          id="email"
          label="Your email"
          type="email"
          value={email}
          onChange={setEmail}
          placeholder="Enter your email"
        />
        <button
          type="submit"
          disabled={loading}
          className={`p-2.5 sm:p-3 ${
            loading
              ? "bg-[#3c5045]/50 hover:bg-[#3c5045]/50"
              : "bg-[#3c5045] hover:bg-[#4a6154]"
          } transition-colors text-[#d3e4cd] text-xs sm:text-sm rounded-xl mt-2 font-medium shadow-md flex items-center justify-center gap-2 disabled:cursor-not-allowed`}
        >
          {loading && (
            <span className="animate-spin size-3 sm:size-4 border-2 border-current border-t-transparent rounded-full"></span>
          )}
          {loading ? "Signing up..." : "Sign up with FIDO2"}
        </button>
      </form>
      {/* Dont have account */}
      <div className="text-center mt-3 sm:mt-4">
        <p className="text-[#a0b9a6]/70 text-[10px] sm:text-xs">
          Already have an account?
          <button
            className="text-[#8fb996] ml-1 hover:underline"
            onClick={() => {
              setScreen(1);
            }}
          >
            Log in
          </button>
        </p>
      </div>
    </div>
  );
};

const LoginForm = ({
  setScreen,
}: {
  setScreen: React.Dispatch<React.SetStateAction<number>>;
}) => {
  const [email, setEmail] = useState<string>("");
  const [loading, setLoading] = useState<boolean>(false);

  const onSubmit = async (e: FormEvent) => {
    e.preventDefault();
    setLoading(true);
    await Login({ email });
    setLoading(false);
  };

  return (
    <div className="flex flex-col items-center py-8 md:py-16 px-4">
      <div className="text-center">
        <h1 className="font-bold text-xl sm:text-2xl text-[#d3e4cd]">
          Welcome back
        </h1>
        <p className="text-[#a0b9a6] text-xs sm:text-sm mt-2 sm:mt-3">
          FIDO2 Authentication Demo
        </p>
      </div>

      {/* Form */}
      <form
        onSubmit={onSubmit}
        className="flex flex-col gap-3 sm:gap-4 w-full max-w-xs mt-8 sm:mt-16"
      >
        <Input
          id="email"
          label="Your email"
          type="email"
          value={email}
          onChange={setEmail}
          placeholder="Enter your email"
        />
        <button
          type="submit"
          disabled={loading}
          className={`p-2.5 sm:p-3 ${
            loading
              ? "bg-[#3c5045]/50 hover:bg-[#3c5045]/50"
              : "bg-[#3c5045] hover:bg-[#4a6154]"
          } transition-colors text-[#d3e4cd] text-xs sm:text-sm rounded-xl mt-2 font-medium shadow-md flex items-center justify-center gap-2 disabled:cursor-not-allowed`}
        >
          {loading && (
            <span className="animate-spin size-3 sm:size-4 border-2 border-current border-t-transparent rounded-full"></span>
          )}
          {loading ? "Logging in..." : "Log in with FIDO2"}
        </button>
      </form>
      {/* Dont have account */}
      <div className="text-center mt-3 sm:mt-4">
        <p className="text-[#a0b9a6]/70 text-[10px] sm:text-xs">
          Don't have an account?
          <button
            className="text-[#8fb996] ml-1 hover:underline"
            onClick={() => {
              setScreen(0);
            }}
          >
            Sign up
          </button>
        </p>
      </div>
    </div>
  );
};

export default App;
