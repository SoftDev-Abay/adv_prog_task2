import { React, useState, useRef } from "react";
import ModelCard from "./ModelCard";
import { useAuthContext } from "../contex/AuthContext";
import axios from "axios";


const LoginModal = ({ setIsOpenLoginModal }) => {
  const { storeAndSetUser } = useAuthContext();
  const [currentProcces, setCurrentProcces] = useState("login");

  const closeModal = () => {
    setIsOpenLoginModal(false);
  };

  const logInEmailInputRef = useRef();
  const logInpasswordInputRef = useRef();

  const signUpEmailInputRef = useRef();
  const signUppasswordInputRef = useRef();
  const signUpConfirmPasswordInputRef = useRef();
  const signUpUsernameInputRef = useRef();

  const onLogin = async () => {
    const email = logInEmailInputRef.current.value;
    const password = logInpasswordInputRef.current.value;

    // after server is ready

    const res = await axios.post("http://localhost:3000/auth/login", {
      email,
      password,
    });
    console.log(res)
    if (res.status === "success") {
      storeAndSetUser(res.data.data.user);
      closeModal();
    } else {
      alert("Wrong email or password");
    }



    storeAndSetUser({ email, password, username: "test" });
    closeModal();
  };

  const onSignUp = async () => {
    const email = signUpEmailInputRef.current.value;
    const password = signUppasswordInputRef.current.value;
    const confirmPassword = signUpConfirmPasswordInputRef.current.value;
    const username = signUpUsernameInputRef.current.value;
    const phone_num = "87719001989"

    // after server is ready
    const res = await axios.post("http://localhost:3000/auth/signup", {
      email,
      password,
      username,
      phone_num,
    });
    console.log(res.statusText)
    console.log(res.status)
    if (res.status === "success") {
      alert("Account created successfully");
    } else {
      alert("Wrong email or password");
    }
    alert("Account created successfully");
  };

  const footerLoginElemenet = (
    <div>
      <p
        className="text-gray-500 text-center text-sm cursor-pointer"
        onClick={() => {
          setCurrentProcces("signup");
        }}
      >
        Don't have an account? <span className="text-red-500">Sign up</span>
      </p>
    </div>
  );

  const footerSignUpElemenet = (
    <div>
      <p
        className="text-gray-500 text-center text-sm cursor-pointer"
        onClick={() => {
          setCurrentProcces("login");
        }}
      >
        Already have an account? <span className="text-red-500">Login</span>
      </p>
    </div>
  );

  return (
    <>
      <ModelCard
        title="Login"
        closeModal={closeModal}
        footer={
          currentProcces === "login"
            ? footerLoginElemenet
            : footerSignUpElemenet
        }
      >
        {currentProcces === "login" ? (
          <div className="">
            <h1 className="custom-modal-heading">Welcome back to Airbnb!</h1>
            <p className="custom-modal-subtext"> Login to your account.</p>
            <div className="space-y-4 mt-6">
              <input
                type="email"
                className="custom-modal-input"
                placeholder="Email"
                ref={logInEmailInputRef}
              />
              <input
                type="text"
                className="custom-modal-input"
                placeholder="Passowrd"
                ref={logInpasswordInputRef}
              />
            </div>
          </div>
        ) : (
          <div className="">
            <h1 className="custom-modal-heading">
              Your journey starts with Airbnb!
            </h1>
            <p className="custom-modal-subtext">
              Sign up to start using Airbnb.
            </p>
            <div className="space-y-4 mt-6">
              <input
                type="email"
                className="custom-modal-input"
                placeholder="Email"
                ref={signUpEmailInputRef}
              />
              <input
                type="text"
                className="custom-modal-input"
                placeholder="Username"
                ref={signUpUsernameInputRef}
              />
              <input
                type="text"
                className="custom-modal-input"
                placeholder="Passowrd"
                ref={signUppasswordInputRef}
              />
              <input
                type="text"
                className="custom-modal-input"
                placeholder="Confirm Passowrd"
                ref={signUpConfirmPasswordInputRef}
              />
            </div>
          </div>
        )}

        <button
          onClick={currentProcces === "login" ? onLogin : onSignUp}
          className="mt-10 w-full text-white bg-red-500 hover:bg-red-600 focus:ring-4 focus:outline-none focus:ring-red-300 font-medium rounded-lg text-sm px-5 py-3 text-center "
        >
          Continue
        </button>
      </ModelCard>
    </>
  );
};

export default LoginModal;
