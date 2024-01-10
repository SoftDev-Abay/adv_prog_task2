import React from "react";
import logo from "../assets/logo.svg";
const Navbar = ({ setIsOpenAddBuildingModal }) => {
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
          <span className="text-2xl text-custom_red font-semibold">airbnb</span>
        </div>
        {/* Find your perfect vacation rental. */}
        <div className="flex items-center gap-5">
          <div
            className="cursor-pointer"
            onClick={() => {
              setIsOpenAddBuildingModal(true);
            }}
          >
            Add new home
          </div>

          <button
            className=" 
            font-medium
          bg-custom_red
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
    </div>
  );
};

export default Navbar;
