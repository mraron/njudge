import { useTranslation } from "react-i18next";
import { Navigate, useNavigate } from "react-router-dom";
import { useContext, useState } from "react";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import RoundedFrame from "../../components/container/RoundedFrame";
import TextBox from "../../components/input/TextBox";
import SVGTitleComponent from "../../components/svg/SVGTitleComponent";
import { routeMap } from "../../config/RouteConfig";
import { register } from "../../util/auth";
import UserContext from "../../contexts/user/UserContext";

function RegisterFrame() {
    const { t } = useTranslation();
    const [username, setUsername] = useState("");
    const [password, setPassword] = useState("");
    const [passwordConfirm, setPasswordConfirm] = useState("");
    const [email, setEmail] = useState("");
    const navigate = useNavigate();
    const titleComponent = (
        <SVGTitleComponent
            svg={
                <FontAwesomeIcon icon="fa-user-plus" className="w-5 h-5 mr-3" />
            }
            title={t("register.register")}
        />
    );
    const handleRegister = (event) => {
        event.preventDefault();
        register(username, email, password, passwordConfirm).then((resp) => {
            if (resp.success) {
                window.flash("flash.successful_registration", "success");
                navigate(routeMap.home);
            } else {
                window.flash(resp.message, "failure");
            }
        });
    };
    return (
        <RoundedFrame titleComponent={titleComponent}>
            <form method="POST">
                <div className="px-10 pt-8 pb-6 border-b border-bordercol">
                    <div className="mb-4 relative">
                        <TextBox
                            id="username"
                            label={t("register.username")}
                            initText={username}
                            onChange={setUsername}
                        />
                    </div>
                    <TextBox
                        id="email"
                        label={t("register.email")}
                        initText={email}
                        onChange={setEmail}
                    />
                </div>
                <div className="px-10 pt-4 pb-8">
                    <div className="mb-4">
                        <TextBox
                            id="password"
                            label={t("register.password")}
                            type="password"
                            initText={password}
                            onChange={setPassword}
                        />
                    </div>
                    <div className="mb-6">
                        <TextBox
                            id="passwordConfirm"
                            label={t("register.confirm_password")}
                            type="password"
                            initText={passwordConfirm}
                            onChange={setPasswordConfirm}
                        />
                    </div>
                    <div className="flex justify-center">
                        <button
                            type="submit"
                            className="btn-indigo padding-btn-default w-40"
                            onClick={handleRegister}>
                            {t("register.register")}
                        </button>
                    </div>
                </div>
            </form>
        </RoundedFrame>
    );
}

function Register() {
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
                    <RegisterFrame />
                </div>
            </div>
        </div>
    );
}

export default Register;
