import { useState } from "react"

function TextBox({ id, label, type = "", initText = "", onChange, onFocus, onBlur }) {
    const [focused, setFocused] = useState(false)

    const handleChange = (event) => {
        onChange?.(event.target.value)
    }
    const handleFocus = () => {
        setFocused(true)
        onFocus?.()
    }
    const handleBlur = () => {
        setFocused(false)
        onBlur?.()
    }
    return (
        <div>
            <label htmlFor={id} className="text-label">
                {label}
            </label>
            <div
                className={`border-b-1 ${
                    focused ? "border-indigo-500 dark:border-indigo-600" : "border-transparent"
                } w-full mt-1`}>
                <div
                    className={`border-b-1 ${
                        focused ? "border-indigo-500 dark:border-indigo-600" : "border-bordefcol"
                    } w-full`}>
                    <input
                        id={id}
                        type={type}
                        value={initText}
                        onChange={handleChange}
                        onFocus={handleFocus}
                        onBlur={handleBlur}
                        className={`py-[0.45rem] px-2 bg-grey-850 border border-b-0 text-sm ${
                            focused ? "border-borstrcol" : "border-bordefcol"
                        } w-full outline-none`}
                    />
                </div>
            </div>
        </div>
    )
}

export default TextBox
