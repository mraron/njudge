import RoundedFrame from '../../components/container/RoundedFrame';
import TextBox from '../../components/input/TextBox';
import {SVGAvatar} from '../../svg/SVGs';
import SVGTitleComponent from '../../svg/SVGTitleComponent';
import {Navigate} from "react-router-dom";
import {routeMap} from "../../config/RouteConfig";
import {useContext} from "react";
import UserContext from "../../contexts/user/UserContext";
import {useTranslation} from "react-i18next";

function RegisterFrame() {
    const {t} = useTranslation()
    const titleComponent = <SVGTitleComponent svg={<SVGAvatar cls="w-[1.1rem] h-[1.1rem] mr-2"/>}
                                              title={t("register.register")}/>
    return (
        <RoundedFrame titleComponent={titleComponent}>
            <div className="px-10 pt-8 pb-6 border-b border-default">
                <div className="mb-4 relative">
                    <TextBox id="username" label={t("register.username")}/>
                </div>
                <TextBox id="email" label={t("register.email")}/>
            </div>
            <div className="px-10 pt-4 pb-8">
                <div className="mb-4">
                    <TextBox id="password" label={t("register.password")}/>
                </div>
                <div className="mb-6">
                    <TextBox id="passwordConfirm" label={t("register.confirm_password")}/>
                </div>
                <div className="flex justify-center">
                    <button className="btn-indigo w-40">{t("register.register")}</button>
                </div>
            </div>
        </RoundedFrame>
    )
}

function Register() {
    const {userData, isLoggedIn} = useContext(UserContext)
    if (isLoggedIn) {
        return <Navigate to={routeMap.profile.replace(":user", encodeURIComponent(userData.username))}/>
    }
    return (
        <div className="text-white">
            <div className="w-full flex justify-center">
                <div className="flex justify-center w-full sm:max-w-md">
                    <div className="w-full px-4">
                        <RegisterFrame/>
                    </div>
                </div>
            </div>
        </div>
    );
}

export default Register;