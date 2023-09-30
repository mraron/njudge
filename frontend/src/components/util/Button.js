import { TERipple } from "tw-elements-react";

function Button({ type, color, minWidth, fullWidth, onClick, children }) {
    const rippleColor = color === "gray" ? "#808080" : "#000000";
    return (
        <TERipple
            rippleColor={rippleColor}
            className={`rounded-md overflow-hidden ${
                fullWidth ? "w-full" : ""
            }`}>
            <button
                type={type}
                className={`${
                    color === "indigo" ? "btn-indigo" : "btn-gray"
                } padding-btn-default ${fullWidth ? "w-full" : ""}`}
                style={{ minWidth: minWidth }}
                onClick={onClick}>
                {children}
            </button>
        </TERipple>
    );
}

export default Button;
