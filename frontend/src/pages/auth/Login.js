import { useContext, useState } from "react";
import { Navigate, useNavigate } from "react-router-dom";
import { useTranslation } from "react-i18next";
import { SVGAvatar, SVGGoogle } from "../../components/svg/SVGs";
import RoundedFrame from "../../components/container/RoundedFrame";
import TextBox from "../../components/input/TextBox";
import SVGTitleComponent from "../../components/svg/SVGTitleComponent";
import { login } from "../../util/auth";
import { routeMap } from "../../config/RouteConfig";
import UserContext from "../../contexts/user/UserContext";

function LoginFrame() {
    const { t } = useTranslation();
    const navigate = useNavigate();
    const [username, setUsername] = useState("");
    const [password, setPassword] = useState("");
    const titleComponent = (
        <SVGTitleComponent
            svg={<SVGAvatar cls="w-5 h-5 mr-2" />}
            title={t("login.login")}
        />
    );
    const handleLogin = (event) => {
        event.preventDefault();
        login(username, password).then((resp) => {
            if (resp.success) {
                window.flash("flash.successful_login", "success");
                navigate(routeMap.home);
            } else {
                window.flash("flash.unsuccessful_login", "failure");
            }
        });
    };
    return (
        <RoundedFrame titleComponent={titleComponent}>
            <form method="POST">
                <div className="px-10 py-8">
                    <div className="mb-4">
                        <TextBox
                            id="userName"
                            label={t("login.username")}
                            initText={username}
                            onChange={setUsername}
                        />
                    </div>
                    <div className="mb-6">
                        <TextBox
                            id="password"
                            label={t("login.password")}
                            initText={password}
                            type="password"
                            onChange={setPassword}
                        />
                    </div>
                    <div className="flex justify-center mb-2">
                        <button
                            type="submit"
                            className="btn-indigo padding-btn-default mr-2 w-1/2"
                            onClick={handleLogin}>
                            {t("login.login")}
                        </button>
                        <button className="relative btn-gray padding-btn-default flex items-center justify-between w-1/2">
                            <div className="h-full flex items-center absolute left-2.5">
                                <SVGGoogle />
                            </div>
                            <div className="w-full flex justify-center">
                                <span>Google</span>
                            </div>
                        </button>
                    </div>
                </div>
            </form>
        </RoundedFrame>
    );
}

function Login() {
    const { userData, isLoggedIn } = useContext(UserContext);
    if (isLoggedIn) {
        return (
            <Navigate
                to={routeMap.profile.replace(
                    ":user",
                    encodeURIComponent(userData.username),
                )}
            />
        );
    }
    return (
        <div className="w-full flex justify-center">
            <div className="flex justify-center w-full sm:max-w-md">
                <div className="w-full px-4">
                    <LoginFrame />
                </div>
            </div>
        </div>
    );
}

export default Login;
