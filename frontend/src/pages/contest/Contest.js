import { Outlet, useParams } from "react-router-dom"
import { useTranslation } from "react-i18next"
import { routeMap } from "../../config/RouteConfig"
import TabFrame from "../../components/container/TabFrame"

function Contest() {
    const { t } = useTranslation()
    let routeLabels = ["contest.problems", "contest.submissions", "contest.ranklist"]
    let routePatterns = [routeMap.contest, routeMap.contestSubmissions, routeMap.contestRanklist]
    const { contest } = useParams()
    const routes = routePatterns.map((item) => item.replace(":contest", contest))
    return (
        <div className="flex justify-center">
            <div className="w-full max-w-7xl">
                <div className="w-full px-3">
                    <TabFrame routes={routes} routePatterns={routePatterns} routeLabels={routeLabels.map(t)}>
                        <div className="relative w-full">
                            <Outlet />
                        </div>
                    </TabFrame>
                </div>
            </div>
        </div>
    )
}

export default Contest
