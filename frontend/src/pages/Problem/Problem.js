import TabFrame from '../../components/TabFrame'
import { Outlet } from 'react-router-dom';

const routes = [
    ["Leírás", "/problemset/main/task/"], 
    ["Beküld", "/problemset/main/task/submit/"],
    ["Beküldések", "/problemset/main/task/status/"],
    ["Eredmények", "/problemset/main/task/ranklist/"]];

function Problem() {
	return (
        <div className="text-white">
            <div className="w-full flex justify-center">
                <div className="flex justify-center w-full max-w-7xl">
                    <div className="w-full px-4">
                        <TabFrame routes={routes}>
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