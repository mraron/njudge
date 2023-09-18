import {matchPath} from "react-router-dom";

function checkData(data, pathname) {
    return matchPath(data.route, pathname);

}

export default checkData