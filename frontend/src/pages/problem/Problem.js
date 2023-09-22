import React from "react";
import TabFrame from '../../components/container/TabFrame'
import {Outlet, useParams} from 'react-router-dom';
import {routeMap} from "../../config/RouteConfig";
import {useTranslation} from "react-i18next";

const routeLabels = [
    "problem.statement",
    "problem.submit",
    "problem.submissions",
    "problem.ranklist"
]
const routePatterns = [
    routeMap.problem,
    routeMap.problemSubmit,
    routeMap.problemSubmissions,
    routeMap.problemRanklist
]

function Problem({data}) {
    const {t} = useTranslation()
    const {problem} = useParams()
    const routes = routePatterns.map(item => item.replace(":problem", problem))

    return (
        <div className="flex justify-center">
            <div className="w-full max-w-7xl">
                <div className="w-full px-4">
                    <TabFrame routes={routes} routeLabels={routeLabels.map(t)} routePatterns={routePatterns}>
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