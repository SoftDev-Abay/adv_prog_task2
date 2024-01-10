import countries from "../../assets/countries.json";

const FormComponentLocation = () => {
  return (
    <div className="">
      <h1 className="leading-relaxed text-xl font-semibold dark:text-gray-400">
        Where is your place located?
      </h1>
      <p className="text-base leading-relaxed mt-2 text-gray-500 dark:text-gray-400">
        Exact location.
      </p>
      <div className="space-y-4 mt-6">
        <select
          type="text"
          className="border-solid border border-gray-200 rounded-md w-full px-5 py-3.5 text-lg text-gray-700 placeholder-gray-400 focus:outline-none focus:border-transparent focus:ring-4 focus:ring-gray-300 "
        >
          {countries.map((country) => {
            return <option value={country}>{country}</option>;
          })}
        </select>

        <input
          type="text"
          className="border-solid border border-gray-200 rounded-md w-full px-5 py-3.5 text-lg text-gray-700 placeholder-gray-400 focus:outline-none focus:border-transparent focus:ring-4 focus:ring-gray-300 "
          placeholder="City"
        />
        <input
          type="text"
          className="border-solid border border-gray-200 rounded-md w-full px-5 py-3.5 text-lg text-gray-700 placeholder-gray-400 focus:outline-none focus:border-transparent focus:ring-4 focus:ring-gray-300 "
          placeholder="Address"
        />
      </div>
    </div>
  );
};

export default FormComponentLocation;
