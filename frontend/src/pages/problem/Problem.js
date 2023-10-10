import { Outlet, useParams } from "react-router-dom"
import { useTranslation } from "react-i18next"
import { routeMap } from "../../config/RouteConfig"
import TabFrame from "../../components/container/TabFrame"
import WidePage from "../wrappers/WidePage"

const routeLabels = ["problem.statement", "problem.submit", "problem.submissions", "problem.ranklist"]
const routePatterns = [routeMap.problem, routeMap.problemSubmit, routeMap.problemSubmissions, routeMap.problemRanklist]

function Problem() {
    const { t } = useTranslation()
    const { problem, problemset } = useParams()
    const routes = routePatterns.map((item) => item.replace(":problemset", problemset).replace(":problem", problem))
    return (
        <WidePage>
            <TabFrame routes={routes} routeLabels={routeLabels.map(t)} routePatterns={routePatterns}>
                <div className="w-full">
                    <Outlet />
                </div>
            </TabFrame>
        </WidePage>
    )
}

export default Problem
