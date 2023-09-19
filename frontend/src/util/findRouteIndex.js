import {matchPath} from "react-router-dom";

export function findRouteIndex(patterns, pathname) {
    return patterns.findIndex(item => matchPath(item, pathname));
}