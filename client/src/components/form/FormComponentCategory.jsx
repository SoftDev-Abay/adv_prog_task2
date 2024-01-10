import { categories } from "../Categories";

const FormComponentCategory = () => {
  return (
    <div>
      <h1 className="leading-relaxed text-xl font-semibold dark:text-gray-400">
        What kind of place are you listing?
      </h1>
      <p className="text-base leading-relaxed mt-2 text-gray-500 dark:text-gray-400">
        Choose a category.
      </p>
      <div className="grid grid-cols-2 gap-4 mt-6">
        {categories.map((category) => {
          const { label, icon: Icon } = category;
          return (
            <div
              tabIndex={-1}
              className="px-5 py-4 flex border-solid border-gray-200 border rounded-md text-gray-600 focus:border-transparent focus:ring-4 focus:ring-gray-300 hover:bg-gray-50 hover:text-black   cursor-pointer"
            >
              <div className=" flex flex-col   gap-1 font-semibold cursor-pointer ">
                <Icon className="text-4xl " />
                <span className="text-sm ">{label}</span>
              </div>
            </div>
          );
        })}
      </div>
    </div>
  );
};

export default FormComponentCategory;
