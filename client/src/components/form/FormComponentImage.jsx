import React, { useState } from "react";

const FormComponentImage = () => {
  const [imageUrl, setImageUrl] = useState("");
  return (
    <div className="">
      <h1 className="leading-relaxed text-xl font-semibold dark:text-gray-400">
        Add a phono of your place
      </h1>
      <p className="text-base leading-relaxed mt-2 text-gray-500 dark:text-gray-400">
        Link of the img.
      </p>
      <div className="space-y-4 mt-6">
        <input
          type="url"
          className="border-solid border border-gray-200 rounded-md w-full px-5 py-3.5 text-lg text-gray-700 placeholder-gray-400 focus:outline-none focus:border-transparent focus:ring-4 focus:ring-gray-300 "
          placeholder="URL"
          value={imageUrl}
          onChange={(e) => setImageUrl(e.target.value)}
        />
      </div>
      <div
        className="
        mt-6 min-h-[200px] flex justify-center items-center border-solid border border-gray-200 rounded-md w-full px-5 py-3.5 text-lg text-gray-700 placeholder-gray-400 focus:outline-none focus:border-transparent focus:ring-4 focus:ring-gray-300
        "
      >
        {imageUrl.length > 0 ? (
          <img
            className="w-full rounded-md object-contain"
            src={imageUrl}
            alt=""
          />
        ) : (
          <div className="text-gray-400">No image provided</div>
        )}
      </div>
    </div>
  );
};

export default FormComponentImage;
