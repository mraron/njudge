import { useState } from "react";

export function BasicInputComponent({ id, initText, type, focused, onChange, onFocus, onBlur }) {
    return (
        <div className={`border-1 border-b-0 ${focused? "border-grey-575": "border-grey-650"} w-full`}>
            <input autoComplete="off" id={id} value={initText} type={type} onChange={onChange} onFocus={onFocus} onBlur={onBlur} className="py-1.5 px-2 bg-grey-850 border-b-2 border-default w-full focus:border-indigo-600 outline-none transition-all duration-200" />
        </div>
    )
}

export function SVGInputComponent(svg) {
    return function({ id, initText, type, focused, onChange, onFocus, onBlur }) {
        return (
            <div className="flex w-full">
                <div
                    className={`flex items-center justify-center ${focused ? "bg-grey-650 border-grey-575" : "bg-grey-750 border-grey-650"} border-1 px-2 py-1 w-10 transition duration-200`}>
                    {svg}
                </div>
                <div
                    className={`border-1 border-b-0 border-l-0 ${focused ? "border-grey-575" : "border-grey-650"} mt-1 w-full`}>
                    <input autoComplete="off" id={id} value={initText} type={type} onChange={onChange} onFocus={onFocus}
                           onBlur={onBlur}
                           className="py-1.5 px-2 bg-grey-850 border-b-2 border-default w-full focus:border-indigo-600 outline-none transition-all duration-200"/>
                </div>
            </div>
        )
    }
}

function TextBox({ id, label, type, initText, inputComponent: InputComponent, onChange, onFocus, onBlur }) {
    initText ||= ""; type ||= "text"; InputComponent ||= BasicInputComponent;

    const [focused, setFocused] = useState(false);
    const handleChange = (event) => onChange? onChange(event.target.value): {};
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
            <div className="mt-1">
                <InputComponent id={id} initText={initText} focused={focused} onChange={handleChange} onFocus={handleFocus} onBlur={handleBlur} />
            </div>
        </div>
    )
}

export default TextBox;