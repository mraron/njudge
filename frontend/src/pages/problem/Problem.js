import { Outlet, useParams } from "react-router-dom"
import { useTranslation } from "react-i18next"
import { routeMap } from "../../config/RouteConfig"
import TabFrame from "../../components/container/TabFrame"

const routeLabels = [
    "problem.statement",
    "problem.submit",
    "problem.submissions",
    "problem.ranklist",
]
const routePatterns = [
    routeMap.problem,
    routeMap.problemSubmit,
    routeMap.problemSubmissions,
    routeMap.problemRanklist,
]

function Problem() {
    const { t } = useTranslation()
    const { problem, problemset } = useParams()
    const routes = routePatterns.map((item) =>
        item.replace(":problemset", problemset).replace(":problem", problem),
    )
    return (
        <div className="flex justify-center">
            <div className="w-full max-w-7xl">
                <div className="w-full px-4">
                    <TabFrame
                        routes={routes}
                        routeLabels={routeLabels.map(t)}
                        routePatterns={routePatterns}>
                        <div className="relative w-full">
                            <Outlet />
                        </div>
                    </TabFrame>
                </div>
            </div>
        </div>
    )
}

export default Problem
