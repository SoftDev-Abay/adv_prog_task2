import { CiCirclePlus, CiCircleMinus } from "react-icons/ci";
import { useState } from "react";
const FormComponentNumbers = () => {
  const [guests, setGuests] = useState(0);
  const [rooms, setRooms] = useState(0);
  const [bathrooms, setBathrooms] = useState(0);

  return (
    <div>
      <h1 className="leading-relaxed text-xl font-semibold dark:text-gray-400">
        Basic information about your place.
      </h1>
      <p className="text-base leading-relaxed mt-2 text-gray-500 dark:text-gray-400">
        Get the numbers right.
      </p>
      <div className="space-y-4 mt-6">
        <div className="pb-8 border-b-2  flex items-center justify-between">
          <div className="flex flex-col">
            <span
              className=" 
                        text-md
                        
                        font-medium
                        "
            >
              Guests
            </span>
            <span className="text-sm mt-2 text-gray-500">
              How many guests can your place accommodate?
            </span>
          </div>
          <div className="flex items-center gap-4 text-xl text-gray-500">
            <CiCircleMinus
              className="text-4xl cursor-pointer"
              onClick={() => {
                if (guests > 0) {
                  setGuests(guests - 1);
                }
              }}
            />
            <span>{guests}</span>
            <CiCirclePlus
              className="text-4xl cursor-pointer"
              onClick={() => {
                setGuests(guests + 1);
              }}
            />
          </div>
        </div>
        <div className="pb-8 border-b-2  flex items-center justify-between">
          <div className="flex flex-col">
            <span
              className=" 
                        text-md
                        
                        font-medium
                        "
            >
              Rooms
            </span>
            <span className="text-sm mt-2 text-gray-500">
              How many rooms does your place have?
            </span>
          </div>
          <div className="flex items-center gap-4 text-xl text-gray-500">
            <CiCircleMinus
              className="text-4xl cursor-pointer"
              onClick={() => {
                if (rooms > 0) {
                  setRooms(rooms - 1);
                }
              }}
            />
            <span>{rooms}</span>
            <CiCirclePlus
              className="text-4xl cursor-pointer"
              onClick={() => {
                setRooms(rooms + 1);
              }}
            />
          </div>
        </div>
        <div className="pb-8 border-b-2  flex items-center justify-between">
          <div className="flex flex-col">
            <span
              className=" 
                        text-md
                        
                        font-medium
                        "
            >
              Bathrooms
            </span>
            <span className="text-sm mt-2 text-gray-500">
              How many bathrooms does your place have?
            </span>
          </div>
          <div className="flex items-center gap-4 text-xl text-gray-500">
            <CiCircleMinus
              className="text-4xl cursor-pointer"
              onClick={() => {
                if (bathrooms > 0) {
                  setBathrooms(bathrooms - 1);
                }
              }}
            />
            <span>{bathrooms}</span>
            <CiCirclePlus
              className="text-4xl cursor-pointer"
              onClick={() => {
                setBathrooms(bathrooms + 1);
              }}
            />
          </div>
        </div>
      </div>
    </div>
  );
};

export default FormComponentNumbers;
