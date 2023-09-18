import RoundedFrame from '../../components/container/RoundedFrame';
import TextBox from '../../components/input/TextBox';
import {SVGAvatar, SVGGoogle} from '../../svg/SVGs';
import SVGTitleComponent from '../../svg/SVGTitleComponent';
import {login} from "../../util/Auth";
import {useState} from "react";
import {useNavigate} from "react-router-dom";
import {routeMap} from "../../config/RouteConfig";

function LoginFrame() {
    const navigate = useNavigate()
    const [username, setUsername] = useState("")
    const [password, setPassword] = useState("")
    const titleComponent = <SVGTitleComponent svg={<SVGAvatar cls="w-[1.1rem] h-[1.1rem] mr-2"/>} title="Belépés"/>
    const handleLogin = () => {
        login(username, password).then(resp => {
            window.flash(resp.message, resp.success? "success": "failure")
        }).then(() => {
            navigate(routeMap.main)
        })
    }
    return (
        <RoundedFrame titleComponent={titleComponent}>
            <div className="px-10 py-8">
                <div className="mb-4">
                    <TextBox id="userName" label="Felhasználónév" initText={username} onChange={(newText) => setUsername(newText)} />
                </div>
                <div className="mb-6">
                    <TextBox id="password" label="Jelszó" initText={password} type="password" onChange={(newText) => setPassword(newText)}/>
                </div>
                <div className="flex justify-center mb-2">
                    <button className="btn-indigo mr-2 w-1/2" onClick={handleLogin}>Belépés</button>
                    <button className="relative btn-gray flex items-center justify-between w-1/2">
                        <div className="h-full flex items-center absolute left-2.5">
                            <SVGGoogle/>
                        </div>
                        <div className="w-full flex justify-center">
                            <span>Google</span>
                        </div>
                    </button>
                </div>
            </div>
        </RoundedFrame>
    )
}

function Login() {
    return (
        <div className="text-white">
            <div className="w-full flex justify-center">
                <div className="flex justify-center w-full sm:max-w-md">
                    <div className="w-full px-4">
                        <LoginFrame/>
                    </div>
                </div>
            </div>
        </div>
    );
}

export default Login;