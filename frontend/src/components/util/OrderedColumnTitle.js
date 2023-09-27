import React from "react";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { useLocation, useNavigate } from "react-router-dom";
import queryString from "query-string";
import UpdateQueryString from "../../util/updateQueryString";

function OrderedColumnTitle({text, label}) {
    const location = useLocation()
    const navigate = useNavigate();
    const qData = queryString.parse(location.search);

    const currentOrder = qData.order === "ASC" &&
        qData.by === label ? 1 : qData.order === "DESC" &&
        qData.by === label ? 2 : 0

    const handleOrderChanged = () => {
        if (currentOrder === 2) {
            UpdateQueryString(location, navigate, [], [], null, ["by", "order"])
        } else {
            const newOrder = currentOrder === 1? 2 : 1
            UpdateQueryString(location, navigate, ["by", "order"], [label, newOrder === 1 ? "ASC" : "DESC"]);
        }
    };
    return (
        <div className="flex items-center link no-underline" onClick={handleOrderChanged}>
            <span>
                {text}
            </span>
            {currentOrder === 0 && <FontAwesomeIcon icon="fa-sort" className="w-3 h-3 ml-1.5"/>}
            {currentOrder === 1 && <FontAwesomeIcon icon="fa-sort-up" className="w-3 h-3 ml-1.5"/>}
            {currentOrder === 2 && <FontAwesomeIcon icon="fa-sort-down" className="w-3 h-3 ml-1.5"/>}
        </div>
    )
}

export default OrderedColumnTitle