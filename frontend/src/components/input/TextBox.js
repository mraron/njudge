import {useState} from "react";

function TextBox({id, label, type, initText, onChange, onFocus, onBlur}) {
    initText ||= "";
    type ||= "text";

    const [focused, setFocused] = useState(false);
    const handleChange = (event) => {
        if (onChange) {
            onChange(event.target.value)
        }
    }
    const handleFocus = () => {
        setFocused(true)
        if (onFocus) {
            onFocus()
        }
    }
    const handleBlur = () => {
        setFocused(false)
        if (onBlur) {
            onBlur()
        }
    }
    return (
        <div>
            <label htmlFor={id} className="text-label">{label}</label>
            <div
                className={`border-b-1 ${focused ? "border-indigo-600" : "border-transparent"} w-full mt-1 transition duration-200`}>
                <div
                    className={`border-b-1 ${focused ? "border-indigo-600" : "border-grey-650"} w-full transition duration-200`}>
                    <input id={id} value={initText} type={type} onChange={handleChange} onFocus={handleFocus}
                           onBlur={handleBlur}
                           className={`py-1.5 px-2 bg-grey-850 border border-b-0 ${focused ? "border-grey-575" : "border-grey-650"} w-full outline-none transition-all duration-200`}/>
                </div>
            </div>
        </div>
    )
}

export default TextBox;