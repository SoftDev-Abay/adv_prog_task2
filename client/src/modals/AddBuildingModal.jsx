import { React, useState } from "react";
import FormComponentDescription from "../components/form/FormComponentDescription";
import FormComponentImage from "../components/form/FormComponentImage";
import FormComponentCategory from "../components/form/FormComponentCategory";
import FormComponentNumbers from "../components/form/FormComponentNumbers";
import FormComponentLocation from "../components/form/FormComponentLocation";
import axios from "axios";

const AddBuildingModal = ({ setIsOpenAddBuildingModal }) => {
  const [formComponentIndex, setFormComponentIndex] = useState(0);
  const [imageUrl, setImageUrl] = useState("");
  const [category, setCategory] = useState("");
  const [description, setDescription] = useState("");
  const [title, setTitle] = useState("");
  const [roomsNum, setRoomsNum] = useState(0);
  const [bathroomsNum, setBathroomsNum] = useState(0);
  const [guestsNum, setGuestsNum] = useState(0);
  const [coutry, setCoutry] = useState("");
  const [address, setAddress] = useState("");
  const [city, setCity] = useState("");

  const FormComponents = [
    <FormComponentDescription
      description={description}
      setDescription={setDescription}
      title={title}
      setTitle={setTitle}
    />,
    <FormComponentImage imageUrl={imageUrl} setImageUrl={setImageUrl} />,
    <FormComponentCategory category={category} setCategory={setCategory} />,
    <FormComponentNumbers
      guests={guestsNum}
      setGuests={setGuestsNum}
      rooms={roomsNum}
      setRooms={setRoomsNum}
      bathrooms={bathroomsNum}
      setBathrooms={setBathroomsNum}
    />,
    <FormComponentLocation
      country={coutry}
      setCountry={setCoutry}
      address={address}
      setAddress={setAddress}
      city={city}
      setCity={setCity}
    />,
  ];

  const sendForm = () => {

    const payload = {
      imageUrl,
      category,
      description,
      title,
      roomsNum,
      bathroomsNum,
      guestsNum,
      coutry,
      address,
      city,
    };
    console.log(payload);
    const requestOptions = {
      method: "POST",
      headers: {
          'Content-Type': 'application/json'
      },
      body: JSON.stringify(payload),
  }
    fetch(`http://localhost:3000/add/building`, requestOptions)
  };

  const checkIfFormIsValid = () => {
    if (formComponentIndex === 0) {
      return title.length > 0 && description.length > 0;
    } else if (formComponentIndex === 1) {
      return imageUrl.length > 0;
    } else if (formComponentIndex === 2) {
      return category.length > 0;
    } else if (formComponentIndex === 3) {
      return guestsNum > 0 && roomsNum > 0 && bathroomsNum > 0;
    } else if (formComponentIndex === 4) {
      return coutry.length > 0 && address.length > 0 && city.length > 0;
    }
  };

  return (
    <>
      <div
        id="default-modal"
        tabIndex={-1} // This is important for accessibility
        aria-hidden="false"
        className="flex bg-black bg-opacity-50 overflow-y-auto overflow-x-hidden fixed top-0 right-0 left-0 z-50 justify-center items-center w-full md:inset-0 h-[100%] max-h-full"
      >
        <div className="relative p-4 w-full max-w-2xl max-h-full">
          {/* <!-- Modal content --> */}
          <div className="relative bg-white rounded-lg shadow dark:bg-gray-700">
            {/* <!-- Modal header --> */}
            <div className="flex items-center justify-between p-4 md:p-5 border-b rounded-t dark:border-gray-600">
              <h3 className="text-xl font-semibold text-gray-900 dark:text-white">
                Add new home
              </h3>
              <button
                type="button"
                className="text-gray-400 bg-transparent hover:bg-gray-200 hover:text-gray-900 rounded-lg text-sm w-8 h-8 ms-auto inline-flex justify-center items-center dark:hover:bg-gray-600 dark:hover:text-white"
                data-modal-hide="default-modal"
                onClick={() => {
                  setIsOpenAddBuildingModal(false);
                }}
              >
                <svg
                  className="w-3 h-3"
                  aria-hidden="true"
                  xmlns="http://www.w3.org/2000/svg"
                  fill="none"
                  viewBox="0 0 14 14"
                >
                  <path
                    stroke="currentColor"
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    stroke-width="2"
                    d="m1 1 6 6m0 0 6 6M7 7l6-6M7 7l-6 6"
                  />
                </svg>
                <span className="sr-only">Close modal</span>
              </button>
            </div>
            {/* <!-- Modal body --> */}

            <div className="p-4 md:p-5 mb-6 space-y-8">
              {FormComponents.map((FormComponent, index) => {
                return (
                  <div
                    className={`${
                      formComponentIndex === index ? "" : "hidden"
                    }`}
                  >
                    {FormComponent}
                  </div>
                );
              })}
            </div>
            {/* <!-- Modal footer --> */}
            <div className="flex justify-between items-center p-4 md:p-5 border-t border-gray-200 rounded-b dark:border-gray-600">
              <button
                data-modal-hide="default-modal"
                type="button"
                className="ms-3 text-gray-500 bg-white hover:bg-gray-50  rounded-lg border border-gray-200 text-sm font-medium px-5 py-2.5 hover:text-gray-900 focus:z-10"
                onClick={() => {
                  if (formComponentIndex === 0) {
                    setIsOpenAddBuildingModal(false);
                  } else {
                    setFormComponentIndex(formComponentIndex - 1);
                  }
                }}
              >
                Back
              </button>

              {formComponentIndex === FormComponents.length - 1 ? (
                <button
                  data-modal-hide="default-modal"
                  type="button"
                  className="text-white bg-red-500 hover:bg-red-600 focus:ring-4 focus:outline-none focus:ring-red-300 font-medium rounded-lg text-sm px-5 py-2.5 text-center "
                  onClick={() => {
                    checkIfFormIsValid() && sendForm();
                    setIsOpenAddBuildingModal(false);
                  }}
                >
                  Add home
                </button>
              ) : (
                <button
                  data-modal-hide="default-modal"
                  type="button"
                  className="text-white bg-red-500 hover:bg-red-600 focus:ring-4 focus:outline-none focus:ring-red-300 font-medium rounded-lg text-sm px-5 py-2.5 text-center "
                  onClick={() => {
                    checkIfFormIsValid() &&
                      setFormComponentIndex(formComponentIndex + 1);
                  }}
                >
                  Next
                </button>
              )}
            </div>
          </div>
        </div>
      </div>
    </>
  );
};

export default AddBuildingModal;
