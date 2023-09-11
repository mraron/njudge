function Checkbox({ id, label }) {
    return (
        <label htmlFor={id} className="flex items-start max-w-fit">
            <input id={id} className="appearance-none bg-grey-850 text-white border-1 border-default rounded w-5 h-5 shrink-0 checked:bg-indigo-600 checked:border-indigo-600 checkmark hover:bg-grey-800 hover:border-grey-600 checked:hover:bg-indigo-500 checked:hover:border-indigo-500 transition duration-200" type="checkbox" />
            <span className="text-label ml-3">{label}</span>
        </label>
    )
}

export default Checkbox;