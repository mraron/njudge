import { TERipple } from "tw-elements-react";

function Button({ type, theme, minWidth, fullWidth, onClick, children }) {
    return (
        <TERipple className={`${fullWidth ? "w-full" : ""}`}>
            <button
                type={type}
                className={`${
                    theme === "indigo" ? "btn-indigo" : "btn-gray"
                } min-w-[${minWidth}] padding-btn-default ${
                    fullWidth ? "w-full" : ""
                }`}
                onClick={onClick}>
                {children}
            </button>
        </TERipple>
    );
}

export default Button;
