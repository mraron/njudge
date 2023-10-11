import { TERipple } from "tw-elements-react"

function Button({ type, color, minWidth, fullWidth, cls, onClick, children, ariaLabel }) {
    const rippleColor = color === "gray" ? "#808080" : "#000000"
    return (
        <TERipple
            rippleColor={rippleColor}
            className={`rounded-md ${minWidth ? "shrink-0" : ""} overflow-hidden ${fullWidth ? "w-full" : ""}`}>
            <button
                type={type}
                className={`${color === "indigo" ? "btn-indigo" : "btn-gray"} padding-btn-default ${
                    fullWidth ? "w-full" : ""
                } h-full ${cls}`}
                style={{ minWidth: minWidth }}
                onClick={onClick}
                aria-label={ariaLabel}>
                {children}
            </button>
        </TERipple>
    )
}

export default Button
