import { useState } from "react";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";

function Checkbox({ id, label, initChecked, onChange }) {
    const [checked, setChecked] = useState(initChecked);
    const [hovered, setHovered] = useState(false);
    const handleChange = (event) => {
        setChecked(event.target.checked);
        if (onChange) {
            onChange(event.target.checked);
        }
    };
    return (
        <label
            htmlFor={id}
            className="flex items-start justify-center max-w-fit space-x-3"
            onMouseOver={() => setHovered(true)}
            onMouseLeave={() => setHovered(false)}>
            <div
                className={`flex border items-center justify-center ${
                    checked
                        ? `border-transparent ${
                              hovered ? "bg-indigo-500" : "bg-indigo-600"
                          }`
                        : `${
                              hovered ? "bg-grey-825" : "bg-grey-850"
                          } border-bordefcol`
                } w-5 h-5 rounded-sm`}>
                <FontAwesomeIcon
                    icon="fa-check"
                    className={`w-3.5 h-3.5 text-white ${
                        checked ? "opacity-100" : "opacity-0"
                    }`}
                />
                <input
                    id={id}
                    onChange={handleChange}
                    className="appearance-none"
                    type="checkbox"
                    checked={initChecked}
                />
            </div>
            <span className="text-label">{label}</span>
        </label>
    );
}

export default Checkbox;
