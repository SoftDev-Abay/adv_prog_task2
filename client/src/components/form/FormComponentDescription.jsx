const FormComponentDescription = ({
  description,
  setDescription,
  title,
  setTitle,
}) => {
  return (
    <div className="">
      <h1 className="leading-relaxed text-xl font-semibold dark:text-gray-400">
        How would you describe your place?
      </h1>
      <p className="text-base leading-relaxed mt-2 text-gray-500 dark:text-gray-400">
        Short and consice.
      </p>
      <div className="space-y-4 mt-6">
        <input
          type="text"
          className="border-solid border border-gray-200 rounded-md w-full px-5 py-3.5 text-lg text-gray-700 placeholder-gray-400 focus:outline-none focus:border-transparent focus:ring-4 focus:ring-gray-300 "
          placeholder="Title"
          value={title}
          onChange={(e) => {
            setTitle(e.target.value);
          }}
        />
        <input
          type="text"
          className="border-solid border border-gray-200 rounded-md w-full px-5 py-3.5 text-lg text-gray-700 placeholder-gray-400 focus:outline-none focus:border-transparent focus:ring-4 focus:ring-gray-300 "
          placeholder="Description"
          value={description}
          onChange={(e) => {
            setDescription(e.target.value);
          }}
        />
      </div>
    </div>
  );
};

export default FormComponentDescription;
