import { useState } from "react";

function TextBox({
    id,
    label,
    type = "",
    initText = "",
    onChange,
    onFocus,
    onBlur,
}) {
    const [focused, setFocused] = useState(false);
    const handleChange = (event) => {
        if (onChange) {
            onChange(event.target.value);
        }
    };
    const handleFocus = () => {
        setFocused(true);
        if (onFocus) {
            onFocus();
        }
    };
    const handleBlur = () => {
        setFocused(false);
        if (onBlur) {
            onBlur();
        }
    };
    return (
        <div>
            <label htmlFor={id} className="text-label">
                {label}
            </label>
            <div
                className={`border-b-1 ${
                    focused
                        ? "border-indigo-600"
                        : "border-transparent"
                } w-full mt-1`}>
                <div
                    className={`border-b-1 ${
                        focused
                            ? "border-indigo-600"
                            : "border-bordercol"
                    } w-full`}>
                    <input
                        id={id}
                        type={type}
                        value={initText}
                        onChange={handleChange}
                        onFocus={handleFocus}
                        onBlur={handleBlur}
                        className={`py-1.5 px-2 bg-grey-850 border border-b-0 ${
                            focused ? "border-grey-575" : "border-grey-650"
                        } w-full outline-none`}
                    />
                </div>
            </div>
        </div>
    );
}

export default TextBox;
