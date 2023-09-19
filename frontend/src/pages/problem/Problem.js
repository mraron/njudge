import React from "react";
import TabFrame from '../../components/container/TabFrame'
import {Outlet, useParams} from 'react-router-dom';
import {routeMap} from "../../config/RouteConfig";

const routeLabels = [
    "Leírás",
    "Beküld",
    "Beküldések",
    "Eredmények"
]
const routePatterns = [
    routeMap.problem,
    routeMap.problemSubmit,
    routeMap.problemSubmissions,
    routeMap.problemRanklist
]

function Problem() {
    const {problem} = useParams()
    const routes = routePatterns.map(item => item.replace(":problem", problem))

    return (
        <div className="flex justify-center">
            <div className="w-full max-w-7xl">
                <div className="w-full px-4">
                    <TabFrame routes={routes} routeLabels={routeLabels} routePatterns={routePatterns}>
                        <div className="relative w-full">
                            <Outlet/>
                        </div>
                    </TabFrame>
                </div>
            </div>
        </div>
    );
}

export default Problem;