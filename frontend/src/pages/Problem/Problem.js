import TabFrame from '../../components/TabFrame'
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
        <div className="text-white">
            <div className="w-full flex justify-center">
                <div className="flex justify-center w-full max-w-7xl">
                    <div className="w-full px-4">
                        <TabFrame routes={routes} routeLabels={routeLabels} routePatterns={routePatterns}>
                            <div className="w-full">
                                <Outlet />
                            </div>
                        </TabFrame>
                    </div>
                </div>
            </div>
        </div>
	);
}

export default Problem;