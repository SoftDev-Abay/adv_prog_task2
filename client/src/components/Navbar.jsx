import React from "react";
import logo from "../assets/logo.svg";
const Navbar = () => {
  return (
    <div
      className="px-7 min-[500px]:px-20  w-full border-b-[2px] 
    border-gray-200
    py-4

  
  border-solid "
    >
      <div className=" mx-auto flex  justify-between items-center">
        <div className="flex items-center gap-1">
          <img src={logo} width={50} height={50} />
          <span className="text-2xl text-red font-semibold">airbnb</span>
        </div>
        {/* Find your perfect vacation rental. */}
        <button
          className=" 
            font-medium
          bg-red
          text-white
          px-5
          py-3
          rounded-md
          hover:opacity-80
          transition 
          duration-200
          ease-in-out

          "
        >
          Sign up
        </button>
      </div>
    </div>
  );
};

export default Navbar;
