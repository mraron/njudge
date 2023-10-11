import { Outlet, useParams } from "react-router-dom";
import { useTranslation } from "react-i18next";
import { routeMap } from "../../config/RouteConfig";
import TabFrame from "../../components/container/TabFrame";
import WidePage from "../wrappers/WidePage";

function Contest() {
    const { t } = useTranslation()
    let routeLabels = ["contest.problems", "contest.submissions", "contest.ranklist"]
    let routePatterns = [routeMap.contest, routeMap.contestSubmissions, routeMap.contestRanklist]
    const { contest } = useParams()
    const routes = routePatterns.map((item) => item.replace(":contest", contest))
    return (
        <WidePage>
            <TabFrame routes={routes} routePatterns={routePatterns} routeLabels={routeLabels.map(t)}>
                <div className="w-full">
                    <Outlet />
                </div>
            </TabFrame>
        </WidePage>
    )
}

export default Contest
